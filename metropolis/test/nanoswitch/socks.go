package main

import (
	"context"
	"fmt"
	"net"

	"source.monogon.dev/metropolis/pkg/socksproxy"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// SOCKSPort is the port at which nanoswitch listens for SOCKS conenctions.
//
// ONCHANGE(//metropolis/test/launch/cluster:cluster.go): port must be kept in sync
const SOCKSPort uint16 = 1080

// socksHandler implements a socksproxy.Handler which permits and logs
// connections to the nanoswitch network.
type socksHandler struct{}

func (s *socksHandler) Connect(ctx context.Context, req *socksproxy.ConnectRequest) *socksproxy.ConnectResponse {
	logger := supervisor.Logger(ctx)
	target := net.JoinHostPort(req.Address.String(), fmt.Sprintf("%d", req.Port))

	if len(req.Address) != 4 {
		logger.Warningf("Connect %s: wrong address type", target)
		return &socksproxy.ConnectResponse{
			Error: socksproxy.ReplyAddressTypeNotSupported,
		}
	}

	addr := req.Address
	switchCIDR := net.IPNet{
		IP:   switchIP.Mask(switchSubnetMask),
		Mask: switchSubnetMask,
	}
	if !switchCIDR.Contains(addr) || switchCIDR.IP.Equal(addr) {
		logger.Warningf("Connect %s: not in switch network", target)
		return &socksproxy.ConnectResponse{
			Error: socksproxy.ReplyNetworkUnreachable,
		}
	}

	con, err := net.Dial("tcp", target)
	if err != nil {
		logger.Warningf("Connect %s: dial failed: %v", target, err)
		return &socksproxy.ConnectResponse{
			Error: socksproxy.ReplyHostUnreachable,
		}
	}
	res, err := socksproxy.ConnectResponseFromConn(con)
	if err != nil {
		logger.Warningf("Connect %s: could not make SOCKS response: %v", target, err)
		con.Close()
		return &socksproxy.ConnectResponse{
			Error: socksproxy.ReplyGeneralFailure,
		}
	}
	logger.Infof("Connect %s: established", target)
	return res
}

// runSOCKSProxy starts a SOCKS proxy to the nanoswitchnetwork at SOCKSPort.
func runSOCKSProxy(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", SOCKSPort))
	if err != nil {
		return fmt.Errorf("failed to listen on :%d : %v", SOCKSPort, err)
	}

	h := &socksHandler{}
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	return socksproxy.Serve(ctx, h, lis)
}
