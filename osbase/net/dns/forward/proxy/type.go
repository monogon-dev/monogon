// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package proxy

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"net"
)

type transportType int

const (
	typeUDP transportType = iota
	typeTCP
	typeTLS
	typeTotalCount // keep this last
)

func stringToTransportType(s string) transportType {
	switch s {
	case "udp":
		return typeUDP
	case "tcp":
		return typeTCP
	case "tcp-tls":
		return typeTLS
	}

	return typeUDP
}

func (t *Transport) transportTypeFromConn(pc *persistConn) transportType {
	if _, ok := pc.c.Conn.(*net.UDPConn); ok {
		return typeUDP
	}

	if t.tlsConfig == nil {
		return typeTCP
	}

	return typeTLS
}
