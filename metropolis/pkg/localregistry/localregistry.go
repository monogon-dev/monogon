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
	"regexp"
	"strconv"

	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/docker/distribution/reference"
	"github.com/opencontainers/go-digest"
	"google.golang.org/protobuf/encoding/prototext"

	"source.monogon.dev/metropolis/pkg/localregistry/spec"
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

func blobFromBazel(s *Server, bd *spec.BlobDescriptor, mediaType string) (distribution.Descriptor, error) {
	digestRaw, err := os.ReadFile(bd.DigestPath)
	if err != nil {
		return distribution.Descriptor{}, fmt.Errorf("while opening digest file: %w", err)
	}
	stat, err := os.Stat(bd.FilePath)
	if err != nil {
		return distribution.Descriptor{}, fmt.Errorf("while stat'ing blob file: %w", err)
	}
	digest := digest.Digest("sha256:" + string(digestRaw))
	s.blobs[digest] = blobMeta{filePath: bd.FilePath, mediaType: mediaType, contentLength: stat.Size()}
	return distribution.Descriptor{
		MediaType: mediaType,
		Size:      stat.Size(),
		Digest:    digest,
	}, nil
}

func FromBazelManifest(m []byte) (*Server, error) {
	var manifest spec.Manifest
	if err := prototext.Unmarshal(m, &manifest); err != nil {
		log.Fatalf("failed to parse manifest: %v", err)
	}
	s := Server{
		manifests: make(map[string][]byte),
		blobs:     make(map[digest.Digest]blobMeta),
	}
	for _, i := range manifest.Images {
		imageManifest := schema2.Manifest{
			Versioned: schema2.SchemaVersion,
		}
		var err error
		imageManifest.Config, err = blobFromBazel(&s, i.Config, schema2.MediaTypeImageConfig)
		if err != nil {
			return nil, fmt.Errorf("while creating blob spec for %q: %w", i.Name, err)
		}
		for _, l := range i.Layers {
			ml, err := blobFromBazel(&s, l, schema2.MediaTypeLayer)
			if err != nil {
				return nil, fmt.Errorf("while creating blob spec for %q: %w", i.Name, err)
			}
			imageManifest.Layers = append(imageManifest.Layers, ml)
		}
		s.manifests[i.Name], err = json.Marshal(imageManifest)
		if err != nil {
			return nil, fmt.Errorf("while marshaling image %q manifest: %w", i.Name, err)
		}
		// For Digest lookups
		s.manifests[string(digest.Canonical.FromBytes(s.manifests[i.Name]))] = s.manifests[i.Name]
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
