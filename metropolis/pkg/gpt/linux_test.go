package gpt

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"golang.org/x/sys/unix"
)

var testUUID = uuid.MustParse("85c0b60f-caf9-40dd-86fa-f8797e26238d")

func TestKernelInterop(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}
	ram0, err := os.OpenFile("/dev/ram0", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	g := Table{
		ID:         uuid.NewSHA1(testUUID, []byte("test")),
		BlockSize:  512,
		BlockCount: 2048,
		BootCode:   []byte("just some test code"),
		Partitions: []*Partition{
			nil,
			// This emoji is very complex and exercises UTF16 surrogate encoding
			// and composing.
			{Name: "Test 🏃‍♂️", FirstBlock: 10, LastBlock: 19, Type: PartitionTypeEFISystem, ID: uuid.NewSHA1(testUUID, []byte("test1")), Attributes: AttrNoBlockIOProto},
			nil,
			{Name: "Test2", FirstBlock: 20, LastBlock: 49, Type: PartitionTypeEFISystem, ID: uuid.NewSHA1(testUUID, []byte("test2")), Attributes: 0},
		},
	}
	if err := Write(ram0, &g); err != nil {
		t.Fatalf("Failed to write GPT: %v", err)
	}

	if err := unix.IoctlSetInt(int(ram0.Fd()), unix.BLKRRPART, 0); err != nil {
		t.Fatalf("Failed to reread partition table: %v", err)
	}
	if _, err := os.Stat("/sys/block/ram0/ram0p2"); err != nil {
		t.Errorf("Expected ram0p2 to exist, got %v", err)
	}
	if _, err := os.Stat("/sys/block/ram0/ram0p4"); err != nil {
		t.Errorf("Expected ram0p4 to exist, got %v", err)
	}
	if _, err := os.Stat("/sys/block/ram0/ram0p1"); err == nil {
		t.Error("Expected ram0p1 not to exist, but it exists")
	}
}
