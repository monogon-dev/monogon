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

package containerd

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	ctr "github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"

	"git.monogon.dev/source/nexantic.git/metropolis/node/core/localstorage"
	"git.monogon.dev/source/nexantic.git/metropolis/pkg/supervisor"
)

const (
	preseedNamespacesDir = "/containerd/preseed/"
)

type Service struct {
	EphemeralVolume *localstorage.EphemeralContainerdDirectory
}

func (s *Service) Run(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "/containerd/bin/containerd", "--config", "/containerd/conf/config.toml")
	cmd.Env = []string{"PATH=/containerd/bin", "TMPDIR=" + s.EphemeralVolume.Tmp.FullPath()}

	runscFifo, err := os.OpenFile(s.EphemeralVolume.RunSCLogsFIFO.FullPath(), os.O_CREATE|os.O_RDONLY, os.ModeNamedPipe|0777)
	if err != nil {
		return err
	}

	if err := supervisor.Run(ctx, "runsc", s.logPump(runscFifo)); err != nil {
		return fmt.Errorf("failed to start runsc log pump: %w", err)
	}

	if err := supervisor.Run(ctx, "preseed", s.runPreseed); err != nil {
		return fmt.Errorf("failed to start preseed runnable: %w", err)
	}
	return supervisor.RunCommand(ctx, cmd)
}

// logPump returns a runnable that pipes data from a file/FIFO into its raw logger.
// TODO(q3k): refactor this out to a generic function in supervisor or logtree.
func (s *Service) logPump(fifo *os.File) supervisor.Runnable {
	return func(ctx context.Context) error {
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		for {
			// Quit if requested.
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			n, err := io.Copy(supervisor.RawLogger(ctx), fifo)
			if n == 0 && err == nil {
				// Hack because pipes/FIFOs can return zero reads when nobody is writing. To avoid busy-looping,
				// sleep a bit before retrying. This does not loose data since the FIFO internal buffer will
				// stall writes when it becomes full. 10ms maximum stall in a non-latency critical process (reading
				// debug logs) is not an issue for us.
				time.Sleep(10 * time.Millisecond)
			} else if err != nil {
				return fmt.Errorf("log pump failed: %v", err)
			}
		}
	}
}

// runPreseed loads OCI bundles in tar form from preseedNamespacesDir into containerd at startup.
// This can be run multiple times, containerd will automatically dedup the layers.
// containerd uses namespaces to keep images (and everything else) separate so to define where the images will be loaded
// to they need to be in a folder named after the namespace they should be loaded into.
// containerd's CRI plugin (which is built as part of containerd) uses a hardcoded namespace ("k8s.io") for everything
// accessed through CRI, so if an image should be available on K8s it needs to be in that namespace.
// As an example if image helloworld should be loaded for use with Kubernetes, the OCI bundle needs to be at
// <preseedNamespacesDir>/k8s.io/helloworld.tar. No tagging beyond what's in the bundle is performed.
func (s *Service) runPreseed(ctx context.Context) error {
	client, err := ctr.New(s.EphemeralVolume.ClientSocket.FullPath())
	if err != nil {
		return fmt.Errorf("failed to connect to containerd: %w", err)
	}
	logger := supervisor.Logger(ctx)
	preseedNamespaceDirs, err := ioutil.ReadDir(preseedNamespacesDir)
	if err != nil {
		return fmt.Errorf("failed to open preseed dir: %w", err)
	}
	for _, dir := range preseedNamespaceDirs {
		if !dir.IsDir() {
			logger.Warningf("Non-Directory %q found in preseed folder, ignoring", dir.Name())
			continue
		}
		namespace := dir.Name()
		images, err := ioutil.ReadDir(filepath.Join(preseedNamespacesDir, namespace))
		if err != nil {
			return fmt.Errorf("failed to list namespace preseed directory for ns \"%v\": %w", namespace, err)
		}
		ctxWithNS := namespaces.WithNamespace(ctx, namespace)
		for _, image := range images {
			if image.IsDir() {
				logger.Warningf("Directory %q found in preseed namespaced folder, ignoring", image.Name())
				continue
			}
			imageFile, err := os.Open(filepath.Join(preseedNamespacesDir, namespace, image.Name()))
			if err != nil {
				return fmt.Errorf("failed to open preseed image \"%v\": %w", image.Name(), err)
			}
			// defer in this loop is fine since we're never going to preseed more than ~1M images which is where our
			// file descriptor limit is.
			defer imageFile.Close()
			importedImages, err := client.Import(ctxWithNS, imageFile)
			if err != nil {
				return fmt.Errorf("failed to import preseed image: %w", err)
			}
			var importedImageNames []string
			for _, img := range importedImages {
				importedImageNames = append(importedImageNames, img.Name)
			}
			logger.Infof("Successfully imported preseeded bundle %s/%s into containerd", namespace, strings.Join(importedImageNames, ","))
		}
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
