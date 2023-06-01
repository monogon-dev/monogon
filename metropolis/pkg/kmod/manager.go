package kmod

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"

	kmodpb "source.monogon.dev/metropolis/pkg/kmod/spec"
)

// Manager contains all the logic and metadata required to efficiently load
// Linux kernel modules. It has internal loading tracking, thus it's recommended
// to only keep one Manager instance running per kernel. It can recreate its
// internal state quite well, but there are edge cases where the kernel makes
// it hard to do so (MODULE_STATE_UNFORMED) thus if possible that single
// instance should be kept alive. It currently does not support unloading
// modules, but that can be added to the existing design if deemed necessary.
type Manager struct {
	// Directory the modules are loaded from. The path in each module Meta
	// message is relative to this.
	modulesPath string
	meta        *kmodpb.Meta
	// Extra map to quickly find module indexes from names
	moduleIndexes map[string]uint32

	// mutex protects loadedModules, rest is read-only
	// This cannot use a RWMutex as locks cannot be upgraded
	mutex         sync.Mutex
	loadedModules map[uint32]bool
}

// NewManager instantiates a kernel module loading manager. Please take a look
// at the additional considerations on the Manager type itself.
func NewManager(meta *kmodpb.Meta, modulesPath string) (*Manager, error) {
	modIndexes := make(map[string]uint32)
	for i, m := range meta.Modules {
		modIndexes[m.Name] = uint32(i)
	}
	modulesFile, err := os.Open("/proc/modules")
	if err != nil {
		return nil, err
	}
	loadedModules := make(map[uint32]bool)
	s := bufio.NewScanner(modulesFile)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) == 0 {
			// Skip invalid lines
			continue
		}
		modIdx, ok := modIndexes[fields[0]]
		if !ok {
			// Certain modules are only available as built-in and are thus not
			// part of the module metadata. They do not need to be handled by
			// this code, ignore them.
			continue
		}
		loadedModules[modIdx] = true
	}
	return &Manager{
		modulesPath:   modulesPath,
		meta:          meta,
		moduleIndexes: modIndexes,
		loadedModules: loadedModules,
	}, nil
}

// NewManagerFromPath instantiates a new kernel module loading manager from a
// path containing a meta.pb file containing a kmod.Meta message as well as the
// kernel modules within. Please take a look at the additional considerations on
// the Manager type itself.
func NewManagerFromPath(modulesPath string) (*Manager, error) {
	moduleMetaRaw, err := os.ReadFile(filepath.Join(modulesPath, "meta.pb"))
	if err != nil {
		return nil, fmt.Errorf("error reading module metadata: %w", err)
	}
	var moduleMeta kmodpb.Meta
	if err := proto.Unmarshal(moduleMetaRaw, &moduleMeta); err != nil {
		return nil, fmt.Errorf("error decoding module metadata: %w", err)
	}
	return NewManager(&moduleMeta, modulesPath)
}

// ErrNotFound is returned when an attempt is made to load a module which does
// not exist according to the loaded metadata.
type ErrNotFound struct {
	Name string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("module %q does not exist", e.Name)
}

// LoadModule loads the module with the given name. If the module is already
// loaded or  built-in, it returns no error. If it failed loading the module or
// the module does not exist, it returns an error.
func (s *Manager) LoadModule(name string) error {
	modIdx, ok := s.moduleIndexes[name]
	if !ok {
		return &ErrNotFound{Name: name}
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.loadModuleRecursive(modIdx)
}

// LoadModulesForDevice loads all modules whose device match expressions
// (modaliases) match the given device modalias. It only returns an error if
// a module which matched the device or one of its dependencies caused an error
// when loading. A device modalias string which matches nothing is not an error.
func (s *Manager) LoadModulesForDevice(devModalias string) error {
	matches := make(map[uint32]bool)
	lookupModulesRec(s.meta.ModuleDeviceMatches, devModalias, matches)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for m := range matches {
		if err := s.loadModuleRecursive(m); err != nil {
			return err
		}
	}
	return nil
}

// Caller is REQUIRED to hold s.mutex!
func (s *Manager) loadModuleRecursive(modIdx uint32) error {
	if s.loadedModules[modIdx] {
		return nil
	}
	modMeta := s.meta.Modules[modIdx]
	if modMeta.Path == "" {
		// Module is built-in, dependency always satisfied
		return nil
	}
	for _, dep := range modMeta.Depends {
		if err := s.loadModuleRecursive(dep); err != nil {
			// Pass though as is, recursion can otherwise cause
			// extremely large errors
			return err
		}
	}
	module, err := os.Open(filepath.Join(s.modulesPath, modMeta.Path))
	if err != nil {
		return fmt.Errorf("error opening kernel module: %w", err)
	}
	defer module.Close()
	err = LoadModule(module, "", 0)
	if err != nil && err != unix.EEXIST {
		return fmt.Errorf("error loading kernel module %v: %w", modMeta.Name, err)
	}
	s.loadedModules[modIdx] = true
	return nil
}
