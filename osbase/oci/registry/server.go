// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"internal/sync"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	ocispecv1 "github.com/opencontainers/image-spec/specs-go/v1"

	"source.monogon.dev/osbase/oci"
	"source.monogon.dev/osbase/structfs"
)

var (
	manifestsEp = regexp.MustCompile("^/v2/(" + repositoryExpr + ")/manifests/(" + tagExpr + "|" + digestExpr + ")$")
	blobsEp     = regexp.MustCompile("^/v2/(" + repositoryExpr + ")/blobs/(" + digestExpr + ")$")
)

// Server is an OCI registry server.
type Server struct {
	mu           sync.Mutex
	repositories map[string]*serverRepository
}

type serverRepository struct {
	tags      map[string]string
	manifests map[string]serverManifest
	blobs     map[string]structfs.Blob
}

type serverManifest struct {
	contentType string
	content     []byte
}

func NewServer() *Server {
	return &Server{
		repositories: make(map[string]*serverRepository),
	}
}

// AddImage adds an image to the server in the specified repository.
//
// If the tag is empty, the image can only be fetched by digest.
func (s *Server) AddImage(repository string, tag string, image *oci.Image) error {
	if !repositoryRegexp.MatchString(repository) {
		return fmt.Errorf("invalid repository %q", repository)
	}
	if tag != "" && !tagRegexp.MatchString(tag) {
		return fmt.Errorf("invalid tag %q", tag)
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	repo := s.repositories[repository]
	if repo == nil {
		repo = &serverRepository{
			tags:      make(map[string]string),
			manifests: make(map[string]serverManifest),
			blobs:     make(map[string]structfs.Blob),
		}
		s.repositories[repository] = repo
	}
	if _, ok := repo.manifests[image.ManifestDigest]; !ok {
		for descriptor := range image.Descriptors() {
			repo.blobs[string(descriptor.Digest)] = image.StructfsBlob(descriptor)
		}
		repo.manifests[image.ManifestDigest] = serverManifest{
			contentType: ocispecv1.MediaTypeImageManifest,
			content:     image.RawManifest,
		}
	}
	if tag != "" {
		repo.tags[tag] = image.ManifestDigest
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" && req.Method != "HEAD" {
		http.Error(w, "Registry is read-only, only GET and HEAD are allowed", http.StatusMethodNotAllowed)
		return
	}

	if req.URL.Path == "/v2/" {
		w.WriteHeader(http.StatusOK)
	} else if matches := manifestsEp.FindStringSubmatch(req.URL.Path); len(matches) > 0 {
		repository := matches[1]
		reference := matches[2]
		s.mu.Lock()
		repo := s.repositories[repository]
		if repo == nil {
			s.mu.Unlock()
			serveError(w, "NAME_UNKNOWN", fmt.Sprintf("Unknown repository: %s", repository), http.StatusNotFound)
			return
		}
		digest := reference
		if !strings.ContainsRune(reference, ':') {
			var ok bool
			digest, ok = repo.tags[reference]
			if !ok {
				s.mu.Unlock()
				serveError(w, "MANIFEST_UNKNOWN", fmt.Sprintf("Unknown tag: %s", reference), http.StatusNotFound)
				return
			}
		}
		manifest, ok := repo.manifests[digest]
		s.mu.Unlock()
		if !ok {
			serveError(w, "MANIFEST_UNKNOWN", fmt.Sprintf("Unknown manifest: %s", digest), http.StatusNotFound)
			return
		}

		w.Header().Set("Docker-Content-Digest", digest)
		w.Header().Set("Etag", fmt.Sprintf(`"%s"`, digest))
		w.Header().Set("Content-Type", manifest.contentType)
		w.Header().Set("X-Content-Type-Options", "nosniff")
		http.ServeContent(w, req, "", time.Time{}, bytes.NewReader(manifest.content))
	} else if matches := blobsEp.FindStringSubmatch(req.URL.Path); len(matches) > 0 {
		repository := matches[1]
		digest := matches[2]
		s.mu.Lock()
		repo := s.repositories[repository]
		if repo == nil {
			s.mu.Unlock()
			serveError(w, "NAME_UNKNOWN", fmt.Sprintf("Unknown repository: %s", repository), http.StatusNotFound)
			return
		}
		blob, ok := repo.blobs[digest]
		s.mu.Unlock()
		if !ok {
			serveError(w, "BLOB_UNKNOWN", fmt.Sprintf("Unknown blob: %s", digest), http.StatusNotFound)
			return
		}

		content, err := blob.Open()
		if err != nil {
			http.Error(w, "Failed to open blob", http.StatusInternalServerError)
			return
		}
		defer content.Close()
		w.Header().Set("Docker-Content-Digest", digest)
		w.Header().Set("Etag", fmt.Sprintf(`"%s"`, digest))
		w.Header().Set("Content-Type", "application/octet-stream")
		if contentSeeker, ok := content.(io.ReadSeeker); ok {
			http.ServeContent(w, req, "", time.Time{}, contentSeeker)
		} else {
			// Range requests are not supported.
			w.Header().Set("Content-Length", strconv.FormatInt(blob.Size(), 10))
			w.WriteHeader(http.StatusOK)
			if req.Method != "HEAD" {
				io.CopyN(w, content, blob.Size())
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func serveError(w http.ResponseWriter, code string, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	content, err := json.Marshal(&ErrorBody{Errors: []ErrorInfo{{
		Code:    code,
		Message: message,
	}}})
	if err == nil {
		w.Write(content)
	}
}
