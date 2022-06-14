// The pstore package provides functions for interfacing with the Linux kernel's
// pstore (persistent storage) system.
// Documentation for pstore itself can be found at
// https://docs.kernel.org/admin-guide/abi-testing.html#abi-sys-fs-pstore.
package pstore

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// CanonicalMountPath contains the canonical mount path of the pstore filesystem
const CanonicalMountPath = "/sys/fs/pstore"

// pstoreDmesgHeader contains parsed header data from a pstore header.
type pstoreDmesgHeader struct {
	Reason  string
	Counter uint64
	Part    uint64
}

var headerRegexp = regexp.MustCompile("^([^#]+)#([0-9]+) Part([0-9]+)$")

// parseDmesgHeader parses textual pstore entry headers as assembled by
// @linux//fs/pstore/platform.c:pstore_dump back into a structured format.
// The input must be the first line of a file with the terminating \n stripped.
func parseDmesgHeader(hdr string) (*pstoreDmesgHeader, error) {
	parts := headerRegexp.FindStringSubmatch(hdr)
	if parts == nil {
		return nil, errors.New("unable to parse pstore entry header")
	}
	counter, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pstore header count: %w", err)
	}
	part, err := strconv.ParseUint(parts[3], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pstore header part: %w", err)
	}
	return &pstoreDmesgHeader{
		Reason:  parts[1],
		Counter: counter,
		Part:    part,
	}, nil
}

// A reassembled kernel message buffer dump from pstore.
type KmsgDump struct {
	// The reason why the dump was created. Common values include "Panic" and
	// "Oops", but depending on the setting `printk.always_kmsg_dump` and
	// potential future reasons this is likely unbounded.
	Reason string
	// The CLOCK_REALTIME value of the first entry in the dump (which is the
	// closest to the actual time the dump happened). This can be zero or
	// garbage if the RTC hasn't been initialized or the system has no working
	// clock source.
	OccurredAt time.Time
	// A counter counting up for every dump created. Can be used to order dumps
	// when the OccurredAt value is not usable due to system issues.
	Counter uint64
	// A list of kernel log lines in oldest-to-newest order, i.e. the oldest
	// message comes first. The actual cause is generally reported last.
	Lines []string
}

var dmesgFileRegexp = regexp.MustCompile("^dmesg-.*-([0-9]+)")

type pstoreDmesgFile struct {
	hdr   pstoreDmesgHeader
	ctime time.Time
	lines []string
}

// GetKmsgDumps returns a list of events where the kernel has dumped its kmsg
// (kernel log) buffer into pstore because of a kernel oops or panic.
func GetKmsgDumps() ([]KmsgDump, error) {
	return getKmsgDumpsFromFS(os.DirFS(CanonicalMountPath))
}

// f is injected here for testing
func getKmsgDumpsFromFS(f fs.FS) ([]KmsgDump, error) {
	var events []KmsgDump
	eventMap := make(map[string][]pstoreDmesgFile)
	pstoreEntries, err := fs.ReadDir(f, ".")
	if err != nil {
		return events, fmt.Errorf("failed to list files in pstore: %w", err)
	}
	for _, entry := range pstoreEntries {
		if !dmesgFileRegexp.MatchString(entry.Name()) {
			continue
		}
		f, err := f.Open(entry.Name())
		if err != nil {
			return events, fmt.Errorf("failed to open pstore entry file: %w", err)
		}
		// This only closes after all files have been read, but the number of
		// files is heavily bound by very small amounts of pstore space.
		defer f.Close()
		finfo, err := f.Stat()
		if err != nil {
			return events, fmt.Errorf("failed to stat pstore entry file: %w", err)
		}
		s := bufio.NewScanner(f)
		if !s.Scan() {
			return events, fmt.Errorf("cannot read first line header of pstore entry %q: %w", entry.Name(), s.Err())
		}
		hdr, err := parseDmesgHeader(s.Text())
		if err != nil {
			return events, fmt.Errorf("failed to parse header of file %q: %w", entry.Name(), err)
		}
		var lines []string
		for s.Scan() {
			lines = append(lines, s.Text())
		}
		// Same textual encoding is used in the header itself, so this
		// is as unique as it gets.
		key := fmt.Sprintf("%v#%d", hdr.Reason, hdr.Counter)
		eventMap[key] = append(eventMap[key], pstoreDmesgFile{hdr: *hdr, ctime: finfo.ModTime(), lines: lines})
	}

	for _, event := range eventMap {
		sort.Slice(event, func(i, j int) bool {
			return event[i].hdr.Part > event[j].hdr.Part
		})
		ev := KmsgDump{
			Counter: event[len(event)-1].hdr.Counter,
			Reason:  event[len(event)-1].hdr.Reason,
			// Entries get created in reverse order, so the most accurate
			// timestamp is the first one.
			OccurredAt: event[len(event)-1].ctime,
		}
		for _, entry := range event {
			ev.Lines = append(ev.Lines, entry.lines...)
		}
		events = append(events, ev)
	}
	sort.Slice(events, func(i, j int) bool {
		return !events[i].OccurredAt.Before(events[j].OccurredAt)
	})
	return events, nil
}

// ClearAll clears out all existing entries from the pstore. This should be done
// after every start (after the relevant data has been read out) to ensure that
// there is always space to store new pstore entries and to minimize the risk
// of breaking badly-programmed firmware.
func ClearAll() error {
	pstoreEntries, err := os.ReadDir(CanonicalMountPath)
	if err != nil {
		return fmt.Errorf("failed to list files in pstore: %w", err)
	}
	for _, entry := range pstoreEntries {
		if err := os.Remove(filepath.Join(CanonicalMountPath, entry.Name())); err != nil {
			return fmt.Errorf("failed to clear pstore entry: %w", err)
		}
	}
	return nil
}
