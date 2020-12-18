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

package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_patchKconfig(t *testing.T) {
	type args struct {
		inFile    string
		overrides map[string]string
	}
	tests := []struct {
		name        string
		args        args
		wantOutFile string
		wantErr     bool
	}{
		{
			"passthroughExtend",
			args{inFile: "# TEST=y\n\n", overrides: map[string]string{"TEST": "n"}},
			"# TEST=y\n\nTEST=n\n",
			false,
		},
		{
			"patch",
			args{inFile: "TEST=y\nTEST_NO=n\n", overrides: map[string]string{"TEST": "n"}},
			"TEST=n\nTEST_NO=n\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outFile := &bytes.Buffer{}
			if err := patchKconfig(strings.NewReader(tt.args.inFile), outFile, tt.args.overrides); (err != nil) != tt.wantErr {
				t.Errorf("patchKconfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutFile := outFile.String(); gotOutFile != tt.wantOutFile {
				t.Errorf("patchKconfig() = %v, want %v", gotOutFile, tt.wantOutFile)
			}
		})
	}
}
