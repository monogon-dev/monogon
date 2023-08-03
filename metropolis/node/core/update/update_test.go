package update

import (
	"testing"

	"github.com/google/uuid"

	"source.monogon.dev/metropolis/pkg/efivarfs"
	"source.monogon.dev/metropolis/pkg/gpt"
)

func TestFindBootEntry(t *testing.T) {
	testUUID1 := uuid.MustParse("85cb7a0c-d31d-4b65-1111-919b069915f1")
	testUUID2 := uuid.MustParse("d3086aa2-0327-4634-2222-5c6c8488cae3")
	cases := []struct {
		name        string
		slot        Slot
		espid       uuid.UUID
		entries     map[int]*efivarfs.LoadOption
		expectedOk  bool
		expectedIdx int
	}{
		{
			name:       "NoEntries",
			slot:       SlotA,
			espid:      testUUID1,
			entries:    make(map[int]*efivarfs.LoadOption),
			expectedOk: false,
		},
		{
			name:  "FindSimple",
			slot:  SlotB,
			espid: testUUID1,
			entries: map[int]*efivarfs.LoadOption{
				5: &efivarfs.LoadOption{
					Description: "Other Entry",
					FilePath: efivarfs.DevicePath{
						&efivarfs.HardDrivePath{
							PartitionNumber: 1,
							PartitionMatch: efivarfs.PartitionMBR{
								DiskSignature: [4]byte{1, 2, 3, 4},
							},
						},
						efivarfs.FilePath("EFI/something/else.efi"),
					},
				},
				6: &efivarfs.LoadOption{
					Description: "Completely different entry",
					FilePath: efivarfs.DevicePath{
						&efivarfs.UnknownPath{
							// Vendor-specific subtype
							TypeVal:    1,
							SubTypeVal: 4,
							DataVal:    []byte{1, 2, 3, 4},
						},
						efivarfs.FilePath("EFI/something"),
						efivarfs.FilePath("else.efi"),
					},
				},
				16: &efivarfs.LoadOption{
					Description: "Target Entry",
					FilePath: efivarfs.DevicePath{
						&efivarfs.HardDrivePath{
							PartitionNumber: 2,
							PartitionMatch: efivarfs.PartitionGPT{
								PartitionUUID: testUUID1,
							},
						},
						efivarfs.FilePath("/EFI/metropolis/boot-b.efi"),
					},
				},
			},
			expectedOk:  true,
			expectedIdx: 16,
		},
		{
			name:  "FindViaESPUUID",
			slot:  SlotA,
			espid: testUUID1,
			entries: map[int]*efivarfs.LoadOption{
				6: &efivarfs.LoadOption{
					Description: "Other ESP UUID",
					FilePath: efivarfs.DevicePath{
						&efivarfs.HardDrivePath{
							PartitionNumber: 2,
							PartitionMatch: efivarfs.PartitionGPT{
								PartitionUUID: testUUID2,
							},
						},
						efivarfs.FilePath("/EFI/metropolis/boot-a.efi"),
					},
				},
				7: &efivarfs.LoadOption{
					Description: "Target Entry",
					FilePath: efivarfs.DevicePath{
						&efivarfs.HardDrivePath{
							PartitionNumber: 2,
							PartitionMatch: efivarfs.PartitionGPT{
								PartitionUUID: testUUID1,
							},
						},
						efivarfs.FilePath("/EFI/metropolis/boot-a.efi"),
					},
				},
			},
			expectedOk:  true,
			expectedIdx: 7,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := Service{
				ESPPart: &gpt.Partition{
					ID: c.espid,
				},
			}
			idx, ok := s.findBootEntry(c.entries, c.slot)
			if ok != c.expectedOk {
				t.Fatalf("expected ok %t, got %t", c.expectedOk, ok)
			}
			if idx != c.expectedIdx {
				t.Fatalf("expected idx %d, got %d", c.expectedIdx, idx)
			}
		})
	}
}
