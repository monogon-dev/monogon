// fwprune is a buildsystem utility that filters linux-firmware repository
// contents to include only files required by the built-in kernel modules,
// that are specified in modules.builtin.modinfo.
// (see: https://www.kernel.org/doc/Documentation/kbuild/kbuild.txt)
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// fwPaths returns a slice of filesystem paths relative to the root of the
// linux-firmware repository, pointing at firmware files, according to contents
// of the kernel build side effect: modules.builtin.modinfo.
func fwPaths(mi []byte) []string {
	// Use a map pset to deduplicate firmware paths.
	pset := make(map[string]bool)
	// Get a slice of entries of the form "unix.license=GPL" from mi. Then extract
	// firmware information from it.
	entries := bytes.Split(mi, []byte{0})
	for _, entry := range entries {
		// Skip empty entries.
		if len(entry) == 0 {
			continue
		}
		// Parse the entries. Split each entry into a key-value pair, separated
		// by "=".
		kv := strings.SplitN(string(entry), "=", 2)
		key, value := kv[0], kv[1]
		// Split the key into a module.attribute] pair, such as "unix.license".
		ma := strings.SplitN(key, ".", 2)
		// Skip, if it's not a firmware entry, according to the attribute.
		if ma[1] != "firmware" {
			continue
		}
		// If it is though, value holds a firmware path.
		pset[value] = true
	}
	// Convert the deduplicated pset to a slice.
	pslice := make([]string, 0, len(pset))
	for p, _ := range pset {
		pslice = append(pslice, p)
	}
	sort.Strings(pslice)
	return pslice
}

// fwDirs returns a slice of filesystem paths relative to the root of
// linux-firmware repository, pointing at directories that need to exist before
// files specified by fwp paths can be created.
func fwDirs(fwp []string) []string {
	// Use a map dset to deduplicate directory paths.
	dset := make(map[string]bool)
	for _, p := range fwp {
		dp := filepath.Dir(p)
		dset[dp] = true
	}
	// Convert dset to a slice.
	dslice := make([]string, 0, len(dset))
	for d, _ := range dset {
		dslice = append(dslice, d)
	}
	sort.Strings(dslice)
	return dslice
}

// copyFile copies a file at filesystem path src to dst. dst must not point to
// an existing file. It may return an IO error.
func copyFile(dst, src string) error {
	i, err := os.Open(src)
	if err != nil {
		return err
	}
	defer i.Close()

	o, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0770)
	if err != nil {
		return err
	}
	defer o.Close()

	if _, err := io.Copy(o, i); err != nil {
		return err
	}
	return nil
}

func main() {
	// The directory at fwdst will be filled with firmware required by the kernel
	// builtins specified in modules.builtin.modinfo [1]. fwsrc must point to the
	// linux-firmware repository [2]. All parameters must be filesystem paths. The
	// necessary parts of the original directory layout will be recreated at fwdst.
	// fwprune will output a list of directories and files it creates.
	// [1] https://www.kernel.org/doc/Documentation/kbuild/kbuild.txt
	// [2] https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git
	if len(os.Args) != 4 {
		// Print usage information, if misused.
		fmt.Println("Usage: fwprune modules.builtin.modinfo fwsrc fwdst")
		os.Exit(1)
	}
	modinfo := os.Args[1]
	fwsrc := os.Args[2]
	fwdst := os.Args[3]

	// Get the firmware file paths.
	mi, err := os.ReadFile(modinfo)
	if err != nil {
		log.Fatalf("While reading modinfo: %v", err)
	}
	fwp := fwPaths(mi)

	// Recreate the necessary parts of the linux-firmware directory tree.
	fwd := fwDirs(fwp)
	for _, rd := range fwd {
		d := filepath.Join(fwdst, rd)
		if err := os.MkdirAll(d, 0770); err != nil {
			log.Fatalf("Couldn't create a subdirectory: %v", err)
		}
		fmt.Println(d)
	}

	// Copy the files specified by fwp.
	for _, p := range fwp {
		dst := filepath.Join(fwdst, p)
		src := filepath.Join(fwsrc, p)

		if err := copyFile(dst, src); err != nil {
			log.Fatalf("Couldn't provide %q: %v", dst, err)
		}
		fmt.Println(p)
	}
}
