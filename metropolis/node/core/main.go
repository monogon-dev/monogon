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

package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"os/signal"
	"runtime/debug"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/cluster"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/kubernetes"
	"source.monogon.dev/metropolis/node/kubernetes/containerd"
	"source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/pkg/supervisor"
	"source.monogon.dev/metropolis/pkg/tpm"
	apb "source.monogon.dev/metropolis/proto/api"
)

var (
	// kubernetesConfig is the static/global part of the Kubernetes service configuration. In the future, this might
	// be configurable by loading it from the EnrolmentConfig. Fow now, it's static and same across all clusters.
	kubernetesConfig = kubernetes.Config{
		ServiceIPRange: net.IPNet{ // TODO(q3k): Decide if configurable / final value
			IP:   net.IP{10, 0, 255, 1},
			Mask: net.IPMask{0xff, 0xff, 0xff, 0x00}, // /24, but Go stores as a literal mask
		},
		ClusterNet: net.IPNet{
			IP:   net.IP{10, 0, 0, 0},
			Mask: net.IPMask{0xff, 0xff, 0x00, 0x00}, // /16
		},
	}
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Init panicked:", r)
			debug.PrintStack()
		}
		unix.Sync()
		// TODO(lorenz): Switch this to Reboot when init panics are less likely
		// Best effort, nothing we can do if this fails except printing the error to the console.
		if err := unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF); err != nil {
			panic(fmt.Sprintf("failed to halt node: %v\n", err))
		}
	}()

	// Set up logger for Metropolis. Currently logs everything to stderr.
	lt := logtree.New()
	reader, err := lt.Read("", logtree.WithChildren(), logtree.WithStream())
	if err != nil {
		panic(fmt.Errorf("could not set up root log reader: %v", err))
	}
	go func() {
		for {
			p := <-reader.Stream
			fmt.Fprintf(os.Stderr, "%s\n", p.String())
		}
	}()

	// Initial logger. Used until we get to a supervisor.
	logger := lt.MustLeveledFor("init")

	// Set up basic mounts
	err = setupMounts(logger)
	if err != nil {
		panic(fmt.Errorf("could not set up basic mounts: %w", err))
	}

	// Linux kernel default is 4096 which is far too low. Raise it to 1M which is what gVisor suggests.
	if err := unix.Setrlimit(unix.RLIMIT_NOFILE, &unix.Rlimit{Cur: 1048576, Max: 1048576}); err != nil {
		logger.Fatalf("Failed to raise rlimits: %v", err)
	}

	logger.Info("Starting Metropolis node init")

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel)

	if err := tpm.Initialize(logger); err != nil {
		logger.Fatalf("Failed to initialize TPM 2.0: %v", err)
	}

	networkSvc := network.New()

	// This function initializes a headless Delve if this is a debug build or does nothing if it's not
	initializeDebugger(networkSvc)

	// Prepare local storage.
	root := &localstorage.Root{}
	if err := declarative.PlaceFS(root, "/"); err != nil {
		panic(fmt.Errorf("when placing root FS: %w", err))
	}

	// trapdoor is a channel used to signal to the init service that a very low-level, unrecoverable failure
	// occured. This causes a GURU MEDITATION ERROR visible to the end user.
	trapdoor := make(chan struct{})

	// Make context for supervisor. We cancel it when we reach the trapdoor.
	ctxS, ctxC := context.WithCancel(context.Background())

	// Start root initialization code as a supervisor one-shot runnable. This means waiting for the network, starting
	// the cluster manager, and then starting all services related to the node's roles.
	// TODO(q3k): move this to a separate 'init' service.
	supervisor.New(ctxS, func(ctx context.Context) error {
		logger := supervisor.Logger(ctx)

		// Start storage and network - we need this to get anything else done.
		if err := root.Start(ctx); err != nil {
			return fmt.Errorf("cannot start root FS: %w", err)
		}
		if err := supervisor.Run(ctx, "network", networkSvc.Run); err != nil {
			return fmt.Errorf("when starting network: %w", err)
		}

		// Start cluster manager. This kicks off cluster membership machinery, which will either start
		// a new cluster, enroll into one or join one.
		m := cluster.NewManager(root, networkSvc)
		if err := supervisor.Run(ctx, "enrolment", m.Run); err != nil {
			return fmt.Errorf("when starting enrolment: %w", err)
		}

		// Wait until the cluster manager settles.
		node, err := m.Wait()
		if err != nil {
			close(trapdoor)
			return fmt.Errorf("enrolment failed, aborting: %w", err)
		}

		// We are now in a cluster. We can thus access our 'node' object and start all services that
		// we should be running.

		logger.Info("Enrolment success, continuing startup.")
		logger.Info(fmt.Sprintf("This node (%s) has roles:", node.String()))
		if cm := node.ConsensusMember(); cm != nil {
			// There's no need to start anything for when we are a consensus member - the cluster
			// manager does this for us if necessary (as creating/enrolling/joining a cluster is
			// pretty tied into cluster lifecycle management).
			logger.Info(fmt.Sprintf(" - etcd consensus member"))
		}
		if kw := node.KubernetesWorker(); kw != nil {
			logger.Info(fmt.Sprintf(" - kubernetes worker"))
		}

		// If we're supposed to be a kubernetes worker, start kubernetes services and containerd.
		// In the future, this might be split further into kubernetes control plane and data plane
		// roles.
		var containerdSvc *containerd.Service
		var kubeSvc *kubernetes.Service
		if kw := node.KubernetesWorker(); kw != nil {
			logger.Info("Starting Kubernetes worker services...")

			// Ensure Kubernetes PKI objects exist in etcd.
			kpki := pki.New(lt.MustLeveledFor("pki.kubernetes"), node.KV)
			if err := kpki.EnsureAll(ctx); err != nil {
				return fmt.Errorf("failed to ensure kubernetes PKI present: %w", err)
			}

			containerdSvc = &containerd.Service{
				EphemeralVolume: &root.Ephemeral.Containerd,
			}
			if err := supervisor.Run(ctx, "containerd", containerdSvc.Run); err != nil {
				return fmt.Errorf("failed to start containerd service: %w", err)
			}

			kubernetesConfig.KPKI = kpki
			kubernetesConfig.Root = root
			kubernetesConfig.Network = networkSvc
			kubeSvc = kubernetes.New(kubernetesConfig)
			if err := supervisor.Run(ctx, "kubernetes", kubeSvc.Run); err != nil {
				return fmt.Errorf("failed to start kubernetes service: %w", err)
			}

		}

		// Start the node debug service.
		dbg := &debugService{
			cluster:    m,
			logtree:    lt,
			kubernetes: kubeSvc,
			traceLock:  make(chan struct{}, 1),
		}
		dbgSrv := grpc.NewServer()
		apb.RegisterNodeDebugServiceServer(dbgSrv, dbg)
		dbgLis, err := net.Listen("tcp", fmt.Sprintf(":%d", common.DebugServicePort))
		if err != nil {
			return fmt.Errorf("failed to listen on debug service: %w", err)
		}
		if err := supervisor.Run(ctx, "debug", supervisor.GRPCServer(dbgSrv, dbgLis, false)); err != nil {
			return fmt.Errorf("failed to start debug service: %w", err)
		}

		supervisor.Signal(ctx, supervisor.SignalHealthy)
		supervisor.Signal(ctx, supervisor.SignalDone)
		return nil
	}, supervisor.WithExistingLogtree(lt))

	// We're PID1, so orphaned processes get reparented to us to clean up
	for {
		select {
		case <-trapdoor:
			// If the trapdoor got closed, we got stuck early enough in the boot process that we can't do anything about
			// it. Display a generic error message until we handle error conditions better.
			ctxC()
			log.Printf("                  ########################")
			log.Printf("                  # GURU MEDIATION ERROR #")
			log.Printf("                  ########################")
			log.Printf("")
			log.Printf("Metropolis encountered an uncorrectable error and this node must be")
			log.Printf("restarted.")
			log.Printf("")
			log.Printf("(Error condition: init trapdoor closed)")
			log.Printf("")
			select {}

		case sig := <-signalChannel:
			switch sig {
			case unix.SIGCHLD:
				var status unix.WaitStatus
				var rusage unix.Rusage
				for {
					res, err := unix.Wait4(-1, &status, unix.WNOHANG, &rusage)
					if err != nil && err != unix.ECHILD {
						logger.Errorf("Failed to wait on orphaned child: %v", err)
						break
					}
					if res <= 0 {
						break
					}
				}
			case unix.SIGURG:
				// Go 1.14 introduced asynchronous preemption, which uses SIGURG.
				// In order not to break backwards compatibility in the unlikely case
				// of an application actually using SIGURG on its own, they're not filtering them.
				// (https://github.com/golang/go/issues/37942)
				logger.V(5).Info("Ignoring SIGURG")
			// TODO(lorenz): We can probably get more than just SIGCHLD as init, but I can't think
			// of any others right now, just log them in case we hit any of them.
			default:
				logger.Warningf("Got unexpected signal %s", sig.String())
			}
		}
	}
}

// nodeCertificate creates a node key/certificate for a foreign node. This is duplicated code with localstorage's
// PKIDirectory EnsureSelfSigned, but is temporary (and specific to 'golden tickets').
func (s *debugService) nodeCertificate() (cert, key []byte, err error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		err = fmt.Errorf("failed to generate key: %w", err)
		return
	}

	key, err = x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		err = fmt.Errorf("failed to marshal key: %w", err)
		return
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = fmt.Errorf("failed to generate serial number: %w", err)
		return
	}

	template := localstorage.CertificateForNode(pubKey)
	template.SerialNumber = serialNumber

	cert, err = x509.CreateCertificate(rand.Reader, &template, &template, pubKey, privKey)
	if err != nil {
		err = fmt.Errorf("could not sign certificate: %w", err)
		return
	}
	return
}
