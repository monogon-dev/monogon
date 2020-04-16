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

package kubernetes

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"regexp"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes/wrappers"
	"go.uber.org/zap"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pluginregistration "k8s.io/kubelet/pkg/apis/pluginregistration/v1"

	"git.monogon.dev/source/nexantic.git/core/internal/storage"
	"git.monogon.dev/source/nexantic.git/core/pkg/fsquota"
)

// Derived from K8s spec for acceptable names, but shortened to 130 characters to avoid issues with
// maximum path length. We don't provision longer names so this applies only if you manually create
// a volume with a name of more than 130 characters.
var acceptableNames = regexp.MustCompile("^[a-z][a-bz0-9-.]{0,128}[a-z0-9]$")

const volumeDir = "volumes"

const pluginSocketPath = "/data/kubernetes/kubelet/plugins/com.smalltown.vfs.sock"

type csiServer struct {
	manager *storage.Manager
	logger  *zap.Logger
}

func runCSIPlugin(storMan *storage.Manager) supervisor.Runnable {
	return func(ctx context.Context) error {
		s := &csiServer{
			manager: storMan,
			logger:  supervisor.Logger(ctx),
		}
		pluginListener, err := net.ListenUnix("unix", &net.UnixAddr{Name: pluginSocketPath, Net: "unix"})
		if err != nil {
			return fmt.Errorf("failed to listen on CSI socket: %w", err)
		}
		pluginListener.SetUnlinkOnClose(true)

		pluginServer := grpc.NewServer()
		csi.RegisterIdentityServer(pluginServer, s)
		csi.RegisterNodeServer(pluginServer, s)
		// Enable graceful shutdown since we don't have long-running RPCs and most of them shouldn't and can't be
		// cancelled anyways.
		if err := supervisor.Run(ctx, "csi-node", supervisor.GRPCServer(pluginServer, pluginListener, true)); err != nil {
			return err
		}

		registrationListener, err := net.ListenUnix("unix", &net.UnixAddr{
			Name: "/data/kubernetes/kubelet/plugins_registry/com.smalltown.vfs-reg.sock",
			Net:  "unix",
		})
		if err != nil {
			return fmt.Errorf("failed to listen on CSI registration socket: %w", err)
		}
		registrationListener.SetUnlinkOnClose(true)

		registrationServer := grpc.NewServer()
		pluginregistration.RegisterRegistrationServer(registrationServer, s)
		if err := supervisor.Run(ctx, "registration", supervisor.GRPCServer(registrationServer, registrationListener, true)); err != nil {
			return err
		}
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		supervisor.Signal(ctx, supervisor.SignalDone)
		return nil
	}
}

func (*csiServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeStageVolume not supported")
}
func (*csiServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeUnstageVolume not supported")
}

