// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package localregistry implements a read-only OCI Distribution / Docker
// V2 container image registry backed by local layers.
package localregistry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest/manifestlist"
	"github.com/docker/distribution/manifest/ocischema"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/docker/distribution/reference"
	"github.com/opencontainers/go-digest"
	"google.golang.org/protobuf/encoding/prototext"

	"source.monogon.dev/metropolis/test/localregistry/spec"
)

type Server struct {
	manifests map[string][]byte
	blobs     map[digest.Digest]blobMeta
}

type blobMeta struct {
	filePath      string
	mediaType     string
	contentLength int64
}

func manifestDescriptorFromBazel(image *spec.Image) (manifestlist.ManifestDescriptor, error) {
	indexPath, err := runfiles.Rlocation(filepath.Join("_main", image.Path, "index.json"))
	if err != nil {
		return manifestlist.ManifestDescriptor{}, fmt.Errorf("while locating manifest list file: %w", err)
	}

	manifestListRaw, err := os.ReadFile(indexPath)
	if err != nil {
		return manifestlist.ManifestDescriptor{}, fmt.Errorf("while opening manifest list file: %w", err)
	}

	var imageManifestList manifestlist.ManifestList
	if err := json.Unmarshal(manifestListRaw, &imageManifestList); err != nil {
		return manifestlist.ManifestDescriptor{}, fmt.Errorf("while unmarshaling manifest list for %q: %w", image.Name, err)
	}

	if len(imageManifestList.Manifests) != 1 {
		return manifestlist.ManifestDescriptor{}, fmt.Errorf("unexpected manifest list length > 1")
	}

	return imageManifestList.Manifests[0], nil
}

func manifestFromBazel(s *Server, image *spec.Image, md manifestlist.ManifestDescriptor) (ocischema.Manifest, error) {
	manifestPath, err := runfiles.Rlocation(filepath.Join("_main", image.Path, "blobs", md.Digest.Algorithm().String(), md.Digest.Hex()))
	if err != nil {
		return ocischema.Manifest{}, fmt.Errorf("while locating manifest file: %w", err)
	}
	manifestRaw, err := os.ReadFile(manifestPath)
	if err != nil {
		return ocischema.Manifest{}, fmt.Errorf("while opening manifest file: %w", err)
	}

	var imageManifest ocischema.Manifest
	if err := json.Unmarshal(manifestRaw, &imageManifest); err != nil {
		return ocischema.Manifest{}, fmt.Errorf("while unmarshaling manifest for %q: %w", image.Name, err)
	}

	// For Digest lookups
	s.manifests[image.Name] = manifestRaw
	s.manifests[md.Digest.String()] = manifestRaw

	return imageManifest, nil
}

func addBazelBlobFromDescriptor(s *Server, image *spec.Image, dd distribution.Descriptor) error {
	path, err := runfiles.Rlocation(filepath.Join("_main", image.Path, "blobs", dd.Digest.Algorithm().String(), dd.Digest.Hex()))
	if err != nil {
		return fmt.Errorf("while locating blob: %w", err)
	}
	s.blobs[dd.Digest] = blobMeta{filePath: path, mediaType: dd.MediaType, contentLength: dd.Size}
	return nil
}

func FromBazelManifest(mb []byte) (*Server, error) {
	var bazelManifest spec.Manifest
	if err := prototext.Unmarshal(mb, &bazelManifest); err != nil {
		log.Fatalf("failed to parse manifest: %v", err)
	}
	s := Server{
		manifests: make(map[string][]byte),
		blobs:     make(map[digest.Digest]blobMeta),
	}
	for _, i := range bazelManifest.Images {
		md, err := manifestDescriptorFromBazel(i)
		if err != nil {
			return nil, err
		}

		if err := addBazelBlobFromDescriptor(&s, i, md.Descriptor); err != nil {
			return nil, err
		}

		m, err := manifestFromBazel(&s, i, md)
		if err != nil {
			return nil, err
		}

		if err := addBazelBlobFromDescriptor(&s, i, m.Config); err != nil {
			return nil, err
		}
		for _, l := range m.Layers {
			if err := addBazelBlobFromDescriptor(&s, i, l); err != nil {
				return nil, err
			}
		}
	}
	return &s, nil
}

var (
	versionCheckEp = regexp.MustCompile(`^/v2/$`)
	manifestEp     = regexp.MustCompile("^/v2/(" + reference.NameRegexp.String() + ")/manifests/(" + reference.TagRegexp.String() + "|" + digest.DigestRegexp.String() + ")$")
	blobEp         = regexp.MustCompile("^/v2/(" + reference.NameRegexp.String() + ")/blobs/(" + digest.DigestRegexp.String() + ")$")
)

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet && req.Method != http.MethodHead {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Registry is read-only, only GET and HEAD are allowed\n")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if versionCheckEp.MatchString(req.URL.Path) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{}")
		return
	} else if matches := manifestEp.FindStringSubmatch(req.URL.Path); len(matches) > 0 {
		name := matches[1]
		manifest, ok := s.manifests[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Image not found")
			return
		}
		w.Header().Set("Content-Type", schema2.MediaTypeManifest)
		w.Header().Set("Content-Length", strconv.FormatInt(int64(len(manifest)), 10))
		w.WriteHeader(http.StatusOK)
		io.Copy(w, bytes.NewReader(manifest))
	} else if matches := blobEp.FindStringSubmatch(req.URL.Path); len(matches) > 0 {
		bm, ok := s.blobs[digest.Digest(matches[2])]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Blob not found")
			return
		}
		w.Header().Set("Content-Type", bm.mediaType)
		w.Header().Set("Content-Length", strconv.FormatInt(bm.contentLength, 10))
		http.ServeFile(w, req, bm.filePath)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
