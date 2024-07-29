module source.monogon.dev

go 1.22.0

// Kubernetes is not fully consumable as a module, fix that
replace (
	k8s.io/api => k8s.io/api v0.30.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.30.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.30.2
	k8s.io/apiserver => k8s.io/apiserver v0.30.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.30.2
	k8s.io/client-go => k8s.io/client-go v0.30.2
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.30.2
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.30.2
	k8s.io/code-generator => k8s.io/code-generator v0.30.2
	k8s.io/component-base => k8s.io/component-base v0.30.2
	k8s.io/component-helpers => k8s.io/component-helpers v0.30.2
	k8s.io/controller-manager => k8s.io/controller-manager v0.30.2
	k8s.io/cri-api => k8s.io/cri-api v0.30.2
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.30.2
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.30.2
	k8s.io/endpointslice => k8s.io/endpointslice v0.30.2
	k8s.io/kms => k8s.io/kms v0.30.2
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.30.2
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.30.2
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.30.2
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.30.2
	k8s.io/kubectl => k8s.io/kubectl v0.30.2
	k8s.io/kubelet => k8s.io/kubelet v0.30.2
	k8s.io/kubernetes => k8s.io/kubernetes v1.30.2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.30.2
	k8s.io/metrics => k8s.io/metrics v0.30.2
	k8s.io/mount-utils => k8s.io/mount-utils v0.30.2
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.30.2
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.30.2
)

// Override version for Bazel support
replace github.com/mwitkow/go-proto-validators => github.com/mwitkow/go-proto-validators v0.3.2

// bazeldnf currently comes with a go-rpmutils patch
replace github.com/sassoftware/go-rpmutils v0.1.1 => github.com/rmohr/go-rpmutils v0.1.2-0.20201215123907-5acf7436c00d

// Our psample patches
replace github.com/vishvananda/netlink => github.com/monogon-dev/netlink v0.0.0-20230125113930-88977c3ff4b3

// Pin buildtools version to an up to date version, as the bazeldnf version
// is outdated and gazelle needs it.
replace github.com/bazelbuild/buildtools => github.com/bazelbuild/buildtools v0.0.0-20231103205921-433ea8554e82

// Our privflags implementation, going upstream with https://github.com/mdlayher/ethtool/pull/22
replace github.com/mdlayher/ethtool => github.com/monogon-dev/ethtool v0.0.0-20231122193313-e9c21a3a83cb

// Fixes https://github.com/prometheus/node_exporter/issues/2849
replace github.com/jsimonetti/rtnetlink => github.com/jsimonetti/rtnetlink v1.4.0

// Upgrade to fix missing constant in x/sys v0.14.0
// https://github.com/cilium/ebpf/releases/tag/v0.12.3
replace github.com/cilium/ebpf => github.com/cilium/ebpf v0.12.3

// Update to the latest version before io_k8s_apiextensions_apiserver isn't
// compatible anymore because of a breaking change.
replace github.com/google/cel-go => github.com/google/cel-go v0.17.8

// Update to the latest version to prevent additional imports
// to appear in our dependency graph: https://github.com/golang/go/issues/37175
replace golang.org/x/exp => golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa

// Replace with our patched library to support hardware listings for a whole
// organization at once.
replace github.com/packethost/packngo => github.com/monogon-dev/packngo v0.0.0-20240122175436-ecbd9eb00ddb

