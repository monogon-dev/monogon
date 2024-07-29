// Package proxy implements a forwarding proxy. It caches an upstream net.Conn
// for some time, so if the same client returns the upstream's Conn will be
// precached. Depending on how you benchmark this looks to be 50% faster than
// just opening a new connection for every client.
// It works with UDP and TCP and uses inband healthchecking.
package proxy

// Taken and modified from CoreDNS, under Apache 2.0.

import (
	"errors"
	"io"
	"strings"
	"sync/atomic"
	"time"

	"github.com/miekg/dns"
)

// AdvertiseUDPSize is the maximum message size that we advertise in the OPT RR
// of UDP messages. This is calculated as the minimum IPv6 MTU (1280) minus
// size of IPv6 (40) and UDP (8) headers.
const AdvertiseUDPSize = 1232

// ErrCachedClosed means cached connection was closed by peer.
var ErrCachedClosed = errors.New("cached connection was closed by peer")

// limitTimeout is a utility function to auto-tune timeout values.
// Average observed time is moved towards the last observed delay moderated by
// a weight next timeout to use will be the double of the computed average,
// limited by min and max frame.
func limitTimeout(currentAvg *int64, minValue time.Duration, maxValue time.Duration) time.Duration {
	rt := time.Duration(atomic.LoadInt64(currentAvg))
	if rt < minValue {
		return minValue
	}
	if rt < maxValue/2 {
		return 2 * rt
	}
	return maxValue
}

func averageTimeout(currentAvg *int64, observedDuration time.Duration, weight int64) {
	dt := time.Duration(atomic.LoadInt64(currentAvg))
	atomic.AddInt64(currentAvg, int64(observedDuration-dt)/weight)
}

func (t *Transport) dialTimeout() time.Duration {
	return limitTimeout(&t.avgDialTime, minDialTimeout, maxDialTimeout)
}

func (t *Transport) updateDialTimeout(newDialTime time.Duration) {
	averageTimeout(&t.avgDialTime, newDialTime, cumulativeAvgWeight)
}

// Dial dials the address configured in transport,
// potentially reusing a connection or creating a new one.
func (t *Transport) Dial(proto string) (*persistConn, bool, error) {
	// If tls has been configured; use it.
	if t.tlsConfig != nil {
		proto = "tcp-tls"
	}

	t.dial <- proto
	pc := <-t.ret

	if pc != nil {
		return pc, true, nil
	}

	reqTime := time.Now()
	timeout := t.dialTimeout()
	if proto == "tcp-tls" {
		conn, err := dns.DialTimeoutWithTLS("tcp", t.addr, t.tlsConfig, timeout)
		t.updateDialTimeout(time.Since(reqTime))
		return &persistConn{c: conn}, false, err
	}
	conn, err := dns.DialTimeout(proto, t.addr, timeout)
	t.updateDialTimeout(time.Since(reqTime))
	return &persistConn{c: conn}, false, err
}

// Connect selects an upstream, sends the request and waits for a response.
func (p *Proxy) Connect(m *dns.Msg, useTCP bool) (*dns.Msg, error) {
	proto := "udp"
	if useTCP {
		proto = "tcp"
	}

	pc, cached, err := p.transport.Dial(proto)
	if err != nil {
		return nil, err
	}

	pc.c.UDPSize = AdvertiseUDPSize

	pc.c.SetWriteDeadline(time.Now().Add(p.writeTimeout))
	m.Id = dns.Id()

	if err := pc.c.WriteMsg(m); err != nil {
		pc.c.Close() // not giving it back
		if err == io.EOF && cached {
			return nil, ErrCachedClosed
		}
		return nil, err
	}

	var ret *dns.Msg
	pc.c.SetReadDeadline(time.Now().Add(p.readTimeout))
	for {
		ret, err = pc.c.ReadMsg()
		if err != nil {
			if ret != nil && (m.Id == ret.Id) && p.transport.transportTypeFromConn(pc) == typeUDP && shouldTruncateResponse(err) {
				// For UDP, if the error is an overflow, we probably have an upstream
				// misbehaving in some way.
				// (e.g. sending >512 byte responses without an eDNS0 OPT RR).
				// Instead of returning an error, return an empty response
				// with TC bit set. This will make the client retry over TCP
				// (if that's supported) or at least receive a clean error.
				// The connection is still good so we break before the close.

				// Truncate the response.
				ret = truncateResponse(ret)
				break
			}

			pc.c.Close() // not giving it back
			if err == io.EOF && cached {
				return nil, ErrCachedClosed
			}
			return ret, err
		}
		// drop out-of-order responses
		if m.Id == ret.Id {
			break
		}
	}

	p.transport.Yield(pc)

	return ret, nil
}

const cumulativeAvgWeight = 4

// Function to determine if a response should be truncated.
func shouldTruncateResponse(err error) bool {
	// This is to handle a scenario in which upstream sets the TC bit,
	// but doesn't truncate the response and we get ErrBuf instead of overflow.
	if errors.Is(err, dns.ErrBuf) {
		return true
	} else if strings.Contains(err.Error(), "overflow") {
		return true
	}
	return false
}

// Function to return an empty response with TC (truncated) bit set.
func truncateResponse(response *dns.Msg) *dns.Msg {
	// Clear out Answer, Extra, and Ns sections
	response.Answer = nil
	response.Extra = nil
	response.Ns = nil

	// Set TC bit to indicate truncation.
	response.Truncated = true
	return response
}
