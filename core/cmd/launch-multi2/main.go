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
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"

	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"git.monogon.dev/source/nexantic.git/core/internal/launch"
	apb "git.monogon.dev/source/nexantic.git/core/proto/api"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sigs
		cancel()
	}()
	sw0, vm0, err := launch.NewSocketPair()
	if err != nil {
		log.Fatalf("Failed to create network pipe: %v\n", err)
	}
	sw1, vm1, err := launch.NewSocketPair()
	if err != nil {
		log.Fatalf("Failed to create network pipe: %v\n", err)
	}

	go func() {
		if err := launch.Launch(ctx, launch.Options{ConnectToSocket: vm0, SerialPort: os.Stdout}); err != nil {
			log.Fatalf("Failed to launch vm0: %v", err)
		}
	}()
	nanoswitchPortMap := make(launch.PortMap)
	identityPorts := []uint16{
		common.ExternalServicePort,
		common.DebugServicePort,
		common.KubernetesAPIPort,
	}
	for _, port := range identityPorts {
		nanoswitchPortMap[port] = port
	}
	go func() {
		opts := []grpcretry.CallOption{
			grpcretry.WithBackoff(grpcretry.BackoffExponential(100 * time.Millisecond)),
		}
		conn, err := nanoswitchPortMap.DialGRPC(common.ExternalServicePort, grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(opts...)))
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		cmc := api.NewClusterManagementClient(conn)
		res, err := cmc.NewEnrolmentConfig(context.Background(), &api.NewEnrolmentConfigRequest{
			Name: "test",
		}, grpcretry.WithMax(10))
		if err != nil {
			log.Fatalf("Failed to get enrolment config: %v", err)
		}
		if err := launch.Launch(ctx, launch.Options{ConnectToSocket: vm1, EnrolmentConfig: res.EnrolmentConfig, SerialPort: os.Stdout}); err != nil {
			log.Fatalf("Failed to launch vm1: %v", err)
		}
	}()
	if err := launch.RunMicroVM(ctx, &launch.MicroVMOptions{
		SerialPort:             os.Stdout,
		KernelPath:             "core/tools/ktest/linux-testing.elf",
		InitramfsPath:          "core/cmd/nanoswitch/initramfs.lz4",
		ExtraNetworkInterfaces: []*os.File{sw0, sw1},
		PortMap:                nanoswitchPortMap,
	}); err != nil {
		log.Fatalf("Failed to launch nanoswitch: %v", err)
	}
}
