// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package osimage

// Media types which appear in the OCI image manifest.
const (
	ArtifactTypeOSImage          = "application/vnd.monogon.os.image.v1"
	MediaTypeOSImageConfig       = "application/vnd.monogon.os.image.config.v1+json"
	MediaTypePayloadUncompressed = "application/octet-stream"
	MediaTypePayloadZstd         = "application/zstd"
)

// ConfigVersion is the current version in [Config.FormatVersion].
const ConfigVersion = "1"

// Config contains metadata of an OS image.
type Config struct {
	// FormatVersion should be incremented when making breaking changes to the
	// image format. Readers must stop when they see an unknown version.
	FormatVersion string `json:"format_version"`
	// Payloads describes the payloads contained in the image. It has the same
	// length and order as the layers list in the image manifest.
	Payloads []PayloadInfo `json:"payloads"`
}

type PayloadInfo struct {
	// Name of this payload, for example "system" or "kernel.efi". Must consist of
	// at least one and at most 80 characters in the set [0-9A-Za-z._-], and may
	// not start with [._] or end with [.]. Each name must be unique.
	Name string `json:"name"`
	// Size is the uncompressed size in bytes of the payload.
	Size int64 `json:"size"`
	// HashChunkSize is the size of each hash chunk, except for the last chunk
	// which may be smaller.
	HashChunkSize int64 `json:"hash_chunk_size"`
	// ChunkHashesSHA256 contains the sha256 hash of each chunk. Chunks are
	// obtained by dividing the uncompressed payload into chunks of size
	// HashChunkSize, with the last chunk containing the remainder.
	// Hashes are encoded as base64 without padding.
	ChunkHashesSHA256 []string `json:"chunk_hashes_sha256"`
}