func (s *csiServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	if !acceptableNames.MatchString(req.VolumeId) {
		return nil, status.Error(codes.InvalidArgument, "invalid characters in volume id")
	}
	volumePath, err := s.manager.GetPathInPlace(storage.PlaceData, filepath.Join(volumeDir, req.VolumeId))
	if err != nil {
		return nil, status.Error(codes.Unavailable, "persistent data storage not available")
	}
	switch req.VolumeCapability.AccessMode.Mode {
	case csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER:
	case csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY:
	default:
		return nil, status.Error(codes.InvalidArgument, "unsupported access mode")
	}
	switch req.VolumeCapability.AccessType.(type) {
	case *csi.VolumeCapability_Mount:
	default:
		return nil, status.Error(codes.InvalidArgument, "unsupported access type")
	}

	err = unix.Mount(volumePath, req.TargetPath, "", unix.MS_BIND, "")
	if err == unix.ENOENT {
		return nil, status.Error(codes.NotFound, "volume not found")
	} else if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to bind-mount volume: %v", err)
	}
	if req.Readonly {
		err := unix.Mount(volumePath, req.TargetPath, "", unix.MS_BIND|unix.MS_REMOUNT|unix.MS_RDONLY, "")
		if err != nil {
			_ = unix.Unmount(req.TargetPath, 0) // Best-effort
			return nil, status.Errorf(codes.Unavailable, "failed to remount volume: %v", err)
		}
	}
	return &csi.NodePublishVolumeResponse{}, nil
}
func (*csiServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	if err := unix.Unmount(req.TargetPath, 0); err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to unmount volume: %v", err)
	}
	return &csi.NodeUnpublishVolumeResponse{}, nil
}
func (*csiServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	quota, err := fsquota.GetQuota(req.VolumePath)
	if os.IsNotExist(err) {
		return nil, status.Error(codes.NotFound, "volume does not exist at this path")
	} else if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to get quota: %v", err)
	}

	return &csi.NodeGetVolumeStatsResponse{
		Usage: []*csi.VolumeUsage{
			{
				Total:     int64(quota.Bytes),
				Unit:      csi.VolumeUsage_BYTES,
				Used:      int64(quota.BytesUsed),
				Available: int64(quota.Bytes - quota.BytesUsed),
			},
			{
				Total:     int64(quota.Inodes),
				Unit:      csi.VolumeUsage_INODES,
				Used:      int64(quota.InodesUsed),
				Available: int64(quota.Inodes - quota.InodesUsed),
			},
		},
	}, nil
}
func (*csiServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	if req.CapacityRange.LimitBytes <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid expanded volume size: at or below zero bytes")
	}
	if err := fsquota.SetQuota(req.VolumePath, uint64(req.CapacityRange.LimitBytes), 0); err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to update quota: %v", err)
	}
	return &csi.NodeExpandVolumeResponse{CapacityBytes: req.CapacityRange.LimitBytes}, nil
}

func rpcCapability(cap csi.NodeServiceCapability_RPC_Type) *csi.NodeServiceCapability {
	return &csi.NodeServiceCapability{
		Type: &csi.NodeServiceCapability_Rpc{
			Rpc: &csi.NodeServiceCapability_RPC{Type: cap},
		},
	}
}
func (*csiServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			rpcCapability(csi.NodeServiceCapability_RPC_EXPAND_VOLUME),
			rpcCapability(csi.NodeServiceCapability_RPC_GET_VOLUME_STATS),
		},
	}, nil
}
func (*csiServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to get node identity: %v", err)
	}
	return &csi.NodeGetInfoResponse{
		NodeId: hostname,
	}, nil
}

// CSI Identity endpoints
func (*csiServer) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return &csi.GetPluginInfoResponse{
		Name:          "com.smalltown.vfs",
		VendorVersion: "0.0.1", // TODO(lorenz): Maybe stamp?
	}, nil
}
func (*csiServer) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	return &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_VolumeExpansion_{
					VolumeExpansion: &csi.PluginCapability_VolumeExpansion{
						Type: csi.PluginCapability_VolumeExpansion_ONLINE,
					},
				},
			},
		},
	}, nil
}

func (s *csiServer) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	_, err := s.manager.GetPathInPlace(storage.PlaceData, volumeDir)
	ready := err == nil
	return &csi.ProbeResponse{Ready: &wrappers.BoolValue{Value: ready}}, nil
}

// Registration endpoints
func (s *csiServer) GetInfo(ctx context.Context, req *pluginregistration.InfoRequest) (*pluginregistration.PluginInfo, error) {
	return &pluginregistration.PluginInfo{
		Type:              "CSIPlugin",
		Name:              "com.smalltown.vfs",
		Endpoint:          pluginSocketPath,
		SupportedVersions: []string{"1.2"}, // Keep in sync with container-storage-interface/spec package version
	}, nil
}

func (s *csiServer) NotifyRegistrationStatus(ctx context.Context, req *pluginregistration.RegistrationStatus) (*pluginregistration.RegistrationStatusResponse, error) {
	if req.Error != "" {
		s.logger.Warn("Kubelet failed registering CSI plugin", zap.String("error", req.Error))
	}
	return &pluginregistration.RegistrationStatusResponse{}, nil
}
