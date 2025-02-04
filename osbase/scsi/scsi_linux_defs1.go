// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build ignore

package scsi

// #include <scsi/sg.h>
import "C"

const (
	SG_IO                = C.SG_IO
	SG_DXFER_NONE        = C.SG_DXFER_NONE
	SG_DXFER_TO_DEV      = C.SG_DXFER_TO_DEV
	SG_DXFER_FROM_DEV    = C.SG_DXFER_FROM_DEV
	SG_DXFER_TO_FROM_DEV = C.SG_DXFER_TO_FROM_DEV
)

type sgIOHdr C.sg_io_hdr_t
