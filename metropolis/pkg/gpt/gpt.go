// Package gpt implements reading and writing GUID Partition Tables as specified
// in the UEFI Specification. It only implements up to 128 partitions per table
// (same as most other implementations) as more would require a dynamic table
// size, significantly complicating the code for little gain.
package gpt

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"strings"
	"unicode/utf16"

	"github.com/google/uuid"

	"source.monogon.dev/metropolis/pkg/blockdev"
	"source.monogon.dev/metropolis/pkg/msguid"
)

var gptSignature = [8]byte{'E', 'F', 'I', ' ', 'P', 'A', 'R', 'T'}
var gptRevision uint32 = 0x00010000 // First 2 bytes major, second 2 bytes minor

// See UEFI Specification 2.9 Table 5-5
type header struct {
	Signature   [8]byte
	Revision    uint32
	HeaderSize  uint32
	HeaderCRC32 uint32
	_           [4]byte

	HeaderBlock          uint64
	AlternateHeaderBlock uint64
	FirstUsableBlock     uint64
	LastUsableBlock      uint64

	ID [16]byte

	PartitionEntriesStartBlock uint64
	PartitionEntryCount        uint32
	PartitionEntrySize         uint32
	PartitionEntriesCRC32      uint32
}

// See UEFI Specification 2.9 Table 5-6
type partition struct {
	Type       [16]byte
	ID         [16]byte
	FirstBlock uint64
	LastBlock  uint64
	Attributes uint64
	Name       [36]uint16
}

var (
	PartitionTypeEFISystem = uuid.MustParse("C12A7328-F81F-11D2-BA4B-00A0C93EC93B")
)

// Attribute is a bitfield of attributes set on a partition. Bits 0 to 47 are
// reserved for UEFI specification use and all current assignments are in the
// following const block. Bits 48 to 64 are available for per-Type use by
// the organization controlling the partition Type.
type Attribute uint64

const (
	// AttrRequiredPartition indicates that this partition is required for the
	// platform to function. Mostly used by vendors to mark things like recovery
	// partitions.
	AttrRequiredPartition = 1 << 0
	// AttrNoBlockIOProto indicates that EFI firmware must not provide an EFI
	// block device (EFI_BLOCK_IO_PROTOCOL) for this partition.
	AttrNoBlockIOProto = 1 << 1
	// AttrLegacyBIOSBootable indicates to special-purpose software outside of
	// UEFI that this partition can be booted using a traditional PC BIOS.
	// Don't use this unless you know that you need it specifically.
	AttrLegacyBIOSBootable = 1 << 2
)

// PerTypeAttrs returns the top 24 bits which are reserved for custom per-Type
// attributes. The top 8 bits of the returned uint32 are always 0.
func (a Attribute) PerTypeAttrs() uint32 {
	return uint32(a >> 48)
}

// SetPerTypeAttrs sets the top 24 bits which are reserved for custom per-Type
// attributes. It does not touch the lower attributes which are specified by the
// UEFI specification. The top 8 bits of v are silently discarded.
func (a *Attribute) SetPerTypeAttrs(v uint32) {
	*a &= 0x000000FF_FFFFFFFF
	*a |= Attribute(v) << 48
}

type Partition struct {
	// Name of the partition, will be truncated if it expands to more than 36
	// UTF-16 code points. Not all systems can display non-BMP code points.
	Name string
	// Type is the type of Table partition, can either be one of the predefined
	// constants by the UEFI specification or a custom type identifier.
	// Note that the all-zero UUID denotes an empty partition slot, so this
	// MUST be set to something, otherwise it is not treated as a partition.
	Type uuid.UUID
	// ID is a unique identifier for this specific partition. It should be
	// changed when cloning the partition.
	ID uuid.UUID
	// The first logical block of the partition (inclusive)
	FirstBlock uint64
	// The last logical block of the partition (inclusive)
	LastBlock uint64
	// Bitset of attributes of this partition.
	Attributes Attribute

	*blockdev.Section
}

// SizeBlocks returns the size of the partition in blocks
func (p *Partition) SizeBlocks() uint64 {
	return 1 + p.LastBlock - p.FirstBlock
}

// IsUnused checks if the partition is unused, i.e. it is nil or its type is
// the null UUID.
func (p *Partition) IsUnused() bool {
	if p == nil {
		return true
	}
	return p.Type == uuid.Nil
}

