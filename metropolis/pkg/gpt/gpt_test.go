package gpt

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestFreeSpaces(t *testing.T) {
	cases := []struct {
		name            string
		parts           []*Partition
		expected        [][2]int64
		expectedOverlap bool
	}{
		{"Empty", []*Partition{}, [][2]int64{{34, 2015}}, false},
		{"OnePart", []*Partition{
			{Type: PartitionTypeEFISystem, FirstBlock: 200, LastBlock: 1499},
		}, [][2]int64{
			{34, 200},
			{1500, 2015},
		}, false},
		{"TwoOverlappingParts", []*Partition{
			{Type: PartitionTypeEFISystem, FirstBlock: 200, LastBlock: 1499},
			{Type: PartitionTypeEFISystem, FirstBlock: 1000, LastBlock: 1999},
		}, [][2]int64{
			{34, 200},
			{2000, 2015},
		}, true},
		{"Full", []*Partition{
			{Type: PartitionTypeEFISystem, FirstBlock: 34, LastBlock: 999},
			{Type: PartitionTypeEFISystem, FirstBlock: 1000, LastBlock: 2014},
		}, [][2]int64{}, false},
		{"TwoSpacedParts", []*Partition{
			{Type: PartitionTypeEFISystem, FirstBlock: 500, LastBlock: 899},
			{Type: PartitionTypeEFISystem, FirstBlock: 1200, LastBlock: 1799},
		}, [][2]int64{
			{34, 500},
			{900, 1200},
			{1800, 2015},
		}, false},
	}

	// Partitions are created manually as AddPartition calls FreeSpaces itself,
	// which makes the test unreliable as well as making failures very hard to
	// debug.
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			g := Table{
				BlockSize:  512,
				BlockCount: 2048, // 0.5KiB * 2048 = 1MiB
				Partitions: c.parts,
			}
			fs, overlap, err := g.GetFreeSpaces()
			if err != nil {
				t.Fatal(err)
			}
			if overlap != c.expectedOverlap {
				t.Errorf("expected overlap %v, got %v", c.expectedOverlap, overlap)
			}
			if len(fs) != len(c.expected) {
				t.Fatalf("expected %v, got %v", c.expected, fs)
			}
			for i := range fs {
				if fs[i] != c.expected[i] {
					t.Errorf("free space mismatch at pos %d: got [%d, %d), expected [%d, %d)", i, fs[i][0], fs[i][1], c.expected[i][0], c.expected[i][1])
				}
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	if os.Getenv("IN_KTEST") == "true" {
		t.Skip("In ktest")
	}
	g := Table{
		ID:         uuid.NewSHA1(zeroUUID, []byte("test")),
		BlockSize:  512,
		BlockCount: 2048,
		BootCode:   []byte("just some test code"),
		Partitions: []*Partition{
			nil,
			// This emoji is very complex and exercises UTF16 surrogate encoding
			// and composing.
			{Name: "Test 🏃‍♂️", FirstBlock: 10, LastBlock: 19, Type: PartitionTypeEFISystem, ID: uuid.NewSHA1(zeroUUID, []byte("test1")), Attributes: AttrNoBlockIOProto},
			nil,
			{Name: "Test2", FirstBlock: 20, LastBlock: 49, Type: PartitionTypeEFISystem, ID: uuid.NewSHA1(zeroUUID, []byte("test2")), Attributes: 0},
		},
	}
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(f.Name())
	if err := Write(f, &g); err != nil {
		t.Fatalf("Error while wrinting Table: %v", err)
	}
	f.Seek(0, io.SeekStart)
	originalHash := sha256.New()
	if _, err := io.Copy(originalHash, f); err != nil {
		panic(err)
	}

	g2, err := Read(f, 512, 2048)
	if err != nil {
		t.Fatalf("Failed to read back GPT: %v", err)
	}
	if g2.ID != g.ID {
		t.Errorf("Disk UUID changed when reading back: %v", err)
	}
	// Destroy primary GPT
	f.Seek(1*g.BlockSize, io.SeekStart)
	f.Write(make([]byte, 512))
	f.Seek(0, io.SeekStart)
	g3, err := Read(f, 512, 2048)
	if err != nil {
		t.Fatalf("Failed to read back GPT with primary GPT destroyed: %v", err)
	}
	if g3.ID != g.ID {
		t.Errorf("Disk UUID changed when reading back: %v", err)
	}
	f.Seek(0, io.SeekStart)
	if err := Write(f, g3); err != nil {
		t.Fatalf("Error while writing back GPT: %v", err)
	}
	f.Seek(0, io.SeekStart)
	rewrittenHash := sha256.New()
	if _, err := io.Copy(rewrittenHash, f); err != nil {
		panic(err)
	}
	if !bytes.Equal(originalHash.Sum([]byte{}), rewrittenHash.Sum([]byte{})) {
		t.Errorf("Write/Read/Write test was not reproducible: %x != %x", originalHash.Sum([]byte{}), rewrittenHash.Sum([]byte{}))
	}
}
