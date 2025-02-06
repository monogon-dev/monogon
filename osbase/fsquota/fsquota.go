// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package fsquota provides a simplified interface to interact with Linux's
// filesystem qouta API.  It only supports setting quotas on directories, not
// groups or users.  Quotas need to be already enabled on the filesystem to be
// able to use them using this package.  See the quotactl package if you intend
// to use this on a filesystem where quotas need to be enabled manually.
package fsquota

import (
	"errors"
	"fmt"
	"math"
	"os"

	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/fsquota/fsxattrs"
	"source.monogon.dev/osbase/fsquota/quotactl"
)

// SetQuota sets the quota of bytes and/or inodes in a given path. To not set a
// limit, set the corresponding argument to zero. Setting both arguments to
// zero removes the quota entirely.  This function can only be called on an
// empty directory. It can't be used to create a quota below a directory which
// already has a quota since Linux doesn't offer hierarchical quotas.
func SetQuota(path string, maxBytes uint64, maxInodes uint64) error {
	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	defer dir.Close()
	var valid uint32
	if maxBytes > 0 {
		valid |= quotactl.FlagBLimitsValid
	}
	if maxInodes > 0 {
		valid |= quotactl.FlagILimitsValid
	}

	attrs, err := fsxattrs.Get(dir)
	if err != nil {
		return err
	}

	var lastID = attrs.ProjectID
	if lastID == 0 {
		// No project/quota exists for this directory, assign a new project
		// quota.
		// TODO(lorenz): This is racy, but the kernel does not support
		// atomically assigning quotas. So this needs to be added to the
		// kernels setquota interface. Due to the short time window and
		// infrequent calls this should not be an immediate issue.
		for {
			quota, err := quotactl.GetNextQuota(dir, quotactl.QuotaTypeProject, lastID)
			if errors.Is(err, unix.ENOENT) || errors.Is(err, unix.ESRCH) {
				// We have enumerated all quotas, nothing exists here
				break
			} else if err != nil {
				return fmt.Errorf("failed to call GetNextQuota: %w", err)
			}
			if quota.ID > lastID+1 {
				// Take the first ID in the quota ID gap
				lastID++
				break
			}
			lastID++
		}
	}

	// If both limits are zero, this is a delete operation, process it as such
	if maxBytes == 0 && maxInodes == 0 {
		valid = quotactl.FlagBLimitsValid | quotactl.FlagILimitsValid
		attrs.ProjectID = 0
		attrs.Flags &= ^fsxattrs.FlagProjectInherit
	} else {
		attrs.ProjectID = lastID
		attrs.Flags |= fsxattrs.FlagProjectInherit
	}

	if err := fsxattrs.Set(dir, attrs); err != nil {
		return err
	}

	// Always round up to the nearest block size
	bytesLimitBlocks := uint64(math.Ceil(float64(maxBytes) / float64(1024)))

	return quotactl.SetQuota(dir, quotactl.QuotaTypeProject, lastID, &quotactl.Quota{
		BHardLimit: bytesLimitBlocks,
		BSoftLimit: bytesLimitBlocks,
		IHardLimit: maxInodes,
		ISoftLimit: maxInodes,
		Valid:      valid,
	})
}

type Quota struct {
	Bytes      uint64
	BytesUsed  uint64
	Inodes     uint64
	InodesUsed uint64
}

// GetQuota returns the current active quota and its utilization at the given
// path
func GetQuota(path string) (*Quota, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	attrs, err := fsxattrs.Get(dir)
	if err != nil {
		return nil, err
	}
	if attrs.ProjectID == 0 {
		return nil, os.ErrNotExist
	}
	quota, err := quotactl.GetQuota(dir, quotactl.QuotaTypeProject, attrs.ProjectID)
	if err != nil {
		return nil, err
	}
	return &Quota{
		Bytes:      quota.BHardLimit * 1024,
		BytesUsed:  quota.CurSpace,
		Inodes:     quota.IHardLimit,
		InodesUsed: quota.CurInodes,
	}, nil
}
