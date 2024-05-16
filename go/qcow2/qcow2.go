// Package qcow2 implements generating QEMU copy-on-write image files.
package qcow2

// QCOW2 is a peculiar file format, as it's cluster-based (read: block-based) and
// that blocks are a concept both exposed to the guest (user data storage is
// expressed in clusters) and used to describe the layout of metadata structures.
//
// Most notably, the 'refcount' (reference count) mechanism for clusters is not
// only for user-visible clusters but also for clusters within the image that
// contain metadata. This means that users of the format can use the same
// allocation system for clusters both for allocating clusters exposed to the
// user and clusters for metadata/housekeeping.
//
// We only support a small subset of QCOW2 functionality: writing, qcow2 (no v1
// or v3) and no extensions. This keeps the codebase quite lean.
//
// QCOW2 images are made out of:
//
//  1. A file header (occupies first cluster);
//  2. Optionally, a backing file name (right after header, within the first
//     cluster);
//  3. A refcount table (made out of two levels, L1 is continuous and pointed to
//     by header, L2 is pointed to by L1, both must be cluster-aligned and occupy a
//     multiple of clusters each). The refcount table is indexed by cluster address
//     and returns whether a cluster is free or in use (by metadata or user data);
//  4. A cluster mapping table (made out of two levels similarly to the refcount
//     table, but the cluster mapping L1 table does not have to occupy an entire
//     cluster as the file header contains information on how many entries are
//     there). The cluster mapping table is indexed by user address and returns
//     the file cluster address.
//
// Our generated files contain a header/backing file name, a refcount table which
// contains enough entries to cover the metadata of the file itself, and a
// minimum L1 table (enough to cover the whole user data address range, but with
// no cluster mapped). The users of the image file will be able to use the
// cluster mapping L1 table as is, and will likely need to remap the refcount
// table to fit more file cluster addresses as the image gets written. This is
// the same kind of layout qemu-img uses.
//
// Reference: https://github.com/qemu/qemu/blob/master/docs/interop/qcow2.txt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// builder is the internal structure defining the requested parameters for image
// generation.
type builder struct {
	// size of the image in bytes. Note that this is the size visible to the user,
	// not the actual size in bytes of the container (which will be in most cases
	// significantly smaller).
	size uint64
	// backingFile is, if non-zero, the backing file for a given qcow2 image. This
	// makes the qcow2 image function as an overlay on top of the backing file (ie.
	// offers snapshot functionality).
	backingFile string
	// clusterBits is the number of bits that a cluster has. This is usually 16, for
	// 64KiB clusters.
	clusterBits uint32
}

type qcow2Header struct {
	Magic                 [4]byte
	Version               uint32
	BackingFileOffset     uint64
	BackingFileSize       uint32
	ClusterBits           uint32
	Size                  uint64
	CryptMethod           uint32
	L1Size                uint32
	L1TableOffset         uint64
	RefcountTableOffset   uint64
	RefcountTableClusters uint32
	NumberSnapshots       uint32
	SnapshotsOffset       uint64
}

// divRoundUp divides a by b and returns the result rounded up. This is useful to
// answer the question 'how many b-sized elements do I need to contain a'.
func divRoundUp(a, b uint64) uint64 {
	res := a / b
	if a%b != 0 {
		res += 1
	}
	return res
}

// clusterSize returns the cluster size for this builder.
func (b *builder) clusterSize() uint64 {
	return uint64(1 << b.clusterBits)
}

// calculateMappingEntries returns the details of the 'cluster mapping'
// structures (sometimes called 'L1'). This is calculated based on the cluster
// size and target image size.
//
// The values returned are the count of L1 entries needed to cover the entirety
// of the user data, and the size of such an entry within its table. The latter
// is constant.
func (b *builder) calculateMappingEntries() (mappingL1Count uint64, mappingL1EntrySize uint64) {
	clusterSize := b.clusterSize()

	mappingL2EntrySize := uint64(8)
	// Each L2 table is always one cluster.
	mappingL2EntryCount := clusterSize / mappingL2EntrySize
	// How much user data each L2 table covers.
	mappingL2Covering := mappingL2EntryCount * clusterSize
	// How many L1 entries do we need.
	mappingL1Count = divRoundUp(b.size, mappingL2Covering)
	// Each L1 entry is always 64 bits.
	mappingL1EntrySize = uint64(8)

	return
}

// refcountTableProperties contains the properties of a L1 or L2 refcount table
// calculated from the cluster size of a qcow2 image.
type refcountTableProperties struct {
	// bytesCoveredByCluster describes how many bytes of {meta,}data are addressed by
	// a cluster full of metadata entries of this refcount table.
	bytesCoveredByCluster uint64
	// entriesPerCluster describes how many entries of this refcount table fit in a
	// full cluster of metadata.
	entriesPerCluster uint64
}

