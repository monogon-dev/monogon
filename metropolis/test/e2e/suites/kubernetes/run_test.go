package kubernetes

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	podv1 "k8s.io/kubernetes/pkg/api/v1/pod"

	common "source.monogon.dev/metropolis/node"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
	mlaunch "source.monogon.dev/metropolis/test/launch"
	"source.monogon.dev/metropolis/test/localregistry"
	"source.monogon.dev/metropolis/test/util"
)

var (
	// These are filled by bazel at linking time with the canonical path of
	// their corresponding file. Inside the init function we resolve it
	// with the rules_go runfiles package to the real path.
	xTestImagesManifestPath string
)

func init() {
	var err error
	for _, path := range []*string{
		&xTestImagesManifestPath,
	} {
		*path, err = runfiles.Rlocation(*path)
		if err != nil {
			panic(err)
		}
	}
}

const (
	// Timeout for the global test context.
	//
	// Bazel would eventually time out the test after 900s ("large") if, for
	// some reason, the context cancellation fails to abort it.
	globalTestTimeout = 600 * time.Second

	// Timeouts for individual end-to-end tests of different sizes.
	smallTestTimeout = 60 * time.Second
	largeTestTimeout = 120 * time.Second
)

// TestE2EKubernetesLabels verifies that Kubernetes node labels are being updated
// when the cluster state changes.
func TestE2EKubernetesLabels(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), globalTestTimeout)
	defer cancel()

	clusterOptions := mlaunch.ClusterOptions{
		NumNodes: 2,
		InitialClusterConfiguration: &cpb.ClusterConfiguration{
			ClusterDomain:         "cluster.test",
			TpmMode:               cpb.ClusterConfiguration_TPM_MODE_DISABLED,
			StorageSecurityPolicy: cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE,
			KubernetesConfig: &cpb.ClusterConfiguration_KubernetesConfig{
				NodeLabelsToSynchronize: []*cpb.ClusterConfiguration_KubernetesConfig_NodeLabelsToSynchronize{
					{Regexp: `^test\.monogon\.dev/`},
				},
			},
		},
	}
	cluster, err := mlaunch.LaunchCluster(ctx, clusterOptions)
	if err != nil {
		t.Fatalf("LaunchCluster failed: %v", err)
	}
	defer func() {
		err := cluster.Close()
		if err != nil {
			t.Fatalf("cluster Close failed: %v", err)
		}
	}()

	con, err := cluster.CuratorClient()
	if err != nil {
		t.Fatalf("Could not get curator client: %v", err)
	}
	mgmt := apb.NewManagementClient(con)
	clientSet, err := cluster.GetKubeClientSet()
	if err != nil {
		t.Fatal(err)
	}

	getLabelsForNode := func(nid string) common.Labels {
		node, err := clientSet.CoreV1().Nodes().Get(ctx, nid, metav1.GetOptions{})
		if kerrors.IsNotFound(err) {
			return nil
		}
		if err != nil {
			t.Fatalf("Could not get node %s: %v", nid, err)
			return nil
		}
		return common.Labels(node.Labels).Filter(func(k, v string) bool {
			if strings.HasPrefix(k, "node-role.kubernetes.io/") {
				return true
			}
			if strings.HasPrefix(k, "test.monogon.dev/") {
				return true
			}
			return false
		})
	}

	// Nodes should have no labels at first.
	for _, nid := range cluster.NodeIDs {
		if labels := getLabelsForNode(nid); !labels.Equals(nil) {
			t.Errorf("Node %s should have no labels, has %s", nid, labels)
		}
	}
	// Nominate both nodes to be Kubernetes workers.
	for _, nid := range cluster.NodeIDs {
		yes := true
		_, err := mgmt.UpdateNodeRoles(ctx, &apb.UpdateNodeRolesRequest{
			Node: &apb.UpdateNodeRolesRequest_Id{
				Id: nid,
			},
			KubernetesWorker: &yes,
		})
		if err != nil {
			t.Fatalf("Could not make %s a KubernetesWorker: %v", nid, err)
		}
	}

	util.MustTestEventual(t, "Labels added", ctx, time.Second*5, func(ctx context.Context) error {
		// Nodes should have role labels now.
		for _, nid := range cluster.NodeIDs {
			want := common.Labels{
				"node-role.kubernetes.io/KubernetesWorker": "",
			}
			if nid == cluster.NodeIDs[0] {
				want["node-role.kubernetes.io/KubernetesController"] = ""
				want["node-role.kubernetes.io/ConsensusMember"] = ""
			}
			if labels := getLabelsForNode(nid); !want.Equals(labels) {
				return fmt.Errorf("node %s should have labels %s, has %s", nid, want, labels)
			}
		}
		return nil
	})

	// Remove KubernetesWorker from first node again. It will stay in k8s (arguably,
	// this is a bug) but its role label should be removed.
	no := false
	_, err = mgmt.UpdateNodeRoles(ctx, &apb.UpdateNodeRolesRequest{
		Node: &apb.UpdateNodeRolesRequest_Id{
			Id: cluster.NodeIDs[0],
		},
		KubernetesWorker: &no,
	})
	if err != nil {
		t.Fatalf("Could not remove KubernetesWorker from %s: %v", cluster.NodeIDs[0], err)
	}

	util.MustTestEventual(t, "Labels removed", ctx, time.Second*5, func(ctx context.Context) error {
		for _, nid := range cluster.NodeIDs {
			want := make(common.Labels)
			if nid == cluster.NodeIDs[0] {
				want["node-role.kubernetes.io/KubernetesController"] = ""
				want["node-role.kubernetes.io/ConsensusMember"] = ""
			} else {
				want["node-role.kubernetes.io/KubernetesWorker"] = ""
			}
			if labels := getLabelsForNode(nid); !want.Equals(labels) {
				return fmt.Errorf("node %s should have labels %s, has %s", nid, want, labels)
			}
		}
		return nil
	})

	// Add Metropolis node label, ensure it gets reflected on the Kubernetes node.
	_, err = mgmt.UpdateNodeLabels(ctx, &apb.UpdateNodeLabelsRequest{
		Node: &apb.UpdateNodeLabelsRequest_Id{
			Id: cluster.NodeIDs[1],
		},
		Upsert: []*apb.UpdateNodeLabelsRequest_Pair{
			{Key: "test.monogon.dev/foo", Value: "bar"},
		},
	})

	util.MustTestEventual(t, "Metropolis labels added", ctx, time.Second*5, func(ctx context.Context) error {
		if err != nil {
			t.Fatalf("Could not add label to node: %v", err)
		}
		want := common.Labels{
			"node-role.kubernetes.io/KubernetesWorker": "",
			"test.monogon.dev/foo":                     "bar",
		}
		if labels := getLabelsForNode(cluster.NodeIDs[1]); !want.Equals(labels) {
			return fmt.Errorf("node %s should have labels %s, has %s", cluster.NodeIDs[1], want, labels)
		}
		return nil
	})

	// Reconfigure node label rules.
	_, err = mgmt.ConfigureCluster(ctx, &apb.ConfigureClusterRequest{
		BaseConfig: &cpb.ClusterConfiguration{
			KubernetesConfig: &cpb.ClusterConfiguration_KubernetesConfig{
				NodeLabelsToSynchronize: []*cpb.ClusterConfiguration_KubernetesConfig_NodeLabelsToSynchronize{
					{Regexp: `^test\.monogon\.dev/`},
				},
			},
		},
		NewConfig: &cpb.ClusterConfiguration{
			KubernetesConfig: &cpb.ClusterConfiguration_KubernetesConfig{},
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"kubernetes_config.node_labels_to_synchronize"},
		},
	})
	if err != nil {
		t.Fatalf("Could not update cluster configuration: %v", err)
	}

	ci, err := mgmt.GetClusterInfo(ctx, &apb.GetClusterInfoRequest{})
	if err != nil {
		t.Fatalf("Could not get cluster info")
	}
	// See if the config changed.
	if rules := ci.ClusterConfiguration.KubernetesConfig.NodeLabelsToSynchronize; len(rules) != 0 {
		t.Fatalf("Wanted 0 label rules in config after reconfiguration, have %d: %v", len(rules), rules)
	}
	// TODO: ensure new rules get applied, but that will require watching the cluster
	// config for changes in the labelmaker.
}