// New returns an empty table on the given block device.
// It does not read any existing GPT on the disk (use Read for that), nor does
// it write anything until Write is called.
func New(b blockdev.BlockDev) (*Table, error) {
	return &Table{
		b: b,
	}, nil
}

type Table struct {
	// ID is the unique identifier of this specific disk / GPT.
	// If this is left uninitialized/all-zeroes a new random ID is automatically
	// generated when writing.
	ID uuid.UUID

	// Data put at the start of the very first block. Gets loaded and executed
	// by a legacy BIOS bootloader. This can be used to make GPT-partitioned
	// disks bootable by legacy systems or display a nice error message.
	// Maximum length is 440 bytes, if that is exceeded Write returns an error.
	// Should be left empty if the device is not bootable and/or compatibility
	// with BIOS booting is not required. Only useful on x86 systems.
	BootCode []byte

	// Partitions contains the list of partitions in this table. This is
	// artificially limited to 128 partitions. Holes in the partition list are
	// represented as nil values. Call IsUnused before checking any other
	// properties of the partition.
	Partitions []*Partition

	b blockdev.BlockDev
}

type addOptions struct {
	preferEnd        bool
	keepEmptyEntries bool
	alignment        int64
}

// AddOption is a bitset controlling various
type AddOption func(*addOptions)

// WithPreferEnd tries to put the partition as close to the end as possible
// instead of as close to the start.
func WithPreferEnd() AddOption {
	return func(options *addOptions) {
		options.preferEnd = true
	}
}

// WithKeepEmptyEntries does not fill up empty entries which are followed by
// filled ones. It always appends the partition after the last used entry.
// Without this flag, the partition is placed in the first empty entry.
func WithKeepEmptyEntries() AddOption {
	return func(options *addOptions) {
		options.keepEmptyEntries = true
	}
}

// WithAlignment allows aligning the partition start block to a non-default
// value. By default, these are aligned to 1MiB.
// Only use this flag if you are certain you need it, it can cause quite severe
// performance degradation under certain conditions.
func WithAlignment(alignmenet int64) AddOption {
	return func(options *addOptions) {
		options.alignment = alignmenet
	}
}

// AddPartition takes a pointer to a partition and adds it, placing it into
// the first (or last using WithPreferEnd) continuous free space which fits it.
// It writes the placement information (FirstBlock, LastBlock) back to p.
// By default, AddPartition aligns FirstBlock to 1MiB boundaries, but this can
// be overridden using WithAlignment.
func (gpt *Table) AddPartition(p *Partition, size int64, options ...AddOption) error {
	blockSize := gpt.b.BlockSize()
	var opts addOptions
	// Align to 1MiB or the block size, whichever is bigger
	opts.alignment = 1 * 1024 * 1024
	if blockSize > opts.alignment {
		opts.alignment = blockSize
	}
	for _, o := range options {
		o(&opts)
	}
	if opts.alignment%blockSize != 0 {
		return fmt.Errorf("requested alignment (%d bytes) is not an integer multiple of the block size (%d), unable to align", opts.alignment, blockSize)
	}
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	fs, _, err := gpt.GetFreeSpaces()
	if err != nil {
		return fmt.Errorf("unable to determine free space: %v", err)
	}
	if opts.preferEnd {
		// Reverse fs slice to start iteration at the end
		for i, j := 0, len(fs)-1; i < j; i, j = i+1, j-1 {
			fs[i], fs[j] = fs[j], fs[i]
		}
	}
	// Number of blocks the partition should occupy, rounded up.
	blocks := (size + blockSize - 1) / blockSize
	if size == -1 {
		var largestFreeSpace int64
		for _, freeInt := range fs {
			intSz := freeInt[1] - freeInt[0]
			if intSz > largestFreeSpace {
				largestFreeSpace = intSz
			}
		}
		blocks = largestFreeSpace
	}
	var maxFreeBlocks int64
	for _, freeInt := range fs {
		start := freeInt[0]
		end := freeInt[1]
		freeBlocks := end - start
		// Align start properly
		alignTo := opts.alignment / blockSize
		// Go doesn't implement the euclidean modulus, thus this construction
		// is necessary.
		paddingBlocks := ((alignTo - start) % alignTo) % alignTo
		freeBlocks -= paddingBlocks
		start += paddingBlocks
		if maxFreeBlocks < freeBlocks {
			maxFreeBlocks = freeBlocks
		}
		if freeBlocks >= blocks {
			if !opts.preferEnd {
				p.FirstBlock = uint64(start)
				p.LastBlock = uint64(start + blocks - 1)
			} else {
				// Realign FirstBlock. This will always succeed as
				// there is enough space to align to the start.
				moveLeft := (end - blocks - 1) % (opts.alignment / blockSize)
				p.FirstBlock = uint64(end - (blocks + 1 + moveLeft))
				p.LastBlock = uint64(end - (2 + moveLeft))
			}
			newPartPos := -1
			if !opts.keepEmptyEntries {
				for i, part := range gpt.Partitions {
					if part.IsUnused() {
						newPartPos = i
						break
					}
				}
			}
			if newPartPos == -1 {
				gpt.Partitions = append(gpt.Partitions, p)
			} else {
				gpt.Partitions[newPartPos] = p
			}
			p.Section = blockdev.NewSection(gpt.b, int64(p.FirstBlock), int64(p.LastBlock)+1)
			return nil
		}
	}

	return fmt.Errorf("no space for partition of %d blocks, largest continuous free space after alignment is %d blocks", blocks, maxFreeBlocks)
}

