// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package connectivity

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/netip"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/protobuf/encoding/protodelim"
	"google.golang.org/protobuf/types/known/durationpb"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/utils/ptr"

	"source.monogon.dev/metropolis/test/e2e/connectivity/spec"
)

type podState struct {
	name      string
	namespace string
	ipv4      netip.Addr
	ipv6      netip.Addr
	reqMu     sync.Mutex
	in        io.Writer
	out       *bufio.Reader
}

func (p *podState) do(req *spec.Request) (*spec.Response, error) {
	// This could be made much faster by introducing a request/response routing
	// layer which would allow concurrent requests and send the response to the
	// correct request handler. For simplicity this hasn't been done.
	p.reqMu.Lock()
	defer p.reqMu.Unlock()
	if _, err := protodelim.MarshalTo(p.in, req); err != nil {
		return nil, fmt.Errorf("while sending request to pod: %w", err)
	}
	var res spec.Response
	if err := protodelim.UnmarshalFrom(p.out, &res); err != nil {
		return nil, fmt.Errorf("while reading response from pod: %w", err)
	}
	return &res, nil
}

type Tester struct {
	cancel    context.CancelFunc
	clientSet kubernetes.Interface
	pods      []podState
	lastToken atomic.Uint64
}

// Convenience aliases
const (
	ExpectedSuccess = spec.TestResponse_SUCCESS
	ExpectedTimeout = spec.TestResponse_CONNECTION_TIMEOUT
	ExpectedReject  = spec.TestResponse_CONNECTION_REJECTED
)

func (te *Tester) GetToken() uint64 {
	return te.lastToken.Add(1)
}

func (te *Tester) TestPodConnectivity(t *testing.T, srcPod, dstPod int, port int, expected spec.TestResponse_Result) {
	te.TestPodConnectivityEventual(t, srcPod, dstPod, port, expected, 0)
}

