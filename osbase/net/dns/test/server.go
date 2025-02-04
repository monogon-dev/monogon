// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package test

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

// A Server is an DNS server listening on a system-chosen port on the local
// loopback interface, for use in end-to-end DNS tests.
type Server struct {
	Addr string // Address where the server listening.

	s1 *dns.Server // udp
	s2 *dns.Server // tcp
}

// NewServer starts and returns a new Server. The caller should call Close when
// finished, to shut it down.
func NewServer(f dns.HandlerFunc) *Server {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	s1 := &dns.Server{Handler: f} // udp
	s2 := &dns.Server{Handler: f} // tcp

	for i := 0; i < 5; i++ { // 5 attempts
		s2.Listener, _ = net.Listen("tcp", "[::1]:0")
		if s2.Listener == nil {
			continue
		}

		s1.PacketConn, _ = net.ListenPacket("udp", s2.Listener.Addr().String())
		if s1.PacketConn != nil {
			break
		}

		// perhaps UDP port is in use, try again
		s2.Listener.Close()
		s2.Listener = nil
	}
	if s2.Listener == nil {
		panic("dnstest.NewServer(): failed to create new server")
	}

	s1.NotifyStartedFunc = func() { close(ch1) }
	s2.NotifyStartedFunc = func() { close(ch2) }
	go s1.ActivateAndServe()
	go s2.ActivateAndServe()

	<-ch1
	<-ch2

	return &Server{s1: s1, s2: s2, Addr: s2.Listener.Addr().String()}
}

// Close shuts down the server.
func (s *Server) Close() {
	s.s1.Shutdown()
	s.s2.Shutdown()
}

// RR parses s as a DNS resource record.
func RR(s string) dns.RR {
	rr, err := dns.NewRR(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse rr %q: %v", s, err))
	}
	return rr
}
