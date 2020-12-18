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

	common "git.monogon.dev/source/nexantic.git/metropolis/node"
	apb "git.monogon.dev/source/nexantic.git/metropolis/proto/api"
	"git.monogon.dev/source/nexantic.git/metropolis/test/launch"
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
		conn, err := nanoswitchPortMap.DialGRPC(common.DebugServicePort, grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(opts...)))
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		debug := apb.NewNodeDebugServiceClient(conn)
		res, err := debug.GetGoldenTicket(ctx, &apb.GetGoldenTicketRequest{
			// HACK: this is assigned by DHCP, and we assume that everything goes well.
			ExternalIp: "10.1.0.3",
		}, grpcretry.WithMax(10))
		if err != nil {
			log.Fatalf("Failed to get golden ticket: %v", err)
		}

		ec := &apb.EnrolmentConfig{
			GoldenTicket: res.Ticket,
		}

		if err := launch.Launch(ctx, launch.Options{ConnectToSocket: vm1, EnrolmentConfig: ec, SerialPort: os.Stdout}); err != nil {
			log.Fatalf("Failed to launch vm1: %v", err)
		}
	}()
	if err := launch.RunMicroVM(ctx, &launch.MicroVMOptions{
		SerialPort:             os.Stdout,
		KernelPath:             "metropolis/test/ktest/linux-testing.elf",
		InitramfsPath:          "metropolis/test/nanoswitch/initramfs.lz4",
		ExtraNetworkInterfaces: []*os.File{sw0, sw1},
		PortMap:                nanoswitchPortMap,
	}); err != nil {
		log.Fatalf("Failed to launch nanoswitch: %v", err)
	}
}