// FirstUsableBlock returns the first usable (i.e. a partition can start there)
// block.
func (gpt *Table) FirstUsableBlock() int64 {
	blockSize := gpt.b.BlockSize()
	partitionEntryBlocks := (16384 + blockSize - 1) / blockSize
	return 2 + partitionEntryBlocks
}

// LastUsableBlock returns the last usable (i.e. a partition can end there)
// block. This block is inclusive.
func (gpt *Table) LastUsableBlock() int64 {
	blockSize := gpt.b.BlockSize()
	partitionEntryBlocks := (16384 + blockSize - 1) / blockSize
	return gpt.b.BlockCount() - (2 + partitionEntryBlocks)
}

// GetFreeSpaces returns a slice of tuples, each containing a half-closed
// interval of logical blocks not occupied by the GPT itself or any partition.
// The returned intervals are always in ascending order as well as
// non-overlapping. It also returns if it detected any overlaps between
// partitions or partitions and the GPT. It returns an error if and only if any
// partition has its FirstBlock before the LastBlock or exceeds the amount of
// blocks on the block device.
//
// Note that the most common use cases for this function are covered by
// AddPartition, you're encouraged to use it instead.
func (gpt *Table) GetFreeSpaces() ([][2]int64, bool, error) {
	// This implements an efficient algorithm for finding free intervals given
	// a set of potentially overlapping occupying intervals. It uses O(n*log n)
	// time for n being the amount of intervals, i.e. partitions. It uses O(n)
	// additional memory. This makes it de facto infinitely scalable in the
	// context of partition tables as the size of the block device is not part
	// of its cyclomatic complexity and O(n*log n) is tiny for even very big
	// partition tables.

	blockCount := gpt.b.BlockCount()

	// startBlocks contains the start blocks (inclusive) of all occupied
	// intervals.
	var startBlocks []int64
	// endBlocks contains the end blocks (exclusive!) of all occupied intervals.
	// The interval at index i is given by [startBlock[i], endBlock[i]).
	var endBlocks []int64

	// Reserve the primary GPT interval including the protective MBR.
	startBlocks = append(startBlocks, 0)
	endBlocks = append(endBlocks, gpt.FirstUsableBlock())

	// Reserve the alternate GPT interval (needs +1 for exclusive interval)
	startBlocks = append(startBlocks, gpt.LastUsableBlock()+1)
	endBlocks = append(endBlocks, blockCount)

	for i, part := range gpt.Partitions {
		if part.IsUnused() {
			continue
		}
		// Bail if partition does not contain a valid interval. These are open
		// intervals, thus part.FirstBlock == part.LastBlock denotes a valid
		// partition with a size of one block.
		if part.FirstBlock > part.LastBlock {
			return nil, false, fmt.Errorf("partition %d has a LastBlock smaller than its FirstBlock, its interval is [%d, %d]", i, part.FirstBlock, part.LastBlock)
		}
		if part.FirstBlock >= uint64(blockCount) || part.LastBlock >= uint64(blockCount) {
			return nil, false, fmt.Errorf("partition %d exceeds the block count of the block device", i)
		}
		startBlocks = append(startBlocks, int64(part.FirstBlock))
		// Algorithm needs open-closed intervals, thus add +1 to the end.
		endBlocks = append(endBlocks, int64(part.LastBlock)+1)
	}
	// Sort both sets of blocks independently in ascending order. Note that it
	// is now no longer possible to extract the original intervals. Integers
	// have no identity thus it doesn't matter if the sort is stable or not.
	sort.Slice(startBlocks, func(i, j int) bool { return startBlocks[i] < startBlocks[j] })
	sort.Slice(endBlocks, func(i, j int) bool { return endBlocks[i] < endBlocks[j] })

	var freeSpaces [][2]int64

	// currentIntervals contains the number of intervals which contain the
	// position currently being iterated over. If currentIntervals is ever
	// bigger than 1, there is overlap within the given intervals.
	currentIntervals := 0
	var hasOverlap bool

	// Iterate for as long as there are interval boundaries to be processed.
	for len(startBlocks) != 0 || len(endBlocks) != 0 {
		// Short-circuit boundary processing. If an interval ends at x and the
		// next one starts at x (this is using half-open intervals), it would
		// otherwise perform useless processing as well as create an empty free
		// interval which would then need to be filtered back out.
		if len(startBlocks) != 0 && len(endBlocks) != 0 && startBlocks[0] == endBlocks[0] {
			startBlocks = startBlocks[1:]
			endBlocks = endBlocks[1:]
			continue
		}
		// Pick the lowest boundary from either startBlocks or endBlocks,
		// preferring endBlocks if they are equal. Don't try to pick from empty
		// slices.
		if (len(startBlocks) != 0 && len(endBlocks) != 0 && startBlocks[0] < endBlocks[0]) || len(endBlocks) == 0 {
			// If currentIntervals == 0 a free space region ends here.
			// Since this algorithm creates the free space interval at the end
			// of an occupied interval, for the first interval there is no free
			// space entry. But in this case it's fine to just ignore it as the
			// first interval always starts at 0 because of the GPT.
			if currentIntervals == 0 && len(freeSpaces) != 0 {
				freeSpaces[len(freeSpaces)-1][1] = startBlocks[0]
			}
			// This is the start of an interval, increase the number of active
			// intervals.
			currentIntervals++
			hasOverlap = hasOverlap || currentIntervals > 1
			// Drop processed startBlock from slice.
			startBlocks = startBlocks[1:]
		} else {
			// This is the end of an interval, decrease the number of active
			// intervals.
			currentIntervals--
			// If currentIntervals == 0 a free space region starts here.
			// Same as with the startBlocks, ignore a potential free block after
			// the final range as the GPT occupies the last blocks anyway.
			if currentIntervals == 0 && len(startBlocks) != 0 {
				freeSpaces = append(freeSpaces, [2]int64{endBlocks[0], 0})
			}
			endBlocks = endBlocks[1:]
		}
	}
	return freeSpaces, hasOverlap, nil
}