// calculateRefcountProperties returns recountTableProperties for L1 and L2
// refcount tables for the qcow2 image being built. These are static properties
// calculated based on the cluster size, not on the target size of the image.
func (b *builder) calculateRefcountProperties() (l1 refcountTableProperties, l2 refcountTableProperties) {
	clusterSize := b.clusterSize()

	// Size of refcount l2 structure entry (16 bits):
	refcountL2EntrySize := uint64(2)
	// Number of entries in refcount l2 structure:
	refcountL2EntryCount := clusterSize / refcountL2EntrySize
	// Storage size covered by every refcount l2 structure (each entry points towards
	// a cluster):
	refcountL2Covering := refcountL2EntryCount * clusterSize

	// Size of refcount l1 structure entry (64 bits):
	refcountL1EntrySize := uint64(8)
	// Number of entries in refcount l1 structure:
	refcountL1EntryCount := clusterSize / refcountL1EntrySize
	// Storage size covered by every refcount l1 structure:
	refcountL1Covering := refcountL1EntryCount * refcountL2Covering

	l1.bytesCoveredByCluster = refcountL1Covering
	l1.entriesPerCluster = refcountL1EntryCount

	l2.bytesCoveredByCluster = refcountL2Covering
	l2.entriesPerCluster = refcountL2EntryCount
	return
}

// build writes out a qcow2 image (based on the builder's configuration) to a
// Writer.
func (b *builder) build(w io.Writer) error {
	clusterSize := b.clusterSize()
	// Minimum supported by QEMU and this codebase.
	if b.clusterSize() < 512 {
		return fmt.Errorf("cluster size too small")
	}

	// Size of a serialized qcow2Header.
	headerSize := 72
	// And however long the backing file name is (we stuff it after the header,
	// within the first cluster).
	headerSize += len(b.backingFile)
	// Make sure the above fits.
	if uint64(headerSize) > clusterSize {
		return fmt.Errorf("cluster size too small")
	}

	// Cluster mapping structures are just based on the backing storage size and
	// don't need to cover metadata. This makes it easy to calculate:
	mappingL1Count, mappingL1EntrySize := b.calculateMappingEntries()
	mappingL1Clusters := divRoundUp(mappingL1Count*mappingL1EntrySize, clusterSize)

	// For refcount structures we need enough to cover just the metadata of our
	// image. Unfortunately that includes the refcount structures themselves. We
	// choose a naïve iterative approach where we start out with zero metadata
	// structures and keep adding more until we cover everything that's necessary.

	// Coverage, in bytes, for an L1 and L2 refcount cluster full of entries.
	refcountL1, refcountL2 := b.calculateRefcountProperties()

	// How many clusters of metadata to put in the final image. We start at zero and
	// iterate upwards.
	refcountL1Clusters := uint64(0)
	refcountL2Clusters := uint64(0)

	// How many bytes the current metadata covers. This is also zero at first.
	maximumL2Covering := refcountL2Clusters * refcountL2.bytesCoveredByCluster
	maximumL1Covering := refcountL1Clusters * refcountL1.bytesCoveredByCluster

	// The size of our metadata in bytes.
	sizeMetadata := (1 + mappingL1Clusters + refcountL1Clusters + refcountL2Clusters) * clusterSize

	// Keep adding L1/L2 metadata clusters until we can cover all of our own
	// metadata. This is an iterative approach to the aforementioned problem. Even
	// with large images (100TB) at the default cluster size this loop only takes one
	// iteration to stabilize. Smaller cluster sizes might take longer.
	for {
		changed := false

		// Keep adding L2 metadata clusters until we cover all the metadata we know so far.
		for maximumL2Covering < sizeMetadata {
			refcountL2Clusters += 1
			// Recalculate.
			maximumL2Covering = refcountL2Clusters * refcountL2.bytesCoveredByCluster
			sizeMetadata += clusterSize
			changed = true
		}
		// Keep adding L1 metadata clusters until we cover all the metadata we know so far.
		for maximumL1Covering < sizeMetadata {
			refcountL1Clusters += 1
			// Recalculate.
			maximumL1Covering = refcountL1Clusters * refcountL1.bytesCoveredByCluster
			sizeMetadata += clusterSize
			changed = true
		}

		// If no changes were introduced, it means we stabilized at some l1/l2 metadata
		// cluster count. Exit.
		if !changed {
			break
		}
	}

	// Now that we have calculated everything, start writing out the header.
	h := qcow2Header{
		Version:     2,
		ClusterBits: b.clusterBits,
		Size:        b.size,
		// Unencrypted.
		CryptMethod: 0,

		L1Size: uint32(mappingL1Count),
		// L1 table starts after the header and refcount tables.
		L1TableOffset: (1 + refcountL1Clusters + refcountL2Clusters) * clusterSize,

		// Refcount table starts right after header cluster.
		RefcountTableOffset:   1 * clusterSize,
		RefcountTableClusters: uint32(refcountL1Clusters),

		// No snapshots.
		NumberSnapshots: 0,
		SnapshotsOffset: 0,
	}
	copy(h.Magic[:], "QFI\xfb")
	if b.backingFile != "" {
		// Backing file path is right after header.
		h.BackingFileOffset = 72
		h.BackingFileSize = uint32(len(b.backingFile))
	}
	if err := binary.Write(w, binary.BigEndian, &h); err != nil {
		return err
	}

	// Write out backing file name right after the header.
	if b.backingFile != "" {
		if _, err := w.Write([]byte(b.backingFile)); err != nil {
			return err
		}
	}

	// Align to next cluster with zeroes.
	if _, err := w.Write(bytes.Repeat([]byte{0}, int(clusterSize)-headerSize)); err != nil {
		return err
	}

	// Write L1 refcount table.
	for i := uint64(0); i < refcountL1Clusters; i++ {
		for j := uint64(0); j < refcountL1.entriesPerCluster; j++ {
			// Index of corresponding l2 table.
			ix := j + i*refcountL1.entriesPerCluster

			var data uint64
			if ix < refcountL2Clusters {
				// If this is an allocated L2 table, write its offset. Otherwise write zero.
				data = (1 + refcountL1Clusters + ix) * clusterSize
			}
			if err := binary.Write(w, binary.BigEndian, &data); err != nil {
				return err
			}
		}
	}
	// Write L2 refcount table.
	for i := uint64(0); i < refcountL2Clusters; i++ {
		for j := uint64(0); j < refcountL2.entriesPerCluster; j++ {
			// Index of corresponding cluster.
			ix := j + i*refcountL2.entriesPerCluster

			var data uint16
			if ix < (1 + refcountL1Clusters + refcountL2Clusters + mappingL1Clusters) {
				// If this is an in-use cluster, mark it as such.
				data = 1
			}
			if err := binary.Write(w, binary.BigEndian, &data); err != nil {
				return err
			}
		}
	}
	// Write L1 mapping table.
	for i := uint64(0); i < mappingL1Count; i++ {
		// No user data yet allocated.
		data := uint64(0)
		if err := binary.Write(w, binary.BigEndian, &data); err != nil {
			return err
		}
	}

	return nil
}

