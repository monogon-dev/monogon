#!/bin/bash
set -euo pipefail

# gazelle.sh regenerates third_party/go/repositories.bzl. This in turn allows Go code to access deps.
# New deps for Smalltown should be defined here and this scripts should then be rerun (it takes a while).

# TODO(T725): This is a temporary workaround (and proof of concept) before we have an alternative Go implementation with a bit more logic in it than this.
# For some basic documentation, see //third_party/go/README.md.

function gazelle() {
    echo "gazelle update-repos $@"
    scripts/bin/bazel run //:gazelle -- update-repos -to_macro=third_party/go/repositories.bzl%go_repositories "$@" 1>/dev/null 2>&1
}

function gazelleGHRelease() {
    local owner="$1"
    local repo="$2"
    local v="$3"
    echo " resolving ${owner}/${repo}..."
    local commit=$(git ls-remote --tags git://github.com/${owner}/${repo} | awk '{ if ($2 == "refs/tags/'${v}'") { print $1 } }')
    echo " resolved ${owner}/${repo} ${v} to ${commit}"
    gazelle github.com/${owner}/${repo}@${commit}
}

function gazelleK8sStaging() {
    local repo="$1"
    local v="kubernetes-$2"
    shift 2
    local args="$@"

    echo " resolving k8s.io/${repo}..."
    local commit=$(git ls-remote --tags git://github.com/kubernetes/${repo} | awk '{ if ($2 == "refs/tags/'${v}'") { print $1 } }')
    echo " resolved k8s.io/${repo} ${v} to ${commit}"
    gazelle $args k8s.io/${repo}@${commit}
}

# Start with an empty repository list, with a well-known Kubernetes version.
# Also add some other deps that need patches.
cat > third_party/go/repositories.bzl <<EOF
load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "io_k8s_kubernetes",
        importpath = "k8s.io/kubernetes",
        version = "v1.19.0-alpha.2",
        sum = "h1:kTsLVxmg/z3Fexcvu75zzGEHOYQ17jzIJFWhfQQnXDE=",
        build_tags = ["providerless"],
        build_file_proto_mode = "disable",
        patches = [
            "//third_party/go/patches:k8s-kubernetes.patch",
            "//third_party/go/patches:k8s-kubernetes-build.patch",
        ],
        patch_args = ["-p1"],
    )

    go_repository(
        name = "com_github_google_cadvisor",
        importpath = "github.com/google/cadvisor",
        sum = "h1:au7bcM+rjGXLBSfqjofcSONBre8tlIy94jEbp40BCOQ=",
        version = "v0.36.1-0.20200323171535-8af10c683a96",
        patches = [
            "//third_party/go/patches:cadvisor.patch",
            "//third_party/go/patches:cadvisor-build.patch",
       ],
        patch_args = ["-p1"],
    )

    go_repository(
        name = "io_k8s_client_go",
        importpath = "k8s.io/client-go",
        sum = "h1:yqKw4cTUQraZK3fcVCMeSa+lqKwcjZ5wtcOIPnxQno4=",
        version = "v0.19.0-alpha.2",
        patches = [
            "//third_party/go/patches:k8s-client-go.patch",
            "//third_party/go/patches:k8s-client-go-build.patch",
        ],
        patch_args = ["-p1"],
    )

    # patches for pure mode
    go_repository(
        name = "com_github_containernetworking_plugins",
        importpath = "github.com/containernetworking/plugins",
        sum = "h1:5lnwfsAYO+V7yXhysJKy3E1A2Gy9oVut031zfdOzI9w=",
        version = "v0.8.2",
        patches = [
            "//third_party/go/patches:cni-plugins-build.patch",
        ],
        patch_args = ["-p1"],
    )

    # required by containerd, but doesn't folow gazelle name convention
    go_repository(                                                                                     
        name = "com_github_opencontainers_runtime-spec",
        importpath = "github.com/opencontainers/runtime-spec",                      
        sum = "h1:d9F+LNYwMyi3BDN4GzZdaSiq4otb8duVEWyZjeUtOQI=",                                       
        version = "v0.1.2-0.20171211145439-b2d941ef6a78",
    )

    go_repository(
        name = "com_github_google_gvisor",
        importpath = "github.com/google/gvisor",
        sum = "h1:rpcz7X//b7LHYEa8FwGlviAPLkFHz46+RW3ur+kiyhg=",
        version = "v0.0.0-20200325151121-d8c4eff3f77b",
        patches = [
            "//third_party/go/patches:gvisor.patch",
        ],
        patch_args = ["-p1"],
    )
    go_repository(
        name = "com_github_google_gvisor_containerd_shim",
        importpath = "github.com/google/gvisor-containerd-shim",
        sum = "h1:RdBNQHpoQ3ekzfXYIV4+nQJ3a2xLnIHuZJkM40OEtyA=",
        version = "v0.0.4",
        patches = [
            "//third_party/go/patches:gvisor-containerd-shim.patch",
            "//third_party/go/patches:gvisor-containerd-shim-build.patch",
            # Patches below are being upstreamed
            "//third_party/go/patches:gvisor-containerd-shim-nogo.patch",
            "//third_party/go/patches:gvisor-shim-root.patch",
        ],
        patch_args = ["-p1"],
    )

    # containerd, Not an actual release, pinned to commit 8e685f78cf66e2901b2fbed2fdddd64449a74ab9 that has support for the required build tags.
    # Also patched for pure mode and some other issues
    go_repository(
        name = "com_github_containerd_containerd",
        build_file_proto_mode = "disable",
        build_tags = ["no_zfs", "no_aufs", "no_devicemapper", "no_btrfs"],
        importpath = "github.com/containerd/containerd",
        sum = "h1:IeFaEbvx6mQe9K1cXG2K7zynPwge3YUrQlLTyiNiveU=",
        version = "v1.3.1-0.20200218165203-8e685f78cf66",
        patches = [
            "//third_party/go/patches:containerd-build.patch",
        ],
        patch_args = ["-p1"],
    )

EOF

# ONCHANGE: bump //build/print-workspace-status.sh accordingly.
k8sVersion="1.19.0-alpha.2"

# k8s repo-infra, Fairly old verison, but anything newer seems to fail.
gazelle k8s.io/repo-infra@df02ded38f9506e5bbcbf21702034b4fef815f2f


# k8s staging/ dumps into external repos
gazelleK8sStaging component-base $k8sVersion
gazelleK8sStaging cloud-provider $k8sVersion
gazelleK8sStaging kubelet $k8sVersion -build_file_proto_mode=disable
gazelleK8sStaging csi-translation-lib $k8sVersion
gazelleK8sStaging legacy-cloud-providers $k8sVersion
gazelleK8sStaging apiextensions-apiserver $k8sVersion -build_file_proto_mode=disable
gazelleK8sStaging kube-aggregator $k8sVersion -build_file_proto_mode=disable
gazelleK8sStaging metrics $k8sVersion -build_file_proto_mode=disable
gazelleK8sStaging kube-scheduler $k8sVersion
gazelleK8sStaging kube-controller-manager $k8sVersion
gazelleK8sStaging cluster-bootstrap $k8sVersion
gazelleK8sStaging kube-proxy $k8sVersion
gazelleK8sStaging kubectl $k8sVersion
gazelleK8sStaging cli-runtime $k8sVersion
gazelleK8sStaging api $k8sVersion -build_file_proto_mode=disable
gazelleK8sStaging apimachinery $k8sVersion -build_file_proto_mode=disable
gazelleK8sStaging apiserver $k8sVersion -build_file_proto_mode=disable
gazelleK8sStaging cri-api $k8sVersion -build_file_proto_mode=disable
gazelleK8sStaging sample-apiserver $k8sVersion

# other k8s deps
gazelle github.com/spf13/pflag@2e9d26c8c37aae03e3f9d4e90b7116f5accb7cab # v1.0.5
gazelle github.com/spf13/cobra@f2b07da1e2c38d5f12845a4f607e2e1018cbb1f5 # v0.0.5
gazelle github.com/blang/semver@b38d23b8782a487059e8fc8773e9a5b228a77cb6 # v3.5.0
gazelle github.com/coreos/go-systemd@95778dfbb74eb7e4dbaf43bf7d71809650ef8076
gazelle k8s.io/utils@a9aa75ae1b89
#gazelle github.com/google/cadvisor@8af10c683a96f85e8ceecfc2b90931aca0ed7448
gazelle github.com/spf13/afero@588a75ec4f32903aa5e39a2619ba6a4631e28424 # v1.2.2
gazelle github.com/armon/circbuf@bbbad097214e2918d8543d5201d12bfd7bca254d
gazelle github.com/container-storage-interface/spec@314ac542302938640c59b6fb501c635f27015326 # v1.2.0
gazelle github.com/docker/go-connections@3ede32e2033de7505e6500d6c868c2b9ed9f169d # v0.3.0
gazelle github.com/moby/term@672ec06f55cd
gazelle k8s.io/kube-openapi@e1beb1bd0f35
gazelle github.com/go-openapi/spec@2223ab324566e4ace63ab69b9c8fff1b81a40eeb # v0.19.3
gazelleGHRelease evanphx json-patch v4.2.0
gazelle -build_file_proto_mode=disable go.etcd.io/etcd@e694b7bb087538c146e188a29753967d189d202b # v3.4.7
gazelleGHRelease go-openapi jsonpointer v0.19.3
gazelleGHRelease go-openapi swag v0.19.5
gazelleGHRelease go-openapi jsonreference v0.19.3
gazelleGHRelease cyphar filepath-securejoin v0.2.2
gazelleGHRelease mailru easyjson v0.7.0
gazelleGHRelease PuerkitoBio purell v1.1.1
gazelle github.com/PuerkitoBio/urlesc@de5bf2ad4578
gazelle github.com/mxk/go-flowrate@cca7078d478f
gazelle cloud.google.com/go@8c41231e01b2085512d98153bcffb847ff9b4b9f # v0.38.0
gazelleGHRelease karrick godirwalk v1.7.5
gazelleGHRelease euank go-kmsg-parser v2.0.0
gazelle github.com/mindprince/gonvml@9ebdce4bb989
gazelle github.com/checkpoint-restore/go-criu@bdb7599cd87b
gazelle github.com/mrunalp/fileutils@7d4729fb3618
gazelle github.com/vishvananda/netns@0a2b9b5464df
gazelle github.com/munnerz/goautoneg@a7dc8b61c822
gazelle github.com/NYTimes/gziphandler@56545f4a5d46
gazelleGHRelease vishvananda netlink v1.1.0
gazelleGHRelease googleapis gnostic v0.4.1
gazelle go.uber.org/zap@feeb9a050b31b40eec6f2470e7599eeeadfe5bdd # v1.15.0
gazelleGHRelease coreos go-semver v0.3.0
gazelle go.uber.org/multierr@3c4937480c32f4c13a875a1829af76c98ca3d40a # v1.1.0
gazelle go.uber.org/atomic@1ea20fb1cbb1cc08cbd0d913a96dead89aa18289 # v1.3.2
gazelle github.com/coreos/pkg@97fdf19511ea
gazelleGHRelease morikuni aec v1.0.0 
gazelleGHRelease lithammer dedent v1.1.0
gazelleGHRelease grpc-ecosystem go-grpc-prometheus v1.2.0
gazelle gopkg.in/natefinch/lumberjack.v2@dd45e6a67c53f673bb49ca8a001fd3a63ceb640e # v2.0.0
gazelleGHRelease robfig cron v1.1.0
gazelleGHRelease coreos go-oidc v2.1.0
gazelle gonum.org/v1/gonum@402b1e2868774b0eee0ec7c85bb9a7b36cf650ae # v0.6.2
gazelle k8s.io/heapster@eea8c9657d0176b52226706765714c6ff6e432cc # v1.2.0
gazelle gopkg.in/square/go-jose.v2@e94fb177d3668d35ab39c61cbb2f311550557e83 # v2.2.2
gazelle github.com/pquerna/cachecontrol@0dec1b30a021
gazelleGHRelease go-openapi validate v0.19.5 
gazelleGHRelease go-openapi strfmt v0.19.3
gazelle k8s.io/gengo@e0e292d8aa12
gazelle golang.org/x/mod@ed3ec21bb8e252814c380df79a80f366440ddb2d # v0.2.0
gazelle golang.org/x/xerrors@9bdfabe68543c54f90421aeb9a60ef8061b5b544
gazelleGHRelease grpc-ecosystem grpc-gateway v1.9.5 
gazelleGHRelease soheilhy cmux v0.1.4
gazelle github.com/tmc/grpc-websocket-proxy@89b8d40f7ca8
gazelleGHRelease go-openapi errors v0.19.2
gazelleGHRelease go-openapi analysis v0.19.5
gazelleGHRelease go-openapi loads v0.19.4
gazelleGHRelease dustin go-humanize v1.0.0
gazelleGHRelease mitchellh mapstructure v1.1.2
gazelleGHRelease go-openapi runtime v0.19.4
gazelle github.com/asaskevich/govalidator@f61b66f89f4a
gazelle go.mongodb.org/mongo-driver@1261197350f3ad46a907489aee7ecc49b39efb82 # v1.1.2
gazelleGHRelease jonboulle clockwork v0.1.0
gazelleGHRelease grpc-ecosystem go-grpc-middleware v1.1.0 
gazelle github.com/xiang90/probing@43a291ad63a2
gazelleGHRelease gorilla websocket v1.4.0
gazelleGHRelease google btree v1.0.0
gazelleGHRelease dgrijalva jwt-go v3.2.0
gazelleGHRelease go-stack stack v1.8.0
gazelle github.com/MakeNowJust/heredoc@bb23615498cd
gazelle github.com/daviddengcn/go-colortext@511bcaf42ccd
gazelle github.com/liggitt/tabwriter@89fcab3d43de
gazelle sigs.k8s.io/structured-merge-diff/v3@877aee05330847a873a1a8998b40e12a1e0fde25 # v3.0.0
gazelle sigs.k8s.io/apiserver-network-proxy/konnectivity-client@v0.0.7


# gRPC/proto deps (https://github.com/bazelbuild/rules_go/blob/master/go/workspace.rst#id8)
# bump down from 1.28.1 to 1.26.0 because https://github.com/etcd-io/etcd/issues/11563
gazelle google.golang.org/grpc@f5b0812e6fe574d90da76b205e9eb51f6ddb1919 # 1.26.0
gazelle golang.org/x/net@d3edc9973b7eb1fb302b0ff2c62357091cea9a30 # master
gazelle golang.org/x/text@f21a4dfb5e38f5895301dc265a8def02365cc3d0 # 0.3.0
gazelle github.com/golang/groupcache@02826c3e79038b59d737d3b1c0a1d937f71a4433

# containerd, Not an actual release, pinned to commit 8e685f78cf66e2901b2fbed2fdddd64449a74ab9 that has support for the required build tags.
gazelle \
    -build_tags no_zfs,no_aufs,no_devicemapper,no_btrfs \
    -build_file_proto_mode=disable \
    github.com/containerd/containerd@8e685f78cf66e2901b2fbed2fdddd64449a74ab9
# containerd deps (taken from github.com/containerd/containerd/vendor.conf)
gazelle github.com/beorn7/perks@37c8de3658fcb183f997c4e13e8337516ab753e6 # v1.0.1
gazelle github.com/BurntSushi/toml@3012a1dbe2e4bd1391d42b32f0577cb7bbc7f005 # v0.3.1
gazelle github.com/cespare/xxhash/v2@d7df74196a9e781ede915320c11c378c1b2f3a1f # v2.1.1
gazelle github.com/containerd/btrfs@153935315f4ab9be5bf03650a1341454b05efa5d
gazelle -build_file_proto_mode=disable github.com/containerd/cgroups@7347743e5d1e8500d9f27c8e748e689ed991d92b
gazelle github.com/containerd/console@8375c3424e4d7b114e8a90a4a40c8e1b40d1d4e6
gazelle github.com/containerd/continuity@0ec596719c75bfd42908850990acea594b7593ac
gazelle github.com/containerd/fifo@bda0ff6ed73c67bfb5e62bc9c697f146b7fd7f13
gazelle github.com/containerd/go-runc@a5c2862aed5e6358b305b0e16bfce58e0549b1cd
gazelle github.com/containerd/ttrpc@92c8520ef9f86600c650dd540266a007bf03670f
gazelle github.com/containerd/typeurl@a93fcdb778cd272c6e9b3028b2f42d813e785d40
gazelle github.com/coreos/go-systemd/v22@2d78030078ef61b3cae27f42ad6d0e46db51b339 # v22.0.0
gazelle github.com/cpuguy83/go-md2man@7762f7e404f8416dfa1d9bb6a8c192aa9acb4d19 # v1.0.10
gazelle github.com/docker/go-events@9461782956ad83b30282bf90e31fa6a70c255ba9
gazelle github.com/docker/go-metrics@b619b3592b65de4f087d9f16863a7e6ff905973c # v0.0.1
gazelle github.com/docker/go-units@519db1ee28dcc9fd2474ae59fca29a810482bfb1 # v0.4.0
gazelle github.com/godbus/dbus/v5@37bf87eef99d69c4f1d3528bd66e3a87dc201472 # v5.0.3
gazelle -build_file_proto_mode=disable github.com/gogo/googleapis@d31c731455cb061f42baff3bda55bad0118b126b # v1.2.0
gazelle github.com/gogo/protobuf@ba06b47c162d49f2af050fb4c75bcbc86a159d5c # v1.2.1
gazelle github.com/google/go-cmp@5a6f75716e1203a923a78c9efb94089d857df0f6 # v0.4.0
gazelle github.com/google/uuid@0cd6bf5da1e1c83f8b45653022c74f71af0538a4 # v1.1.1
gazelle github.com/hashicorp/errwrap@8a6fb523712970c966eefc6b39ed2c5e74880354 # v1.0.0
gazelle github.com/hashicorp/go-multierror@886a7fbe3eb1c874d46f623bfa70af45f425b3d1 # v1.0.0
gazelle github.com/hashicorp/golang-lru@7f827b33c0f158ec5dfbba01bb0b14a4541fd81d # v0.5.3
gazelle github.com/imdario/mergo@7c29201646fa3de8506f701213473dd407f19646 # v0.3.7
gazelle github.com/konsorten/go-windows-terminal-sequences@5c8c8bd35d3832f5d134ae1e1e375b69a4d25242 # v1.0.1
gazelle github.com/matttproud/golang_protobuf_extensions@c12348ce28de40eed0136aa2b644d0ee0650e56c # v1.0.1
gazelle github.com/Microsoft/go-winio@6c72808b55902eae4c5943626030429ff20f3b63 # v0.4.14
gazelle -build_file_proto_mode=disable github.com/Microsoft/hcsshim@0b571ac85d7c5842b26d2571de4868634a4c39d7 # v0.8.7-24-g0b571ac8
gazelle github.com/opencontainers/go-digest@c9281466c8b2f606084ac71339773efd177436e7
gazelle github.com/opencontainers/image-spec@d60099175f88c47cd379c4738d158884749ed235 # v1.0.1
gazelle github.com/opencontainers/runtime-spec@29686dbc5559d93fb1ef402eeda3e35c38d75af4 # v1.0.1-59-g29686db
gazelle github.com/pkg/errors@ba968bfe8b2f7e042a574c888954fccecfa385b4 # v0.8.1
gazelle github.com/prometheus/client_golang@c42bebe5a5cddfc6b28cd639103369d8a75dfa89 # v1.3.0
gazelle github.com/prometheus/client_model@d1d2010b5beead3fa1c5f271a5cf626e40b3ad6e # v0.1.0
gazelle github.com/prometheus/common@287d3e634a1e550c9e463dd7e5a75a422c614505 # v0.7.0
gazelle github.com/prometheus/procfs@6d489fc7f1d9cd890a250f3ea3431b1744b9623f # v0.0.8
gazelle github.com/russross/blackfriday@05f3235734ad95d0016f6a23902f06461fcf567a # v1.5.2
gazelle github.com/sirupsen/logrus@8bdbc7bcc01dcbb8ec23dc8a28e332258d25251f # v1.4.1
gazelle github.com/syndtr/gocapability@d98352740cb2c55f81556b63d4a1ec64c5a319c2
gazelle github.com/urfave/cli@bfe2e925cfb6d44b40ad3a779165ea7e8aff9212 # v1.22.0
gazelle go.etcd.io/bbolt@a0458a2b35708eef59eb5f620ceb3cd1c01a824d # v1.3.3
gazelle go.opencensus.io@9c377598961b706d1542bd2d84d538b5094d596e # v0.22.0
gazelle golang.org/x/sync@42b317875d0fa942474b76e1b46a6060d720ae6e
gazelle golang.org/x/sys@c990c680b611ac1aeb7d8f2af94a825f98d69720
gazelle google.golang.org/genproto@d80a6e20e776b0b17a324d0ba1ab50a39c8e8944
gazelle gotest.tools@1083505acf35a0bd8a696b26837e1fb3187a7a83 # v2.3.0
gazelle github.com/cilium/ebpf@60c3aa43f488292fe2ee50fb8b833b383ca8ebbb
gazelle sigs.k8s.io/kustomize@a6f65144121d1955266b0cd836ce954c04122dc8 # v2.0.3
gazelle vbom.ml/util@db5cfe13f5cc
gazelle github.com/exponent-io/jsonpath@d6023ce2651d 

# containerd/cri
gazelle -build_file_proto_mode=disable github.com/containerd/cri@c0294ebfe0b4342db85c0faf7727ceb8d8c3afce # master
gazelle github.com/containerd/go-cni@0d360c50b10b350b6bb23863fd4dfb1c232b01c9
gazelle github.com/containernetworking/cni@4cfb7b568922a3c79a23e438dc52fe537fc9687e # v0.7.1
# containerd/cri wants 0.7.6, but that actually doesn't build against containernetworking/cni 0.7.1. Bump to 0.8.2,
# which does.
gazelle github.com/davecgh/go-spew@8991bc29aa16c548c550c7ff78260e27b9ab7c73 # v1.1.1
gazelle github.com/docker/distribution@0d3efadf0154c2b8a4e7b6621fff9809655cc580
gazelle github.com/docker/docker@d1d5f6476656c6aad457e2a91d3436e66b6f2251
gazelle github.com/docker/spdystream@449fdfce4d962303d702fec724ef0ad181c92528
gazelle github.com/emicklei/go-restful@b993709ae1a4f6dd19cfa475232614441b11c9d5 # v2.9.5
gazelle github.com/google/gofuzz@f140a6486e521aad38f5917de355cbf147cc0496 # v1.0.0
gazelle github.com/json-iterator/go@03217c3e97663914aec3faafde50d081f197a0a2 # v1.1.8
gazelle github.com/modern-go/concurrent@bacd9c7ef1dd9b15be4a9909b8ac7a4e313eec94 # 1.0.3
gazelle github.com/modern-go/reflect2@4b7aa43c6742a2c18fdef89dd197aaae7dac7ccd # 1.0.1
gazelle -build_tags=selinux github.com/opencontainers/selinux@5215b1806f52b1fcc2070a8826c542c9d33cd3cf
gazelle github.com/seccomp/libseccomp-golang@689e3c1541a84461afc49c1c87352a6cedf72e9c # v0.9.1
gazelle github.com/stretchr/testify@221dbe5ed46703ee255b1da0dec05086f5035f62 # v1.4.0
gazelle github.com/tchap/go-patricia@666120de432aea38ab06bd5c818f04f4129882c9 # v2.2.6
gazelle golang.org/x/crypto@69ecbb4d6d5dab05e49161c6e77ea40a030884e1
gazelle golang.org/x/oauth2@0f29369cfe4552d0e4bcddc57cc75f4d7e672a33
gazelle golang.org/x/time@9d24e82272b4f38b78bc8cff74fa936d31ccd8ef
gazelle gopkg.in/inf.v0@d2d2541c53f18d2a059457998ce2876cc8e67cbf # v0.9.1
gazelle gopkg.in/yaml.v2@53403b58ad1b561927d19068c655246f2db79d48 # v2.2.8
gazelle k8s.io/klog@2ca9ad30301bf30a8a6e0fa2110db6b8df699a91 # v1.0.0
#k8s.io/kubernetes                                   d224476cd0730baca2b6e357d144171ed74192d6 # v1.17.1
gazelle sigs.k8s.io/yaml@fd68e9863619f6ec2fdd8625fe1f02e7c877e480 # v1.1.0
gazelleGHRelease fatih camelcase v1.0.0
gazelleGHRelease ghodss yaml v1.0.0 
gazelle github.com/gregjones/httpcache@9cad4c3443a7
gazelleGHRelease mitchellh go-wordwrap v1.0.0
gazelleGHRelease peterbourgon diskv v2.0.1
gazelle github.com/Azure/go-ansiterm@d6e3b3328b78
gazelle github.com/chai2010/gettext-go@c6fed771bfd5


# containernetworking/plugins dependencies
gazelle github.com/alexflint/go-filemutex@72bdc8eae2ae
gazelleGHRelease coreos go-iptables v0.4.5
gazelle github.com/safchain/ethtool@42ed695e3de8
gazelle github.com/j-keck/arping@2cf9dc699c56


# gvisor dependencies
gazelleGHRelease grpc grpc v1.26.0
gazelle github.com/google/subcommands@636abe8753b8
gazelle github.com/cenkalti/backoff@2146c93394225c3732078705043ce9f26584d334 # old vesrion
gazelleGHRelease kr pretty v0.2.0
gazelleGHRelease kr pty v1.1.1
gazelle github.com/gofrs/flock@886344bea079
gazelle golang.org/x/time@555d28b269f0

# runc 
gazelleGHRelease opencontainers runc v1.0.0-rc10

# sqlboiler and deps
gazelleGHRelease volatiletech sqlboiler v3.6.1
gazelleGHRelease friendsofgo errors v0.9.2
gazelleGHRelease go-sql-driver mysql v1.5.0
gazelleGHRelease spf13 cast v1.3.1
gazelleGHRelease spf13 cobra v1.0.0
gazelleGHRelease volatiletech null v8.0.0
gazelleGHRelease volatiletech randomize v0.0.1
gazelleGHRelease volatiletech strmangle v0.0.1
gazelleGHRelease lib pq v1.2.0
gazelle github.com/denisenkom/go-mssqldb@bbfc9a55622e
gazelle github.com/ericlagergren/decimal@73749d4874d5
gazelle github.com/rubenv/sql-migrate@9355dd04f4b3
gazelle github.com/glerchundi/sqlboiler-crdb@62014c8c8df1
gazelleGHRelease mitchellh cli v1.0.0
gazelle github.com/olekukonko/tablewriter@a0225b3f23b5
gazelleGHRelease mattn go-sqlite3 v1.12.0
gazelle gopkg.in/gorp.v1@v1.7.2
gazelle github.com/volatiletech/inflect@e7201282ae8d
gazelleGHRelease mattn go-runewidth v0.0.2
gazelle github.com/golang-sql/civil@cb61b32ac6fe
gazelleGHRelease posener complete v1.1.1
gazelleGHRelease mattn go-isatty v0.0.4
gazelle github.com/armon/go-radix@7fddfc383310
gazelleGHRelease bgentry speakeasy v0.1.0
gazelleGHRelease fatih color v1.7.0
gazelleGHRelease spf13 viper v1.3.2
gazelleGHRelease hashicorp hcl v1.0.0
gazelleGHRelease spf13 jwalterweatherman v1.0.0
gazelleGHRelease magiconair properties v1.8.0
gazelleGHRelease mattn go-colorable v0.0.9

# our own deps
gazelle github.com/google/go-tpm@ae6dd98980d4
gazelle github.com/google/go-tpm-tools@f8c04ff88181
gazelle github.com/insomniacslk/dhcp@5dd7202f19711228cb4a51aa8b3415421c2edefe
gazelle github.com/cenkalti/backoff/v4@18fe4ce5a8550e0d0919b680ad3c080a5455bddf # v4.0.2
gazelle github.com/rekby/gpt@a930afbc6edcc89c83d39b79e52025698156178d
gazelle github.com/yalue/native_endian@51013b03be4fd97b0aabf29a6923e60359294186
# used by incomniacslk/dhcp
gazelle github.com/mdlayher/ethernet@0394541c37b7f86a10e0b49492f6d4f605c34163
gazelle github.com/mdlayher/raw@50f2db8cc0658568575938a39dbaa46172921d98
# used by insomniacslk/dhcp for pkg/uio
gazelleGHRelease u-root u-root v6.0.0
gazelleGHRelease diskfs go-diskfs v1.0.0
# used by diskfs/go-diskfs
gazelle gopkg.in/djherbis/times.v1@847c5208d8924cea0acea3376ff62aede93afe39
gazelleGHRelease kevinburke go-bindata v3.16.0

# delta & deltagen deps
gazelleGHRelease gofrs uuid v3.2.0
gazelleGHRelease lyft protoc-gen-star v0.4.14

