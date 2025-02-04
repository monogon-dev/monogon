// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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