// GenerateOption structures are passed to the Generate call.
type GenerateOption struct {
	fileSize        *uint64
	backingFilePath *string
}

// GenerateWithFileSize will generate an image with a given user size in bytes.
func GenerateWithFileSize(fileSize uint64) *GenerateOption {
	return &GenerateOption{
		fileSize: &fileSize,
	}
}

// GenerateWithBackingFile will generate an image size backed by a given file
// path. If GenerateWithFileSize is not given, the path will also be checked at
// generation time and used to determine the size of the user image.
func GenerateWithBackingFile(path string) *GenerateOption {
	return &GenerateOption{
		backingFilePath: &path,
	}
}

// Generate builds a QCOW2 image and writes it to the given Writer.
//
// At least one of GenerateWithFileSize or GenerateWithBackingFile must be given.
func Generate(w io.Writer, opts ...*GenerateOption) error {
	var size uint64
	var haveSize bool

	var backingFile string
	var haveBackingFile bool

	for _, opt := range opts {
		if opt.fileSize != nil {
			if haveSize {
				return fmt.Errorf("cannot specify GenerateWithFileSize twice")
			}
			haveSize = true

			size = *opt.fileSize
		}
		if opt.backingFilePath != nil {
			if haveBackingFile {
				return fmt.Errorf("cannot specify GenerateWithBackingFile twice")
			}
			haveBackingFile = true

			backingFile = *opt.backingFilePath
		}
	}

	if !haveSize && !haveBackingFile {
		return fmt.Errorf("must be called with GenerateWithFileSize or GenerateWithBackingFileAndSize")
	}

	if haveBackingFile && !haveSize {
		st, err := os.Stat(backingFile)
		if err != nil {
			return fmt.Errorf("cannot read backing file: %w", err)
		}
		size = uint64(st.Size())
	}

	b := builder{
		size:        size,
		backingFile: backingFile,
		// 64k clusters (standard generated by qemu-img).
		clusterBits: uint32(0x10),
	}
	return b.build(w)
}
