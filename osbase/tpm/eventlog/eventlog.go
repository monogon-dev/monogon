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

// Taken and pruned from go-attestation revision
// 2453c8f39a4ff46009f6a9db6fb7c6cca789d9a1 under Apache 2.0

package eventlog

import (
	"bytes"
	"crypto"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/google/go-tpm/tpm2"
)

// HashAlg identifies a hashing Algorithm.
type HashAlg uint8

// Valid hash algorithms.
var (
	HashSHA1   = HashAlg(tpm2.AlgSHA1)
	HashSHA256 = HashAlg(tpm2.AlgSHA256)
)

func (a HashAlg) cryptoHash() crypto.Hash {
	switch a {
	case HashSHA1:
		return crypto.SHA1
	case HashSHA256:
		return crypto.SHA256
	}
	return 0
}

func (a HashAlg) goTPMAlg() tpm2.Algorithm {
	switch a {
	case HashSHA1:
		return tpm2.AlgSHA1
	case HashSHA256:
		return tpm2.AlgSHA256
	}
	return 0
}

// String returns a human-friendly representation of the hash algorithm.
func (a HashAlg) String() string {
	switch a {
	case HashSHA1:
		return "SHA1"
	case HashSHA256:
		return "SHA256"
	}
	return fmt.Sprintf("HashAlg<%d>", int(a))
}

// ReplayError describes the parsed events that failed to verify against
// a particular PCR.
type ReplayError struct {
	Events      []Event
	invalidPCRs []int
}

func (e ReplayError) affected(pcr int) bool {
	for _, p := range e.invalidPCRs {
		if p == pcr {
			return true
		}
	}
	return false
}

// Error returns a human-friendly description of replay failures.
func (e ReplayError) Error() string {
	return fmt.Sprintf("event log failed to verify: the following registers failed to replay: %v", e.invalidPCRs)
}

// TPM algorithms. See the TPM 2.0 specification section 6.3.
//
//   https://trustedcomputinggroup.org/wp-content/uploads/TPM-Rev-2.0-Part-2-Structures-01.38.pdf#page=42
const (
	algSHA1   uint16 = 0x0004
	algSHA256 uint16 = 0x000B
)

// EventType indicates what kind of data an event is reporting.
type EventType uint32

// Event is a single event from a TCG event log. This reports descrete items such
// as BIOs measurements or EFI states.
type Event struct {
	// order of the event in the event log.
	sequence int

	// PCR index of the event.
	Index int
	// Type of the event.
	Type EventType

	// Data of the event. For certain kinds of events, this must match the event
	// digest to be valid.
	Data []byte
	// Digest is the verified digest of the event data. While an event can have
	// multiple for different hash values, this is the one that was matched to the
	// PCR value.
	Digest []byte

	// TODO(ericchiang): Provide examples or links for which event types must
	// match their data to their digest.
}

func (e *Event) digestEquals(b []byte) error {
	if len(e.Digest) == 0 {
		return errors.New("no digests present")
	}

	switch len(e.Digest) {
	case crypto.SHA256.Size():
		s := sha256.Sum256(b)
		if bytes.Equal(s[:], e.Digest) {
			return nil
		}
	case crypto.SHA1.Size():
		s := sha1.Sum(b)
		if bytes.Equal(s[:], e.Digest) {
			return nil
		}
	default:
		return fmt.Errorf("cannot compare hash of length %d", len(e.Digest))
	}

	return fmt.Errorf("digest (len %d) does not match", len(e.Digest))
}

// EventLog is a parsed measurement log. This contains unverified data representing
// boot events that must be replayed against PCR values to determine authenticity.
type EventLog struct {
	// Algs holds the set of algorithms that the event log uses.
	Algs []HashAlg

	rawEvents []rawEvent
}

func (e *EventLog) clone() *EventLog {
	out := EventLog{
		Algs:      make([]HashAlg, len(e.Algs)),
		rawEvents: make([]rawEvent, len(e.rawEvents)),
	}
	copy(out.Algs, e.Algs)
	copy(out.rawEvents, e.rawEvents)
	return &out
}

type elWorkaround struct {
	id          string
	affectedPCR int
	apply       func(e *EventLog) error
}

// inject3 appends two new events into the event log.
func inject3(e *EventLog, pcr int, data1, data2, data3 string) error {
	if err := inject(e, pcr, data1); err != nil {
		return err
	}
	if err := inject(e, pcr, data2); err != nil {
		return err
	}
	return inject(e, pcr, data3)
}

// inject2 appends two new events into the event log.
func inject2(e *EventLog, pcr int, data1, data2 string) error {
	if err := inject(e, pcr, data1); err != nil {
		return err
	}
	return inject(e, pcr, data2)
}

