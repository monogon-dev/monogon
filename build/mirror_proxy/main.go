// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	flagUser             string
	flagPass             string
	flagMirrorBucketName string
	flagCredentialsFile  string
)

func main() {
	flag.StringVar(&flagUser, "username", "", "Username required to enable s3 upload")
	flag.StringVar(&flagPass, "password", "", "Password required to enable s3 upload")
	flag.StringVar(&flagCredentialsFile, "credentials_file", "", "Credentials file to use for GCS")
	flag.StringVar(&flagMirrorBucketName, "bucket_name", "monogon-bazel-mirror", "Name of GCS bucket to mirror into.")
	flag.Parse()

	if flagUser == "" || flagPass == "" {
		log.Fatalf("Missing username or password flag")
	}

	if flagCredentialsFile == "" {
		log.Fatalf("Missing credentials flag")
	}

	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(flagCredentialsFile))
	if err != nil {
		log.Fatalf("Could not build google cloud storage client: %v", err)
	}

	bucketClient := client.Bucket(flagMirrorBucketName)
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		mirrorHandler(bucketClient, w, r)
	}

	log.Panic(http.ListenAndServe(":80", http.HandlerFunc(handlerFunc)))
}

func mirrorHandler(m *storage.BucketHandle, w http.ResponseWriter, r *http.Request) {
	targetPath := strings.TrimPrefix(r.URL.Path, "/")
	targetURL := "https://" + targetPath
	if len(r.URL.Query()) != 0 {
		targetURL += "?" + r.URL.Query().Encode()
	}

	if r.Method != http.MethodGet {
		log.Printf("%s: invalid method %q: %v", r.RemoteAddr, targetURL, r.Method)
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	if len(r.URL.Query()) != 0 {
		log.Printf("%s: invalid query url: %q", r.RemoteAddr, targetURL)
		http.Error(w, "URLs with query parameters are not supported", http.StatusNotAcceptable)
		return
	}

	obj := m.Object(targetPath)
	objR, err := obj.NewReader(r.Context())
	if err != nil && !errors.Is(err, storage.ErrObjectNotExist) {
		log.Printf("%s: fetching %q from bucket: %v", r.RemoteAddr, obj.ObjectName(), err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// If not found and not authenticated, return 404
	if errors.Is(err, storage.ErrObjectNotExist) && !isAuthenticated(r) {
		http.Error(w, "object not found in mirror", http.StatusNotFound)
		return
	}

	// If found, return mirror content
	if err == nil {
		log.Printf("%s: serving cached object %q", r.RemoteAddr, targetURL)

		w.Header().Set("Content-Type", objR.Attrs.ContentType)
		w.Header().Set("Content-Length", fmt.Sprintf("%d", objR.Attrs.Size))
		w.WriteHeader(http.StatusOK)

		_, _ = io.Copy(w, objR)
		return
	}

	// If I am not reading the logic wrong, this should not happen, but
	// better to be sure.
	if !isAuthenticated(r) {
		http.Error(w, "upstream fetch requires authentication", http.StatusUnauthorized)
		return
	}

	// If not found, try download.
	outReq, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		log.Printf("%s: forwarding to %q failed: %v", r.RemoteAddr, targetURL, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	copyHeader(outReq.Header, r.Header)
	outReq.Header.Del("Authorization") // Don't forward our basic auth

	res, err := http.DefaultClient.Do(outReq)
	if err != nil {
		log.Printf("%s: forwarding to %q failed: %v", r.RemoteAddr, targetURL, err)
		http.Error(w, "could not reach endpoint", http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	// If not StatusOK, return upstream error
	if res.StatusCode != http.StatusOK {
		log.Printf("%s: serving upstream error %q: %s", r.RemoteAddr, targetURL, res.Status)

		copyHeader(w.Header(), res.Header)
		w.WriteHeader(res.StatusCode)

		_, _ = io.Copy(w, res.Body)
		return
	}

	var outW io.Writer = w
	if objR == nil {
		// If not exist and authenticated, create

		log.Printf("%s: populating object %q", r.RemoteAddr, targetURL)
		objW := obj.If(storage.Conditions{DoesNotExist: true}).NewWriter(r.Context())
		defer objW.Close()

		outW = io.MultiWriter(outW, objW)
	} else if res.ContentLength != -1 && res.ContentLength != objR.Attrs.Size {
		// If diff and authenticated, update

		log.Printf("%s: replacing object %q: size differs (orig, mirror) %d != %d", r.RemoteAddr, targetURL, res.ContentLength, objR.Attrs.Size)
		objW := obj.If(storage.Conditions{GenerationMatch: objR.Attrs.Generation}).NewWriter(r.Context())
		defer objW.Close()

		outW = io.MultiWriter(outW, objW)
	} else {
		// If same and authenticated, return cached
		log.Printf("%s: serving cached object %q", r.RemoteAddr, targetURL)

		w.Header().Set("Content-Type", objR.Attrs.ContentType)
		w.Header().Set("Content-Length", fmt.Sprintf("%d", objR.Attrs.Size))
		w.WriteHeader(http.StatusOK)

		_, _ = io.Copy(w, objR)
		return
	}

	copyHeader(w.Header(), res.Header)
	w.WriteHeader(res.StatusCode)

	_, _ = io.Copy(outW, res.Body)
}

func isAuthenticated(r *http.Request) bool {
	user, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}

	return user == flagUser && pass == flagPass
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
