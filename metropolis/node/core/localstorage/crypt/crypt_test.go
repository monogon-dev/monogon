package crypt

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

// TestMapUnmap performs a round-trip test for all modes, making sure we can
// intialize, map, unmap, map again and unmap again and that data isn't getting
// corrupted.
func TestMapUnmap(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}

	init := func(name string, key []byte, mode Mode) string {
		t.Helper()

		target, err := Init(name, "/dev/ram0", key, mode)
		if err != nil {
			t.Fatalf("Init failed: %v", err)
		}
		return target
	}

	unmap := func(name string, mode Mode) {
		t.Helper()
		if err := Unmap(name, mode); err != nil {
			t.Fatalf("Unmap failed: %v", err)
		}

	}

	map_ := func(name string, key []byte, mode Mode) string {
		t.Helper()
		target, err := Map(name, "/dev/ram0", key, mode)
		if err != nil {
			t.Fatalf("Map fialed: %v", err)
		}
		return target
	}

	writeWitness := func(target string, i int) string {
		t.Helper()

		file, err := os.OpenFile(target, os.O_WRONLY, 0644)
		if err != nil {
			t.Fatalf("opening initialized crypt failed: %v", err)
		}
		defer file.Close()

		witness := fmt.Sprintf("this is test %d", i)
		_, err = fmt.Fprintf(file, "%s", witness)
		if err != nil {
			t.Fatalf("writing to initialized crypt failed; %v", err)
		}
		return witness
	}

	checkWitness := func(target, witness string) {
		t.Helper()

		file, err := os.OpenFile(target, 0, 0644)
		if err != nil {
			t.Fatalf("opening mapped crypt failed: %v", err)
		}
		defer file.Close()

		buf := make([]byte, len(witness))
		_, err = file.Read(buf)
		if err != nil {
			t.Fatalf("reading mapped crypt failed: %v", err)
		}
		defer file.Close()

		if want, got := witness, string(buf); want != got {
			t.Fatalf("read data differs, wanted %q, got %q", want, got)
		}
		file.Close()
	}

	for i, mode := range []Mode{
		ModeInsecure,
		ModeEncrypted,
		ModeAuthenticated,
		ModeEncryptedAuthenticated,
	} {
		t.Run(string(mode), func(t *testing.T) {
			name := fmt.Sprintf("test-%d", i)
			key := bytes.Repeat([]byte("a"), 32)
			if mode == ModeInsecure {
				key = nil
			}

			target := init(name, key, mode)
			witness := writeWitness(target, i)
			unmap(name, mode)

			if target != "/dev/ram0" {
				if _, err := os.Stat(target); !os.IsNotExist(err) {
					t.Fatalf("Unmount didn't remove %s", target)
				}
			}

			target2 := map_(name, key, mode)
			if target != target2 {
				t.Fatalf("Init mounted at %s, first Map mounted at %s", target, target2)
			}

			checkWitness(target, witness)
			unmap(name, mode)

			target3 := map_(name, key, mode)
			if target != target3 {
				t.Fatalf("Init mounted at %s, second Map mounted at %s", target, target2)
			}
			checkWitness(target, witness)
			unmap(name, mode)
		})
	}
}