// Overhead returns the number of blocks the GPT partitioning itself consumes,
// i.e. aren't usable for user data.
func Overhead(blockSize int64) int64 {
	// 3 blocks + 2x 16384 bytes (partition entry space)
	partitionEntryBlocks := (16384 + blockSize - 1) / blockSize
	return 3 + (2 * partitionEntryBlocks)
}

// Write writes the two GPTs, first the alternate, then the primary to the
// block device. If gpt.ID or any of the partition IDs are the all-zero UUID,
// new random ones are generated and written back. If the output is supposed
// to be reproducible, generate the UUIDs beforehand.
func (gpt *Table) Write() error {
	blockSize := gpt.b.BlockSize()
	blockCount := gpt.b.BlockCount()
	if blockSize < 512 {
		return errors.New("block size is smaller than 512 bytes, this is unsupported")
	}
	// Layout looks as follows:
	// Block 0: Protective MBR
	// Block 1: GPT Header
	// Block 2-(16384 bytes): GPT partition entries
	// Block (16384 bytes)-n: GPT partition entries alternate copy
	// Block n: GPT Header alternate copy
	partitionEntryCount := 128
	if len(gpt.Partitions) > partitionEntryCount {
		return errors.New("bigger-than default GPTs (>128 partitions) are unimplemented")
	}

	partitionEntryBlocks := (16384 + blockSize - 1) / blockSize
	if blockCount < 3+(2*partitionEntryBlocks) {
		return errors.New("not enough blocks to write GPT")
	}

	if gpt.ID == uuid.Nil {
		gpt.ID = uuid.New()
	}

	partSize := binary.Size(partition{})
	var partitionEntriesData bytes.Buffer
	for i := 0; i < partitionEntryCount; i++ {
		if len(gpt.Partitions) <= i || gpt.Partitions[i] == nil {
			// Write an empty entry
			partitionEntriesData.Write(make([]byte, partSize))
			continue
		}
		p := gpt.Partitions[i]
		if p.ID == uuid.Nil {
			p.ID = uuid.New()
		}
		rawP := partition{
			Type:       msguid.From(p.Type),
			ID:         msguid.From(p.ID),
			FirstBlock: p.FirstBlock,
			LastBlock:  p.LastBlock,
			Attributes: uint64(p.Attributes),
		}
		nameUTF16 := utf16.Encode([]rune(p.Name))
		// copy will automatically truncate if target is too short
		copy(rawP.Name[:], nameUTF16)
		binary.Write(&partitionEntriesData, binary.LittleEndian, rawP)
	}

	hdr := header{
		Signature:  gptSignature,
		Revision:   gptRevision,
		HeaderSize: uint32(binary.Size(&header{})),
		ID:         msguid.From(gpt.ID),

		PartitionEntryCount: uint32(partitionEntryCount),
		PartitionEntrySize:  uint32(partSize),

		FirstUsableBlock: uint64(2 + partitionEntryBlocks),
		LastUsableBlock:  uint64(blockCount - (2 + partitionEntryBlocks)),
	}
	hdr.PartitionEntriesCRC32 = crc32.ChecksumIEEE(partitionEntriesData.Bytes())

	hdrChecksum := crc32.NewIEEE()

	// Write alternate header first, as otherwise resizes are unsafe. If the
	// alternate is currently not at the end of the block device, it cannot
	// be found. Thus if the write operation is aborted abnormally, the
	// primary GPT is corrupted and the alternate cannot be found because it
	// is not at its canonical location. Rewriting the alternate first avoids
	// this problem.

	// Alternate header
	hdr.HeaderBlock = uint64(blockCount - 1)
	hdr.AlternateHeaderBlock = 1
	hdr.PartitionEntriesStartBlock = uint64(blockCount - (1 + partitionEntryBlocks))

	hdrChecksum.Reset()
	hdr.HeaderCRC32 = 0
	binary.Write(hdrChecksum, binary.LittleEndian, &hdr)
	hdr.HeaderCRC32 = hdrChecksum.Sum32()

	for partitionEntriesData.Len()%int(blockSize) != 0 {
		partitionEntriesData.WriteByte(0x00)
	}
	if _, err := gpt.b.WriteAt(partitionEntriesData.Bytes(), int64(hdr.PartitionEntriesStartBlock)*blockSize); err != nil {
		return fmt.Errorf("failed to write alternate partition entries: %w", err)
	}

	var hdrRaw bytes.Buffer
	if err := binary.Write(&hdrRaw, binary.LittleEndian, &hdr); err != nil {
		return fmt.Errorf("failed to encode alternate header: %w", err)
	}
	for hdrRaw.Len()%int(blockSize) != 0 {
		hdrRaw.WriteByte(0x00)
	}
	if _, err := gpt.b.WriteAt(hdrRaw.Bytes(), (blockCount-1)*blockSize); err != nil {
		return fmt.Errorf("failed to write alternate header: %v", err)
	}

	// Primary header
	hdr.HeaderBlock = 1
	hdr.AlternateHeaderBlock = uint64(blockCount - 1)
	hdr.PartitionEntriesStartBlock = 2

	hdrChecksum.Reset()
	hdr.HeaderCRC32 = 0
	binary.Write(hdrChecksum, binary.LittleEndian, &hdr)
	hdr.HeaderCRC32 = hdrChecksum.Sum32()

	hdrRaw.Reset()

	if err := makeProtectiveMBR(&hdrRaw, blockCount, gpt.BootCode); err != nil {
		return fmt.Errorf("failed creating protective MBR: %w", err)
	}
	for hdrRaw.Len()%int(blockSize) != 0 {
		hdrRaw.WriteByte(0x00)
	}
	if err := binary.Write(&hdrRaw, binary.LittleEndian, &hdr); err != nil {
		panic(err)
	}
	for hdrRaw.Len()%int(blockSize) != 0 {
		hdrRaw.WriteByte(0x00)
	}
	hdrRaw.Write(partitionEntriesData.Bytes())
	for hdrRaw.Len()%int(blockSize) != 0 {
		hdrRaw.WriteByte(0x00)
	}

	if _, err := gpt.b.WriteAt(hdrRaw.Bytes(), 0); err != nil {
		return fmt.Errorf("failed to write primary GPT: %w", err)
	}
	return nil
}

