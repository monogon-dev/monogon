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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec"`
	Status VirtualMachineStatus `json:"status"`
}

type VirtualMachineSpec struct {
	// TODO(lorenz): document
	InitialImage VirtualMachineImage `json:"initialImage"`

	// HypervsisorImage determines what OCI image will run within the pod. This
	// image must communicate using the Metropolis VM hypervisor API and run
	// the actual VM monitor (eg. qemu).
	// Defaults to "" (default image for cluster).
	// +optional
	HypervisorImage string

	// Resources are the resources assigned to the pod backing this virtual
	// machine when running. Non-integer CPU requests and overcommit will
	// result in reduced side-channel attack resistance as CPUs will not be
	// statically assigned
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// VolumeClaimTemplate is used to produce a PersistentVolumeClaim that will
	// be attached to the pod backing this virtual machine when running.
	VolumeClaimTemplate corev1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`

	// VirtualMachineIPs is a list of IP addresses (in CIDR prefix notation)
	// that the virtual machine will receive traffic for when running. These
	// can be either IPv4 or IPv6 addresses.
	VirtualMachineIPs []string `json:"vmIPs"`

	// MigrateStrategy determines what migration strategy is used when the
	// virtual machine needs to be migrated across hosts.
	// Defaults to "Live".
	// +optional
	// +default="Live"
	MigrateStrategy VirtualMachineMigrateStrategy `json:"migrateStrategy"`

	// ExecutionMode determines whether the machine is running or paused.
	// Defaults to "Run".
	// +optional
	// +default="Run"
	ExecutionMode VirtualMachineExecutionMode `json:"mode"`
}

type VirtualMachineImage struct {
	// TODO(lorenz): document
	// +optional
	URL string `json:"url"`
}

type VirtualMachineMigrateStrategy string

const (
	VirtualMachineColdMigrate VirtualMachineMigrateStrategy = "Cold"
	VirtualMachineLiveMigrate VirtualMachineMigrateStrategy = "Live"
)

type VirtualMachineExecutionMode string

const (
	VirtualMachineRun   VirtualMachineExecutionMode = "Run"
	VirtualMachinePause VirtualMachineExecutionMode = "Pause"
)

type VirtualMachineStatus struct {
	Phase           VirtualMachinePhase `json:"phase"`
	ActivePodName   string              `json:"activePodName"`
	VolumeClaimName string              `json:"volumeClaimName"`
}

type VirtualMachinePhase string

const (
	VirtualMachinePhaseCreating    VirtualMachinePhase = "Creating"
	VirtualMachinePhaseRunning     VirtualMachinePhase = "Running"
	VirtualMachinePhaseTerminating VirtualMachinePhase = "Terminating"
	VirtualMachinePhaseStopped     VirtualMachinePhase = "Stopped"
	VirtualMachinePhaseLost        VirtualMachinePhase = "Lost"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []VirtualMachine `json:"items"`
}