// inject appends a new event into the event log.
func inject(e *EventLog, pcr int, data string) error {
	evt := rawEvent{
		data:     []byte(data),
		index:    pcr,
		sequence: e.rawEvents[len(e.rawEvents)-1].sequence + 1,
	}
	for _, alg := range e.Algs {
		h := alg.cryptoHash().New()
		h.Write([]byte(data))
		evt.digests = append(evt.digests, digest{hash: alg.cryptoHash(), data: h.Sum(nil)})
	}
	e.rawEvents = append(e.rawEvents, evt)
	return nil
}

const (
	ebsInvocation = "Exit Boot Services Invocation"
	ebsSuccess    = "Exit Boot Services Returned with Success"
	ebsFailure    = "Exit Boot Services Returned with Failure"
)

var eventlogWorkarounds = []elWorkaround{
	{
		id:          "EBS Invocation + Success",
		affectedPCR: 5,
		apply: func(e *EventLog) error {
			return inject2(e, 5, ebsInvocation, ebsSuccess)
		},
	},
	{
		id:          "EBS Invocation + Failure",
		affectedPCR: 5,
		apply: func(e *EventLog) error {
			return inject2(e, 5, ebsInvocation, ebsFailure)
		},
	},
	{
		id:          "EBS Invocation + Failure + Success",
		affectedPCR: 5,
		apply: func(e *EventLog) error {
			return inject3(e, 5, ebsInvocation, ebsFailure, ebsSuccess)
		},
	},
}

// Verify replays the event log against a TPM's PCR values, returning the
// events which could be matched to a provided PCR value.
// An error is returned if the replayed digest for events with a given PCR
// index do not match any provided value for that PCR index.
func (e *EventLog) Verify(pcrs []PCR) ([]Event, error) {
	events, rErr := replayEvents(e.rawEvents, pcrs)
	if rErr == nil {
		return events, nil
	}
	// If there were any issues replaying the PCRs, try each of the workarounds
	// in turn.
	// TODO(jsonp): Allow workarounds to be combined.
	for _, wkrd := range eventlogWorkarounds {
		if !rErr.affected(wkrd.affectedPCR) {
			continue
		}
		el := e.clone()
		if err := wkrd.apply(el); err != nil {
			return nil, fmt.Errorf("failed applying workaround %q: %w", wkrd.id, err)
		}
		if events, err := replayEvents(el.rawEvents, pcrs); err == nil {
			return events, nil
		}
	}

	return events, rErr
}

// PCR encapsulates the value of a PCR at a point in time.
type PCR struct {
	Index     int
	Digest    []byte
	DigestAlg crypto.Hash
}

func extend(pcr PCR, replay []byte, e rawEvent) (pcrDigest []byte, eventDigest []byte, err error) {
	h := pcr.DigestAlg

	for _, digest := range e.digests {
		if digest.hash != pcr.DigestAlg {
			continue
		}
		if len(digest.data) != len(pcr.Digest) {
			return nil, nil, fmt.Errorf("digest data length (%d) doesn't match PCR digest length (%d)", len(digest.data), len(pcr.Digest))
		}
		hash := h.New()
		if len(replay) != 0 {
			hash.Write(replay)
		} else {
			b := make([]byte, h.Size())
			hash.Write(b)
		}
		hash.Write(digest.data)
		return hash.Sum(nil), digest.data, nil
	}
	return nil, nil, fmt.Errorf("no event digest matches pcr algorithm: %v", pcr.DigestAlg)
}

// replayPCR replays the event log for a specific PCR, using pcr and
// event digests with the algorithm in pcr. An error is returned if the
// replayed values do not match the final PCR digest, or any event tagged
// with that PCR does not posess an event digest with the specified algorithm.
func replayPCR(rawEvents []rawEvent, pcr PCR) ([]Event, bool) {
	var (
		replay    []byte
		outEvents []Event
	)

	for _, e := range rawEvents {
		if e.index != pcr.Index {
			continue
		}

		replayValue, digest, err := extend(pcr, replay, e)
		if err != nil {
			return nil, false
		}
		replay = replayValue
		outEvents = append(outEvents, Event{sequence: e.sequence, Data: e.data, Digest: digest, Index: pcr.Index, Type: e.typ})
	}

	if len(outEvents) > 0 && !bytes.Equal(replay, pcr.Digest) {
		return nil, false
	}
	return outEvents, true
}

type pcrReplayResult struct {
	events     []Event
	successful bool
}