func (te *Tester) TestPodConnectivityEventual(t *testing.T, srcPod, dstPod int, port int, expected spec.TestResponse_Result, timeout time.Duration) {
	token := te.GetToken()
	prefix := fmt.Sprintf("%d -> %d :%d", srcPod, dstPod, port)
	if err := te.startServer(dstPod, &spec.StartServerRequest{
		Address: net.JoinHostPort("0.0.0.0", strconv.Itoa(port)),
		Token:   token,
	}); err != nil {
		t.Fatalf("%v: %v", prefix, err)
		return
	}
	defer func() {
		if err := te.stopServer(dstPod, &spec.StopServerRequest{Token: token}); err != nil {
			t.Log(err)
		}
	}()
	deadline := time.Now().Add(timeout)
	for {
		res, err := te.runTest(srcPod, &spec.TestRequest{
			Token:   token,
			Timeout: durationpb.New(5 * time.Second),
			Address: net.JoinHostPort(te.pods[dstPod].ipv4.String(), strconv.Itoa(port)),
		})
		if err != nil {
			// Unknown errors do not get retried
			t.Fatalf("%v: %v", prefix, err)
			return
		}
		err = checkExpectations(t, res, expected, prefix)
		if err == nil {
			return
		} else if deadline.Before(time.Now()) {
			t.Error(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func checkExpectations(t *testing.T, res *spec.TestResponse, expected spec.TestResponse_Result, prefix string) error {
	switch res.Result {
	case spec.TestResponse_CONNECTION_REJECTED, spec.TestResponse_CONNECTION_TIMEOUT:
		if expected != res.Result {
			return fmt.Errorf("%v expected %v, got %v (%v)", prefix, expected, res.Result, res.ErrorDescription)
		}
	case spec.TestResponse_WRONG_TOKEN:
		return fmt.Errorf("%v connected, but got wrong token", prefix)
	case spec.TestResponse_SUCCESS:
		if expected != ExpectedSuccess {
			return fmt.Errorf("%v expected %v, got %v", prefix, expected, res.Result)
		}
	}
	return nil
}

func (te *Tester) startServer(pod int, req *spec.StartServerRequest) error {
	res, err := te.pods[pod].do(&spec.Request{Req: &spec.Request_StartServer{StartServer: req}})
	if err != nil {
		return fmt.Errorf("test pod communication failure: %w", err)
	}
	startRes := res.GetStartServer()
	if !startRes.Ok {
		return fmt.Errorf("test server could not be started: %v", startRes.ErrorDescription)
	}
	return nil
}

func (te *Tester) stopServer(pod int, req *spec.StopServerRequest) error {
	res, err := te.pods[pod].do(&spec.Request{Req: &spec.Request_StopServer{StopServer: req}})
	if err != nil {
		return fmt.Errorf("test pod communication failure: %w", err)
	}
	stopRes := res.GetStopServer()
	if !stopRes.Ok {
		return fmt.Errorf("test server could not be stopped: %v", stopRes.ErrorDescription)
	}
	return nil
}

func (te *Tester) runTest(pod int, req *spec.TestRequest) (*spec.TestResponse, error) {
	res, err := te.pods[pod].do(&spec.Request{Req: &spec.Request_Test{Test: req}})
	if err != nil {
		return nil, fmt.Errorf("test pod communication failure: %w", err)
	}
	testRes := res.GetTest()
	if testRes.Result == spec.TestResponse_UNKNOWN {
		return nil, fmt.Errorf("test encountered unknown error: %v", testRes.ErrorDescription)
	}
	return testRes, nil
}

type TestSpec struct {
	// Name needs to contain a DNS label-compatible unique name which
	// is used to identify the Kubernetes resources for this test.
	Name string
	// ClientSet needs to be a client for the test cluster.
	ClientSet kubernetes.Interface
	// RESTConfig needs to be the client config for the same test cluster.
	RESTConfig *rest.Config
	// Number of pods to start for testing. Their identities go from 0 to
	// NumPods-1.
	NumPods int
	// ExtraPodConfig is called for every pod to be created and can be used to
	// customize their specification.
	ExtraPodConfig func(i int, pod *corev1.Pod)
}

// SetupTest sets up the K8s resources and communication channels for
// connectivity tests. It registers a cleanup function which automatically
// tears them down again after the test.
func SetupTest(t *testing.T, s *TestSpec) *Tester {
	t.Helper()

	testCtx, testCancel := context.WithCancel(context.Background())

	tester := Tester{
		cancel:    testCancel,
		clientSet: s.ClientSet,
		pods:      make([]podState, s.NumPods),
	}

	// Use a non-zero arbitrary start value to decrease the chance of
	// accidential conflicts.
	tester.lastToken.Store(1234)
	wg := sync.WaitGroup{}
	setupCtx, cancel := context.WithTimeout(testCtx, 60*time.Second)
	defer cancel()
	errChan := make(chan error, s.NumPods)
	for i := 0; i < s.NumPods; i++ {
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%v-%d", s.Name, i),
				Namespace: "default",
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{{
					Name:  "connectivitytester",
					Image: "test.monogon.internal/connectivity/agent:latest",
					Stdin: true,
				}},
				EnableServiceLinks: ptr.To(false),
			},
		}
		if s.ExtraPodConfig != nil {
			s.ExtraPodConfig(i, pod)
		}
		wg.Add(1)
		go func() {
			pod, err := backoff.RetryNotifyWithData(func() (*corev1.Pod, error) {
				ctx, cancel := context.WithTimeout(setupCtx, 5*time.Second)
				defer cancel()
				return s.ClientSet.CoreV1().Pods(pod.Namespace).Create(ctx, pod, metav1.CreateOptions{})
			}, backoff.WithContext(backoff.NewConstantBackOff(1*time.Second), setupCtx), func(err error, d time.Duration) {
				t.Logf("attempted creating pod %d: %v", i, err)
			})
			if err != nil {
				errChan <- fmt.Errorf("creating pod %d failed: %w", i, err)
				wg.Done()
				return
			}
			tester.pods[i].name = pod.Name
			tester.pods[i].namespace = pod.Namespace
			// Wait for pods to be ready and populate IPs
			err = backoff.Retry(func() error {
				podW, err := s.ClientSet.CoreV1().Pods(pod.Namespace).Watch(setupCtx, metav1.SingleObject(pod.ObjectMeta))
				if err != nil {
					return err
				}
				defer podW.Stop()
				podUpdates := podW.ResultChan()
				for event := range podUpdates {
					if event.Type != watch.Modified && event.Type != watch.Added {
						continue
					}
					pod = event.Object.(*corev1.Pod)
					if pod.Status.Phase == corev1.PodRunning {
						for _, ipObj := range pod.Status.PodIPs {
							ip, err := netip.ParseAddr(ipObj.IP)
							if err != nil {
								return backoff.Permanent(fmt.Errorf("while parsing IP address: %w", err))
							}
							if ip.Is4() {
								tester.pods[i].ipv4 = ip
							} else {
								tester.pods[i].ipv6 = ip
							}
						}
						return nil
					}
				}
				return fmt.Errorf("pod watcher failed")
			}, backoff.WithContext(backoff.NewConstantBackOff(1*time.Second), setupCtx))
			if err != nil {
				errChan <- fmt.Errorf("waiting for pod %d to be running failed: %w", i, err)
				wg.Done()
				return
			}

			inR, inW := io.Pipe()
			outR, outW := io.Pipe()
			tester.pods[i].in = inW
			tester.pods[i].out = bufio.NewReader(outR)

			req := s.ClientSet.CoreV1().RESTClient().Post().Resource("pods").Namespace(pod.Namespace).Name(pod.Name).SubResource("attach")
			options := &corev1.PodAttachOptions{
				Container: "connectivitytester",
				Stdin:     true,
				Stdout:    true,
				Stderr:    true,
			}
			req.VersionedParams(options, scheme.ParameterCodec)
			attachHandle, err := remotecommand.NewWebSocketExecutor(s.RESTConfig, "POST", req.URL().String())
			if err != nil {
				panic(err)
			}

			wg.Done()
			err = attachHandle.StreamWithContext(testCtx, remotecommand.StreamOptions{
				Stdin:  inR,
				Stdout: outW,
				Stderr: os.Stderr,
			})
			if err != nil && !errors.Is(err, testCtx.Err()) {
				t.Logf("Stream for pod %d failed: %v", i, err)
			}
		}()
	}
	t.Cleanup(func() {
		tester.cancel()
		var cleanupWg sync.WaitGroup
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		for i := range tester.pods {
			p := &tester.pods[i]
			if p.name == "" {
				continue
			}
			cleanupWg.Add(1)
			go func() {
				defer cleanupWg.Done()
				err := backoff.Retry(func() error {
					ctx, cancel := context.WithTimeout(cleanupCtx, 5*time.Second)
					defer cancel()
					err := s.ClientSet.CoreV1().Pods(p.namespace).Delete(ctx, p.name, metav1.DeleteOptions{})
					if kerrors.IsNotFound(err) {
						return nil
					}
					return err
				}, backoff.WithContext(backoff.NewConstantBackOff(1*time.Second), cleanupCtx))
				if err != nil {
					t.Logf("Cleanup of pod %d failed: %v", i, err)
				}
			}()
		}
		cleanupWg.Wait()
	})
	wg.Wait()
	close(errChan)
	// Process asynchronous errors
	for err := range errChan {
		t.Error(err)
	}
	return &tester
}
