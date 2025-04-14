// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package oci contains tools for handling OCI images.
package oci

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"strings"

	ocispecv1 "github.com/opencontainers/image-spec/specs-go/v1"

	"source.monogon.dev/osbase/structfs"
)

// Image represents an OCI image.
type Image struct {
	// Manifest contains the parsed image manifest.
	Manifest *ocispecv1.Manifest
	// RawManifest contains the bytes of the image manifest.
	RawManifest []byte
	// ManifestDigest contains the computed digest of RawManifest.
	ManifestDigest string

	blobs Blobs
}

// Blobs is the interface which image sources implement to retrieve the content
// of blobs.
type Blobs interface {
	// Blob returns the contents of a blob from its descriptor.
	// It does not verify the contents against the digest.
	Blob(*ocispecv1.Descriptor) (io.ReadCloser, error)
}

// NewImage verifies the manifest against the expected digest if not empty,
// then parses it and returns an [Image].
func NewImage(rawManifest []byte, expectedDigest string, blobs Blobs) (*Image, error) {
	digest := fmt.Sprintf("sha256:%x", sha256.Sum256(rawManifest))
	if expectedDigest != "" && expectedDigest != digest {
		return nil, fmt.Errorf("failed verification of manifest: expected digest %q, computed %q", expectedDigest, digest)
	}

	manifest := &ocispecv1.Manifest{}
	err := json.Unmarshal(rawManifest, &manifest)
	if err != nil {
		return nil, fmt.Errorf("failed to parse image manifest: %w", err)
	}
	if manifest.MediaType != ocispecv1.MediaTypeImageManifest {
		return nil, fmt.Errorf("unexpected manifest media type %q", manifest.MediaType)
	}
	image := &Image{
		Manifest:       manifest,
		RawManifest:    rawManifest,
		ManifestDigest: digest,
		blobs:          blobs,
	}
	for descriptor := range image.Descriptors() {
		if descriptor.Size < 0 {
			return nil, fmt.Errorf("invalid manifest: contains descriptor with negative size")
		}
	}

	return image, nil
}

// Descriptors returns an iterator over all descriptors in the image (config and
// layers).
func (i *Image) Descriptors() iter.Seq[*ocispecv1.Descriptor] {
	return func(yield func(*ocispecv1.Descriptor) bool) {
		if !yield(&i.Manifest.Config) {
			return
		}
		for l := range i.Manifest.Layers {
			if !yield(&i.Manifest.Layers[l]) {
				return
			}
		}
	}
}

// Blob returns the contents of a blob from its descriptor.
// It does not verify the contents against the digest.
func (i *Image) Blob(descriptor *ocispecv1.Descriptor) (io.ReadCloser, error) {
	if int64(len(descriptor.Data)) == descriptor.Size {
		return structfs.Bytes(descriptor.Data).Open()
	} else if len(descriptor.Data) != 0 {
		return nil, fmt.Errorf("descriptor has embedded data of wrong length")
	}
	return i.blobs.Blob(descriptor)
}

// ReadBlobVerified reads a blob into a byte slice and verifies it against the
// digest.
func (i *Image) ReadBlobVerified(descriptor *ocispecv1.Descriptor) ([]byte, error) {
	if descriptor.Size < 0 {
		return nil, fmt.Errorf("invalid descriptor size %d", descriptor.Size)
	}
	if descriptor.Size > 50*1024*1024 {
		return nil, fmt.Errorf("refusing to read blob of size %d into memory", descriptor.Size)
	}
	expectedDigest := string(descriptor.Digest)
	if _, _, err := ParseDigest(expectedDigest); err != nil {
		return nil, err
	}
	blob, err := i.Blob(descriptor)
	if err != nil {
		return nil, err
	}
	defer blob.Close()
	content := make([]byte, descriptor.Size)
	_, err = io.ReadFull(blob, content)
	if err != nil {
		return nil, err
	}
	digest := fmt.Sprintf("sha256:%x", sha256.Sum256(content))
	if expectedDigest != digest {
		return nil, fmt.Errorf("failed verification of blob: expected digest %q, computed %q", expectedDigest, digest)
	}
	return content, nil
}

// StructfsBlob wraps an image and descriptor into a [structfs.Blob].
func (i *Image) StructfsBlob(descriptor *ocispecv1.Descriptor) structfs.Blob {
	return &structfsBlob{
		image:      i,
		descriptor: descriptor,
	}
}

type structfsBlob struct {
	image      *Image
	descriptor *ocispecv1.Descriptor
}

func (b *structfsBlob) Open() (io.ReadCloser, error) {
	return b.image.Blob(b.descriptor)
}

func (b *structfsBlob) Size() int64 {
	return b.descriptor.Size
}

// ParseDigest splits a digest into its components. It returns an error if the
// algorithm is not supported, or if encoded is not valid for the algorithm.
func ParseDigest(digest string) (algorithm string, encoded string, err error) {
	algorithm, encoded, ok := strings.Cut(digest, ":")
	if !ok {
		return "", "", fmt.Errorf("invalid digest")
	}
	switch algorithm {
	case "sha256":
		rest := strings.TrimLeft(encoded, "0123456789abcdef")
		if len(rest) != 0 {
			return "", "", fmt.Errorf("invalid character in sha256 digest")
		}
		if len(encoded) != sha256.Size*2 {
			return "", "", fmt.Errorf("invalid sha256 digest length")
		}
	default:
		return "", "", fmt.Errorf("unknown digest algorithm %q", algorithm)
	}
	return
}