// Read reads a Table from a block device.
func Read(r blockdev.BlockDev) (*Table, error) {
	if Overhead(r.BlockSize()) > r.BlockCount() {
		return nil, errors.New("disk cannot contain a GPT as the block count is too small to store one")
	}
	zeroBlock := make([]byte, r.BlockSize())
	if _, err := r.ReadAt(zeroBlock, 0); err != nil {
		return nil, fmt.Errorf("failed to read first block: %w", err)
	}

	var m mbr
	if err := binary.Read(bytes.NewReader(zeroBlock[:512]), binary.LittleEndian, &m); err != nil {
		panic(err) // Read is from memory and with enough data
	}
	// The UEFI standard says that the only acceptable MBR for a GPT-partitioned
	// device is a pure protective MBR with one partition of type 0xEE covering
	// the entire disk. But reality is sadly not so simple. People have come up
	// with hacks like Hybrid MBR which is basically a way to expose partitions
	// as both GPT partitions and MBR partitions. There are also GPTs without
	// any MBR at all.
	// Following the standard strictly when reading means that this library
	// would fail to read valid GPT disks where such schemes are employed.
	// On the other hand just looking at the GPT signature is also dangerous
	// as not all tools clear the second block where the GPT resides when
	// writing an MBR, which results in reading a wrong/obsolete GPT.
	// As a pragmatic solution this library treats any disk as GPT-formatted if
	// the first block does not contain an MBR signature or at least one MBR
	// partition has type 0xEE (GPT). It does however not care in which slot
	// this partition is or if it begins at the start of the disk.
	//
	// Note that the block signatures for MBR and FAT are shared. This is a
	// historical artifact from DOS. It is not reliably possible to
	// differentiate the two as either has boot code where the other has meta-
	// data and both lack any checksums. Because the MBR partition table is at
	// the very end of the FAT bootcode section the following code always
	// assumes that it is dealing with an MBR. This is both more likely and
	// the 0xEE marker is rarer and thus more specific than FATs 0x00, 0x80 and
	// 0x02.
	var bootCode []byte
	hasDOSBootSig := m.Signature == mbrSignature
	if hasDOSBootSig {
		var isGPT bool
		for _, p := range m.PartitionRecords {
			if p.Type == 0xEE {
				isGPT = true
			}
		}
		// Note that there is a small but non-zero chance that isGPT is true
		// for a raw FAT filesystem if the bootcode contains a "valid" MBR.
		// The next error message mentions that possibility.
		if !isGPT {
			return nil, errors.New("block device contains an MBR table without a GPT marker or a raw FAT filesystem")
		}
		// Trim right zeroes away as they are padded back when writing. This
		// makes BootCode empty when it is all-zeros, making it easier to work
		// with while still round-tripping correctly.
		bootCode = bytes.TrimRight(m.BootCode[:], "\x00")
	}
	// Read the primary GPT. If it is damaged and/or broken, read the alternate.
	primaryGPT, err := readSingleGPT(r, 1)
	if err != nil {
		alternateGPT, err2 := readSingleGPT(r, r.BlockCount()-1)
		if err2 != nil {
			return nil, fmt.Errorf("failed to read both GPTs: primary GPT (%v), secondary GPT (%v)", err, err2)
		}
		alternateGPT.BootCode = bootCode
		return alternateGPT, nil
	}
	primaryGPT.BootCode = bootCode
	return primaryGPT, nil
}

