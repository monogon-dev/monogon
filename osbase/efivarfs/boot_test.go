package efivarfs

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

// Generated with old working marshaler and manually double-checked
var ref, _ = hex.DecodeString(
	"010000004a004500780061006d0070006c006500000004012a00010000000" +
		"500000000000000080000000000000014b8a76bad9dd11180b400c04fd430" +
		"c8020204041c005c0074006500730074005c0061002e00650066006900000" +
		"07fff0400",
)

func TestEncoding(t *testing.T) {
	opt := LoadOption{
		Description: "Example",
		FilePath: DevicePath{
			&HardDrivePath{
				PartitionNumber:     1,
				PartitionStartBlock: 5,
				PartitionSizeBlocks: 8,
				PartitionMatch: PartitionGPT{
					PartitionUUID: uuid.NameSpaceX500,
				},
			},
			FilePath("/test/a.efi"),
		},
	}
	got, err := opt.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(ref, got) {
		t.Fatalf("expected %x, got %x", ref, got)
	}
	got2, err := UnmarshalLoadOption(got)
	if err != nil {
		t.Fatalf("failed to unmarshal marshaled LoadOption: %v", err)
	}
	diff := cmp.Diff(&opt, got2)
	if diff != "" {
		t.Errorf("marshal/unmarshal wasn't transparent: %v", diff)
	}
}

func FuzzDecode(f *testing.F) {
	f.Add(ref)
	f.Fuzz(func(t *testing.T, a []byte) {
		// Just try to see if it crashes
		_, _ = UnmarshalLoadOption(a)
	})
}
