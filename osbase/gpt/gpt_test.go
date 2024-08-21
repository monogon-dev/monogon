package gpt

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"

	"source.monogon.dev/osbase/blockdev"
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
			d := blockdev.MustNewMemory(512, 2048) // 1MiB
			g, err := New(d)
			if err != nil {
				panic(err)
			}
			g.Partitions = c.parts
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

func TestAddPartition(t *testing.T) {
	if os.Getenv("IN_KTEST") == "true" {
		t.Skip("In ktest")
	}
	cases := []struct {
		name        string
		parts       []*Partition
		addSize     int64
		addOptions  []AddOption
		expectErr   string
		expectParts []*Partition
	}{
		{
			name:        "empty-basic",
			addSize:     9 * 512,
			expectParts: []*Partition{{Name: "added", FirstBlock: 2048, LastBlock: 2048 + 9 - 1}},
		},
		{
			name:        "empty-fill",
			addSize:     -1,
			expectParts: []*Partition{{Name: "added", FirstBlock: 2048, LastBlock: 5*2048 - 16384/512 - 2}},
		},
		{
			name:        "empty-end",
			addSize:     9 * 512,
			addOptions:  []AddOption{WithPreferEnd()},
			expectParts: []*Partition{{Name: "added", FirstBlock: 4 * 2048, LastBlock: 4*2048 + 9 - 1}},
		},
		{
			name:       "empty-align0",
			addSize:    9 * 512,
			addOptions: []AddOption{WithAlignment(0)},
			expectErr:  "must be positive",
		},
		{
			name:        "empty-align512",
			addSize:     9 * 512,
			addOptions:  []AddOption{WithAlignment(512)},
			expectParts: []*Partition{{Name: "added", FirstBlock: 2 + 16384/512, LastBlock: 2 + 16384/512 + 9 - 1}},
		},
		{
			name:      "empty-zero-size",
			addSize:   0,
			expectErr: "must be positive",
		},
		{
			name:        "full-fill",
			parts:       []*Partition{{Name: "full", FirstBlock: 2048, LastBlock: 5*2048 - 16384/512 - 2}},
			addSize:     -1,
			expectErr:   "no free space",
			expectParts: []*Partition{{Name: "full", FirstBlock: 2048, LastBlock: 5*2048 - 16384/512 - 2}},
		},
		{
			name:    "haveone-basic",
			parts:   []*Partition{{Name: "one", FirstBlock: 2048, LastBlock: 2048 + 5}},
			addSize: 9 * 512,
			expectParts: []*Partition{
				{Name: "one", FirstBlock: 2048, LastBlock: 2048 + 5},
				{Name: "added", FirstBlock: 2 * 2048, LastBlock: 2*2048 + 9 - 1},
			},
		},
		{
			name:    "havemiddle-basic",
			parts:   []*Partition{{Name: "middle", FirstBlock: 2 * 2048, LastBlock: 2*2048 + 5}},
			addSize: 9 * 512,
			expectParts: []*Partition{
				{Name: "middle", FirstBlock: 2 * 2048, LastBlock: 2*2048 + 5},
				{Name: "added", FirstBlock: 2048, LastBlock: 2048 + 9 - 1},
			},
		},
		{
			name:       "havemiddle-end",
			parts:      []*Partition{{Name: "middle", FirstBlock: 2 * 2048, LastBlock: 2*2048 + 5}},
			addSize:    9 * 512,
			addOptions: []AddOption{WithPreferEnd()},
			expectParts: []*Partition{
				{Name: "middle", FirstBlock: 2 * 2048, LastBlock: 2*2048 + 5},
				{Name: "added", FirstBlock: 4 * 2048, LastBlock: 4*2048 + 9 - 1},
			},
		},
		{
			name:       "end-aligned",
			parts:      []*Partition{{Name: "endalign", FirstBlock: 4 * 2048, LastBlock: 4*2048 + 8}},
			addSize:    2048 * 512,
			addOptions: []AddOption{WithPreferEnd()},
			expectParts: []*Partition{
				{Name: "endalign", FirstBlock: 4 * 2048, LastBlock: 4*2048 + 8},
				{Name: "added", FirstBlock: 3 * 2048, LastBlock: 4*2048 - 1},
			},
		},
		{
			name: "emptyslots-basic",
			parts: []*Partition{
				{Name: "one", FirstBlock: 2048, LastBlock: 2048},
				nil, nil,
				{Name: "two", FirstBlock: 2048 + 1, LastBlock: 2048 + 1},
			},
			addSize: 9 * 512,
			expectParts: []*Partition{
				{Name: "one", FirstBlock: 2048, LastBlock: 2048},
				{Name: "added", FirstBlock: 2 * 2048, LastBlock: 2*2048 + 9 - 1},
				nil,
				{Name: "two", FirstBlock: 2048 + 1, LastBlock: 2048 + 1},
			},
		},
		{
			name: "emptyslots-keep",
			parts: []*Partition{
				{Name: "one", FirstBlock: 2048, LastBlock: 2048},
				nil, nil,
				{Name: "two", FirstBlock: 2048 + 1, LastBlock: 2048 + 1},
			},
			addSize:    9 * 512,
			addOptions: []AddOption{WithKeepEmptyEntries()},
			expectParts: []*Partition{
				{Name: "one", FirstBlock: 2048, LastBlock: 2048},
				nil, nil,
				{Name: "two", FirstBlock: 2048 + 1, LastBlock: 2048 + 1},
				{Name: "added", FirstBlock: 2 * 2048, LastBlock: 2*2048 + 9 - 1},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, part := range c.parts {
				if part != nil {
					part.Type = PartitionTypeEFISystem
				}
			}
			addPart := &Partition{Name: "added", Type: PartitionTypeEFISystem}
			d := blockdev.MustNewMemory(512, 5*2048) // 5MiB
			g, err := New(d)
			if err != nil {
				panic(err)
			}
			g.Partitions = c.parts
			err = g.AddPartition(addPart, c.addSize, c.addOptions...)
			if (err == nil) != (c.expectErr == "") || (err != nil && !strings.Contains(err.Error(), c.expectErr)) {
				t.Errorf("expected %q, got %v", c.expectErr, err)
			}
			_, overlap, err := g.GetFreeSpaces()
			if err != nil {
				t.Fatal(err)
			}
			if overlap {
				t.Errorf("partitions overlap")
			}
			if len(g.Partitions) != len(c.expectParts) {
				t.Fatalf("expected %+v, got %+v", c.expectParts, g.Partitions)
			}
			for i := range g.Partitions {
				gotPart, wantPart := g.Partitions[i], c.expectParts[i]
				if (gotPart == nil) != (wantPart == nil) {
					t.Errorf("partition %d: got %+v, expected %+v", i, gotPart, wantPart)
				}
				if gotPart == nil || wantPart == nil {
					continue
				}
				if gotPart.Name != wantPart.Name {
					t.Errorf("partition %d: got Name %q, expected %q", i, gotPart.Name, wantPart.Name)
				}
				if gotPart.FirstBlock != wantPart.FirstBlock {
					t.Errorf("partition %d: got FirstBlock %d, expected %d", i, gotPart.FirstBlock, wantPart.FirstBlock)
				}
				if gotPart.LastBlock != wantPart.LastBlock {
					t.Errorf("partition %d: got LastBlock %d, expected %d", i, gotPart.LastBlock, wantPart.LastBlock)
				}
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	if os.Getenv("IN_KTEST") == "true" {
		t.Skip("In ktest")
	}
	d := blockdev.MustNewMemory(512, 2048) // 1 MiB

	g := Table{
		ID:       uuid.NewSHA1(uuid.Nil, []byte("test")),
		BootCode: []byte("just some test code"),
		Partitions: []*Partition{
			nil,
			// This emoji is very complex and exercises UTF16 surrogate encoding
			// and composing.
			{Name: "Test 🏃‍♂️", FirstBlock: 10, LastBlock: 19, Type: PartitionTypeEFISystem, ID: uuid.NewSHA1(uuid.Nil, []byte("test1")), Attributes: AttrNoBlockIOProto},
			nil,
			{Name: "Test2", FirstBlock: 20, LastBlock: 49, Type: PartitionTypeEFISystem, ID: uuid.NewSHA1(uuid.Nil, []byte("test2")), Attributes: 0},
		},
		b: d,
	}
	if err := g.Write(); err != nil {
		t.Fatalf("Error while writing Table: %v", err)
	}

	originalHash := sha256.New()
	sr1 := io.NewSectionReader(d, 0, d.BlockSize()*d.BlockCount())
	if _, err := io.CopyBuffer(originalHash, sr1, make([]byte, d.OptimalBlockSize())); err != nil {
		panic(err)
	}

	g2, err := Read(d)
	if err != nil {
		t.Fatalf("Failed to read back GPT: %v", err)
	}
	if g2.ID != g.ID {
		t.Errorf("Disk UUID changed when reading back: %v", err)
	}
	// Destroy primary GPT
	d.Zero(1*d.BlockSize(), 5*d.BlockSize())
	g3, err := Read(d)
	if err != nil {
		t.Fatalf("Failed to read back GPT with primary GPT destroyed: %v", err)
	}
	if g3.ID != g.ID {
		t.Errorf("Disk UUID changed when reading back: %v", err)
	}
	if err := g3.Write(); err != nil {
		t.Fatalf("Error while writing back GPT: %v", err)
	}
	rewrittenHash := sha256.New()
	sr2 := io.NewSectionReader(d, 0, d.BlockSize()*d.BlockCount())
	if _, err := io.CopyBuffer(rewrittenHash, sr2, make([]byte, d.OptimalBlockSize())); err != nil {
		panic(err)
	}
	if !bytes.Equal(originalHash.Sum(nil), rewrittenHash.Sum(nil)) {
		t.Errorf("Write/Read/Write test was not reproducible: %x != %x", originalHash.Sum(nil), rewrittenHash.Sum(nil))
	}
}
