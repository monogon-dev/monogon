// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package pstore

import (
	"fmt"
	"testing"
	"testing/fstest"
	"time"
)

func TestParseHeader(t *testing.T) {
	var cases = []struct {
		input       string
		expectedOut *pstoreDmesgHeader
	}{
		{"Panic#2 Part30", &pstoreDmesgHeader{"Panic", 2, 30}},
		{"Oops#1 Part5", &pstoreDmesgHeader{"Oops", 1, 5}},
		// Random kernel output that is similar, but definitely not a dump header
		{"<4>[2501503.489317] Oops: 0010 [#1] SMP NOPTI", nil},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test#%d", i+1), func(t *testing.T) {
			out, err := parseDmesgHeader(c.input)
			switch {
			case err != nil && c.expectedOut != nil:
				t.Errorf("Failed parsing %q: %v", c.input, err)
			case err == nil && c.expectedOut == nil:
				t.Errorf("Successfully parsed %q, expected error", c.input)
			case err != nil && c.expectedOut == nil:
			case err == nil && c.expectedOut != nil:
				if out.Part != c.expectedOut.Part {
					t.Errorf("Expected part to be %d, got %d", c.expectedOut.Part, out.Part)
				}
				if out.Counter != c.expectedOut.Counter {
					t.Errorf("Expected counter to be %d, got %d", c.expectedOut.Counter, out.Counter)
				}
				if out.Reason != c.expectedOut.Reason {
					t.Errorf("Expected reason to be %q, got %q", c.expectedOut.Reason, out.Reason)
				}
			}
		})
	}
}

func TestGetKmsgDumps(t *testing.T) {
	testTime1 := time.Date(2022, 06, 13, 1, 2, 3, 4, time.UTC)
	testTime2 := time.Date(2020, 06, 13, 1, 2, 3, 4, time.UTC)
	testTime3 := time.Date(2010, 06, 13, 1, 2, 3, 4, time.UTC)
	cases := []struct {
		name          string
		inputFS       fstest.MapFS
		expectErr     bool
		expectedDumps []KmsgDump
	}{
		{"EmptyPstore", map[string]*fstest.MapFile{}, false, []KmsgDump{}},
		{"SingleDumpSingleFile", map[string]*fstest.MapFile{
			"dmesg-efi-165467917816002": {ModTime: testTime1, Data: []byte("Panic#2 Part1\ntest1\ntest2")},
			"yolo-efi-165467917816002":  {ModTime: testTime1, Data: []byte("something totally unrelated")},
		}, false, []KmsgDump{{
			Reason:     "Panic",
			OccurredAt: testTime1,
			Counter:    2,
			Lines: []string{
				"test1",
				"test2",
			},
		}}},
		{"SingleDumpMultipleFiles", map[string]*fstest.MapFile{
			"dmesg-efi-165467917816002": {ModTime: testTime1, Data: []byte("Panic#2 Part1\ntest2\ntest3")},
			"dmesg-efi-165467917817002": {ModTime: testTime2, Data: []byte("Panic#2 Part2\ntest1")},
		}, false, []KmsgDump{{
			Reason:     "Panic",
			OccurredAt: testTime1,
			Counter:    2,
			Lines: []string{
				"test1",
				"test2",
				"test3",
			},
		}}},
		{"MultipleDumpsMultipleFiles", map[string]*fstest.MapFile{
			"dmesg-efi-165467917816002": {ModTime: testTime1, Data: []byte("Panic#2 Part1\ntest2\ntest3")},
			"dmesg-efi-165467917817002": {ModTime: testTime2, Data: []byte("Panic#2 Part2\ntest1")},
			"dmesg-efi-265467917816002": {ModTime: testTime3, Data: []byte("Oops#1 Part1\noops3")},
			"dmesg-efi-265467917817002": {ModTime: testTime2, Data: []byte("Oops#1 Part2\noops1\noops2")},
		}, false, []KmsgDump{{
			Reason:     "Panic",
			OccurredAt: testTime1,
			Counter:    2,
			Lines: []string{
				"test1",
				"test2",
				"test3",
			},
		}, {
			Reason:     "Oops",
			OccurredAt: testTime3,
			Counter:    1,
			Lines: []string{
				"oops1",
				"oops2",
				"oops3",
			},
		}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			dumps, err := getKmsgDumpsFromFS(c.inputFS)
			switch {
			case err == nil && c.expectErr:
				t.Fatal("Expected error, but got none")
				return
			case err != nil && !c.expectErr:
				t.Fatalf("Got unexpected error: %v", err)
				return
			case err != nil && c.expectErr:
				// Got expected error
				return
			case err == nil && !c.expectErr:
				if len(dumps) != len(c.expectedDumps) {
					t.Fatalf("Expected %d dumps, got %d", len(c.expectedDumps), len(dumps))
				}
				for i, dump := range dumps {
					if dump.OccurredAt != c.expectedDumps[i].OccurredAt {
						t.Errorf("Dump %d expected to have occurred at %v, got %v", i, c.expectedDumps[i].OccurredAt, dump.OccurredAt)
					}
					if dump.Reason != c.expectedDumps[i].Reason {
						t.Errorf("Expected reason in dump %d to be %v, got %v", i, c.expectedDumps[i].Reason, dump.Reason)
					}
					if dump.Counter != c.expectedDumps[i].Counter {
						t.Errorf("Expected counter in dump %d to be %d, got %d", i, c.expectedDumps[i].Counter, dump.Counter)
					}
					if len(dump.Lines) != len(c.expectedDumps[i].Lines) {
						t.Errorf("Expected number of lines in dump %d to be %d, got %d", i, len(c.expectedDumps[i].Lines), len(dump.Lines))
					}
					for j := range dump.Lines {
						if dump.Lines[j] != c.expectedDumps[i].Lines[j] {
							t.Errorf("Expected line %d in dump %d to be %q, got %q", i, j, c.expectedDumps[i].Lines[j], dump.Lines[j])
						}
					}
				}
			}
		})
	}
}
