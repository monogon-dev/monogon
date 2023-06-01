// fwprune is a buildsystem utility that filters linux-firmware repository
// contents to include only files required by the built-in kernel modules,
// that are specified in modules.builtin.modinfo.
// (see: https://www.kernel.org/doc/Documentation/kbuild/kbuild.txt)
package main

import (
	"debug/elf"
	"flag"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/build/fsspec"
	"source.monogon.dev/metropolis/pkg/kmod"
)

// linkRegexp parses the Link: lines in the WHENCE file. This does not have
// an official grammar, the regexp has been written in an approximation of
// the original parsing algorithm at @linux-firmware//:copy_firmware.sh.
var linkRegexp = regexp.MustCompile(`(?m:^Link:\s*([^\s]+)\s+->\s+([^\s+]+)\s*$)`)

var (
	modinfoPath      = flag.String("modinfo", "", "Path to the modules.builtin.modinfo file built with the kernel")
	modulesPath      = flag.String("modules", "", "Path to the directory containing the dynamically loaded kernel modules (.ko files)")
	firmwareListPath = flag.String("firmware-file-list", "", "Path to a file containing a newline-separated list of paths to firmware files")
	whenceFilePath   = flag.String("firmware-whence", "", "Path to the linux-firmware WHENCE file containing aliases for firmware files")
	outMetaPath      = flag.String("out-meta", "", "Path where the resulting module metadata protobuf file should be created")
	outFSSpecPath    = flag.String("out-fsspec", "", "Path where the resulting fsspec should be created")
)

func main() {
	flag.Parse()
	if *modinfoPath == "" || *modulesPath == "" || *firmwareListPath == "" ||
		*whenceFilePath == "" || *outMetaPath == "" || *outFSSpecPath == "" {
		log.Fatal("all flags are required and need to be provided")
	}

	allFirmwareData, err := os.ReadFile(*firmwareListPath)
	if err != nil {
		log.Fatalf("Failed to read firmware source list: %v", err)
	}
	allFirmwarePaths := strings.Split(string(allFirmwareData), "\n")

	// Create a look-up table of all possible suffixes to their full paths as
	// this is much faster at O(n) than calling strings.HasSuffix for every
	// possible combination which is O(n^2).
	// For example a build output at out/a/b/c.bin will be entered into
	// the suffix LUT as build as out/a/b/c.bin, a/b/c.bin, b/c.bin and c.bin.
	// If the firmware then requests b/c.bin, the output path is contained in
	// the suffix LUT.
	suffixLUT := make(map[string]string)
	for _, firmwarePath := range allFirmwarePaths {
		pathParts := strings.Split(firmwarePath, string(os.PathSeparator))
		for i := range pathParts {
			suffixLUT[path.Join(pathParts[i:]...)] = firmwarePath
		}
	}

	// The linux-firmware repo contains a WHENCE file which contains (among
	// other information) aliases for firmware which should be symlinked.
	// Open this file and create a map of aliases in it.
	linkMap := make(map[string]string)
	metadata, err := os.ReadFile(*whenceFilePath)
	if err != nil {
		log.Fatalf("Failed to read metadata file: %v", err)
	}
	linksRaw := linkRegexp.FindAllStringSubmatch(string(metadata), -1)
	for _, link := range linksRaw {
		// For links we know the exact path referenced by kernel drives so
		// a suffix LUT is unnecessary.
		linkMap[link[1]] = link[2]
	}

	// Collect module metadata (modinfo) from both built-in modules via the
	// kbuild-generated metadata file as well as from the loadable modules by
	// walking them.
	var files []*fsspec.File
	var symlinks []*fsspec.SymbolicLink

	mi, err := os.Open(*modinfoPath)
	if err != nil {
		log.Fatalf("While reading modinfo: %v", err)
	}
	modMeta, err := kmod.GetBuiltinModulesInfo(mi)
	if err != nil {
		log.Fatalf("Failed to read modules modinfo data: %v", err)
	}

	err = filepath.WalkDir(*modulesPath, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() {
			return nil
		}
		mod, err := elf.Open(p)
		if err != nil {
			log.Fatal(err)
		}
		defer mod.Close()
		out, err := kmod.GetModuleInfo(mod)
		if err != nil {
			log.Fatal(err)
		}
		relPath, err := filepath.Rel(*modulesPath, p)
		if err != nil {
			return err
		}
		// Add path information for MakeMetaFromModuleInfo.
		out["path"] = []string{relPath}
		modMeta = append(modMeta, out)
		files = append(files, &fsspec.File{
			Path:       path.Join("/lib/modules", relPath),
			SourcePath: filepath.Join(*modulesPath, relPath),
			Mode:       0555,
		})
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking modules: %v", err)
	}

	// Generate loading metadata from all known modules.
	meta, err := kmod.MakeMetaFromModuleInfo(modMeta)
	if err != nil {
		log.Fatal(err)
	}
	metaRaw, err := proto.Marshal(meta)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(*outMetaPath, metaRaw, 0640); err != nil {
		log.Fatal(err)
	}
	files = append(files, &fsspec.File{
		Path:       "/lib/modules/meta.pb",
		SourcePath: *outMetaPath,
		Mode:       0444,
	})

	// Create set of all firmware paths required by modules
	fwset := make(map[string]bool)
	for _, m := range modMeta {
		if len(m["path"]) == 0 && len(m.Firmware()) > 0 {
			log.Fatalf("Module %v is built-in, but requires firmware. Linux does not support this in all configurations.", m.Name())
		}
		for _, fw := range m.Firmware() {
			fwset[fw] = true
		}
	}

	// Convert set to list and sort for determinism
	fwp := make([]string, 0, len(fwset))
	for p := range fwset {
		fwp = append(fwp, p)
	}
	sort.Strings(fwp)

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
	if err := os.WriteFile(*outFSSpecPath, fsspecRaw, 0644); err != nil {
		log.Fatalf("failed writing output: %v", err)
	}
}