require (
	4d63.com/gocheckcompilerdirectives v1.2.1
	cloud.google.com/go/storage v1.36.0
	github.com/adrg/xdg v0.4.0
	github.com/bazelbuild/rules_go v0.49.0
	github.com/cavaliergopher/cpio v1.0.1
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/cockroachdb/cockroach-go/v2 v2.2.10
	github.com/container-storage-interface/spec v1.8.0
	github.com/containerd/containerd v1.7.15
	github.com/containernetworking/plugins v1.2.0
	github.com/corverroos/commentwrap v0.0.0-20191204065359-2926638be44c
	github.com/diskfs/go-diskfs v1.2.0
	github.com/docker/distribution v2.8.2+incompatible
	github.com/go-delve/delve v1.8.2
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/google/cel-go v0.18.1
	github.com/google/certificate-transparency-go v1.1.2
	github.com/google/go-cmp v0.6.0
	github.com/google/go-tpm v0.3.3
	github.com/google/go-tpm-tools v0.3.5
	github.com/google/gopacket v1.1.19
	github.com/google/nftables v0.0.0-20220221214239-211824995dcb
	github.com/google/uuid v1.6.0
	github.com/iancoleman/strcase v0.3.0
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/insomniacslk/dhcp v0.0.0-20231016090811-6a2c8fbdcc1c
	github.com/joho/godotenv v1.4.0
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/kevinburke/go-bindata v3.23.0+incompatible
	github.com/lib/pq v1.10.9
	github.com/mattn/go-shellwords v1.0.12
	github.com/mdlayher/arp v0.0.0-20220512170110-6706a2966875
	github.com/mdlayher/ethernet v0.0.0-20220221185849-529eae5b6118
	github.com/mdlayher/ethtool v0.1.0
	github.com/mdlayher/genetlink v1.3.2
	github.com/mdlayher/kobject v0.0.0-20200520190114-19ca17470d7d
	github.com/mdlayher/netlink v1.7.2
	github.com/mdlayher/packet v1.1.2
	github.com/miekg/dns v1.1.58
	github.com/mitchellh/go-wordwrap v1.0.1
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/runc v1.1.12
	github.com/packethost/packngo v0.29.0
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.13.1
	github.com/prometheus/client_golang v1.17.0
	github.com/prometheus/node_exporter v1.7.0
	github.com/pseudomuto/protoc-gen-doc v1.5.0
	github.com/rmohr/bazeldnf v0.5.4
	github.com/sbezverk/nfproxy v0.0.0-20210112155058-0d98b4a69f0c
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/sqlc-dev/sqlc v1.23.0
	github.com/stretchr/testify v1.8.4
	github.com/vishvananda/netlink v1.2.1-beta.2
	github.com/yalue/native_endian v1.0.2
	go.etcd.io/etcd/api/v3 v3.5.13
	go.etcd.io/etcd/client/pkg/v3 v3.5.13
	go.etcd.io/etcd/client/v3 v3.5.13
	go.etcd.io/etcd/server/v3 v3.5.13
	go.etcd.io/etcd/tests/v3 v3.5.13
	go.starlark.net v0.0.0-20230525235612-a134d8f9ddca
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.26.0
	golang.org/x/crypto v0.21.0
	golang.org/x/mod v0.16.0
	golang.org/x/net v0.23.0
	golang.org/x/sync v0.6.0
	golang.org/x/sys v0.18.0
	golang.org/x/text v0.14.0
	golang.org/x/time v0.5.0
	golang.org/x/tools v0.18.0
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20220208144051-fde48d68ee68
	google.golang.org/genproto/googleapis/api v0.0.0-20240415180920-8c6c420018be
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
	gvisor.dev/gvisor v0.0.0-20230911190645-2e1d76499fd5
	honnef.co/go/tools v0.4.2
	k8s.io/api v0.30.2
	k8s.io/apimachinery v0.30.2
	k8s.io/apiserver v0.30.2
	k8s.io/cli-runtime v0.30.2
	k8s.io/client-go v0.30.2
	k8s.io/code-generator v0.30.2
	k8s.io/component-base v0.30.2
	k8s.io/klog/v2 v2.120.1
	k8s.io/kubectl v0.0.0
	k8s.io/kubelet v0.30.2
	k8s.io/kubernetes v1.20.1
	k8s.io/pod-security-admission v0.0.0
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b
)