func replayEvents(rawEvents []rawEvent, pcrs []PCR) ([]Event, *ReplayError) {
	var (
		invalidReplays []int
		verifiedEvents []Event
		allPCRReplays  = map[int][]pcrReplayResult{}
	)

	// Replay the event log for every PCR and digest algorithm combination.
	for _, pcr := range pcrs {
		events, ok := replayPCR(rawEvents, pcr)
		allPCRReplays[pcr.Index] = append(allPCRReplays[pcr.Index], pcrReplayResult{events, ok})
	}

	// Record PCR indices which do not have any successful replay. Record the
	// events for a successful replay.
pcrLoop:
	for i, replaysForPCR := range allPCRReplays {
		for _, replay := range replaysForPCR {
			if replay.successful {
				// We consider the PCR verified at this stage: The replay of values with
				// one digest algorithm matched a provided value.
				// As such, we save the PCR's events, and proceed to the next PCR.
				verifiedEvents = append(verifiedEvents, replay.events...)
				continue pcrLoop
			}
		}
		invalidReplays = append(invalidReplays, i)
	}

	if len(invalidReplays) > 0 {
		events := make([]Event, 0, len(rawEvents))
		for _, e := range rawEvents {
			events = append(events, Event{e.sequence, e.index, e.typ, e.data, nil})
		}
		return nil, &ReplayError{
			Events:      events,
			invalidPCRs: invalidReplays,
		}
	}

	sort.Slice(verifiedEvents, func(i int, j int) bool {
		return verifiedEvents[i].sequence < verifiedEvents[j].sequence
	})
	return verifiedEvents, nil
}

// EV_NO_ACTION is a special event type that indicates information to the
// parser instead of holding a measurement. For TPM 2.0, this event type is
// used to signal switching from SHA1 format to a variable length digest.
//
//   https://trustedcomputinggroup.org/wp-content/uploads/TCG_PCClientSpecPlat_TPM_2p0_1p04_pub.pdf#page=110
const eventTypeNoAction = 0x03

// ParseEventLog parses an unverified measurement log.
func ParseEventLog(measurementLog []byte) (*EventLog, error) {
	var specID *specIDEvent
	r := bytes.NewBuffer(measurementLog)
	parseFn := parseRawEvent
	var el EventLog
	e, err := parseFn(r, specID)
	if err != nil {
		return nil, fmt.Errorf("parse first event: %w", err)
	}
	if e.typ == eventTypeNoAction {
		specID, err = parseSpecIDEvent(e.data)
		if err != nil {
			return nil, fmt.Errorf("failed to parse spec ID event: %w", err)
		}
		for _, alg := range specID.algs {
			switch tpm2.Algorithm(alg.ID) {
			case tpm2.AlgSHA1:
				el.Algs = append(el.Algs, HashSHA1)
			case tpm2.AlgSHA256:
				el.Algs = append(el.Algs, HashSHA256)
			}
		}
		if len(el.Algs) == 0 {
			return nil, fmt.Errorf("measurement log didn't use sha1 or sha256 digests")
		}
		// Switch to parsing crypto agile events. Don't include this in the
		// replayed events since it intentionally doesn't extend the PCRs.
		//
		// Note that this doesn't actually guarentee that events have SHA256
		// digests.
		parseFn = parseRawEvent2
	} else {
		el.Algs = []HashAlg{HashSHA1}
		el.rawEvents = append(el.rawEvents, e)
	}
	sequence := 1
	for r.Len() != 0 {
		e, err := parseFn(r, specID)
		if err != nil {
			return nil, err
		}
		e.sequence = sequence
		sequence++
		el.rawEvents = append(el.rawEvents, e)
	}
	return &el, nil
}

type specIDEvent struct {
	algs []specAlgSize
}

type specAlgSize struct {
	ID   uint16
	Size uint16
}

// Expected values for various Spec ID Event fields.
//   https://trustedcomputinggroup.org/wp-content/uploads/EFI-Protocol-Specification-rev13-160330final.pdf#page=19
var wantSignature = [16]byte{0x53, 0x70,
	0x65, 0x63, 0x20, 0x49,
	0x44, 0x20, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x30,
	0x33, 0x00} // "Spec ID Event03\0"

const (
	wantMajor  = 2
	wantMinor  = 0
	wantErrata = 0
)

