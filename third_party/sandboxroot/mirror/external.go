package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/cenkalti/backoff/v4"
	"k8s.io/klog/v2"
)

// rpmDef is a definition of an RPM dependency, containing the internal
// bazeldnf/bazel name of the dependency, an expected SHA256 sum of the RPM file,
// and a list of URLs of where that file should be downloaded from. This
// structure is parsed from repositories.bzl.
type rpmDef struct {
	name   string
	sha256 string
	mpath  string
	urls   []*url.URL
}

// newRPMDef builds and validates an rpmDef based on raw data from
// repositories.bzl.
func newRPMDef(name string, sha256 string, urls []string) (*rpmDef, error) {
	if len(urls) < 1 {
		return nil, fmt.Errorf("needs at least one URL")
	}
	var urlsP []*url.URL

	// Look through all URLs and make sure they're valid Fedora mirror paths, and
	// that all the mirror paths are the same.
	path := ""
	for _, us := range urls {
		u, err := url.Parse(us)
		if err != nil {
			return nil, fmt.Errorf("url invalid %w", err)
		}

		mpath, err := getFedoraMirrorPath(u)
		if err != nil {
			return nil, fmt.Errorf("unexpected url %s: %w", us, err)
		}

		// If this isn't the first mirror path we've seen, make sure they're the same.
		if path == "" {
			path = mpath
		} else {
			if path != mpath {
				return nil, fmt.Errorf("url path difference, %s vs %s", path, mpath)
			}
		}
		urlsP = append(urlsP, u)
	}
	return &rpmDef{
		name:   name,
		sha256: sha256,
		urls:   urlsP,
		mpath:  path,
	}, nil
}

// getFedoraMirrorPath takes a full URL to a mirrored RPM and returns its
// mirror-root-relative path, ie. the path which starts with fedora/linux/....
func getFedoraMirrorPath(u *url.URL) (string, error) {
	parts := strings.Split(u.Path, "/")

	// Find fedora/linux/...
	found := false
	for i, p := range parts {
		if p == "fedora" && (i+1) < len(parts) && parts[i+1] == "linux" {
			parts = parts[i:]
			found = true
			break
		}
	}
	if !found || len(parts) < 7 {
		return "", fmt.Errorf("does not look like a fedora mirror URL")
	}
	// Make sure the rest of the path makes some vague sense.
	switch parts[2] {
	case "releases", "updates":
	default:
		return "", fmt.Errorf("unexpected category %q", parts[2])
	}
	switch parts[4] {
	case "Everything":
	default:
		return "", fmt.Errorf("unexpected category %q", parts[3])
	}
	switch parts[5] {
	case "x86_64":
	default:
		return "", fmt.Errorf("unexpected architecture %q", parts[5])
	}

	// Return the path rebuilt and starting at fedora/linux/...
	return strings.Join(parts, "/"), nil
}

// validateOurs checks if our mirror has a copy of this RPM. If deep is true, the
// file will be downloaded and its SHA256 verified. Otherwise, a simple HEAD
// request is used.
func (e *rpmDef) validateOurs(ctx context.Context, deep bool) (bool, error) {
	ctxT, ctxC := context.WithTimeout(ctx, 2*time.Second)
	defer ctxC()

	url := ourMirrorURL(e.mpath)

	bo := backoff.NewExponentialBackOff()
	var found bool
	err := backoff.Retry(func() error {
		method := "HEAD"
		if deep {
			method = "GET"
		}
		req, err := http.NewRequestWithContext(ctxT, method, url, nil)
		if err != nil {
			return backoff.Permanent(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode == 200 {
			found = true
		} else {
			found = false
		}

		if !deep || !found {
			return nil
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		h := sha256.New()
		h.Write(data)
		got := hex.EncodeToString(h.Sum(nil))
		want := strings.ToLower(e.sha256)
		if want != got {
			log.Printf("SHA256 mismatch: wanted %s, got %s", want, got)
			found = false
		}
		return nil

	}, backoff.WithContext(bo, ctxT))
	if err != nil {
		return false, err
	}
	return found, nil
}

// mirrorToOurs attempts to download this RPM from a mirror that's not ours and
// upload it to our mirror via the given bucket.
func (e *rpmDef) mirrorToOurs(ctx context.Context, bucket *storage.BucketHandle) error {
	log.Printf("Mirroring %s ...", e.name)
	for _, source := range e.urls {
		// Skip our own mirror as a source.
		if strings.HasPrefix(source.String(), ourMirrorURL()) {
			continue
		}

		log.Printf("  Getting %s ...", source)
		data, err := e.get(ctx, source.String())
		if err != nil {
			klog.Errorf("  Failed: %v", err)
			continue
		}

		objName := filepath.Join(flagMirrorBucketSubdir, e.mpath)
		obj := bucket.Object(objName)
		log.Printf("  Uploading to %s...", objName)
		wr := obj.NewWriter(ctx)
		if _, err := wr.Write(data); err != nil {
			return fmt.Errorf("Write failed: %w", err)
		}
		if err := wr.Close(); err != nil {
			return fmt.Errorf("Close failed: %w", err)
		}
		return nil
	}
	return fmt.Errorf("all mirrors failed")
}

// get downloads the given RPM from the given URL and checks its SHA256.
func (e *rpmDef) get(ctx context.Context, url string) ([]byte, error) {
	ctxT, ctxC := context.WithTimeout(ctx, 60*time.Second)
	defer ctxC()

	bo := backoff.NewExponentialBackOff()
	var data []byte
	err := backoff.Retry(func() error {
		req, err := http.NewRequestWithContext(ctxT, "GET", url, nil)
		if err != nil {
			return backoff.Permanent(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return nil
	}, backoff.WithContext(bo, ctxT))
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write(data)
	got := hex.EncodeToString(h.Sum(nil))
	want := strings.ToLower(e.sha256)
	if want != got {
		return nil, fmt.Errorf("sha256 mismatch: wanted %s, got %s", want, got)
	}
	return data, nil
}
