// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package oci

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go"
	ocispecv1 "github.com/opencontainers/image-spec/specs-go/v1"

	"source.monogon.dev/osbase/structfs"
)

// ReadLayout reads an image from an OS path to an OCI layout directory.
func ReadLayout(path string) (*Image, error) {
	// Read the oci-layout marker file.
	layoutBytes, err := os.ReadFile(filepath.Join(path, "oci-layout"))
	if err != nil {
		return nil, err
	}
	layout := ocispecv1.ImageLayout{}
	err = json.Unmarshal(layoutBytes, &layout)
	if err != nil {
		return nil, fmt.Errorf("failed to parse oci-layout: %w", err)
	}
	if layout.Version != "1.0.0" {
		return nil, fmt.Errorf("unknown oci-layout version %q", layout.Version)
	}

	// Read the index.
	imageIndexBytes, err := os.ReadFile(filepath.Join(path, "index.json"))
	if err != nil {
		return nil, err
	}
	imageIndex := ocispecv1.Index{}
	err = json.Unmarshal(imageIndexBytes, &imageIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse index.json: %w", err)
	}
	if imageIndex.MediaType != ocispecv1.MediaTypeImageIndex {
		return nil, fmt.Errorf("unknown index.json mediaType %q", imageIndex.MediaType)
	}
	if len(imageIndex.Manifests) == 0 {
		return nil, fmt.Errorf("index.json contains no manifests")
	}
	if len(imageIndex.Manifests) != 1 {
		return nil, fmt.Errorf("index.json files containing multiple manifests are not supported")
	}
	manifestDescriptor := &imageIndex.Manifests[0]
	if manifestDescriptor.MediaType != ocispecv1.MediaTypeImageManifest {
		return nil, fmt.Errorf("unexpected manifest media type %q", manifestDescriptor.MediaType)
	}

	// Read the image manifest.
	imageManifestPath, err := layoutBlobPath(path, manifestDescriptor)
	if err != nil {
		return nil, err
	}
	imageManifestBytes, err := os.ReadFile(imageManifestPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read image manifest: %w", err)
	}

	blobs := &layoutBlobs{path: path}
	return NewImage(imageManifestBytes, string(manifestDescriptor.Digest), blobs)
}

type layoutBlobs struct {
	path string
}

func (r *layoutBlobs) Blob(descriptor *ocispecv1.Descriptor) (io.ReadCloser, error) {
	blobPath, err := layoutBlobPath(r.path, descriptor)
	if err != nil {
		return nil, err
	}
	return os.Open(blobPath)
}

func layoutBlobPath(layoutPath string, descriptor *ocispecv1.Descriptor) (string, error) {
	algorithm, encoded, err := ParseDigest(string(descriptor.Digest))
	if err != nil {
		return "", fmt.Errorf("failed to parse digest in image manifest: %w", err)
	}
	return filepath.Join(layoutPath, "blobs", algorithm, encoded), nil
}

// CreateLayout builds an OCI layout from an Image.
func CreateLayout(image *Image) (structfs.Tree, error) {
	// Build the index.
	artifactType := image.Manifest.Config.MediaType
	if artifactType == ocispecv1.MediaTypeImageConfig {
		artifactType = ""
	}
	imageIndex := ocispecv1.Index{
		Versioned: ocispec.Versioned{SchemaVersion: 2},
		MediaType: ocispecv1.MediaTypeImageIndex,
		Manifests: []ocispecv1.Descriptor{{
			MediaType:    ocispecv1.MediaTypeImageManifest,
			ArtifactType: artifactType,
			Digest:       digest.Digest(image.ManifestDigest),
			Size:         int64(len(image.RawManifest)),
		}},
	}
	imageIndexBytes, err := json.MarshalIndent(imageIndex, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal image index: %w", err)
	}
	imageIndexBytes = append(imageIndexBytes, '\n')

	root := structfs.Tree{
		structfs.File("oci-layout", structfs.Bytes(`{"imageLayoutVersion": "1.0.0"}`+"\n")),
		structfs.File("index.json", structfs.Bytes(imageIndexBytes)),
	}

	algorithm, encoded, err := ParseDigest(image.ManifestDigest)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest digest: %w", err)
	}
	imageManifestPath := path.Join("blobs", algorithm, encoded)
	err = root.PlaceFile(imageManifestPath, structfs.Bytes(image.RawManifest))
	if err != nil {
		return nil, err
	}

	hasBlob := map[string]bool{}
	for descriptor := range image.Descriptors() {
		algorithm, encoded, err := ParseDigest(string(descriptor.Digest))
		if err != nil {
			return nil, fmt.Errorf("failed to parse digest in image manifest: %w", err)
		}
		blobPath := path.Join("blobs", algorithm, encoded)
		if hasBlob[blobPath] {
			// If multiple blobs have the same hash, we only need the first one.
			continue
		}
		hasBlob[blobPath] = true
		err = root.PlaceFile(blobPath, image.StructfsBlob(descriptor))
		if err != nil {
			return nil, err
		}
	}

	return root, nil
}
