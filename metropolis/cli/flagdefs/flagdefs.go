// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package flagdefs contains shared flag definitions for Metropolis.
// The usage is the same as for the standard flags in the [flag] package,
// except that the [flag.FlagSet] needs to be passed in the first parameter.
// Pass [flag.CommandLine] to use the default FlagSet.
// There are also separate functions for use with the [pflag] package.
package flagdefs

import (
	"errors"
	"flag"
	"strings"

	"github.com/spf13/pflag"

	cpb "source.monogon.dev/metropolis/proto/common"
)

// tpmModeValue implements the [flag.Value] and [pflag.Value] interfaces.
type tpmModeValue cpb.ClusterConfiguration_TPMMode

func (v *tpmModeValue) Set(val string) error {
	var tpmMode cpb.ClusterConfiguration_TPMMode
	switch strings.ToLower(val) {
	case "required", "require":
		tpmMode = cpb.ClusterConfiguration_TPM_MODE_REQUIRED
	case "best-effort", "besteffort":
		tpmMode = cpb.ClusterConfiguration_TPM_MODE_BEST_EFFORT
	case "disabled", "disable":
		tpmMode = cpb.ClusterConfiguration_TPM_MODE_DISABLED
	default:
		return errors.New("must be one of: required, best-effort, disabled")
	}
	*v = tpmModeValue(tpmMode)
	return nil
}

func (v *tpmModeValue) String() string {
	switch cpb.ClusterConfiguration_TPMMode(*v) {
	case cpb.ClusterConfiguration_TPM_MODE_REQUIRED:
		return "required"
	case cpb.ClusterConfiguration_TPM_MODE_BEST_EFFORT:
		return "best-effort"
	case cpb.ClusterConfiguration_TPM_MODE_DISABLED:
		return "disabled"
	default:
		return ""
	}
}

func (*tpmModeValue) Type() string {
	return "tpmMode"
}

// TPMModeVar defines a TPMMode flag with specified name, default value, and
// usage string. The argument p points to a TPMMode variable in which to store
// the value of the flag.
func TPMModeVar(flags *flag.FlagSet, p *cpb.ClusterConfiguration_TPMMode, name string, value cpb.ClusterConfiguration_TPMMode, usage string) {
	*p = value
	flags.Var((*tpmModeValue)(p), name, usage+" (one of: required, best-effort, disabled)")
}

// TPMMode defines a TPMMode flag with specified name, default value, and
// usage string. The return value is the address of a TPMMode variable that
// stores the value of the flag.
func TPMMode(flags *flag.FlagSet, name string, value cpb.ClusterConfiguration_TPMMode, usage string) *cpb.ClusterConfiguration_TPMMode {
	val := new(cpb.ClusterConfiguration_TPMMode)
	TPMModeVar(flags, val, name, value, usage)
	return val
}

// TPMModeVarPflag defines a TPMMode flag with specified name, default value,
// and usage string. The argument p points to a TPMMode variable in which to
// store the value of the flag.
func TPMModeVarPflag(flags *pflag.FlagSet, p *cpb.ClusterConfiguration_TPMMode, name string, value cpb.ClusterConfiguration_TPMMode, usage string) {
	*p = value
	flags.Var((*tpmModeValue)(p), name, usage+" (one of: required, best-effort, disabled)")
}

// TPMModePflag defines a TPMMode flag with specified name, default value, and
// usage string. The return value is the address of a TPMMode variable that
// stores the value of the flag.
func TPMModePflag(flags *pflag.FlagSet, name string, value cpb.ClusterConfiguration_TPMMode, usage string) *cpb.ClusterConfiguration_TPMMode {
	val := new(cpb.ClusterConfiguration_TPMMode)
	TPMModeVarPflag(flags, val, name, value, usage)
	return val
}

// storageSecurityPolicyValue implements the [flag.Value] and [pflag.Value]
// interfaces.
type storageSecurityPolicyValue cpb.ClusterConfiguration_StorageSecurityPolicy

func (v *storageSecurityPolicyValue) Set(val string) error {
	var storageSecurityPolicy cpb.ClusterConfiguration_StorageSecurityPolicy
	switch strings.ToLower(val) {
	case "permissive":
		storageSecurityPolicy = cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_PERMISSIVE
	case "needs-encryption":
		storageSecurityPolicy = cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION
	case "needs-encryption-and-authentication":
		storageSecurityPolicy = cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION_AND_AUTHENTICATION
	case "needs-insecure":
		storageSecurityPolicy = cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE
	default:
		return errors.New("must be one of: permissive, needs-encryption, needs-encryption-and-authentication, needs-insecure")
	}
	*v = storageSecurityPolicyValue(storageSecurityPolicy)
	return nil
}

func (v *storageSecurityPolicyValue) String() string {
	switch cpb.ClusterConfiguration_StorageSecurityPolicy(*v) {
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_PERMISSIVE:
		return "permissive"
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION:
		return "needs-encryption"
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_ENCRYPTION_AND_AUTHENTICATION:
		return "needs-encryption-and-authentication"
	case cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE:
		return "needs-insecure"
	default:
		return ""
	}
}

func (*storageSecurityPolicyValue) Type() string {
	return "storageSecurityPolicy"
}

// StorageSecurityPolicyVar defines a StorageSecurityPolicy flag with specified
// name, default value, and usage string. The argument p points to a
// StorageSecurityPolicy variable in which to store the value of the flag.
func StorageSecurityPolicyVar(flags *flag.FlagSet, p *cpb.ClusterConfiguration_StorageSecurityPolicy, name string, value cpb.ClusterConfiguration_StorageSecurityPolicy, usage string) {
	*p = value
	flags.Var((*storageSecurityPolicyValue)(p), name, usage+" (one of: permissive, needs-encryption, needs-encryption-and-authentication, needs-insecure)")
}

// StorageSecurityPolicy defines a StorageSecurityPolicy flag with specified
// name, default value, and usage string. The return value is the address of a
// StorageSecurityPolicy variable that stores the value of the flag.
func StorageSecurityPolicy(flags *flag.FlagSet, name string, value cpb.ClusterConfiguration_StorageSecurityPolicy, usage string) *cpb.ClusterConfiguration_StorageSecurityPolicy {
	val := new(cpb.ClusterConfiguration_StorageSecurityPolicy)
	StorageSecurityPolicyVar(flags, val, name, value, usage)
	return val
}

// StorageSecurityPolicyVarPflag defines a StorageSecurityPolicy flag with
// specified name, default value, and usage string. The argument p points to a
// StorageSecurityPolicy variable in which to store the value of the flag.
func StorageSecurityPolicyVarPflag(flags *pflag.FlagSet, p *cpb.ClusterConfiguration_StorageSecurityPolicy, name string, value cpb.ClusterConfiguration_StorageSecurityPolicy, usage string) {
	*p = value
	flags.Var((*storageSecurityPolicyValue)(p), name, usage+" (one of: permissive, needs-encryption, needs-encryption-and-authentication, needs-insecure)")
}

// StorageSecurityPolicyPflag defines a StorageSecurityPolicy flag with
// specified name, default value, and usage string. The return value is the
// address of a StorageSecurityPolicy variable that stores the value of the
// flag.
func StorageSecurityPolicyPflag(flags *pflag.FlagSet, name string, value cpb.ClusterConfiguration_StorageSecurityPolicy, usage string) *cpb.ClusterConfiguration_StorageSecurityPolicy {
	val := new(cpb.ClusterConfiguration_StorageSecurityPolicy)
	StorageSecurityPolicyVarPflag(flags, val, name, value, usage)
	return val
}