require (
	cloud.google.com/go v0.112.0 // indirect
	cloud.google.com/go/compute v1.24.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.6 // indirect
	dario.cat/mergo v1.0.0 // indirect
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20230811130428-ced1acdcaa24 // indirect
	github.com/AdamKorcz/go-118-fuzz-build v0.0.0-20230306123547-8075edf89bb0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/GoogleCloudPlatform/k8s-cloud-provider v1.18.1-0.20220218231025-f11817397a1b // indirect
	github.com/JeffAshton/win_pdh v0.0.0-20161109143554-76bb4ee9f0ab // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/Microsoft/hcsshim v0.11.4 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/alecthomas/kingpin/v2 v2.4.0 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/alexflint/go-filemutex v1.2.0 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230321174746-8dcc6526cfb1 // indirect
	github.com/armon/circbuf v0.0.0-20150827004946-bbbad097214e // indirect
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/bazelbuild/buildtools v0.0.0-20201023142455-8a8e1e724705 // indirect
	github.com/beevik/ntp v1.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/bytecodealliance/wasmtime-go/v14 v14.0.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/checkpoint-restore/go-criu/v5 v5.3.0 // indirect
	github.com/cilium/ebpf v0.12.3 // indirect
	github.com/containerd/aufs v1.0.0 // indirect
	github.com/containerd/btrfs/v2 v2.0.0 // indirect
	github.com/containerd/cgroups v1.1.0 // indirect
	github.com/containerd/cgroups/v3 v3.0.2 // indirect
	github.com/containerd/console v1.0.4 // indirect
	github.com/containerd/continuity v0.4.2 // indirect
	github.com/containerd/fifo v1.1.0 // indirect
	github.com/containerd/go-cni v1.1.9 // indirect
	github.com/containerd/go-runc v1.0.0 // indirect
	github.com/containerd/imgcrypt v1.1.7 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/nri v0.6.0 // indirect
	github.com/containerd/ttrpc v1.2.3 // indirect
	github.com/containerd/typeurl v1.0.2 // indirect
	github.com/containerd/typeurl/v2 v2.1.1 // indirect
	github.com/containerd/zfs v1.1.0 // indirect
	github.com/containernetworking/cni v1.1.2 // indirect
	github.com/containers/ocicrypt v1.1.6 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/coreos/go-oidc v2.2.1+incompatible // indirect
	github.com/coreos/go-semver v0.3.1
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cosiner/argv v0.1.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/creack/pty v1.1.18 // indirect
	github.com/crillab/gophersat v1.3.1 // indirect
	github.com/cubicdaiya/gonp v1.0.4 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/daviddengcn/go-colortext v1.0.0 // indirect
	github.com/dennwc/btrfs v0.0.0-20230312211831-a1f570bd01a1 // indirect
	github.com/dennwc/ioctl v1.0.0 // indirect
	github.com/derekparker/trie v0.0.0-20200317170641-1fdf38b7b0e9 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/ema/qdisc v1.0.0 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/fvbommel/sortorder v1.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-delve/liner v1.2.2-1 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.7.0-rc.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/cadvisor v0.49.0 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-dap v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20230509042627-b1315fad0c5a // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/subcommands v1.0.2-0.20190508160503-636abe8753b8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-envparse v0.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hodgesds/perf-utils v0.7.0 // indirect
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/illumos/go-kstat v0.0.0-20210513183136-173c9b0a9973 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/intel/goresctrl v0.3.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/jsimonetti/rtnetlink v1.3.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/klauspost/compress v1.17.2
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/libopenstorage/openstorage v1.0.0 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/lufia/iostat v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/mattn/go-xmlrpc v0.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/mdlayher/socket v0.5.0 // indirect
	github.com/mdlayher/wifi v0.1.0 // indirect
	github.com/miekg/pkcs11 v1.1.1 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/mistifyio/go-zfs/v3 v3.0.1 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/sys/signal v0.7.0 // indirect
	github.com/moby/sys/symlink v0.2.0 // indirect
	github.com/moby/sys/user v0.1.0 // indirect
	github.com/moby/term v0.0.0-20221205130635-1aeaba878587 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/mrunalp/fileutils v0.5.1 // indirect
	github.com/muesli/reflow v0.0.0-20191128061954-86f094cbed14 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/mwitkow/go-proto-validators v0.2.0 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/opencontainers/image-spec v1.1.0-rc2.0.20221005185240-3a7f492d3f1b // indirect
	github.com/opencontainers/runtime-spec v1.1.0 // indirect
	github.com/opencontainers/runtime-tools v0.9.1-0.20221107090550-2e043c6bd626 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pganalyze/pg_query_go/v4 v4.2.3 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pingcap/errors v0.11.5-0.20210425183316-da1aaba5fb63 // indirect
	github.com/pingcap/failpoint v0.0.0-20220801062533-2eaa32854a6c // indirect
	github.com/pingcap/log v1.1.0 // indirect
	github.com/pingcap/tidb/parser v0.0.0-20231010133155-38cb4f3312be // indirect
	github.com/pkg/xattr v0.4.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/prometheus-community/go-runit v0.1.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/exporter-toolkit v0.10.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/pseudomuto/protokit v0.2.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/riza-io/grpc-go v0.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rs/cors v1.8.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/safchain/ethtool v0.3.0 // indirect
	github.com/sassoftware/go-rpmutils v0.1.1 // indirect
	github.com/sbezverk/nftableslib v0.0.0-20210111145735-b08b2d804e1f // indirect
	github.com/seccomp/libseccomp-golang v0.10.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/stefanberger/go-pkcs11uri v0.0.0-20201008174630-78d3cae3a980 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/tchap/go-patricia/v2 v2.3.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75 // indirect
	github.com/u-root/uio v0.0.0-20230220225925-ffce2a382923 // indirect
	github.com/ulikunitz/xz v0.5.7 // indirect
	github.com/urfave/cli v1.22.12 // indirect
	github.com/vishvananda/netns v0.0.4 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/xlab/treeprint v1.2.0 // indirect
	go.etcd.io/bbolt v1.3.9 // indirect
	go.etcd.io/etcd/client/v2 v2.305.13 // indirect
	go.etcd.io/etcd/pkg/v3 v3.5.13 // indirect
	go.etcd.io/etcd/raft/v3 v3.5.13 // indirect
	go.mozilla.org/pkcs7 v0.0.0-20200128120323-432b2356ecb1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful v0.42.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.47.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.47.0 // indirect
	go.opentelemetry.io/otel v1.22.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.19.0 // indirect
	go.opentelemetry.io/otel/metric v1.22.0 // indirect
	go.opentelemetry.io/otel/sdk v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.22.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/exp/typeparams v0.0.0-20221208152030-732eee02a75a // indirect
	golang.org/x/oauth2 v0.17.0 // indirect
	golang.org/x/term v0.18.0 // indirect
	golang.zx2c4.com/wireguard v0.0.0-20220202223031-3b95c81cc178 // indirect
	google.golang.org/api v0.162.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	gopkg.in/djherbis/times.v1 v1.2.0 // indirect
	gopkg.in/gcfg.v1 v1.2.3 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	howett.net/plist v1.0.0 // indirect
	k8s.io/apiextensions-apiserver v0.0.0 // indirect
	k8s.io/cloud-provider v0.30.2 // indirect
	k8s.io/cluster-bootstrap v0.0.0 // indirect
	k8s.io/component-helpers v0.30.2 // indirect
	k8s.io/controller-manager v0.30.2 // indirect
	k8s.io/cri-api v0.30.2 // indirect
	k8s.io/csi-translation-lib v0.0.0 // indirect
	k8s.io/dynamic-resource-allocation v0.0.0 // indirect
	k8s.io/endpointslice v0.0.0 // indirect
	k8s.io/gengo/v2 v2.0.0-20240228010128-51d4e06bde70 // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kms v0.30.2 // indirect
	k8s.io/kube-aggregator v0.0.0 // indirect
	k8s.io/kube-controller-manager v0.0.0 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	k8s.io/kube-scheduler v0.0.0 // indirect
	k8s.io/legacy-cloud-providers v0.0.0 // indirect
	k8s.io/metrics v0.30.2 // indirect
	k8s.io/mount-utils v0.0.0 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.29.0 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/kustomize/api v0.13.5-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/kustomize/kustomize/v5 v5.0.4-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/kustomize/kyaml v0.14.3-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
	tags.cncf.io/container-device-interface v0.6.2 // indirect
	tags.cncf.io/container-device-interface/specs-go v0.6.0 // indirect
)