func readSingleGPT(r blockdev.BlockDev, headerBlockPos int64) (*Table, error) {
	hdrBlock := make([]byte, r.BlockSize())
	if _, err := r.ReadAt(hdrBlock, r.BlockSize()*headerBlockPos); err != nil {
		return nil, fmt.Errorf("failed to read GPT header block: %w", err)
	}
	hdrBlockReader := bytes.NewReader(hdrBlock)
	var hdr header
	if err := binary.Read(hdrBlockReader, binary.LittleEndian, &hdr); err != nil {
		panic(err) // Read from memory with enough bytes, should not fail
	}
	if hdr.Signature != gptSignature {
		return nil, errors.New("no GPT signature found")
	}
	if hdr.HeaderSize < uint32(binary.Size(hdr)) {
		return nil, fmt.Errorf("GPT header size is too small, likely corrupted")
	}
	if int64(hdr.HeaderSize) > r.BlockSize() {
		return nil, fmt.Errorf("GPT header size is bigger than block size, likely corrupted")
	}
	// Use reserved bytes to hash, but do not expose them to the user.
	// If someone has a need to process them, they should extend this library
	// with whatever an updated UEFI specification contains.
	// It has been considered to store these in the user-exposed GPT struct to
	// be able to round-trip them cleanly, but there is significant complexity
	// and risk involved in doing so.
	reservedBytes := hdrBlock[binary.Size(hdr):hdr.HeaderSize]
	hdrExpectedCRC := hdr.HeaderCRC32
	hdr.HeaderCRC32 = 0
	hdrCRC := crc32.NewIEEE()
	binary.Write(hdrCRC, binary.LittleEndian, &hdr)
	hdrCRC.Write(reservedBytes)
	if hdrCRC.Sum32() != hdrExpectedCRC {
		return nil, fmt.Errorf("GPT header checksum mismatch, probably corrupted")
	}
	if hdr.HeaderBlock != uint64(headerBlockPos) {
		return nil, errors.New("GPT header indicates wrong block")
	}
	if hdr.PartitionEntrySize < uint32(binary.Size(partition{})) {
		return nil, errors.New("partition entry size too small")
	}
	if hdr.PartitionEntriesStartBlock > uint64(r.BlockCount()) {
		return nil, errors.New("partition entry start block is out of range")
	}
	// Sanity-check total size of the partition entry area. Otherwise, this is a
	// trivial DoS as it could cause allocation of gigabytes of memory.
	// 4MiB is equivalent to around 45k partitions at the current size.
	// I know of no operating system which would handle even a fraction of this.
	if uint64(hdr.PartitionEntryCount)*uint64(hdr.PartitionEntrySize) > 4*1024*1024 {
		return nil, errors.New("partition entry area bigger than 4MiB, refusing to read")
	}
	partitionEntryData := make([]byte, hdr.PartitionEntrySize*hdr.PartitionEntryCount)
	if _, err := r.ReadAt(partitionEntryData, r.BlockSize()*int64(hdr.PartitionEntriesStartBlock)); err != nil {
		return nil, fmt.Errorf("failed to read partition entries: %w", err)
	}
	if crc32.ChecksumIEEE(partitionEntryData) != hdr.PartitionEntriesCRC32 {
		return nil, errors.New("GPT partition entry table checksum mismatch")
	}
	var g Table
	g.ID = msguid.To(hdr.ID)
	for i := uint32(0); i < hdr.PartitionEntryCount; i++ {
		entryReader := bytes.NewReader(partitionEntryData[i*hdr.PartitionEntrySize : (i+1)*hdr.PartitionEntrySize])
		var part partition
		if err := binary.Read(entryReader, binary.LittleEndian, &part); err != nil {
			panic(err) // Should not happen
		}
		// If the partition type is the all-zero UUID, this slot counts as
		// unused.
		if part.Type == uuid.Nil {
			g.Partitions = append(g.Partitions, nil)
			continue
		}
		g.Partitions = append(g.Partitions, &Partition{
			ID:         msguid.To(part.ID),
			Type:       msguid.To(part.Type),
			Name:       strings.TrimRight(string(utf16.Decode(part.Name[:])), "\x00"),
			FirstBlock: part.FirstBlock,
			LastBlock:  part.LastBlock,
			Attributes: Attribute(part.Attributes),
		})
	}
	// Remove long list of nils at the end as it's inconvenient to work with
	// (append doesn't work, debug prints are very long) and it round-trips
	// correctly even without it as it gets zero-padded when writing anyway.
	var maxValidPartition int
	for i, p := range g.Partitions {
		if !p.IsUnused() {
			maxValidPartition = i
		}
	}
	g.Partitions = g.Partitions[:maxValidPartition+1]
	g.b = r
	return &g, nil
}
