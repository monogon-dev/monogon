// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package erofs

import (
	"reflect"
	"testing"
)

func TestEncodeSmallVLEBlock(t *testing.T) {
	type args struct {
		vals    [2]uint16
		blkaddr uint32
	}
	tests := []struct {
		name string
		args args
		want [8]byte
	}{
		{
			name: "Reference",
			args: args{vals: [2]uint16{vleClusterTypeHead | 1527, vleClusterTypeNonhead | 1}, blkaddr: 1},
			want: [8]byte{0xf7, 0x15, 0x01, 0x20, 0x01, 0x00, 0x00, 0x00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeSmallVLEBlock(tt.args.vals, tt.args.blkaddr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeSmallVLEBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeBigVLEBlock(t *testing.T) {
	type args struct {
		vals    [16]uint16
		blkaddr uint32
	}
	tests := []struct {
		name string
		args args
		want [32]byte
	}{
		{
			name: "Reference",
			args: args{
				vals: [16]uint16{
					vleClusterTypeNonhead | 2,
					vleClusterTypeHead | 1460,
					vleClusterTypeNonhead | 1,
					vleClusterTypeNonhead | 2,
					vleClusterTypeHead | 2751,
					vleClusterTypeNonhead | 1,
					vleClusterTypeNonhead | 2,
					vleClusterTypeHead | 940,
					vleClusterTypeNonhead | 1,
					vleClusterTypeHead | 3142,
					vleClusterTypeNonhead | 1,
					vleClusterTypeNonhead | 2,
					vleClusterTypeHead | 1750,
					vleClusterTypeNonhead | 1,
					vleClusterTypeNonhead | 2,
					vleClusterTypeHead | 683,
				},
				blkaddr: 3,
			},
			want: [32]byte{0x02, 0x20, 0x6d, 0x15, 0x00, 0x0a, 0x80, 0xbf, 0x5a, 0x00, 0x28, 0x00, 0xb2, 0x4e, 0x01, 0xa0, 0x11, 0x17, 0x00, 0x0a, 0x80, 0xd6, 0x56, 0x00, 0x28, 0x00, 0xae, 0x4a, 0x03, 0x00, 0x00, 0x00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBigVLEBlock(tt.args.vals, tt.args.blkaddr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeBigVLEBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