// TestE2EKubernetes exercises the Kubernetes functionality of Metropolis.
//
// The tests are performed against an in-memory cluster.
func TestE2EKubernetes(t *testing.T) {
	// Set a global timeout to make sure this terminates
	ctx, cancel := context.WithTimeout(context.Background(), globalTestTimeout)
	defer cancel()

	df, err := os.ReadFile(xTestImagesManifestPath)
	if err != nil {
		t.Fatalf("Reading registry manifest failed: %v", err)
	}
	lr, err := localregistry.FromBazelManifest(df)
	if err != nil {
		t.Fatalf("Creating test image registry failed: %v", err)
	}

	// Launch cluster.
	clusterOptions := mlaunch.ClusterOptions{
		NumNodes:      2,
		LocalRegistry: lr,
		InitialClusterConfiguration: &cpb.ClusterConfiguration{
			ClusterDomain:         "cluster.test",
			TpmMode:               cpb.ClusterConfiguration_TPM_MODE_DISABLED,
			StorageSecurityPolicy: cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE,
		},
	}
	cluster, err := mlaunch.LaunchCluster(ctx, clusterOptions)
	if err != nil {
		t.Fatalf("LaunchCluster failed: %v", err)
	}
	defer func() {
		err := cluster.Close()
		if err != nil {
			t.Fatalf("cluster Close failed: %v", err)
		}
	}()

	clientSet, err := cluster.GetKubeClientSet()
	if err != nil {
		t.Fatal(err)
	}
	util.TestEventual(t, "Add KubernetesWorker roles", ctx, smallTestTimeout, func(ctx context.Context) error {
		// Make everything but the first node into KubernetesWorkers.
		for i := 1; i < clusterOptions.NumNodes; i++ {
			err := cluster.MakeKubernetesWorker(ctx, cluster.NodeIDs[i])
			if err != nil {
				return util.Permanent(fmt.Errorf("MakeKubernetesWorker: %w", err))
			}
		}
		return nil
	})
	util.TestEventual(t, "Node is registered and ready", ctx, largeTestTimeout, func(ctx context.Context) error {
		nodes, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
		if err != nil {
			return err
		}
		if len(nodes.Items) < 1 {
			return errors.New("node not yet registered")
		}
		node := nodes.Items[0]
		for _, cond := range node.Status.Conditions {
			if cond.Type != corev1.NodeReady {
				continue
			}
			if cond.Status != corev1.ConditionTrue {
				return fmt.Errorf("node not ready: %v", cond.Message)
			}
		}
		return nil
	})
	util.TestEventual(t, "Simple deployment", ctx, largeTestTimeout, func(ctx context.Context) error {
		_, err := clientSet.AppsV1().Deployments("default").Create(ctx, makeTestDeploymentSpec("test-deploy-1"), metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "Simple deployment is running", ctx, largeTestTimeout, func(ctx context.Context) error {
		res, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{LabelSelector: "name=test-deploy-1"})
		if err != nil {
			return err
		}
		if len(res.Items) == 0 {
			return errors.New("pod didn't get created")
		}
		pod := res.Items[0]
		if podv1.IsPodAvailable(&pod, 1, metav1.NewTime(time.Now())) {
			return nil
		}
		events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
		if err != nil || len(events.Items) == 0 {
			return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
		} else {
			return fmt.Errorf("pod is not ready: %v", events.Items[0].Message)
		}
	})
	util.TestEventual(t, "Simple deployment with gvisor", ctx, largeTestTimeout, func(ctx context.Context) error {
		deployment := makeTestDeploymentSpec("test-deploy-2")
		gvisorStr := "gvisor"
		deployment.Spec.Template.Spec.RuntimeClassName = &gvisorStr
		_, err := clientSet.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "Simple deployment is running on gvisor", ctx, largeTestTimeout, func(ctx context.Context) error {
		res, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{LabelSelector: "name=test-deploy-2"})
		if err != nil {
			return err
		}
		if len(res.Items) == 0 {
			return errors.New("pod didn't get created")
		}
		pod := res.Items[0]
		if podv1.IsPodAvailable(&pod, 1, metav1.NewTime(time.Now())) {
			return nil
		}
		events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
		if err != nil || len(events.Items) == 0 {
			return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
		} else {
			var errorMsg strings.Builder
			for _, msg := range events.Items {
				errorMsg.WriteString(" | ")
				errorMsg.WriteString(msg.Message)
			}
			return fmt.Errorf("pod is not ready: %s", errorMsg.String())
		}
	})
	util.TestEventual(t, "Simple StatefulSet with PVC", ctx, largeTestTimeout, func(ctx context.Context) error {
		_, err := clientSet.AppsV1().StatefulSets("default").Create(ctx, makeTestStatefulSet("test-statefulset-1", corev1.PersistentVolumeFilesystem), metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "Simple StatefulSet with PVC is running", ctx, largeTestTimeout, func(ctx context.Context) error {
		res, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{LabelSelector: "name=test-statefulset-1"})
		if err != nil {
			return err
		}
		if len(res.Items) == 0 {
			return errors.New("pod didn't get created")
		}
		pod := res.Items[0]
		if podv1.IsPodAvailable(&pod, 1, metav1.NewTime(time.Now())) {
			return nil
		}
		events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
		if err != nil || len(events.Items) == 0 {
			return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
		} else {
			return fmt.Errorf("pod is not ready: %v", events.Items[0].Message)
		}
	})
	util.TestEventual(t, "Simple StatefulSet with Block PVC", ctx, largeTestTimeout, func(ctx context.Context) error {
		_, err := clientSet.AppsV1().StatefulSets("default").Create(ctx, makeTestStatefulSet("test-statefulset-2", corev1.PersistentVolumeBlock), metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "Simple StatefulSet with Block PVC is running", ctx, largeTestTimeout, func(ctx context.Context) error {
		res, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{LabelSelector: "name=test-statefulset-2"})
		if err != nil {
			return err
		}
		if len(res.Items) == 0 {
			return errors.New("pod didn't get created")
		}
		pod := res.Items[0]
		if podv1.IsPodAvailable(&pod, 1, metav1.NewTime(time.Now())) {
			return nil
		}
		events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
		if err != nil || len(events.Items) == 0 {
			return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
		} else {
			return fmt.Errorf("pod is not ready: %v", events.Items[0].Message)
		}
	})
	util.TestEventual(t, "In-cluster self-test job", ctx, smallTestTimeout, func(ctx context.Context) error {
		_, err := clientSet.BatchV1().Jobs("default").Create(ctx, makeSelftestSpec("selftest"), metav1.CreateOptions{})
		return err
	})
	util.TestEventual(t, "In-cluster self-test job passed", ctx, smallTestTimeout, func(ctx context.Context) error {
		res, err := clientSet.BatchV1().Jobs("default").Get(ctx, "selftest", metav1.GetOptions{})
		if err != nil {
			return err
		}
		if res.Status.Failed > 0 {
			pods, err := clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{
				LabelSelector: "job-name=selftest",
			})
			if err != nil {
				return util.Permanent(fmt.Errorf("job failed but failed to find pod: %w", err))
			}
			if len(pods.Items) < 1 {
				return fmt.Errorf("job failed but pod does not exist")
			}
			lines, err := getPodLogLines(ctx, clientSet, pods.Items[0].Name, 1)
			if err != nil {
				return fmt.Errorf("job failed but could not get logs: %w", err)
			}
			if len(lines) > 0 {
				return util.Permanent(fmt.Errorf("job failed, last log line: %s", lines[0]))
			}
			return util.Permanent(fmt.Errorf("job failed, empty log"))
		}
		if res.Status.Succeeded > 0 {
			return nil
		}
		return fmt.Errorf("job still running")
	})
	util.TestEventual(t, "Start NodePort test setup", ctx, smallTestTimeout, func(ctx context.Context) error {
		_, err := clientSet.AppsV1().Deployments("default").Create(ctx, makeHTTPServerDeploymentSpec("nodeport-server"), metav1.CreateOptions{})
		if err != nil && !kerrors.IsAlreadyExists(err) {
			return err
		}
		_, err = clientSet.CoreV1().Services("default").Create(ctx, makeHTTPServerNodePortService("nodeport-server"), metav1.CreateOptions{})
		if err != nil && !kerrors.IsAlreadyExists(err) {
			return err
		}
		return nil
	})
	util.TestEventual(t, "NodePort accessible from all nodes", ctx, smallTestTimeout, func(ctx context.Context) error {
		nodes, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
		if err != nil {
			return err
		}
		// Use a new client for each attempt
		hc := http.Client{
			Timeout: 2 * time.Second,
			Transport: &http.Transport{
				Dial: cluster.SOCKSDialer.Dial,
			},
		}
		for _, n := range nodes.Items {
			var addr string
			for _, a := range n.Status.Addresses {
				if a.Type == corev1.NodeInternalIP {
					addr = a.Address
				}
			}
			u := url.URL{Scheme: "http", Host: addr, Path: "/"}
			res, err := hc.Get(u.String())
			if err != nil {
				return fmt.Errorf("failed getting from node %q: %w", n.Name, err)
			}
			if res.StatusCode != http.StatusOK {
				return fmt.Errorf("getting from node %q: HTTP %d", n.Name, res.StatusCode)
			}
			t.Logf("Got response from %q", n.Name)
		}
		return nil
	})
	util.TestEventual(t, "containerd metrics retrieved", ctx, smallTestTimeout, func(ctx context.Context) error {
		pool := x509.NewCertPool()
		pool.AddCert(cluster.CACertificate)
		cl := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates: []tls.Certificate{cluster.Owner},
					RootCAs:      pool,
				},
				DialContext: func(ctx context.Context, _, addr string) (net.Conn, error) {
					return cluster.DialNode(ctx, addr)
				},
			},
		}
		u := url.URL{
			Scheme: "https",
			Host:   net.JoinHostPort(cluster.NodeIDs[1], common.MetricsPort.PortString()),
			Path:   "/metrics/containerd",
		}
		res, err := cl.Get(u.String())
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return fmt.Errorf("status code %d", res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		needle := "containerd_build_info_total"
		if !strings.Contains(string(body), needle) {
			return util.Permanent(fmt.Errorf("could not find %q in returned response", needle))
		}
		return nil
	})
	if os.Getenv("HAVE_NESTED_KVM") != "" {
		util.TestEventual(t, "Pod for KVM/QEMU smoke test", ctx, smallTestTimeout, func(ctx context.Context) error {
			runcRuntimeClass := "runc"
			_, err := clientSet.CoreV1().Pods("default").Create(ctx, &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "vm-smoketest",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:            "vm-smoketest",
						ImagePullPolicy: corev1.PullNever,
						Image:           "test.monogon.internal/metropolis/vm/smoketest:smoketest_container",
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								"devices.monogon.dev/kvm": *resource.NewQuantity(1, ""),
							},
						},
					}},
					RuntimeClassName: &runcRuntimeClass,
					RestartPolicy:    corev1.RestartPolicyNever,
				},
			}, metav1.CreateOptions{})
			return err
		})
		util.TestEventual(t, "KVM/QEMU smoke test completion", ctx, smallTestTimeout, func(ctx context.Context) error {
			pod, err := clientSet.CoreV1().Pods("default").Get(ctx, "vm-smoketest", metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("failed to get pod: %w", err)
			}
			if pod.Status.Phase == corev1.PodSucceeded {
				return nil
			}
			events, err := clientSet.CoreV1().Events("default").List(ctx, metav1.ListOptions{FieldSelector: fmt.Sprintf("involvedObject.name=%s,involvedObject.namespace=default", pod.Name)})
			if err != nil || len(events.Items) == 0 {
				return fmt.Errorf("pod is not ready: %v", pod.Status.Phase)
			} else {
				return fmt.Errorf("pod is not ready: %v", events.Items[len(events.Items)-1].Message)
			}
		})
	}
}
