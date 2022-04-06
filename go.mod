module source.monogon.dev

go 1.17

// Kubernetes is not fully consumable as a module, fix that
replace (
	k8s.io/api => k8s.io/api v0.23.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.23.4
	k8s.io/apiserver => k8s.io/apiserver v0.23.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.4
	k8s.io/client-go => k8s.io/client-go v0.23.4
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.4
	k8s.io/code-generator => k8s.io/code-generator v0.23.4
	k8s.io/component-base => k8s.io/component-base v0.23.4
	k8s.io/component-helpers => k8s.io/component-helpers v0.23.4
	k8s.io/controller-manager => k8s.io/controller-manager v0.23.4
	k8s.io/cri-api => k8s.io/cri-api v0.23.4
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.4
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.4
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.4
	k8s.io/kubectl => k8s.io/kubectl v0.23.4
	k8s.io/kubelet => k8s.io/kubelet v0.23.4
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.4
	k8s.io/metrics => k8s.io/metrics v0.23.4
	k8s.io/mount-utils => k8s.io/mount-utils v0.23.4
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.4
)

// Pin down opentelementry
// See https://github.com/open-telemetry/opentelemetry-go/issues/2577
replace (
	go.opentelemetry.io/contrib => go.opentelemetry.io/contrib v0.20.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.20.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp => go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.20.0
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/otlp => go.opentelemetry.io/otel/exporters/otlp v0.20.0
	go.opentelemetry.io/otel/metric => go.opentelemetry.io/otel/metric v0.20.0
	go.opentelemetry.io/otel/oteltest => go.opentelemetry.io/otel/oteltest v0.20.0
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v0.20.0
	go.opentelemetry.io/otel/sdk/export/metric => go.opentelemetry.io/otel/sdk/export/metric v0.20.0
	go.opentelemetry.io/otel/sdk/metric => go.opentelemetry.io/otel/sdk/metric v0.20.0
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v0.20.0
	go.opentelemetry.io/proto/otlp => go.opentelemetry.io/proto/otlp v0.7.0
)

// Custom pins for semver breakage best resolved by manual resolution
// Breaking change in 1.15.0 at https://github.com/onsi/ginkgo/pull/736
// K8s still uses 1.12
replace github.com/onsi/ginkgo => github.com/onsi/ginkgo v1.14.2

// Our own patches
replace github.com/containerd/ttrpc => github.com/monogon-dev/ttrpc v1.0.2-0.20210119122237-222b428f008e

// Override version for Bazel support
replace github.com/mwitkow/go-proto-validators => github.com/mwitkow/go-proto-validators v0.3.2

require (
	github.com/adrg/xdg v0.4.0
	github.com/bazelbuild/rules_go v0.30.0
	github.com/cavaliergopher/cpio v1.0.1
	github.com/cenkalti/backoff/v4 v4.1.2
	github.com/container-storage-interface/spec v1.5.0
	github.com/containerd/containerd v1.6.1
	github.com/containernetworking/plugins v1.0.1
	github.com/coredns/coredns v1.9.1
	github.com/corverroos/commentwrap v0.0.0-20191204065359-2926638be44c
	github.com/diskfs/go-diskfs v1.2.0
	github.com/go-delve/delve v1.8.2
	github.com/golang/protobuf v1.5.2
	github.com/google/certificate-transparency-go v1.1.2
	github.com/google/go-cmp v0.5.7
	github.com/google/go-tpm v0.3.3
	github.com/google/go-tpm-tools v0.3.5
	github.com/google/gopacket v1.1.19
	github.com/google/nftables v0.0.0-20220221214239-211824995dcb
	github.com/google/uuid v1.3.0
	github.com/insomniacslk/dhcp v0.0.0-20220119180841-3c283ff8b7dd
	github.com/joho/godotenv v1.4.0
	github.com/mattn/go-shellwords v1.0.12
	github.com/mdlayher/raw v0.1.0
	github.com/opencontainers/runc v1.1.0
	github.com/pierrec/lz4/v4 v4.1.14
	github.com/pkg/errors v0.9.1
	github.com/rekby/gpt v0.0.0-20200614112001-7da10aec5566
	github.com/sbezverk/nfproxy v0.0.0-20210112155058-0d98b4a69f0c
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/vishvananda/netlink v1.1.1-0.20210330154013-f5de75959ad5
	github.com/yalue/native_endian v1.0.2
	go.etcd.io/etcd/api/v3 v3.5.2
	go.etcd.io/etcd/client/pkg/v3 v3.5.2
	go.etcd.io/etcd/client/v3 v3.5.2
	go.etcd.io/etcd/server/v3 v3.5.2
	go.etcd.io/etcd/tests/v3 v3.5.2
	go.uber.org/multierr v1.8.0
	golang.org/x/crypto v0.0.0-20220208050332-20e1d8d225ab
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5
	golang.org/x/text v0.3.7
	golang.org/x/tools v0.1.9
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20220208144051-fde48d68ee68
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
	gvisor.dev/gvisor v0.0.0-20220315202956-f1399ecf1672
	k8s.io/api v0.23.4
	k8s.io/apimachinery v0.23.4
	k8s.io/cli-runtime v0.23.4
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/component-base v0.23.4
	k8s.io/kubectl v0.0.0
	k8s.io/kubelet v0.0.0
	k8s.io/kubernetes v1.23.4
)

require (
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/google/cadvisor v0.44.0 // indirect
	go.starlark.net v0.0.0-20210223155950-e043a3d3c984 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)
