#  Copyright 2020 The Monogon Project Authors.
#
#  SPDX-License-Identifier: Apache-2.0
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

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
        sum = "h1:YqJuHm/xOYP2VIOWPnQO+ix+Ag5KditpdHmIreWYyTY=",
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
        sum = "h1:Cef96rKLuXxeGzERI/0ve9yAzIeTpx0qz9JKFDZALYw=",
        version = "v1.0.2-0.20190207185410-29686dbc5559",
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

    go_repository(
        name = "io_k8s_repo_infra",
        importpath = "k8s.io/repo-infra",
        sum = "h1:PQyAIB6SRdV0a3Vj/VA39L1uANW36k/zg3tOk/Ffh3U=",
        version = "v0.0.0-20190329054012-df02ded38f95",
    )
    go_repository(
        name = "io_k8s_component_base",
        importpath = "k8s.io/component-base",
        sum = "h1:nZfjiRab7LcpScUgnudRCG6UMRVmZ3L0GNZZWHkYjus=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_cloud_provider",
        importpath = "k8s.io/cloud-provider",
        sum = "h1:nFQ/M6B8o+/ICEHbjLFMe4nHgH/8jAHQ1GFw2NJ5Elo=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_kubelet",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kubelet",
        sum = "h1:UPkB1eGbkIWr38J++4Gk7LZjcYeB5JfJBqKzRGfJ/VM=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_csi_translation_lib",
        importpath = "k8s.io/csi-translation-lib",
        sum = "h1:lH3FPZqHFwPthCQKLKNP90LR5oqjAMxYMJNhicDA5d8=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_legacy_cloud_providers",
        importpath = "k8s.io/legacy-cloud-providers",
        sum = "h1:jpu9SqacduO6iKtiCKCovH/uZ0GL1PkbsJndUZKUxtc=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_apiextensions_apiserver",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apiextensions-apiserver",
        sum = "h1:lQjE543mSh4jeBxrvnwz37DCzGHW2UMefX8eCzk8uAU=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_kube_aggregator",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kube-aggregator",
        sum = "h1:Li0htDytvDHRnf7IR9AWGSahhyvD4qVxWIJwsUVgo2w=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_metrics",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/metrics",
        sum = "h1:5/OfIQ5HeJutKUPpjXXdcgFqxmFf01bYfnFRd1li5b8=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_kube_scheduler",
        importpath = "k8s.io/kube-scheduler",
        sum = "h1:EpIJpmI5Nn3mii1aaWg5VFMd9Y0Qt+jCcduVxH92Vk8=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_kube_controller_manager",
        importpath = "k8s.io/kube-controller-manager",
        sum = "h1:E5GkOKLf+ODm2uXQaBqtmf+D4ZJpUUlo8XJoX0nEDL0=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_cluster_bootstrap",
        importpath = "k8s.io/cluster-bootstrap",
        sum = "h1:MHG+0kAEEh4nDQU2iC8NXNILDDIANK12RB8PcAjyej4=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_kube_proxy",
        importpath = "k8s.io/kube-proxy",
        sum = "h1:8awQLk0DLJEXew80mjbFTMNs9EtbtXJElBi7K7BqalE=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_kubectl",
        importpath = "k8s.io/kubectl",
        sum = "h1:ygJWExSY2hnEHt72gJV6DgPDmkdp6xwkQlrZbtmW9EI=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_node_api",
        importpath = "k8s.io/node-api",
        sum = "h1:L5S79H5EO4jDK/ARsypIXeOrirXJmFplj4ZBIvnK+WM=",
        version = "v0.17.5",
    )
    go_repository(
        name = "io_k8s_cli_runtime",
        importpath = "k8s.io/cli-runtime",
        sum = "h1:/cZeGGp0GxuFSUdjz8jlUQP75QJVz99YtXEU1uNW/LI=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_api",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/api",
        sum = "h1:GVZeds8bgQOSdQ/LYcjL7+NstBByZ5L3U/Ks6+E+QRI=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_apimachinery",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apimachinery",
        sum = "h1:N155+ZeSeRnCFyzjYRv3vg9GWJIUm5ElZba66f7qicY=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_apiserver",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apiserver",
        sum = "h1:k1fpzJAPZvtRT9Z8Rc42kciGehIH0GiEmTgEmc46drw=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "io_k8s_cri_api",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/cri-api",
        sum = "h1:JDsPY0mIzxR6JYGWKWhX7NIIXa9giiVQ1X/RE0Mw1GY=",
        version = "v0.19.0-alpha.2",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        sum = "h1:6m/oheQuQ13N9ks4hubMG6BnvwOeaJrqSPLahSnczz8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_blang_semver",
        importpath = "github.com/blang/semver",
        sum = "h1:CGxCgetQ64DKk7rdZ++Vfnb1+ogGNnB17OJKJXD2Cfs=",
        version = "v3.5.0+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_systemd",
        importpath = "github.com/coreos/go-systemd",
        sum = "h1:Wf6HqHfScWJN9/ZjdUKyjop4mf3Qdd+1TvvltAvM3m8=",
        version = "v0.0.0-20190321100706-95778dfbb74e",
    )
    go_repository(
        name = "io_k8s_utils",
        importpath = "k8s.io/utils",
        sum = "h1:d4vVOjXm687F1iLSP2q3lyPPuyvTUt3aVoBpi2DqRsU=",
        version = "v0.0.0-20200324210504-a9aa75ae1b89",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:5jhuqJyZCZf2JRofRvN/nIFgIWNzPa3/Vz8mYylgbWc=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_armon_circbuf",
        importpath = "github.com/armon/circbuf",
        sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
        version = "v0.0.0-20150827004946-bbbad097214e",
    )
    go_repository(
        name = "com_github_container_storage_interface_spec",
        importpath = "github.com/container-storage-interface/spec",
        sum = "h1:bD9KIVgaVKKkQ/UbVUY9kCaH/CJbhNxe0eeB4JeJV2s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_docker_go_connections",
        importpath = "github.com/docker/go-connections",
        sum = "h1:3lOnM9cSzgGwx8VfK/NGOW5fLQ0GjIlCkaktF+n1M6o=",
        version = "v0.3.0",
    )
    go_repository(
        name = "io_k8s_kube_openapi",
        importpath = "k8s.io/kube-openapi",
        sum = "h1:FDWYFE3itI1G8UFOMjUuLbROZExo+Rrfm/Qaf473rm4=",
        version = "v0.0.0-20200403204345-e1beb1bd0f35",
    )
    go_repository(
        name = "com_github_go_openapi_spec",
        importpath = "github.com/go-openapi/spec",
        sum = "h1:0XRyw8kguri6Yw4SxhsQA/atC88yqrk0+G4YhI2wabc=",
        version = "v0.19.3",
    )
    go_repository(
        name = "com_github_evanphx_json_patch",
        importpath = "github.com/evanphx/json-patch",
        sum = "h1:fUDGZCv/7iAN7u0puUVhvKCcsR6vRfwrJatElLBEf0I=",
        version = "v4.2.0+incompatible",
    )
    go_repository(
        name = "io_etcd_go_etcd",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd",
        sum = "h1:C7kWARE8r64ppRadl40yfNo6pag+G6ocvGU2xZ6yNes=",
        version = "v0.5.0-alpha.5.0.20200401174654-e694b7bb0875",
    )
    go_repository(
        name = "com_github_go_openapi_jsonpointer",
        importpath = "github.com/go-openapi/jsonpointer",
        sum = "h1:gihV7YNZK1iK6Tgwwsxo2rJbD1GTbdm72325Bq8FI3w=",
        version = "v0.19.3",
    )
    go_repository(
        name = "com_github_go_openapi_swag",
        importpath = "github.com/go-openapi/swag",
        sum = "h1:lTz6Ys4CmqqCQmZPBlbQENR1/GucA2bzYTE12Pw4tFY=",
        version = "v0.19.5",
    )
    go_repository(
        name = "com_github_go_openapi_jsonreference",
        importpath = "github.com/go-openapi/jsonreference",
        sum = "h1:5cxNfTy0UVC3X8JL5ymxzyoUZmo8iZb+jeTWn7tUa8o=",
        version = "v0.19.3",
    )
    go_repository(
        name = "com_github_cyphar_filepath_securejoin",
        importpath = "github.com/cyphar/filepath-securejoin",
        sum = "h1:jCwT2GTP+PY5nBz3c/YL5PAIbusElVrPujOBSCj8xRg=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:aizVhC/NAAcKWb+5QsU1iNOZb4Yws5UO2I+aIprQITM=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_puerkitobio_purell",
        importpath = "github.com/PuerkitoBio/purell",
        sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_puerkitobio_urlesc",
        importpath = "github.com/PuerkitoBio/urlesc",
        sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
        version = "v0.0.0-20170810143723-de5bf2ad4578",
    )
    go_repository(
        name = "com_github_mxk_go_flowrate",
        importpath = "github.com/mxk/go-flowrate",
        sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
        version = "v0.0.0-20140419014527-cca7078d478f",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:ROfEUZz+Gh5pa62DJWXSaonyu3StP6EA6lPEXPI6mCo=",
        version = "v0.38.0",
    )
    go_repository(
        name = "com_github_karrick_godirwalk",
        importpath = "github.com/karrick/godirwalk",
        sum = "h1:VbzFqwXwNbAZoA6W5odrLr+hKK197CcENcPh6E/gJ0M=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_github_euank_go_kmsg_parser",
        importpath = "github.com/euank/go-kmsg-parser",
        sum = "h1:cHD53+PLQuuQyLZeriD1V/esuG4MuU0Pjs5y6iknohY=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_mindprince_gonvml",
        importpath = "github.com/mindprince/gonvml",
        sum = "h1:PS1dLCGtD8bb9RPKJrc8bS7qHL6JnW1CZvwzH9dPoUs=",
        version = "v0.0.0-20190828220739-9ebdce4bb989",
    )
    go_repository(
        name = "com_github_checkpoint_restore_go_criu",
        importpath = "github.com/checkpoint-restore/go-criu",
        sum = "h1:T4nWG1TXIxeor8mAu5bFguPJgSIGhZqv/f0z55KCrJM=",
        version = "v0.0.0-20190109184317-bdb7599cd87b",
    )
    go_repository(
        name = "com_github_mrunalp_fileutils",
        importpath = "github.com/mrunalp/fileutils",
        sum = "h1:7InQ7/zrOh6SlFjaXFubv0xX0HsuC9qJsdqm7bNQpYM=",
        version = "v0.0.0-20171103030105-7d4729fb3618",
    )
    go_repository(
        name = "com_github_vishvananda_netns",
        importpath = "github.com/vishvananda/netns",
        sum = "h1:OviZH7qLw/7ZovXvuNyL3XQl8UFofeikI1NW1Gypu7k=",
        version = "v0.0.0-20191106174202-0a2b9b5464df",
    )
    go_repository(
        name = "com_github_munnerz_goautoneg",
        importpath = "github.com/munnerz/goautoneg",
        sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
        version = "v0.0.0-20191010083416-a7dc8b61c822",
    )
    go_repository(
        name = "com_github_nytimes_gziphandler",
        importpath = "github.com/NYTimes/gziphandler",
        sum = "h1:lsxEuwrXEAokXB9qhlbKWPpo3KMLZQ5WB5WLQRW1uq0=",
        version = "v0.0.0-20170623195520-56545f4a5d46",
    )
    go_repository(
        name = "com_github_vishvananda_netlink",
        importpath = "github.com/vishvananda/netlink",
        sum = "h1:1iyaYNBLmP6L0220aDnYQpo1QEV4t4hJ+xEEhhJH8j0=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_googleapis_gnostic",
        importpath = "github.com/googleapis/gnostic",
        sum = "h1:DLJCy1n/vrD4HPjOvYcT8aYQXpPIzoRZONaYwyycI+I=",
        version = "v0.4.1",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        sum = "h1:ZZCA22JRF2gQE5FoNmhmrf7jeJJ2uhqDUNRYKm8dvmM=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
        version = "v0.3.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        sum = "h1:HoEmRHQPVSqub6w2z2d2EOVs2fjyFRGyofhKuyDq0QI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        sum = "h1:2Oa65PReHzfn29GpvgsYwloV9AVFHPDk8tYxt2c2tr4=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_coreos_pkg",
        importpath = "github.com/coreos/pkg",
        sum = "h1:n2Ltr3SrfQlf/9nOna1DoGKxLx3qTSI8Ttl6Xrqp6mw=",
        version = "v0.0.0-20180108230652-97fdf19511ea",
    )
    go_repository(
        name = "com_github_morikuni_aec",
        importpath = "github.com/morikuni/aec",
        sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_lithammer_dedent",
        importpath = "github.com/lithammer/dedent",
        sum = "h1:VNzHMVCBNG1j0fh3OrsFRkVUwStdDArbgBWoPAffktY=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
        version = "v1.2.0",
    )
    go_repository(
        name = "in_gopkg_natefinch_lumberjack_v2",
        importpath = "gopkg.in/natefinch/lumberjack.v2",
        sum = "h1:Kjoe/521+yYTO0xv6VymPwLPVq3RES8170ZpJkrFE9Q=",
        version = "v2.0.0-20161104145732-dd45e6a67c53",
    )
    go_repository(
        name = "com_github_robfig_cron",
        importpath = "github.com/robfig/cron",
        sum = "h1:ZjScXvvxeQ63Dbyxy76Fj3AT3Ut0aKsyd2/tl3DTMuQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_coreos_go_oidc",
        importpath = "github.com/coreos/go-oidc",
        sum = "h1:sdJrfw8akMnCuUlaZU3tE/uYXFgfqom8DBE9so9EBsM=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "org_gonum_v1_gonum",
        importpath = "gonum.org/v1/gonum",
        sum = "h1:4r+yNT0+8SWcOkXP+63H2zQbN+USnC73cjGUxnDF94Q=",
        version = "v0.6.2",
    )
    go_repository(
        name = "io_k8s_heapster",
        importpath = "k8s.io/heapster",
        sum = "h1:zq1luJo07HL2HT4hiddH4nron62gVRlAq/DpOV3n31s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "in_gopkg_square_go_jose_v2",
        importpath = "gopkg.in/square/go-jose.v2",
        sum = "h1:orlkJ3myw8CN1nVQHBFfloD+L3egixIa4FvUP6RosSA=",
        version = "v2.2.2",
    )
    go_repository(
        name = "com_github_pquerna_cachecontrol",
        importpath = "github.com/pquerna/cachecontrol",
        sum = "h1:0XM1XL/OFFJjXsYXlG30spTkV/E9+gmd5GD1w2HE8xM=",
        version = "v0.0.0-20171018203845-0dec1b30a021",
    )
    go_repository(
        name = "com_github_go_openapi_validate",
        importpath = "github.com/go-openapi/validate",
        sum = "h1:QhCBKRYqZR+SKo4gl1lPhPahope8/RLt6EVgY8X80w0=",
        version = "v0.19.5",
    )
    go_repository(
        name = "com_github_go_openapi_strfmt",
        importpath = "github.com/go-openapi/strfmt",
        sum = "h1:eRfyY5SkaNJCAwmmMcADjY31ow9+N7MCLW7oRkbsINA=",
        version = "v0.19.3",
    )
    go_repository(
        name = "io_k8s_gengo",
        importpath = "k8s.io/gengo",
        sum = "h1:pZzawYyz6VRNPVYpqGv61LWCimQv1BihyeqFrp50/G4=",
        version = "v0.0.0-20200205140755-e0e292d8aa12",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:KU7oHjnv3XNWfa5COkzUifxZmxp1TyI7ImMXqFxLwvQ=",
        version = "v0.2.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
        version = "v0.0.0-20191204190536-9bdfabe68543",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        sum = "h1:UImYN5qQ8tuGpGE16ZmjvcTtTw24zw1QAp/SlnNrZhI=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_github_soheilhy_cmux",
        importpath = "github.com/soheilhy/cmux",
        sum = "h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_tmc_grpc_websocket_proxy",
        importpath = "github.com/tmc/grpc-websocket-proxy",
        sum = "h1:ndzgwNDnKIqyCvHTXaCqh9KlOWKvBry6nuXMJmonVsE=",
        version = "v0.0.0-20170815181823-89b8d40f7ca8",
    )
    go_repository(
        name = "com_github_go_openapi_errors",
        importpath = "github.com/go-openapi/errors",
        sum = "h1:a2kIyV3w+OS3S97zxUndRVD46+FhGOUBDFY7nmu4CsY=",
        version = "v0.19.2",
    )
    go_repository(
        name = "com_github_go_openapi_analysis",
        importpath = "github.com/go-openapi/analysis",
        sum = "h1:8b2ZgKfKIUTVQpTb77MoRDIMEIwvDVw40o3aOXdfYzI=",
        version = "v0.19.5",
    )
    go_repository(
        name = "com_github_go_openapi_loads",
        importpath = "github.com/go-openapi/loads",
        sum = "h1:5I4CCSqoWzT+82bBkNIvmLc0UOsoKKQ4Fz+3VxOB7SY=",
        version = "v0.19.4",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_go_openapi_runtime",
        importpath = "github.com/go-openapi/runtime",
        sum = "h1:csnOgcgAiuGoM/Po7PEpKDoNulCcF3FGbSnbHfxgjMI=",
        version = "v0.19.4",
    )
    go_repository(
        name = "com_github_asaskevich_govalidator",
        importpath = "github.com/asaskevich/govalidator",
        sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
        version = "v0.0.0-20190424111038-f61b66f89f4a",
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        importpath = "go.mongodb.org/mongo-driver",
        sum = "h1:jxcFYjlkl8xaERsgLo+RNquI0epW6zuy/ZRQs6jnrFA=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_jonboulle_clockwork",
        importpath = "github.com/jonboulle/clockwork",
        sum = "h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        sum = "h1:THDBEeQ9xZ8JEaCLyLQqXMMdRqNr0QAUJTIkQAUtFjg=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_xiang90_probing",
        importpath = "github.com/xiang90/probing",
        sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
        version = "v0.0.0-20190116061207-43a291ad63a2",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:WDFjx/TMzVgy9VdMMQi2K2Emtwi2QcUQsztZ/zLaH/Q=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_google_btree",
        importpath = "github.com/google/btree",
        sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_dgrijalva_jwt_go",
        importpath = "github.com/dgrijalva/jwt-go",
        sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_makenowjust_heredoc",
        importpath = "github.com/MakeNowJust/heredoc",
        sum = "h1:sjQovDkwrZp8u+gxLtPgKGjk5hCxuy2hrRejBTA9xFU=",
        version = "v0.0.0-20170808103936-bb23615498cd",
    )
    go_repository(
        name = "com_github_daviddengcn_go_colortext",
        importpath = "github.com/daviddengcn/go-colortext",
        sum = "h1:uVsMphB1eRx7xB1njzL3fuMdWRN8HtVzoUOItHMwv5c=",
        version = "v0.0.0-20160507010035-511bcaf42ccd",
    )
    go_repository(
        name = "com_github_liggitt_tabwriter",
        importpath = "github.com/liggitt/tabwriter",
        sum = "h1:9TO3cAIGXtEhnIaL+V+BEER86oLrvS+kWobKpbJuye0=",
        version = "v0.0.0-20181228230101-89fcab3d43de",
    )
    go_repository(
        name = "io_k8s_sigs_structured_merge_diff_v3",
        importpath = "sigs.k8s.io/structured-merge-diff/v3",
        sum = "h1:dOmIZBMfhcHS09XZkMyUgkq5trg3/jRyJYFZUiaOp8E=",
        version = "v3.0.0",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:2dTRdpdFEEhJYQD8EMLB61nnrzSCTbG38PhqdhvOltg=",
        version = "v1.26.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:3G+cUijn7XD+S4eJFddp53Pv7+slrESplyjG25HgL+k=",
        version = "v0.0.0-20200324143707-d3edc9973b7e",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:LbsanbbD6LieFkXbj9YNNBupiGHJgFeLpO0j0Fza1h8=",
        version = "v0.0.0-20160516000752-02826c3e7903",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
        version = "v2.1.1",
    )
    go_repository(
        name = "com_github_containerd_btrfs",
        importpath = "github.com/containerd/btrfs",
        sum = "h1:u5X1yvVEsXLcuTWYsFSpTgQKRvo2VTB5gOHcERpF9ZI=",
        version = "v0.0.0-20200117014249-153935315f4a",
    )
    go_repository(
        name = "com_github_containerd_cgroups",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/cgroups",
        sum = "h1:sL8rdngVdYA2SLRwj6sSZ1cLDpBkFBd7IZVp0M2Lboc=",
        version = "v0.0.0-20200113070643-7347743e5d1e",
    )
    go_repository(
        name = "com_github_containerd_console",
        importpath = "github.com/containerd/console",
        sum = "h1:fU3UuQapBs+zLJu82NhR11Rif1ny2zfMMAyPJzSN5tQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_continuity",
        importpath = "github.com/containerd/continuity",
        sum = "h1:hBSbT5nWoYGwpmUa8TCsSVFVSdTyFoNlz85rNkH4OGk=",
        version = "v0.0.0-20200107062522-0ec596719c75",
    )
    go_repository(
        name = "com_github_containerd_fifo",
        importpath = "github.com/containerd/fifo",
        sum = "h1:KFbqHhDeaHM7IfFtXHfUHMDaUStpM2YwBR+iJCIOsKk=",
        version = "v0.0.0-20190816180239-bda0ff6ed73c",
    )
    go_repository(
        name = "com_github_containerd_go_runc",
        importpath = "github.com/containerd/go-runc",
        sum = "h1:9aJwmidmB33rxuib1NxR5NT4nvDMA9/S2sDR/D3tE5U=",
        version = "v0.0.0-20191206163734-a5c2862aed5e",
    )
    go_repository(
        name = "com_github_containerd_ttrpc",
        importpath = "github.com/containerd/ttrpc",
        sum = "h1:NY8Zk2i7TpkLxrkOASo+KTFq9iNCEmMH2/ZG9OuOw6k=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_typeurl",
        importpath = "github.com/containerd/typeurl",
        sum = "h1:7LMH7LfEmpWeCkGcIputvd4P0Rnd0LrIv1Jk2s5oobs=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:XJIw/+VlJ+87J+doOxznsAWIdmWuViOVhkQamW5YV28=",
        version = "v22.0.0",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man",
        importpath = "github.com/cpuguy83/go-md2man",
        sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
        version = "v1.0.10",
    )
    go_repository(
        name = "com_github_docker_go_events",
        importpath = "github.com/docker/go-events",
        sum = "h1:VXIse57M5C6ezDuCPyq6QmMvEJ2xclYKZ35SfkXdm3E=",
        version = "v0.0.0-20170721190031-9461782956ad",
    )
    go_repository(
        name = "com_github_docker_go_metrics",
        importpath = "github.com/docker/go-metrics",
        sum = "h1:AgB/0SvBxihN0X8OR4SjsblXkbMvalQ8cjmtKQ2rQV8=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_docker_go_units",
        importpath = "github.com/docker/go-units",
        sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        importpath = "github.com/godbus/dbus/v5",
        sum = "h1:ZqHaoEF7TBzh4jzPmqVhE/5A1z9of6orkAe5uHoAeME=",
        version = "v5.0.3",
    )
    go_repository(
        name = "com_github_gogo_googleapis",
        build_file_proto_mode = "disable",
        importpath = "github.com/gogo/googleapis",
        sum = "h1:Z0v3OJDotX9ZBpdz2V+AI7F4fITSZhVE5mg6GQppwMM=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:/s5zKNz0uPFCZ5hddgPdo2TK2TVrUNMn0OOX8/aZMTE=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:xsAVV57WRhGj6kEIi8ReJzQlHHqcBYCElAvkovg3B/4=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        importpath = "github.com/hashicorp/errwrap",
        sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        importpath = "github.com/hashicorp/go-multierror",
        sum = "h1:iVjPR7a6H0tWELX5NxNe7bYopibicUzc7uPribsnS6o=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:YPkqC67at8FYaadspW/6uE0COsBxS2656RLEr8Bppgk=",
        version = "v0.5.3",
    )
    go_repository(
        name = "com_github_imdario_mergo",
        importpath = "github.com/imdario/mergo",
        sum = "h1:Y+UAYTZ7gDEuOfhxKWy+dvb5dRQ6rJjFSdX2HZY1/gI=",
        version = "v0.3.7",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        sum = "h1:mweAR1A6xJ3oS2pRaGiHgQ4OO8tzTaLawm8vnODuwDk=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        importpath = "github.com/Microsoft/go-winio",
        sum = "h1:+hMXMk01us9KgxGb7ftKQt2Xpf5hH/yky+TDA+qxleU=",
        version = "v0.4.14",
    )
    go_repository(
        name = "com_github_microsoft_hcsshim",
        build_file_proto_mode = "disable",
        importpath = "github.com/Microsoft/hcsshim",
        sum = "h1:GDxLeqRF1hCkdTFNncrs8fZNcB6Fg79G0Q3m38EyySM=",
        version = "v0.8.8-0.20200109000640-0b571ac85d7c",
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        importpath = "github.com/opencontainers/go-digest",
        sum = "h1:2C93eP55foV5f0eNmXbidhKzwUZbs/Gk4PRp1zfeffs=",
        version = "v1.0.0-rc1.0.20180430190053-c9281466c8b2",
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        importpath = "github.com/opencontainers/image-spec",
        sum = "h1:JMemWkRwHx4Zj+fVxWoMCFm/8sYGGrUVojFA6h/TRcI=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:iURUrRGxPUNPdy5/HRSm+Yj6okJ6UtLINN0Q9M4+h3I=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:miYCvYqFXtl/J9FIy8eNpBfYthAEFg+Ys0XyUVEcDsc=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:ElTg5tNp4DqfV7UQjDqv2+RJlNzsDtvNAWccbItceIE=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        sum = "h1:L+1lyG48J1zAQXA3RBX/nG/B3gjlHq0zTt2tlbJLyCY=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:+fpWZdT24pJBiqJdAwYBjPSk+5YmQzYNPYzQsdzLkt8=",
        version = "v0.0.8",
    )
    go_repository(
        name = "com_github_russross_blackfriday",
        importpath = "github.com/russross/blackfriday",
        sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:GL2rEmy6nsikmW0r8opw9JIRScdMF5hA8cOYLH7In1k=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_syndtr_gocapability",
        importpath = "github.com/syndtr/gocapability",
        sum = "h1:b6uOv7YOFK0TYG7HtkIgExQo+2RdLuwRft63jn2HWj8=",
        version = "v0.0.0-20180916011248-d98352740cb2",
    )
    go_repository(
        name = "com_github_urfave_cli",
        importpath = "github.com/urfave/cli",
        sum = "h1:8nz/RUUotroXnOpYzT/Fy3sBp+2XEbXaY641/s3nbFI=",
        version = "v1.22.0",
    )
    go_repository(
        name = "io_etcd_go_bbolt",
        importpath = "go.etcd.io/bbolt",
        sum = "h1:MUGmc65QhB3pIlaQ5bB4LwqSj6GIonVJXpZiaKNyaKk=",
        version = "v1.3.3",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:C9hSCOW830chIVkdja34wa6Ky+IzWllkUinR+BtRZd4=",
        version = "v0.22.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:Bl/8QSvNqXvPGPGXa2z5xUTmV7VDcZyvRZ+QQXkXTZQ=",
        version = "v0.0.0-20181108010431-42b317875d0f",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:wVJP1pATLVPNxCz4R2mTO6HUJgfGE0PmIu2E10RuhCw=",
        version = "v0.0.0-20170523043604-d80a6e20e776",
    )
    go_repository(
        name = "tools_gotest",
        importpath = "gotest.tools",
        sum = "h1:YPidOweaQrSUDfne29Fnuwwo8uoQZuxnrAzZ+Q0pTeE=",
        version = "v1.4.1-0.20181223230014-1083505acf35",
    )
    go_repository(
        name = "com_github_cilium_ebpf",
        importpath = "github.com/cilium/ebpf",
        sum = "h1:kNrHgLQr3ftwQr9JKL3lmyNVlc/7Mjd8lwcbccE5BsI=",
        version = "v0.0.0-20191203103619-60c3aa43f488",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize",
        importpath = "sigs.k8s.io/kustomize",
        sum = "h1:JUufWFNlI44MdtnjUqVnvh29rR37PQFzPbLXqhyOyX0=",
        version = "v2.0.3+incompatible",
    )
    go_repository(
        name = "ml_vbom_util",
        importpath = "vbom.ml/util",
        sum = "h1:MksmcCZQWAQJCTA5T0jgI/0sJ51AVm4Z41MrmfczEoc=",
        version = "v0.0.0-20160121211510-db5cfe13f5cc",
    )
    go_repository(
        name = "com_github_exponent_io_jsonpath",
        importpath = "github.com/exponent-io/jsonpath",
        sum = "h1:105gxyaGwCFad8crR9dcMQWvV9Hvulu6hwUh4tWPJnM=",
        version = "v0.0.0-20151013193312-d6023ce2651d",
    )
    go_repository(
        name = "com_github_containerd_cri",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/cri",
        sum = "h1:tkxzigQGIymwkagfa+zsr1GzlYWJCVh6dUVhEc3fQeo=",
        version = "v1.11.1-0.20200130003317-c0294ebfe0b4",
    )
    go_repository(
        name = "com_github_containerd_go_cni",
        importpath = "github.com/containerd/go-cni",
        sum = "h1:76H5xRcgFYQvHpdlKBiw3CJOeaatmhn6ZETIsNWZJVs=",
        version = "v0.0.0-20190822145629-0d360c50b10b",
    )
    go_repository(
        name = "com_github_containernetworking_cni",
        importpath = "github.com/containernetworking/cni",
        sum = "h1:fE3r16wpSEyaqY4Z4oFrLMmIGfBYIKpPrHK31EJ9FzE=",
        version = "v0.7.1",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_docker_distribution",
        importpath = "github.com/docker/distribution",
        sum = "h1:dvc1KSkIYTVjZgHf/CTC2diTYC8PzhaA5sFISRfNVrE=",
        version = "v2.7.1-0.20190205005809-0d3efadf0154+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker",
        importpath = "github.com/docker/docker",
        sum = "h1:+kIkr4upwOTq7D78hByaTvwFw5F8WRkoGwDgBNJt4SA=",
        version = "v17.12.0-ce-rc1.0.20191121165722-d1d5f6476656+incompatible",
    )
    go_repository(
        name = "com_github_docker_spdystream",
        importpath = "github.com/docker/spdystream",
        sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
        version = "v0.0.0-20160310174837-449fdfce4d96",
    )
    go_repository(
        name = "com_github_emicklei_go_restful",
        importpath = "github.com/emicklei/go-restful",
        sum = "h1:spTtZBk5DYEvbxMVutUuTyh1Ao2r4iyvLdACqsl/Ljk=",
        version = "v2.9.5+incompatible",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        sum = "h1:QiWkFLKq0T7mpzwOTu6BzNDbfTE8OLrYhVKYMLF46Ok=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:Esafd1046DLDQ0W1YjYsBW+p8U2u7vzgW2SQVmlNazg=",
        version = "v0.0.0-20180701023420-4b7aa43c6742",
    )
    go_repository(
        name = "com_github_opencontainers_selinux",
        build_tags = ["selinux"],
        importpath = "github.com/opencontainers/selinux",
        sum = "h1:B8hYj3NxHmjsC3T+tnlZ1UhInqUgnyF1zlGPmzNg2Qk=",
        version = "v1.3.1-0.20190929122143-5215b1806f52",
    )
    go_repository(
        name = "com_github_seccomp_libseccomp_golang",
        importpath = "github.com/seccomp/libseccomp-golang",
        sum = "h1:NJjM5DNFOs0s3kYE1WUOr6G8V97sdt46rlXTMfXGWBo=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:2E4SXV/wtOkTonXsotYi4li6zVWxYlZuYNCXe9XRJyk=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_tchap_go_patricia",
        importpath = "github.com/tchap/go-patricia",
        sum = "h1:JvoDL7JSoIP2HDE8AbDH3zC8QBPxmzYe32HHy5yQ+Ck=",
        version = "v2.2.6+incompatible",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:9FCpayM9Egr1baVnV1SX0H87m+XB0B8S0hAMi99X/3U=",
        version = "v0.0.0-20200128174031-69ecbb4d6d5d",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:SVwTIAaPC2U/AvvLNZ2a7OVsmBpC8L5BlwK1whH3hm0=",
        version = "v0.0.0-20190604053449-0f29369cfe45",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:/5xXl8Y5W96D+TtHSlonuFqGHIWVuyCkGJLwGh9JJFs=",
        version = "v0.0.0-20191024005414-555d28b269f0",
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        importpath = "gopkg.in/inf.v0",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
        version = "v0.9.1",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:obN1ZagJSUGI0Ek/LBmuj4SNLPfIny3KsKFopxRdj10=",
        version = "v2.2.8",
    )
    go_repository(
        name = "io_k8s_klog",
        importpath = "k8s.io/klog",
        sum = "h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "io_k8s_sigs_yaml",
        importpath = "sigs.k8s.io/yaml",
        sum = "h1:4A07+ZFc2wgJwo8YNlQpr1rVlgUDlxXHhPJciaPY5gs=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_fatih_camelcase",
        importpath = "github.com/fatih/camelcase",
        sum = "h1:hxNvNX/xYBp0ovncs8WyWZrOrpBNub/JfaMvbURyft8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gregjones_httpcache",
        importpath = "github.com/gregjones/httpcache",
        sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
        version = "v0.0.0-20180305231024-9cad4c3443a7",
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        importpath = "github.com/mitchellh/go-wordwrap",
        sum = "h1:6GlHJ/LTGMrIJbwgdqdl2eEH8o+Exx/0m8ir9Gns0u4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_peterbourgon_diskv",
        importpath = "github.com/peterbourgon/diskv",
        sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        importpath = "github.com/Azure/go-ansiterm",
        sum = "h1:w+iIsaOQNcT7OZ575w+acHgRric5iCyQh+xv+KJ4HB8=",
        version = "v0.0.0-20170929234023-d6e3b3328b78",
    )
    go_repository(
        name = "com_github_chai2010_gettext_go",
        importpath = "github.com/chai2010/gettext-go",
        sum = "h1:7aWHqerlJ41y6FOsEUvknqgXnGmJyJSbjhAWq5pO4F8=",
        version = "v0.0.0-20160711120539-c6fed771bfd5",
    )
    go_repository(
        name = "com_github_alexflint_go_filemutex",
        importpath = "github.com/alexflint/go-filemutex",
        sum = "h1:AMzIhMUqU3jMrZiTuW0zkYeKlKDAFD+DG20IoO421/Y=",
        version = "v0.0.0-20171022225611-72bdc8eae2ae",
    )
    go_repository(
        name = "com_github_coreos_go_iptables",
        importpath = "github.com/coreos/go-iptables",
        sum = "h1:DpHb9vJrZQEFMcVLFKAAGMUVX0XoRC0ptCthinRYm38=",
        version = "v0.4.5",
    )
    go_repository(
        name = "com_github_safchain_ethtool",
        importpath = "github.com/safchain/ethtool",
        sum = "h1:2c1EFnZHIPCW8qKWgHMH/fX2PkSabFc5mrVzfUNdg5U=",
        version = "v0.0.0-20190326074333-42ed695e3de8",
    )
    go_repository(
        name = "com_github_j_keck_arping",
        importpath = "github.com/j-keck/arping",
        sum = "h1:742eGXur0715JMq73aD95/FU0XpVKXqNuTnEfXsLOYQ=",
        version = "v0.0.0-20160618110441-2cf9dc699c56",
    )
    go_repository(
        name = "com_github_grpc_grpc",
        importpath = "github.com/grpc/grpc",
        sum = "h1:0/fjvIF5JHJdr34/JPEk1DJFFonjW37pDLvuAy9YieQ=",
        version = "v1.26.0",
    )
    go_repository(
        name = "com_github_google_subcommands",
        importpath = "github.com/google/subcommands",
        sum = "h1:8nlgEAjIalk6uj/CGKCdOO8CQqTeysvcW4RFZ6HbkGM=",
        version = "v1.0.2-0.20190508160503-636abe8753b8",
    )
    go_repository(
        name = "com_github_cenkalti_backoff",
        importpath = "github.com/cenkalti/backoff",
        sum = "h1:8eZxmY1yvxGHzdzTEhI09npjMVGzNAdrqzruTX6jcK4=",
        version = "v1.1.1-0.20190506075156-2146c9339422",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:s5hAObm+yFO5uHYt5dYjxi2rXrsnmRpJx4OYvIWUaQs=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gofrs_flock",
        importpath = "github.com/gofrs/flock",
        sum = "h1:JFTFz3HZTGmgMz4E1TabNBNJljROSYgja1b4l50FNVs=",
        version = "v0.6.1-0.20180915234121-886344bea079",
    )
    go_repository(
        name = "com_github_opencontainers_runc",
        importpath = "github.com/opencontainers/runc",
        sum = "h1:AbmCEuSZXVflng0/cboQkpdEOeBsPMjz6tmq4Pv8MZw=",
        version = "v1.0.0-rc10",
    )
    go_repository(
        name = "com_github_volatiletech_sqlboiler",
        importpath = "github.com/volatiletech/sqlboiler",
        sum = "h1:EI9MwhgyXiqKbvZaomzL8scDs/k+21QAyo9z7GXBZQU=",
        version = "v3.6.1+incompatible",
    )
    go_repository(
        name = "com_github_friendsofgo_errors",
        importpath = "github.com/friendsofgo/errors",
        sum = "h1:X6NYxef4efCBdwI7BgS820zFaN7Cphrmb+Pljdzjtgk=",
        version = "v0.9.2",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        sum = "h1:nFm6S0SMdyzrzcmThSipiEubIDy8WEXKNZ0UOgiRpng=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_volatiletech_null",
        importpath = "github.com/volatiletech/null",
        sum = "h1:7wP8m5d/gZ6kW/9GnrLtMCRre2dlEnaQ9Km5OXlK4zg=",
        version = "v8.0.0+incompatible",
    )
    go_repository(
        name = "com_github_volatiletech_randomize",
        importpath = "github.com/volatiletech/randomize",
        sum = "h1:eE5yajattWqTB2/eN8df4dw+8jwAzBtbdo5sbWC4nMk=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_volatiletech_strmangle",
        importpath = "github.com/volatiletech/strmangle",
        sum = "h1:UKQoHmY6be/R3tSvD2nQYrH41k43OJkidwEiC74KIzk=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_lib_pq",
        importpath = "github.com/lib/pq",
        sum = "h1:LXpIM/LZ5xGFhOpXAQUIMM1HdyqzVYM13zNdjCEEcA0=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_denisenkom_go_mssqldb",
        importpath = "github.com/denisenkom/go-mssqldb",
        sum = "h1:LzwWXEScfcTu7vUZNlDDWDARoSGEtvlDKK2BYHowNeE=",
        version = "v0.0.0-20200206145737-bbfc9a55622e",
    )
    go_repository(
        name = "com_github_ericlagergren_decimal",
        importpath = "github.com/ericlagergren/decimal",
        sum = "h1:HQGCJNlqt1dUs/BhtEKmqWd6LWS+DWYVxi9+Jo4r0jE=",
        version = "v0.0.0-20181231230500-73749d4874d5",
    )
    go_repository(
        name = "com_github_rubenv_sql_migrate",
        importpath = "github.com/rubenv/sql-migrate",
        sum = "h1:lwDYefgiwhjuAuVnMVUYknoF+Yg9CBUykYGvYoPCNnQ=",
        version = "v0.0.0-20191025130928-9355dd04f4b3",
    )
    go_repository(
        name = "com_github_glerchundi_sqlboiler_crdb",
        importpath = "github.com/glerchundi/sqlboiler-crdb",
        sum = "h1:qC5Gu83hkXlQVECFs1u/9EdwYjo58RYaxWnJZGa0YN0=",
        version = "v0.0.0-20191112134934-62014c8c8df1",
    )
    go_repository(
        name = "com_github_mitchellh_cli",
        importpath = "github.com/mitchellh/cli",
        sum = "h1:iGBIsUe3+HZ/AD/Vd7DErOt5sU9fa8Uj7A2s1aggv1Y=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_olekukonko_tablewriter",
        importpath = "github.com/olekukonko/tablewriter",
        sum = "h1:58+kh9C6jJVXYjt8IE48G2eWl6BjwU5Gj0gqY84fy78=",
        version = "v0.0.0-20170122224234-a0225b3f23b5",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        importpath = "github.com/mattn/go-sqlite3",
        sum = "h1:u/x3mp++qUxvYfulZ4HKOvVO0JWhk7HtE8lWhbGz/Do=",
        version = "v1.12.0",
    )
    go_repository(
        name = "in_gopkg_gorp_v1",
        importpath = "gopkg.in/gorp.v1",
        sum = "h1:j3DWlAyGVv8whO7AcIWznQ2Yj7yJkn34B8s63GViAAw=",
        version = "v1.7.2",
    )
    go_repository(
        name = "com_github_volatiletech_inflect",
        importpath = "github.com/volatiletech/inflect",
        sum = "h1:gI4/tqP6lCY5k6Sg+4k9qSoBXmPwG+xXgMpK7jivD4M=",
        version = "v0.0.0-20170731032912-e7201282ae8d",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:UnlwIPBGaTZfPQ6T1IGzPI0EkYAQmT9fAEJ/poFC63o=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_golang_sql_civil",
        importpath = "github.com/golang-sql/civil",
        sum = "h1:lXe2qZdvpiX5WZkZR4hgp4KJVfY3nMkvmwbVkpv1rVY=",
        version = "v0.0.0-20190719163853-cb61b32ac6fe",
    )
    go_repository(
        name = "com_github_posener_complete",
        importpath = "github.com/posener/complete",
        sum = "h1:ccV59UEOTzVDnDUEFdT95ZzHVZ+5+158q8+SJb2QV5w=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:bnP0vzxcAdeI1zdubAl5PjU6zsERjGZb7raWodagDYs=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_armon_go_radix",
        importpath = "github.com/armon/go-radix",
        sum = "h1:BUAU3CGlLvorLI26FmByPp2eC2qla6E1Tw+scpcg/to=",
        version = "v0.0.0-20180808171621-7fddfc383310",
    )
    go_repository(
        name = "com_github_bgentry_speakeasy",
        importpath = "github.com/bgentry/speakeasy",
        sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        sum = "h1:DkWD4oS2D8LGGgTQ6IvwJJXSL5Vp2ffcQg58nFV38Ys=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_spf13_viper",
        importpath = "github.com/spf13/viper",
        sum = "h1:VUFqw5KcqRf7i70GOzW7N+Q7+gxVBkSSqiXB12+JQ4M=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        importpath = "github.com/spf13/jwalterweatherman",
        sum = "h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        sum = "h1:LLgXmsheXeRoUOBOjtwPQCWIYqM/LU1ayDtDePerRcY=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:UVL0vNpWh04HeJXV0KLcaT7r06gOH2l4OW6ddYRUIY4=",
        version = "v0.0.9",
    )
    go_repository(
        name = "com_github_google_go_tpm",
        importpath = "github.com/google/go-tpm",
        sum = "h1:GNNkIb6NSjYfw+KvgUFW590mcgsSFihocSrbXct1sEw=",
        version = "v0.1.2-0.20190725015402-ae6dd98980d4",
    )
    go_repository(
        name = "com_github_google_go_tpm_tools",
        importpath = "github.com/google/go-tpm-tools",
        sum = "h1:1Y5W2uh6E7I6hhI6c0WVSbV+Ae15uhemqi3RvSgtZpk=",
        version = "v0.0.0-20190731025042-f8c04ff88181",
    )
    go_repository(
        name = "com_github_insomniacslk_dhcp",
        importpath = "github.com/insomniacslk/dhcp",
        sum = "h1:P1pxzF2xvdnSY12ODSSwjxA4tyEjDEJNn829OXKnqks=",
        version = "v0.0.0-20200402185128-5dd7202f1971",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:JIufpQLbh4DkbQoii76ItQIUFzevQSqOLZca4eamEDs=",
        version = "v4.0.2",
    )
    go_repository(
        name = "com_github_rekby_gpt",
        importpath = "github.com/rekby/gpt",
        sum = "h1:goZGTwEEn8mWLcY012VouWZWkJ8GrXm9tS3VORMxT90=",
        version = "v0.0.0-20200219180433-a930afbc6edc",
    )
    go_repository(
        name = "com_github_yalue_native_endian",
        importpath = "github.com/yalue/native_endian",
        sum = "h1:nsQCScpQ8RRf+wIooqfyyEUINV2cAPuo2uVtHSBbA4M=",
        version = "v0.0.0-20180607135909-51013b03be4f",
    )
    go_repository(
        name = "com_github_mdlayher_ethernet",
        importpath = "github.com/mdlayher/ethernet",
        sum = "h1:lez6TS6aAau+8wXUP3G9I3TGlmPFEq2CTxBaRqY6AGE=",
        version = "v0.0.0-20190606142754-0394541c37b7",
    )
    go_repository(
        name = "com_github_mdlayher_raw",
        importpath = "github.com/mdlayher/raw",
        sum = "h1:aFkJ6lx4FPip+S+Uw4aTegFMct9shDvP+79PsSxpm3w=",
        version = "v0.0.0-20191009151244-50f2db8cc065",
    )
    go_repository(
        name = "com_github_u_root_u_root",
        importpath = "github.com/u-root/u-root",
        sum = "h1:YqPGmRoRyYmeg17KIWFRSyVq6LX5T6GSzawyA6wG6EE=",
        version = "v6.0.0+incompatible",
    )
    go_repository(
        name = "com_github_diskfs_go_diskfs",
        importpath = "github.com/diskfs/go-diskfs",
        sum = "h1:sLQnXItICiYgiHcYNNujKT9kOKnk7diOvZGEKvxrwpc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "in_gopkg_djherbis_times_v1",
        importpath = "gopkg.in/djherbis/times.v1",
        sum = "h1:UCvDKl1L/fmBygl2Y7hubXCnY7t4Yj46ZrBFNUipFbM=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_kevinburke_go_bindata",
        importpath = "github.com/kevinburke/go-bindata",
        sum = "h1:TFzFZop2KxGhqNwsyjgmIh5JOrpG940MZlm5gNbxr8g=",
        version = "v3.16.0+incompatible",
    )
    go_repository(
        name = "com_github_gofrs_uuid",
        importpath = "github.com/gofrs/uuid",
        sum = "h1:y12jRkkFxsd7GpqdSZ+/KCs/fJbqpEXSGd4+jfEaewE=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star",
        importpath = "github.com/lyft/protoc-gen-star",
        sum = "h1:HUkD4H4dYFIgu3Bns/3N6J5GmKHCEGnhYBwNu3fvXgA=",
        version = "v0.4.14",
    )
    go_repository(
        name = "com_github_moby_term",
        importpath = "github.com/moby/term",
        sum = "h1:aY7OQNf2XqY/JQ6qREWamhI/81os/agb2BAGpcx5yWI=",
        version = "v0.0.0-20200312100748-672ec06f55cd",
    )
    go_repository(
        name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
        importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
        sum = "h1:uuHDyjllyzRyCIvvn0OBjiRB0SgBZGqHNYAmjR7fO50=",
        version = "v0.0.7",
    )
    go_repository(
        name = "io_k8s_sample_apiserver",
        importpath = "k8s.io/sample-apiserver",
        sum = "h1:Nw+rJYx+0cb8Kxtxhe87iT73S6CF67396cIf7tU3JZ8=",
        version = "v0.19.0-alpha.2",
    )
