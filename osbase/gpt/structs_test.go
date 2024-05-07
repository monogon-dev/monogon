package gpt

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestStructureSizes(t *testing.T) {
	cases := []struct {
		StructInstance interface{}
		ExpectedSize   int
	}{
		{mbr{}, 512},
		{mbrPartitionRecord{}, 16},
		{header{}, 92},
		{partition{}, 128},
	}
	for _, c := range cases {
		t.Run(reflect.TypeOf(c.StructInstance).String(), func(t *testing.T) {
			actualSize := binary.Size(c.StructInstance)
			if actualSize != c.ExpectedSize {
				t.Errorf("Expected %d bytes, got %d", c.ExpectedSize, actualSize)
			}
		})
	}
}
