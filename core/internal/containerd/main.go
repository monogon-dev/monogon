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
	"time"

	ctr "github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"go.uber.org/zap"

	"git.monogon.dev/source/nexantic.git/core/internal/localstorage"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/pkg/logbuffer"
)

const (
	preseedNamespacesDir = "/containerd/preseed/"
)

type Service struct {
	EphemeralVolume *localstorage.EphemeralContainerdDirectory
	Log             *logbuffer.LogBuffer
	RunscLog        *logbuffer.LogBuffer
}

func (s *Service) Run(ctx context.Context) error {
	if s.Log == nil {
		s.Log = logbuffer.New(5000, 16384)
	}
	if s.RunscLog == nil {
		s.RunscLog = logbuffer.New(5000, 16384)
	}

	logger := supervisor.Logger(ctx)

	cmd := exec.CommandContext(ctx, "/containerd/bin/containerd", "--config", "/containerd/conf/config.toml")
	cmd.Stdout = s.Log
	cmd.Stderr = s.Log
	cmd.Env = []string{"PATH=/containerd/bin", "TMPDIR=" + s.EphemeralVolume.Tmp.FullPath()}

	runscFifo, err := os.OpenFile(s.EphemeralVolume.RunSCLogsFIFO.FullPath(), os.O_CREATE|os.O_RDONLY, os.ModeNamedPipe|0777)
	if err != nil {
		return err
	}
	go func() {
		for {
			n, err := io.Copy(s.RunscLog, runscFifo)
			if n == 0 && err == nil {
				// Hack because pipes/FIFOs can return zero reads when nobody is writing. To avoid busy-looping,
				// sleep a bit before retrying. This does not loose data since the FIFO internal buffer will
				// stall writes when it becomes full. 10ms maximum stall in a non-latency critical process (reading
				// debug logs) is not an issue for us.
				time.Sleep(10 * time.Millisecond)
			} else if err != nil {
				logger.Error("gVisor log pump failed, stopping it", zap.Error(err))
				return // It's likely that this will busy-loop printing errors if it encounters one, so bail
			}
		}
	}()

	if err := supervisor.Run(ctx, "preseed", s.runPreseed); err != nil {
		return fmt.Errorf("failed to start preseed runnable: %w", err)
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	err = cmd.Run()
	fmt.Fprintf(s.Log, "containerd stopped: %v\n", err)
	return err
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
			logger.Warn("Non-Directory found in preseed folder, ignoring", zap.String("name", dir.Name()))
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
				logger.Warn("Directory found in preseed namespaced folder, ignoring", zap.String("name", image.Name()))
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
			logger.Info("Successfully imported preseeded bundle into containerd",
				zap.String("namespace", namespace), zap.Strings("images", importedImageNames))
		}
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