// parseSpecIDEvent parses a TCG_EfiSpecIDEventStruct structure from the reader.
//   https://trustedcomputinggroup.org/wp-content/uploads/EFI-Protocol-Specification-rev13-160330final.pdf#page=18
func parseSpecIDEvent(b []byte) (*specIDEvent, error) {
	r := bytes.NewReader(b)
	var header struct {
		Signature     [16]byte
		PlatformClass uint32
		VersionMinor  uint8
		VersionMajor  uint8
		Errata        uint8
		UintnSize     uint8
		NumAlgs       uint32
	}
	if err := binary.Read(r, binary.LittleEndian, &header); err != nil {
		return nil, fmt.Errorf("reading event header: %w", err)
	}
	if header.Signature != wantSignature {
		return nil, fmt.Errorf("invalid spec id signature: %x", header.Signature)
	}
	if header.VersionMajor != wantMajor {
		return nil, fmt.Errorf("invalid spec major version, got %02x, wanted %02x",
			header.VersionMajor, wantMajor)
	}
	if header.VersionMinor != wantMinor {
		return nil, fmt.Errorf("invalid spec minor version, got %02x, wanted %02x",
			header.VersionMajor, wantMinor)
	}

	// TODO(ericchiang): Check errata? Or do we expect that to change in ways
	// we're okay with?

	var specAlg specAlgSize
	var e specIDEvent
	for i := 0; i < int(header.NumAlgs); i++ {
		if err := binary.Read(r, binary.LittleEndian, &specAlg); err != nil {
			return nil, fmt.Errorf("reading algorithm: %w", err)
		}
		e.algs = append(e.algs, specAlg)
	}

	var vendorInfoSize uint8
	if err := binary.Read(r, binary.LittleEndian, &vendorInfoSize); err != nil {
		return nil, fmt.Errorf("reading vender info size: %w", err)
	}
	if r.Len() != int(vendorInfoSize) {
		return nil, fmt.Errorf("reading vendor info, expected %d remaining bytes, got %d", vendorInfoSize, r.Len())
	}
	return &e, nil
}

type digest struct {
	hash crypto.Hash
	data []byte
}

type rawEvent struct {
	sequence int
	index    int
	typ      EventType
	data     []byte
	digests  []digest
}

// TPM 1.2 event log format. See "5.1 SHA1 Event Log Entry Format"
//   https://trustedcomputinggroup.org/wp-content/uploads/EFI-Protocol-Specification-rev13-160330final.pdf#page=15
type rawEventHeader struct {
	PCRIndex  uint32
	Type      uint32
	Digest    [20]byte
	EventSize uint32
}

type eventSizeErr struct {
	eventSize uint32
	logSize   int
}

func (e *eventSizeErr) Error() string {
	return fmt.Sprintf("event data size (%d bytes) is greater than remaining measurement log (%d bytes)", e.eventSize, e.logSize)
}

func parseRawEvent(r *bytes.Buffer, specID *specIDEvent) (event rawEvent, err error) {
	var h rawEventHeader
	if err = binary.Read(r, binary.LittleEndian, &h); err != nil {
		return event, err
	}
	if h.EventSize == 0 {
		return event, errors.New("event data size is 0")
	}
	if h.EventSize > uint32(r.Len()) {
		return event, &eventSizeErr{h.EventSize, r.Len()}
	}

	data := make([]byte, int(h.EventSize))
	if _, err := io.ReadFull(r, data); err != nil {
		return event, err
	}

	digests := []digest{{hash: crypto.SHA1, data: h.Digest[:]}}

	return rawEvent{
		typ:     EventType(h.Type),
		data:    data,
		index:   int(h.PCRIndex),
		digests: digests,
	}, nil
}

// TPM 2.0 event log format. See "5.2 Crypto Agile Log Entry Format"
//   https://trustedcomputinggroup.org/wp-content/uploads/EFI-Protocol-Specification-rev13-160330final.pdf#page=15
type rawEvent2Header struct {
	PCRIndex uint32
	Type     uint32
}

func parseRawEvent2(r *bytes.Buffer, specID *specIDEvent) (event rawEvent, err error) {
	var h rawEvent2Header

	if err = binary.Read(r, binary.LittleEndian, &h); err != nil {
		return event, err
	}
	event.typ = EventType(h.Type)
	event.index = int(h.PCRIndex)

	// parse the event digests
	var numDigests uint32
	if err := binary.Read(r, binary.LittleEndian, &numDigests); err != nil {
		return event, err
	}

	for i := 0; i < int(numDigests); i++ {
		var algID uint16
		if err := binary.Read(r, binary.LittleEndian, &algID); err != nil {
			return event, err
		}
		var digest digest

		for _, alg := range specID.algs {
			if alg.ID != algID {
				continue
			}
			if uint16(r.Len()) < alg.Size {
				return event, fmt.Errorf("reading digest: %w", io.ErrUnexpectedEOF)
			}
			digest.data = make([]byte, alg.Size)
			digest.hash = HashAlg(alg.ID).cryptoHash()
		}
		if len(digest.data) == 0 {
			return event, fmt.Errorf("unknown algorithm ID %x", algID)
		}
		if _, err := io.ReadFull(r, digest.data); err != nil {
			return event, err
		}
		event.digests = append(event.digests, digest)
	}

	// parse event data
	var eventSize uint32
	if err = binary.Read(r, binary.LittleEndian, &eventSize); err != nil {
		return event, err
	}
	if eventSize == 0 {
		return event, errors.New("event data size is 0")
	}
	if eventSize > uint32(r.Len()) {
		return event, &eventSizeErr{eventSize, r.Len()}
	}
	event.data = make([]byte, int(eventSize))
	if _, err := io.ReadFull(r, event.data); err != nil {
		return event, err
	}
	return event, err
}
