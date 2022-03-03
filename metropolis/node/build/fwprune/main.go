// fwprune is a buildsystem utility that filters linux-firmware repository
// contents to include only files required by the built-in kernel modules,
// that are specified in modules.builtin.modinfo.
// (see: https://www.kernel.org/doc/Documentation/kbuild/kbuild.txt)
package main

import (
	"bytes"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"

	"google.golang.org/protobuf/encoding/prototext"

	"source.monogon.dev/metropolis/node/build/fsspec"
)

var (
	// linkRegexp parses the Link: lines in the WHENCE file. This does not have
	// an official grammar, the regexp has been written in an approximation of
	// the original parsing algorithm at @linux-firmware//:copy_firmware.sh.
	linkRegexp = regexp.MustCompile(`(?m:^Link:\s*([^\s]+)\s+->\s+([^\s+]+)\s*$)`)
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

// fwprune takes a modinfo file from the kernel and extracts a list of all
// firmware files requested by all modules in that file. It then takes all
// available firmware file paths (newline-separated in the firmwareList file)
// and tries to match the requested file paths as a suffix of them.
// For example if a module requests firmware foo/bar.bin and in the firmware list
// there is a file at path build-out/x/y/foo/bar.bin it will use that file.
// It also parses links from the linux-firmware metadata file and uses them
// for matching requested firmware.
// Finally it generates an fsspec placing each file under its requested path
// under /lib/firmware.
func main() {
	if len(os.Args) != 5 {
		log.Fatal("Usage: fwprune modules.builtin.modinfo firmwareListPath metadataFilePath outSpec")
	}
	modinfo := os.Args[1]
	firmwareListPath := os.Args[2]
	metadataFilePath := os.Args[3]
	outSpec := os.Args[4]

	allFirmwareData, err := os.ReadFile(firmwareListPath)
	if err != nil {
		log.Fatalf("Failed to read firmware source list: %v", err)
	}
	allFirmwarePaths := strings.Split(string(allFirmwareData), "\n")

	// Create a look-up table of all possible suffixes to their full paths as
	// this is much faster at O(n) than calling strings.HasSuffix for every
	// possible combination which is O(n^2).
	suffixLUT := make(map[string]string)
	for _, firmwarePath := range allFirmwarePaths {
		pathParts := strings.Split(firmwarePath, string(os.PathSeparator))
		for i := range pathParts {
			suffixLUT[path.Join(pathParts[i:len(pathParts)]...)] = firmwarePath
		}
	}

	linkMap := make(map[string]string)
	metadata, err := os.ReadFile(metadataFilePath)
	if err != nil {
		log.Fatalf("Failed to read metadata file: %v", err)
	}
	linksRaw := linkRegexp.FindAllStringSubmatch(string(metadata), -1)
	for _, link := range linksRaw {
		// For links we know the exact path referenced by kernel drives so
		// a suffix LUT is unnecessary.
		linkMap[link[1]] = link[2]
	}

	// Get the firmware file paths used by modules according to modinfo data
	mi, err := os.ReadFile(modinfo)
	if err != nil {
		log.Fatalf("While reading modinfo: %v", err)
	}
	fwp := fwPaths(mi)

	var files []*fsspec.File
	var symlinks []*fsspec.SymbolicLink

	// This function is called for every requested firmware file and adds and
	// resolves symlinks until it finds the target file and adds that too.
	populatedPaths := make(map[string]bool)
	var chaseReference func(string)
	chaseReference = func(p string) {
		if populatedPaths[p] {
			// Bail if path is already populated. Because of the DAG-like
			// property of links in filesystems everything transitively pointed
			// to by anything at this path has already been included.
			return
		}
		placedPath := path.Join("/lib/firmware", p)
		if linkTarget := linkMap[p]; linkTarget != "" {
			symlinks = append(symlinks, &fsspec.SymbolicLink{
				Path:       placedPath,
				TargetPath: linkTarget,
			})
			populatedPaths[placedPath] = true
			// Symlinks are relative to their place, resolve them to be relative
			// to the firmware root directory.
			chaseReference(path.Join(path.Dir(p), linkTarget))
			return
		}
		sourcePath := suffixLUT[p]
		if sourcePath == "" {
			// This should not be fatal as sometimes linux-firmware cannot
			// ship all firmware usable by the kernel for mostly legal reasons.
			log.Printf("WARNING: Requested firmware %q not found", p)
			return
		}
		files = append(files, &fsspec.File{
			Path:       path.Join("/lib/firmware", p),
			Mode:       0444,
			SourcePath: sourcePath,
		})
		populatedPaths[path.Join("/lib/firmware", p)] = true
	}

	for _, p := range fwp {
		chaseReference(p)
	}
	// Format output in a both human- and machine-readable form
	marshalOpts := prototext.MarshalOptions{Multiline: true, Indent: "  "}
	fsspecRaw, err := marshalOpts.Marshal(&fsspec.FSSpec{File: files, SymbolicLink: symlinks})
	if err := os.WriteFile(outSpec, fsspecRaw, 0644); err != nil {
		log.Fatalf("failed writing output: %v", err)
	}
}
