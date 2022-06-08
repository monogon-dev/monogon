package consensus

import (
	"context"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"

	"source.monogon.dev/metropolis/pkg/event/memory"
)

type testServiceHandle struct {
	s memory.Value
}

// TestServiceHandle builds a somewhat functioning ServiceHandle from a bare
// etcd connection, effectively creating a fake Consensus service. This must
// only be used in test code to perform dependency injection of a etcd client
// into code which expects a Consensus service instance, eg. for testing the
// Curator.
//
// The 'somewhat functioning' description above should serve as a hint to the
// API stability and backwards/forwards compatibility of this function: there is
// none.
func TestServiceHandle(t *testing.T, cl *clientv3.Client) ServiceHandle {
	ca := pkiCA()

	tsh := testServiceHandle{}
	st := &Status{
		cl:                        cl,
		ca:                        ca,
		noClusterMemberManagement: true,
	}
	etcdPKI, err := st.pkiClient()
	if err != nil {
		t.Fatalf("failed to get PKI etcd client: %v", err)
	}
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()
	if _, err := ca.Ensure(ctx, etcdPKI); err != nil {
		t.Fatalf("failed to ensure PKI CA: %v", err)
	}
	tsh.s.Set(st)
	return &tsh
}

func (h *testServiceHandle) Watch() Watcher {
	return Watcher{h.s.Watch()}
}
