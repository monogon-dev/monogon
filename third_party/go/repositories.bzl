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
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        version = "v0.51.0",
        sum = "h1:PvKAVQWCtlGUSlZkGW3QLelKaWq7KYv/MW1EboG8bfM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        importpath = "github.com/Azure/go-ansiterm",
        version = "v0.0.0-20170929234023-d6e3b3328b78",
        sum = "h1:w+iIsaOQNcT7OZ575w+acHgRric5iCyQh+xv+KJ4HB8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        version = "v0.3.1",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_makenowjust_heredoc",
        importpath = "github.com/MakeNowJust/heredoc",
        version = "v0.0.0-20170808103936-bb23615498cd",
        sum = "h1:sjQovDkwrZp8u+gxLtPgKGjk5hCxuy2hrRejBTA9xFU=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        importpath = "github.com/Microsoft/go-winio",
        version = "v0.4.14",
        sum = "h1:+hMXMk01us9KgxGb7ftKQt2Xpf5hH/yky+TDA+qxleU=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_microsoft_hcsshim",
        importpath = "github.com/Microsoft/hcsshim",
        version = "v0.8.10",
        sum = "h1:k5wTrpnVU2/xv8ZuzGkbXVd3js5zJ8RnumPo5RxiIxU=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_nytimes_gziphandler",
        importpath = "github.com/NYTimes/gziphandler",
        version = "v0.0.0-20170623195520-56545f4a5d46",
        sum = "h1:lsxEuwrXEAokXB9qhlbKWPpo3KMLZQ5WB5WLQRW1uq0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_puerkitobio_purell",
        importpath = "github.com/PuerkitoBio/purell",
        version = "v1.1.1",
        sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_puerkitobio_urlesc",
        importpath = "github.com/PuerkitoBio/urlesc",
        version = "v0.0.0-20170810143723-de5bf2ad4578",
        sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_alexflint_go_filemutex",
        importpath = "github.com/alexflint/go-filemutex",
        version = "v0.0.0-20171022225611-72bdc8eae2ae",
        sum = "h1:AMzIhMUqU3jMrZiTuW0zkYeKlKDAFD+DG20IoO421/Y=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_armon_circbuf",
        importpath = "github.com/armon/circbuf",
        version = "v0.0.0-20150827004946-bbbad097214e",
        sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_asaskevich_govalidator",
        importpath = "github.com/asaskevich/govalidator",
        version = "v0.0.0-20190424111038-f61b66f89f4a",
        sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        version = "v1.0.1",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_bgentry_speakeasy",
        importpath = "github.com/bgentry/speakeasy",
        version = "v0.1.0",
        sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_blang_semver",
        importpath = "github.com/blang/semver",
        version = "v3.5.0+incompatible",
        sum = "h1:CGxCgetQ64DKk7rdZ++Vfnb1+ogGNnB17OJKJXD2Cfs=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_caddyserver_caddy",
        importpath = "github.com/caddyserver/caddy",
        version = "v1.0.5",
        sum = "h1:5B1Hs0UF2x2tggr2X9jL2qOZtDXbIWQb9YLbmlxHSuM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_cenkalti_backoff",
        importpath = "github.com/cenkalti/backoff",
        version = "v1.1.1-0.20190506075156-2146c9339422",
        sum = "h1:8eZxmY1yvxGHzdzTEhI09npjMVGzNAdrqzruTX6jcK4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        version = "v4.0.2",
        sum = "h1:JIufpQLbh4DkbQoii76ItQIUFzevQSqOLZca4eamEDs=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        version = "v2.1.1",
        sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_chai2010_gettext_go",
        importpath = "github.com/chai2010/gettext-go",
        version = "v0.0.0-20160711120539-c6fed771bfd5",
        sum = "h1:7aWHqerlJ41y6FOsEUvknqgXnGmJyJSbjhAWq5pO4F8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_checkpoint_restore_go_criu_v4",
        importpath = "github.com/checkpoint-restore/go-criu/v4",
        version = "v4.1.0",
        sum = "h1:WW2B2uxx9KWF6bGlHqhm8Okiafwwx7Y2kcpn8lCpjgo=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_cilium_ebpf",
        importpath = "github.com/cilium/ebpf",
        version = "v0.0.0-20200702112145-1c8d4c9ef775",
        sum = "h1:cHzBGGVew0ezFsq2grfy2RsB8hO/eNyBgOLHBCqfR1U=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_container_storage_interface_spec",
        importpath = "github.com/container-storage-interface/spec",
        version = "v1.2.0",
        sum = "h1:bD9KIVgaVKKkQ/UbVUY9kCaH/CJbhNxe0eeB4JeJV2s=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_btrfs",
        importpath = "github.com/containerd/btrfs",
        version = "v0.0.0-20201111183144-404b9149801e",
        sum = "h1:chFw/cg0TDyK43qm8DKbblny2WHc4ML+j1KOkdEp9pI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_cgroups",
        importpath = "github.com/containerd/cgroups",
        version = "v0.0.0-20200710171044-318312a37340",
        sum = "h1:9atoWyI9RtXFwf7UDbme/6M8Ud0rFrx+Q3ZWgSnsxtw=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_console",
        importpath = "github.com/containerd/console",
        version = "v1.0.0",
        sum = "h1:fU3UuQapBs+zLJu82NhR11Rif1ny2zfMMAyPJzSN5tQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_containerd",
        importpath = "github.com/containerd/containerd",
        version = "v1.4.3",
        sum = "h1:ijQT13JedHSHrQGWFcGEwzcNKrAGIiZ+jSD5QQG07SY=",
        build_file_proto_mode = "disable",
        build_tags = [
            "no_zfs",
            "no_aufs",
            "no_devicemapper",
            "no_btrfs",
        ],
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_continuity",
        importpath = "github.com/containerd/continuity",
        version = "v0.0.0-20200710164510-efbc4488d8fe",
        sum = "h1:PEmIrUvwG9Yyv+0WKZqjXfSFDeZjs/q15g0m08BYS9k=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_cri",
        importpath = "github.com/containerd/cri",
        version = "v1.19.1-0.20201126003523-adc0b6a578ed",
        sum = "h1:M2yIwrNSafh4rW/yXAiAlSqpydW7vjvDjZ0ClMb+EMQ=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_fifo",
        importpath = "github.com/containerd/fifo",
        version = "v0.0.0-20200410184934-f15a3290365b",
        sum = "h1:qUtCegLdOUVfVJOw+KDg6eJyE1TGvLlkGEd1091kSSQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_go_cni",
        importpath = "github.com/containerd/go-cni",
        version = "v1.0.1",
        sum = "h1:VXr2EkOPD0v1gu7CKfof6XzEIDzsE/dI9yj/W7PSWLs=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_go_runc",
        importpath = "github.com/containerd/go-runc",
        version = "v0.0.0-20200220073739-7016d3ce2328",
        sum = "h1:PRTagVMbJcCezLcHXe8UJvR1oBzp2lG3CEumeFOLOds=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_imgcrypt",
        importpath = "github.com/containerd/imgcrypt",
        version = "v1.0.1",
        sum = "h1:IyI3IIP4m6zrNFuNFT7HizGVcuD6BYJFpdM1JvPKCbQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_ttrpc",
        importpath = "github.com/containerd/ttrpc",
        version = "v1.0.2-0.20210119122237-222b428f008e",
        sum = "h1:+Fbjfo26pg4HtkAw9sC/YhUwaAb16355o/J/oHkyCDc=",
        replace = "github.com/monogon-dev/ttrpc",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containerd_typeurl",
        importpath = "github.com/containerd/typeurl",
        version = "v1.0.1",
        sum = "h1:PvuK4E3D5S5q6IqsPDCy928FhP0LUIGcmZ/Yhgp5Djw=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containernetworking_cni",
        importpath = "github.com/containernetworking/cni",
        version = "v0.8.0",
        sum = "h1:BT9lpgGoH4jw3lFC7Odz2prU5ruiYKcgAjMCbgybcKI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containernetworking_plugins",
        importpath = "github.com/containernetworking/plugins",
        version = "v0.8.2",
        sum = "h1:5lnwfsAYO+V7yXhysJKy3E1A2Gy9oVut031zfdOzI9w=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_containers_ocicrypt",
        importpath = "github.com/containers/ocicrypt",
        version = "v1.0.1",
        sum = "h1:EToign46OSLTFWnb2oNj9RG3XDnkOX8r28ZIXUuk5Pc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_coredns_coredns",
        importpath = "github.com/coredns/coredns",
        version = "v1.7.0",
        sum = "h1:Tm2ZSdhTk+4okgjUp4K6KYzvBI2u34cdD4fKQRC4Eeo=",
        pre_patches = [
            "//third_party/go/patches:coredns-remove-unused-plugins.patch",
        ],
        patch_args = ["-p1"],
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_coreos_go_iptables",
        importpath = "github.com/coreos/go-iptables",
        version = "v0.4.2",
        sum = "h1:KH0EwId05JwWIfb96gWvkiT2cbuOu8ygqUaB+yPAwIg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_coreos_go_oidc",
        importpath = "github.com/coreos/go-oidc",
        version = "v2.1.0+incompatible",
        sum = "h1:sdJrfw8akMnCuUlaZU3tE/uYXFgfqom8DBE9so9EBsM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        version = "v0.3.0",
        sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_coreos_go_systemd",
        importpath = "github.com/coreos/go-systemd",
        version = "v0.0.0-20190321100706-95778dfbb74e",
        sum = "h1:Wf6HqHfScWJN9/ZjdUKyjop4mf3Qdd+1TvvltAvM3m8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        importpath = "github.com/coreos/go-systemd/v22",
        version = "v22.1.0",
        sum = "h1:kq/SbG2BCKLkDKkjQf5OWwKWUKj1lgs3lFI4PxnR5lg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_coreos_pkg",
        importpath = "github.com/coreos/pkg",
        version = "v0.0.0-20180928190104-399ea9e2e55f",
        sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_cosiner_argv",
        importpath = "github.com/cosiner/argv",
        version = "v0.0.0-20170225145430-13bacc38a0a5",
        sum = "h1:rIXlvz2IWiupMFlC45cZCXZFvKX/ExBcSLrDy2G0Lp8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        version = "v2.0.0",
        sum = "h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_cyphar_filepath_securejoin",
        importpath = "github.com/cyphar/filepath-securejoin",
        version = "v0.2.2",
        sum = "h1:jCwT2GTP+PY5nBz3c/YL5PAIbusElVrPujOBSCj8xRg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        version = "v1.1.1",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_daviddengcn_go_colortext",
        importpath = "github.com/daviddengcn/go-colortext",
        version = "v0.0.0-20160507010035-511bcaf42ccd",
        sum = "h1:uVsMphB1eRx7xB1njzL3fuMdWRN8HtVzoUOItHMwv5c=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_dgrijalva_jwt_go",
        importpath = "github.com/dgrijalva/jwt-go",
        version = "v3.2.0+incompatible",
        sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_diskfs_go_diskfs",
        importpath = "github.com/diskfs/go-diskfs",
        version = "v1.0.0",
        sum = "h1:sLQnXItICiYgiHcYNNujKT9kOKnk7diOvZGEKvxrwpc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_dnstap_golang_dnstap",
        importpath = "github.com/dnstap/golang-dnstap",
        version = "v0.2.0",
        sum = "h1:+NrmP4mkaTeKYV7xJ5FXpUxRn0RpcgoQcsOCTS8WQPk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_docker_distribution",
        importpath = "github.com/docker/distribution",
        version = "v2.7.1+incompatible",
        sum = "h1:a5mlkVzth6W5A4fOsS3D2EO5BUmsJpcB+cRlLU7cSug=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_docker_docker",
        importpath = "github.com/docker/docker",
        version = "v17.12.0-ce-rc1.0.20200310163718-4634ce647cf2+incompatible",
        sum = "h1:ax4NateCD5bjRTqLvQBlFrSUPOoZRgEXWpJ6Bmu6OO0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_docker_go_connections",
        importpath = "github.com/docker/go-connections",
        version = "v0.4.0",
        sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_docker_go_events",
        importpath = "github.com/docker/go-events",
        version = "v0.0.0-20190806004212-e31b211e4f1c",
        sum = "h1:+pKlWGMw7gf6bQ+oDZB4KHQFypsfjYlq/C4rfL7D3g8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_docker_go_metrics",
        importpath = "github.com/docker/go-metrics",
        version = "v0.0.1",
        sum = "h1:AgB/0SvBxihN0X8OR4SjsblXkbMvalQ8cjmtKQ2rQV8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_docker_go_units",
        importpath = "github.com/docker/go-units",
        version = "v0.4.0",
        sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_docker_spdystream",
        importpath = "github.com/docker/spdystream",
        version = "v0.0.0-20160310174837-449fdfce4d96",
        sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        importpath = "github.com/dustin/go-humanize",
        version = "v1.0.0",
        sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_elazarl_goproxy",
        importpath = "github.com/elazarl/goproxy",
        version = "v0.0.0-20180725130230-947c36da3153",
        sum = "h1:yUdfgN0XgIJw7foRItutHYUIhlcKzcSf5vDpdhQAKTc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_emicklei_go_restful",
        importpath = "github.com/emicklei/go-restful",
        version = "v2.9.5+incompatible",
        sum = "h1:spTtZBk5DYEvbxMVutUuTyh1Ao2r4iyvLdACqsl/Ljk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_euank_go_kmsg_parser",
        importpath = "github.com/euank/go-kmsg-parser",
        version = "v2.0.0+incompatible",
        sum = "h1:cHD53+PLQuuQyLZeriD1V/esuG4MuU0Pjs5y6iknohY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_evanphx_json_patch",
        importpath = "github.com/evanphx/json-patch",
        version = "v4.9.0+incompatible",
        sum = "h1:kLcOMZeuLAJvL2BPWLMIj5oaZQobrkAqrL+WFZwQses=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_exponent_io_jsonpath",
        importpath = "github.com/exponent-io/jsonpath",
        version = "v0.0.0-20151013193312-d6023ce2651d",
        sum = "h1:105gxyaGwCFad8crR9dcMQWvV9Hvulu6hwUh4tWPJnM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_farsightsec_golang_framestream",
        importpath = "github.com/farsightsec/golang-framestream",
        version = "v0.0.0-20190425193708-fa4b164d59b8",
        sum = "h1:/iPdQppoAsTfML+yqFSq2EBChiEMnRkh5WvhFgtWwcU=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_fatih_camelcase",
        importpath = "github.com/fatih/camelcase",
        version = "v1.0.0",
        sum = "h1:hxNvNX/xYBp0ovncs8WyWZrOrpBNub/JfaMvbURyft8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        version = "v1.7.0",
        sum = "h1:DkWD4oS2D8LGGgTQ6IvwJJXSL5Vp2ffcQg58nFV38Ys=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_flynn_go_shlex",
        importpath = "github.com/flynn/go-shlex",
        version = "v0.0.0-20150515145356-3f9db97f8568",
        sum = "h1:BHsljHzVlRcyQhjrss6TZTdY2VfCqZPbv5k3iBFa2ZQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_fullsailor_pkcs7",
        importpath = "github.com/fullsailor/pkcs7",
        version = "v0.0.0-20180613152042-8306686428a5",
        sum = "h1:v+vxrd9XS8uWIXG2RK0BHCnXc30qLVQXVqbK+IOmpXk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        version = "v1.0.0",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_delve_delve",
        importpath = "github.com/go-delve/delve",
        version = "v1.4.1",
        sum = "h1:kZs0umEv+VKnK84kY9/ZXWrakdLTeRTyYjFdgLelZCQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        version = "v0.2.0",
        sum = "h1:QvGt2nLcHH0WK9orKa+ppBPAxREcH364nPUedEpK0TY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_analysis",
        importpath = "github.com/go-openapi/analysis",
        version = "v0.19.5",
        sum = "h1:8b2ZgKfKIUTVQpTb77MoRDIMEIwvDVw40o3aOXdfYzI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_errors",
        importpath = "github.com/go-openapi/errors",
        version = "v0.19.2",
        sum = "h1:a2kIyV3w+OS3S97zxUndRVD46+FhGOUBDFY7nmu4CsY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_jsonpointer",
        importpath = "github.com/go-openapi/jsonpointer",
        version = "v0.19.3",
        sum = "h1:gihV7YNZK1iK6Tgwwsxo2rJbD1GTbdm72325Bq8FI3w=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_jsonreference",
        importpath = "github.com/go-openapi/jsonreference",
        version = "v0.19.3",
        sum = "h1:5cxNfTy0UVC3X8JL5ymxzyoUZmo8iZb+jeTWn7tUa8o=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_loads",
        importpath = "github.com/go-openapi/loads",
        version = "v0.19.4",
        sum = "h1:5I4CCSqoWzT+82bBkNIvmLc0UOsoKKQ4Fz+3VxOB7SY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_runtime",
        importpath = "github.com/go-openapi/runtime",
        version = "v0.19.4",
        sum = "h1:csnOgcgAiuGoM/Po7PEpKDoNulCcF3FGbSnbHfxgjMI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_spec",
        importpath = "github.com/go-openapi/spec",
        version = "v0.19.3",
        sum = "h1:0XRyw8kguri6Yw4SxhsQA/atC88yqrk0+G4YhI2wabc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_strfmt",
        importpath = "github.com/go-openapi/strfmt",
        version = "v0.19.3",
        sum = "h1:eRfyY5SkaNJCAwmmMcADjY31ow9+N7MCLW7oRkbsINA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_swag",
        importpath = "github.com/go-openapi/swag",
        version = "v0.19.5",
        sum = "h1:lTz6Ys4CmqqCQmZPBlbQENR1/GucA2bzYTE12Pw4tFY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_openapi_validate",
        importpath = "github.com/go-openapi/validate",
        version = "v0.19.5",
        sum = "h1:QhCBKRYqZR+SKo4gl1lPhPahope8/RLt6EVgY8X80w0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        version = "v1.8.0",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        importpath = "github.com/godbus/dbus/v5",
        version = "v5.0.3",
        sum = "h1:ZqHaoEF7TBzh4jzPmqVhE/5A1z9of6orkAe5uHoAeME=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_gofrs_flock",
        importpath = "github.com/gofrs/flock",
        version = "v0.6.1-0.20180915234121-886344bea079",
        sum = "h1:JFTFz3HZTGmgMz4E1TabNBNJljROSYgja1b4l50FNVs=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_gogo_googleapis",
        importpath = "github.com/gogo/googleapis",
        version = "v1.3.2",
        sum = "h1:kX1es4djPJrsDhY7aZKJy7aZasdcB5oSOEphMjSB53c=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        version = "v1.3.1",
        sum = "h1:DqDEcV5aeaTmdFBePNpYsp3FlcVH/2ISVVM9Qf8PSls=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        version = "v0.0.0-20191227052852-215e87163ea7",
        sum = "h1:5ZkaAPbicIKTF2I64qf5Fh8Aa83Q/dnOafMYV0OMwjA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_btree",
        importpath = "github.com/google/btree",
        version = "v1.0.0",
        sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_cadvisor",
        importpath = "github.com/google/cadvisor",
        version = "v0.37.3",
        sum = "h1:qsH/np74sg1/tEe+bn+e2JIPFxrw6En3gCVuQdolc74=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_certificate_transparency_go",
        importpath = "github.com/google/certificate-transparency-go",
        version = "v1.1.0",
        sum = "h1:10MlrYzh5wfkToxWI4yJzffsxLfxcEDlOATMx/V9Kzw=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        version = "v0.4.0",
        sum = "h1:xsAVV57WRhGj6kEIi8ReJzQlHHqcBYCElAvkovg3B/4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_go_dap",
        importpath = "github.com/google/go-dap",
        version = "v0.2.0",
        sum = "h1:whjIGQRumwbR40qRU7CEKuFLmePUUc2s4Nt9DoXXxWk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_go_tpm",
        importpath = "github.com/google/go-tpm",
        version = "v0.1.2-0.20190725015402-ae6dd98980d4",
        sum = "h1:GNNkIb6NSjYfw+KvgUFW590mcgsSFihocSrbXct1sEw=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_go_tpm_tools",
        importpath = "github.com/google/go-tpm-tools",
        version = "v0.0.0-20190731025042-f8c04ff88181",
        sum = "h1:1Y5W2uh6E7I6hhI6c0WVSbV+Ae15uhemqi3RvSgtZpk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        version = "v1.1.0",
        sum = "h1:Hsa8mG0dQ46ij8Sl2AYJDUv1oA9/d6Vk+3LG99Oe02g=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_gopacket",
        importpath = "github.com/google/gopacket",
        version = "v1.1.17",
        sum = "h1:rMrlX2ZY2UbvT+sdz3+6J+pp2z+msCq9MxTU6ymxbBY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_gvisor",
        importpath = "github.com/google/gvisor",
        version = "v0.0.0-20201216082428-b645fcd241a8",
        sum = "h1:gNssWp0Zg2Ij2OMz4Gi5ciVLnMVGzqfvPOADTN1ou+E=",
        patches = [
            "//third_party/go/patches:gvisor.patch",
            "//third_party/go/patches:gvisor-build-against-newer-runtime-specs.patch",
        ],
        patch_args = ["-p1"],
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_nftables",
        importpath = "github.com/google/nftables",
        version = "v0.0.0-20200316075819-7127d9d22474",
        sum = "h1:D6bN82zzK92ywYsE+Zjca7EHZCRZbcNTU3At7WdxQ+c=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_subcommands",
        importpath = "github.com/google/subcommands",
        version = "v1.0.2-0.20190508160503-636abe8753b8",
        sum = "h1:8nlgEAjIalk6uj/CGKCdOO8CQqTeysvcW4RFZ6HbkGM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        version = "v1.1.1",
        sum = "h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_googleapis_gnostic",
        importpath = "github.com/googleapis/gnostic",
        version = "v0.4.1",
        sum = "h1:DLJCy1n/vrD4HPjOvYcT8aYQXpPIzoRZONaYwyycI+I=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        importpath = "github.com/gorilla/websocket",
        version = "v1.4.0",
        sum = "h1:WDFjx/TMzVgy9VdMMQi2K2Emtwi2QcUQsztZ/zLaH/Q=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_gregjones_httpcache",
        importpath = "github.com/gregjones/httpcache",
        version = "v0.0.0-20180305231024-9cad4c3443a7",
        sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        version = "v1.0.1-0.20190118093823-f849b5445de4",
        sum = "h1:z53tR0945TRRQO/fLEVPI6SMv7ZflF0TEaTAoU7tOzg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        version = "v1.2.0",
        sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        version = "v1.9.5",
        sum = "h1:UImYN5qQ8tuGpGE16ZmjvcTtTw24zw1QAp/SlnNrZhI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_opentracing",
        importpath = "github.com/grpc-ecosystem/grpc-opentracing",
        version = "v0.0.0-20180507213350-8e809c8a8645",
        sum = "h1:MJG/KsmcqMwFAkh8mTnAwhyKoB+sTAnY4CACC110tbU=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_grpc_grpc",
        importpath = "github.com/grpc/grpc",
        version = "v1.29.1",
        sum = "h1:oDOYav2X6WE7espebiQ//iP9N+/gGygUv6XuuyvkFMc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        importpath = "github.com/hashicorp/errwrap",
        version = "v1.0.0",
        sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        importpath = "github.com/hashicorp/go-multierror",
        version = "v1.0.0",
        sum = "h1:iVjPR7a6H0tWELX5NxNe7bYopibicUzc7uPribsnS6o=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        version = "v0.5.3",
        sum = "h1:YPkqC67at8FYaadspW/6uE0COsBxS2656RLEr8Bppgk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        importpath = "github.com/hashicorp/hcl",
        version = "v1.0.0",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_hpcloud_tail",
        importpath = "github.com/hpcloud/tail",
        version = "v1.0.0",
        sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_imdario_mergo",
        importpath = "github.com/imdario/mergo",
        version = "v0.3.7",
        sum = "h1:Y+UAYTZ7gDEuOfhxKWy+dvb5dRQ6rJjFSdX2HZY1/gI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_infobloxopen_go_trees",
        importpath = "github.com/infobloxopen/go-trees",
        version = "v0.0.0-20190313150506-2af4e13f9062",
        sum = "h1:d3VSuNcgTCn21dNMm8g412Fck/XWFmMj4nJhhHT7ZZ0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_insomniacslk_dhcp",
        importpath = "github.com/insomniacslk/dhcp",
        version = "v0.0.0-20200922210017-67c425063dca",
        sum = "h1:zhwTlFGM8ZkD5J/c43IWkxSJQWzhm20QWou8zajbCck=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_j_keck_arping",
        importpath = "github.com/j-keck/arping",
        version = "v0.0.0-20160618110441-2cf9dc699c56",
        sum = "h1:742eGXur0715JMq73aD95/FU0XpVKXqNuTnEfXsLOYQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_joho_godotenv",
        importpath = "github.com/joho/godotenv",
        version = "v1.3.0",
        sum = "h1:Zjp+RcGpHhGlrMbJzXTrZZPrWj+1vfm90La1wgB6Bhc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_jonboulle_clockwork",
        importpath = "github.com/jonboulle/clockwork",
        version = "v0.1.0",
        sum = "h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        version = "v1.1.10",
        sum = "h1:Kz6Cvnvv2wGdaG/V8yMvfkmNiXq9Ya2KUv4rouJJr68=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_karrick_godirwalk",
        importpath = "github.com/karrick/godirwalk",
        version = "v1.7.5",
        sum = "h1:VbzFqwXwNbAZoA6W5odrLr+hKK197CcENcPh6E/gJ0M=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_kevinburke_go_bindata",
        importpath = "github.com/kevinburke/go-bindata",
        version = "v3.16.0+incompatible",
        sum = "h1:TFzFZop2KxGhqNwsyjgmIh5JOrpG940MZlm5gNbxr8g=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_koneu_natend",
        importpath = "github.com/koneu/natend",
        version = "v0.0.0-20150829182554-ec0926ea948d",
        sum = "h1:MFX8DxRnKMY/2M3H61iSsVbo/n3h0MWGmWNN1UViOU0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        version = "v1.0.3",
        sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        version = "v0.1.0",
        sum = "h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        version = "v1.1.4-0.20190131011033-7dc38fb350b1",
        sum = "h1:zc0R6cOw98cMengLA0fvU55mqbnN7sd/tBMLzSejp+M=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_liggitt_tabwriter",
        importpath = "github.com/liggitt/tabwriter",
        version = "v0.0.0-20181228230101-89fcab3d43de",
        sum = "h1:9TO3cAIGXtEhnIaL+V+BEER86oLrvS+kWobKpbJuye0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_lithammer_dedent",
        importpath = "github.com/lithammer/dedent",
        version = "v1.1.0",
        sum = "h1:VNzHMVCBNG1j0fh3OrsFRkVUwStdDArbgBWoPAffktY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        version = "v1.8.1",
        sum = "h1:ZC2Vc7/ZFkGmsVC9KvOjumD+G5lXy2RtTKyzRKO2BQ4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        importpath = "github.com/mailru/easyjson",
        version = "v0.7.0",
        sum = "h1:aizVhC/NAAcKWb+5QsU1iNOZb4Yws5UO2I+aIprQITM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        version = "v0.0.9",
        sum = "h1:UVL0vNpWh04HeJXV0KLcaT7r06gOH2l4OW6ddYRUIY4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        version = "v0.0.4",
        sum = "h1:bnP0vzxcAdeI1zdubAl5PjU6zsERjGZb7raWodagDYs=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        version = "v0.0.2",
        sum = "h1:UnlwIPBGaTZfPQ6T1IGzPI0EkYAQmT9fAEJ/poFC63o=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        version = "v1.0.1",
        sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mdlayher_ethernet",
        importpath = "github.com/mdlayher/ethernet",
        version = "v0.0.0-20190606142754-0394541c37b7",
        sum = "h1:lez6TS6aAau+8wXUP3G9I3TGlmPFEq2CTxBaRqY6AGE=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mdlayher_genetlink",
        importpath = "github.com/mdlayher/genetlink",
        version = "v1.0.0",
        sum = "h1:OoHN1OdyEIkScEmRgxLEe2M9U8ClMytqA5niynLtfj0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mdlayher_netlink",
        importpath = "github.com/mdlayher/netlink",
        version = "v1.1.0",
        sum = "h1:mpdLgm+brq10nI9zM1BpX1kpDbh3NLl3RSnVq6ZSkfg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mdlayher_raw",
        importpath = "github.com/mdlayher/raw",
        version = "v0.0.0-20191009151244-50f2db8cc065",
        sum = "h1:aFkJ6lx4FPip+S+Uw4aTegFMct9shDvP+79PsSxpm3w=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_miekg_dns",
        importpath = "github.com/miekg/dns",
        version = "v1.1.29",
        sum = "h1:xHBEhR+t5RzcFJjBLJlax2daXOrTYtr9z4WdKEfWFzg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mindprince_gonvml",
        importpath = "github.com/mindprince/gonvml",
        version = "v0.0.0-20190828220739-9ebdce4bb989",
        sum = "h1:PS1dLCGtD8bb9RPKJrc8bS7qHL6JnW1CZvwzH9dPoUs=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mistifyio_go_zfs",
        importpath = "github.com/mistifyio/go-zfs",
        version = "v2.1.2-0.20190413222219-f784269be439+incompatible",
        sum = "h1:aKW/4cBs+yK6gpqU3K/oIwk9Q/XICqd3zOX/UFuvqmk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        importpath = "github.com/mitchellh/go-wordwrap",
        version = "v1.0.0",
        sum = "h1:6GlHJ/LTGMrIJbwgdqdl2eEH8o+Exx/0m8ir9Gns0u4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        version = "v1.1.2",
        sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_moby_sys_mountinfo",
        importpath = "github.com/moby/sys/mountinfo",
        version = "v0.1.3",
        sum = "h1:KIrhRO14+AkwKvG/g2yIpNMOUVZ02xNhOw8KY1WsLOI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_moby_term",
        importpath = "github.com/moby/term",
        version = "v0.0.0-20200312100748-672ec06f55cd",
        sum = "h1:aY7OQNf2XqY/JQ6qREWamhI/81os/agb2BAGpcx5yWI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        version = "v1.0.1",
        sum = "h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mohae_deepcopy",
        importpath = "github.com/mohae/deepcopy",
        version = "v0.0.0-20170308212314-bb9b5e7adda9",
        sum = "h1:Sha2bQdoWE5YQPTlJOL31rmce94/tYi113SlFo1xQ2c=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_morikuni_aec",
        importpath = "github.com/morikuni/aec",
        version = "v1.0.0",
        sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mrunalp_fileutils",
        importpath = "github.com/mrunalp/fileutils",
        version = "v0.0.0-20200520151820-abd8a0e76976",
        sum = "h1:aZQToFSLH8ejFeSkTc3r3L4dPImcj7Ib/KgmkQqbGGg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_munnerz_goautoneg",
        importpath = "github.com/munnerz/goautoneg",
        version = "v0.0.0-20191010083416-a7dc8b61c822",
        sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_mxk_go_flowrate",
        importpath = "github.com/mxk/go-flowrate",
        version = "v0.0.0-20140419014527-cca7078d478f",
        sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_olekukonko_tablewriter",
        importpath = "github.com/olekukonko/tablewriter",
        version = "v0.0.0-20170122224234-a0225b3f23b5",
        sum = "h1:58+kh9C6jJVXYjt8IE48G2eWl6BjwU5Gj0gqY84fy78=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_onsi_ginkgo",
        importpath = "github.com/onsi/ginkgo",
        version = "v1.11.0",
        sum = "h1:JAKSXpt1YjtLA7YpPiqO9ss6sNXEsPfSGdwN0UHqzrw=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_onsi_gomega",
        importpath = "github.com/onsi/gomega",
        version = "v1.7.0",
        sum = "h1:XPnZz8VVBHjVsy1vzJmRwIcSwiUO+JFfrv/xGiigmME=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        importpath = "github.com/opencontainers/go-digest",
        version = "v1.0.0",
        sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        importpath = "github.com/opencontainers/image-spec",
        version = "v1.0.1",
        sum = "h1:JMemWkRwHx4Zj+fVxWoMCFm/8sYGGrUVojFA6h/TRcI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_opencontainers_runc",
        importpath = "github.com/opencontainers/runc",
        version = "v1.0.0-rc92",
        sum = "h1:+IczUKCRzDzFDnw99O/PAqrcBBCoRp9xN3cB1SYSNS4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_opencontainers_runtime_spec",
        importpath = "github.com/opencontainers/runtime-spec",
        version = "v1.0.3-0.20200728170252-4d89ac9fbff6",
        sum = "h1:NhsM2gc769rVWDqJvapK37r+7+CBXI8xHhnfnt8uQsg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_opencontainers_selinux",
        importpath = "github.com/opencontainers/selinux",
        version = "v1.6.0",
        sum = "h1:+bIAS/Za3q5FTwWym4fTB0vObnfCf3G/NC7K6Jx62mY=",
        build_tags = [
            "selinux",
        ],
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        importpath = "github.com/opentracing/opentracing-go",
        version = "v1.1.0",
        sum = "h1:pWlfV3Bxv7k65HYwkikxat0+s3pV4bsqf19k25Ur8rU=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_peterbourgon_diskv",
        importpath = "github.com/peterbourgon/diskv",
        version = "v2.0.1+incompatible",
        sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_peterh_liner",
        importpath = "github.com/peterh/liner",
        version = "v0.0.0-20170317030525-88609521dc4b",
        sum = "h1:8uaXtUkxiy+T/zdLWuxa/PG4so0TPZDZfafFNNSaptE=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        version = "v0.9.1",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_pquerna_cachecontrol",
        importpath = "github.com/pquerna/cachecontrol",
        version = "v0.0.0-20171018203845-0dec1b30a021",
        sum = "h1:0XM1XL/OFFJjXsYXlG30spTkV/E9+gmd5GD1w2HE8xM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        version = "v1.6.0",
        sum = "h1:YVPodQOcK15POxhgARIvnDRVpLcuK8mglnMrWfyrw6A=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        version = "v0.2.0",
        sum = "h1:uq5h0d+GuxiXLJLNABMgp2qUWDPiLvgCzz2dUR+/W/M=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        version = "v0.9.1",
        sum = "h1:KOMtN28tlbam3/7ZKEYKHhKoJZYYj3gMH4uc62x7X7U=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        version = "v0.0.11",
        sum = "h1:DhHlBtkHWPYi8O2y31JkK0TF+DGM+51OopZjH/Ia5qI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_rekby_gpt",
        importpath = "github.com/rekby/gpt",
        version = "v0.0.0-20200219180433-a930afbc6edc",
        sum = "h1:goZGTwEEn8mWLcY012VouWZWkJ8GrXm9tS3VORMxT90=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_robfig_cron",
        importpath = "github.com/robfig/cron",
        version = "v1.1.0",
        sum = "h1:jk4/Hud3TTdcrJgUOBgsqrZBarcxl6ADIjSC2iniwLY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_russross_blackfriday",
        importpath = "github.com/russross/blackfriday",
        version = "v1.5.2",
        sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_russross_blackfriday_v2",
        importpath = "github.com/russross/blackfriday/v2",
        version = "v2.0.1",
        sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_safchain_ethtool",
        importpath = "github.com/safchain/ethtool",
        version = "v0.0.0-20190326074333-42ed695e3de8",
        sum = "h1:2c1EFnZHIPCW8qKWgHMH/fX2PkSabFc5mrVzfUNdg5U=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_sbezverk_nfproxy",
        importpath = "github.com/sbezverk/nfproxy",
        version = "v0.0.0-20200514180651-7fac5f39824e",
        sum = "h1:fJ2lHQ7ZUjmgJbvVQ509ioBmrGHcbvlwfjUieExw/dU=",
        patches = [
            "//third_party/go/patches:nfproxy.patch",
        ],
        patch_args = ["-p1"],
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_sbezverk_nftableslib",
        importpath = "github.com/sbezverk/nftableslib",
        version = "v0.0.0-20200402150358-c20bed91f482",
        sum = "h1:k7gEZ/EwJhHDTRXFUZQlE4/p1cmoha7zL7PWCDG3ZHQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_shurcool_sanitized_anchor_name",
        importpath = "github.com/shurcooL/sanitized_anchor_name",
        version = "v1.0.0",
        sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        version = "v1.6.0",
        sum = "h1:UBcNElsrwanuuMsnGSlYmtmgbb23qDR5dG+6X6Oo89I=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_soheilhy_cmux",
        importpath = "github.com/soheilhy/cmux",
        version = "v0.1.4",
        sum = "h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        version = "v1.2.2",
        sum = "h1:5jhuqJyZCZf2JRofRvN/nIFgIWNzPa3/Vz8mYylgbWc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        version = "v1.3.0",
        sum = "h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        version = "v1.0.0",
        sum = "h1:6m/oheQuQ13N9ks4hubMG6BnvwOeaJrqSPLahSnczz8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        importpath = "github.com/spf13/jwalterweatherman",
        version = "v1.1.0",
        sum = "h1:ue6voC5bR5F8YxI5S67j9i582FU4Qvo2bmqnqMYADFk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        version = "v1.0.5",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_spf13_viper",
        importpath = "github.com/spf13/viper",
        version = "v1.4.0",
        sum = "h1:yXHLWeravcrgGyFSyCgdYpXQ9dR9c/WED3pg1RhxqEU=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        version = "v1.4.0",
        sum = "h1:2E4SXV/wtOkTonXsotYi4li6zVWxYlZuYNCXe9XRJyk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_syndtr_gocapability",
        importpath = "github.com/syndtr/gocapability",
        version = "v0.0.0-20180916011248-d98352740cb2",
        sum = "h1:b6uOv7YOFK0TYG7HtkIgExQo+2RdLuwRft63jn2HWj8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_tchap_go_patricia",
        importpath = "github.com/tchap/go-patricia",
        version = "v2.2.6+incompatible",
        sum = "h1:JvoDL7JSoIP2HDE8AbDH3zC8QBPxmzYe32HHy5yQ+Ck=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_tmc_grpc_websocket_proxy",
        importpath = "github.com/tmc/grpc-websocket-proxy",
        version = "v0.0.0-20190109142713-0ad062ec5ee5",
        sum = "h1:LnC5Kc/wtumK+WB441p7ynQJzVuNRJiqddSIE3IlSEQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_u_root_u_root",
        importpath = "github.com/u-root/u-root",
        version = "v7.0.0+incompatible",
        sum = "h1:u+KSS04pSxJGI5E7WE4Bs9+Zd75QjFv+REkjy/aoAc8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_urfave_cli",
        importpath = "github.com/urfave/cli",
        version = "v1.22.1",
        sum = "h1:+mkCCcOFKPnCmVYVcURKps1Xe+3zP90gSYGNfRkjoIY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_vishvananda_netlink",
        importpath = "github.com/vishvananda/netlink",
        version = "v1.1.0",
        sum = "h1:1iyaYNBLmP6L0220aDnYQpo1QEV4t4hJ+xEEhhJH8j0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_vishvananda_netns",
        importpath = "github.com/vishvananda/netns",
        version = "v0.0.0-20200520041808-52d707b772fe",
        sum = "h1:mjAZxE1nh8yvuwhGHpdDqdhtNu2dgbpk93TwoXuk5so=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_willf_bitset",
        importpath = "github.com/willf/bitset",
        version = "v1.1.11",
        sum = "h1:N7Z7E9UvjW+sGsEl7k/SJrvY2reP1A07MrGuCjIOjRE=",
        build_tags = [
            "selinux",
        ],
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_xiang90_probing",
        importpath = "github.com/xiang90/probing",
        version = "v0.0.0-20190116061207-43a291ad63a2",
        sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_github_yalue_native_endian",
        importpath = "github.com/yalue/native_endian",
        version = "v0.0.0-20180607135909-51013b03be4f",
        sum = "h1:nsQCScpQ8RRf+wIooqfyyEUINV2cAPuo2uVtHSBbA4M=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_etcd_go_bbolt",
        importpath = "go.etcd.io/bbolt",
        version = "v1.3.5",
        sum = "h1:XAzx9gjCb0Rxj7EoqcClPD1d5ZBxZJk0jbuoPHenBt0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_etcd_go_etcd",
        importpath = "go.etcd.io/etcd",
        version = "v0.5.0-alpha.5.0.20200819165624-17cef6e3e9d5",
        sum = "h1:Gqga3zA9tdAcfqobUGjSoCob5L3f8Dt5EuOp3ihNZko=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        importpath = "go.mongodb.org/mongo-driver",
        version = "v1.1.2",
        sum = "h1:jxcFYjlkl8xaERsgLo+RNquI0epW6zuy/ZRQs6jnrFA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        version = "v0.22.0",
        sum = "h1:C9hSCOW830chIVkdja34wa6Ky+IzWllkUinR+BtRZd4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "net_starlark_go",
        importpath = "go.starlark.net",
        version = "v0.0.0-20190702223751-32f345186213",
        sum = "h1:lkYv5AKwvvduv5XWP6szk/bvvgO6aDeUujhZQXIFTes=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        version = "v1.4.0",
        sum = "h1:cxzIVoETapQEqDhQu3QfnvXAV4AlzcvUCxkVUFw3+EU=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        version = "v1.1.0",
        sum = "h1:HoEmRHQPVSqub6w2z2d2EOVs2fjyFRGyofhKuyDq0QI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        version = "v1.15.0",
        sum = "h1:ZZCA22JRF2gQE5FoNmhmrf7jeJJ2uhqDUNRYKm8dvmM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_arch",
        importpath = "golang.org/x/arch",
        version = "v0.0.0-20190927153633-4e8777c89be4",
        sum = "h1:QlVATYS7JBoZMVaf+cNjb90WD/beKVHnIxFKT4QaHVI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        version = "v0.0.0-20200622213623-75b288015ac9",
        sum = "h1:psW17arqaxU48Z5kZ0CQnkZWQJsqcURM6tKiBApRjXI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        version = "v0.3.0",
        sum = "h1:RM4zey1++hCTbCVQfnWeKs9/IEsaBLA8vTkd0WVtmH4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        version = "v0.0.0-20201110031124-69a78807bb2b",
        sum = "h1:uwuIcX0g4Yl1NC5XAz37xsr2lTtcqevgzYNVt49waME=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        version = "v0.0.0-20191202225959-858c2ad4c8b6",
        sum = "h1:pE8b58s1HRDMi8RDc79m0HISf9D4TzseP40cEA6IGfs=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        version = "v0.0.0-20181108010431-42b317875d0f",
        sum = "h1:Bl/8QSvNqXvPGPGXa2z5xUTmV7VDcZyvRZ+QQXkXTZQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        version = "v0.3.0",
        sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        version = "v0.0.0-20191024005414-555d28b269f0",
        sum = "h1:/5xXl8Y5W96D+TtHSlonuFqGHIWVuyCkGJLwGh9JJFs=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        version = "v0.0.0-20201215171152-6307297f4651",
        sum = "h1:bdfqbHwYVvhLEIkESR524rqSsmV06Og3Fgz60LE7vZc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        version = "v0.0.0-20191204190536-9bdfabe68543",
        sum = "h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "com_zx2c4_golang_wireguard_wgctrl",
        importpath = "golang.zx2c4.com/wireguard/wgctrl",
        version = "v0.0.0-20200515170644-ec7f26be9d9e",
        sum = "h1:fqDhK9OlzaaiFjnyaAfR9Q1RPKCK7OCTLlHGP9f74Nk=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_gonum_v1_gonum",
        importpath = "gonum.org/v1/gonum",
        version = "v0.6.2",
        sum = "h1:4r+yNT0+8SWcOkXP+63H2zQbN+USnC73cjGUxnDF94Q=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        version = "v0.0.0-20200224152610-e50cd9704f63",
        sum = "h1:YzfoEYWbODU5Fbt37+h7X16BWQbad7Q4S6gclTKFXM8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        version = "v1.29.1",
        sum = "h1:EC2SB8S04d2r73uptxphDSUG+kTKVgjRPF+N3xpxRB4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "in_gopkg_djherbis_times_v1",
        importpath = "gopkg.in/djherbis/times.v1",
        version = "v1.2.0",
        sum = "h1:UCvDKl1L/fmBygl2Y7hubXCnY7t4Yj46ZrBFNUipFbM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "in_gopkg_fsnotify_v1",
        importpath = "gopkg.in/fsnotify.v1",
        version = "v1.4.7",
        sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        importpath = "gopkg.in/inf.v0",
        version = "v0.9.1",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "in_gopkg_natefinch_lumberjack_v2",
        importpath = "gopkg.in/natefinch/lumberjack.v2",
        version = "v2.0.0",
        sum = "h1:1Lc07Kr7qY4U2YPouBjpCLxpiyxIVoxqXgkXLknAOE8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "in_gopkg_square_go_jose_v2",
        importpath = "gopkg.in/square/go-jose.v2",
        version = "v2.2.2",
        sum = "h1:orlkJ3myw8CN1nVQHBFfloD+L3egixIa4FvUP6RosSA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "in_gopkg_tomb_v1",
        importpath = "gopkg.in/tomb.v1",
        version = "v1.0.0-20141024135613-dd632973f1e7",
        sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        version = "v2.2.8",
        sum = "h1:obN1ZagJSUGI0Ek/LBmuj4SNLPfIny3KsKFopxRdj10=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_api",
        importpath = "k8s.io/api",
        version = "v0.19.7",
        sum = "h1:MpHhls03C2pyzoYcpbe4QqYiiZjdvW+tuWq6TbjV14Y=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_apiextensions_apiserver",
        importpath = "k8s.io/apiextensions-apiserver",
        version = "v0.19.7",
        sum = "h1:aV9DANMSCCYBEMbtoT/5oesrtcciQrjy9yqWVtZZL5A=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_apimachinery",
        importpath = "k8s.io/apimachinery",
        version = "v0.19.8-rc.0",
        sum = "h1:/vt04+wL+Y79Qsu8hAo2K4QJA+AKGkJCYmoTTVrUiPQ=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_apiserver",
        importpath = "k8s.io/apiserver",
        version = "v0.19.7",
        sum = "h1:fOOELJ9TNC6DgKL3GUkQLE/EBMLjwBseTstx2eRP61o=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_cli_runtime",
        importpath = "k8s.io/cli-runtime",
        version = "v0.19.7",
        sum = "h1:VkHsqrQYCD6+yBm2k9lOxLJtfo1tmb/TdYIHQ2RSCsY=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_client_go",
        importpath = "k8s.io/client-go",
        version = "v0.19.7",
        sum = "h1:SoJ4mzZ9LyXBGDe8MmpMznw0CwQ1ITWgsmG7GixvhUU=",
        pre_patches = [
            "//third_party/go/patches:k8s-client-go.patch",
        ],
        patch_args = ["-p1"],
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_cloud_provider",
        importpath = "k8s.io/cloud-provider",
        version = "v0.19.7",
        sum = "h1:01fiPTLkTU/MNKZBcMmeYQ5DWqRS4d3GhYGGGlkjgOw=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_cluster_bootstrap",
        importpath = "k8s.io/cluster-bootstrap",
        version = "v0.19.7",
        sum = "h1:xlI+YfeS5gOVa33WVh1viiPZMDN9j7BAiY0iJkg2LwI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_component_base",
        importpath = "k8s.io/component-base",
        version = "v0.19.7",
        sum = "h1:ZXS2VRWOWBOc2fTd1zjzhi/b/mkqFT9FDqiNsn1cH30=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_cri_api",
        importpath = "k8s.io/cri-api",
        version = "v0.19.8-rc.0",
        sum = "h1:aXNNIIoVcmIB/mlz/otcULQOgnErxnLB4uaWENHKblA=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_csi_translation_lib",
        importpath = "k8s.io/csi-translation-lib",
        version = "v0.19.7",
        sum = "h1:Spr0XWqXufEUQA47axmPTm1xOabdMYG9MUbJVaRRb0g=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_gengo",
        importpath = "k8s.io/gengo",
        version = "v0.0.0-20200428234225-8167cfdcfc14",
        sum = "h1:t4L10Qfx/p7ASH3gXCdIUtPbbIuegCoUJf3TMSFekjw=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_heapster",
        importpath = "k8s.io/heapster",
        version = "v1.2.0-beta.1",
        sum = "h1:lUsE/AHOMHpi3MLlBEkaU8Esxm5QhdyCrv1o7ot0s84=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_klog_v2",
        importpath = "k8s.io/klog/v2",
        version = "v2.2.0",
        sum = "h1:XRvcwJozkgZ1UQJmfMGpvRthQHOvihEhYtDfAaxMz/A=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_kube_aggregator",
        importpath = "k8s.io/kube-aggregator",
        version = "v0.19.7",
        sum = "h1:Eol5vPNFKaDScdVuTh0AofhuSr4cJxP5Vfv8JXW8OAQ=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_kube_controller_manager",
        importpath = "k8s.io/kube-controller-manager",
        version = "v0.19.7",
        sum = "h1:3rNXjHM5LHcv2HiO2JjdV4yW3EN+2tCPaKXWL/Cl8TM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_kube_openapi",
        importpath = "k8s.io/kube-openapi",
        version = "v0.0.0-20200805222855-6aeccd4b50c6",
        sum = "h1:+WnxoVtG8TMiudHBSEtrVL1egv36TkkJm+bA8AxicmQ=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_kube_proxy",
        importpath = "k8s.io/kube-proxy",
        version = "v0.19.7",
        sum = "h1:QQUwEnHA1jawodclndlmK/6Ifc9XVNlUaQ4Vq5RVbI8=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_kube_scheduler",
        importpath = "k8s.io/kube-scheduler",
        version = "v0.19.7",
        sum = "h1:TlQFoH7rATVqU7myNZ4FBgnXdGIwR7iBBNk3ir8Y9WM=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_kubectl",
        importpath = "k8s.io/kubectl",
        version = "v0.19.7",
        sum = "h1:pSsha+MBr9KLhn0IKrRikeAZ7g2oeShIGHLgqAzE3Ak=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_kubelet",
        importpath = "k8s.io/kubelet",
        version = "v0.19.7",
        sum = "h1:cPp0fXN99cxyXeoI3nG2ZBORUvR0liT+bg6ofCybJzw=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_kubernetes",
        importpath = "k8s.io/kubernetes",
        version = "v1.19.7",
        sum = "h1:Yk9W5SL1KR2mwy0nNZwjFXNImfK7ihrbKhXttidNTiE=",
        build_file_proto_mode = "disable",
        build_tags = [
            "providerless",
        ],
        patches = [
            "//third_party/go/patches:k8s-kubernetes.patch",
            "//third_party/go/patches:k8s-kubernetes-build.patch",
            "//third_party/go/patches:k8s-native-metrics.patch",
            "//third_party/go/patches:k8s-use-native.patch",
            "//third_party/go/patches:k8s-revert-seccomp-runtime-default.patch",
        ],
        pre_patches = [
            "//third_party/go/patches:k8s-e2e-tests-providerless.patch",
            "//third_party/go/patches:k8s-fix-paths.patch",
            "//third_party/go/patches:k8s-fix-logs-path.patch",
        ],
        patch_args = ["-p1"],
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_legacy_cloud_providers",
        importpath = "k8s.io/legacy-cloud-providers",
        version = "v0.19.7",
        sum = "h1:YJ/l/8/Hn56I9m1cudK8aNypRA/NvI/hYhg8fo/CTus=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_metrics",
        importpath = "k8s.io/metrics",
        version = "v0.19.7",
        sum = "h1:fpTtFhNtS0DwJiYGGsL4YoSjHlLw8qugkgw3EXSWaUA=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_repo_infra",
        importpath = "k8s.io/repo-infra",
        version = "v0.1.4-0.20210105022653-a3483874bd37",
        sum = "h1:0GPavEcPKBA0rYl7f6dO0mXYmx7t9RaXD3be2g23Ps4=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_sample_apiserver",
        importpath = "k8s.io/sample-apiserver",
        version = "v0.19.7",
        sum = "h1:ZWD6dsvqpqhWj3jKRb19/m/bo/0r+TRgjkX+h5m7f4g=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_utils",
        importpath = "k8s.io/utils",
        version = "v0.0.0-20200729134348-d5654de09c73",
        sum = "h1:uJmqzgNWG7XyClnU/mLPBWwfKKF1K8Hf8whTseBgJcg=",
        patches = [
            "//third_party/go/patches:k8s-native-mounter.patch",
        ],
        patch_args = ["-p1"],
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
        importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
        version = "v0.0.9",
        sum = "h1:rusRLrDhjBp6aYtl9sGEvQJr6faoHoDLd0YcUBTZguI=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_sigs_kustomize",
        importpath = "sigs.k8s.io/kustomize",
        version = "v2.0.3+incompatible",
        sum = "h1:JUufWFNlI44MdtnjUqVnvh29rR37PQFzPbLXqhyOyX0=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_sigs_structured_merge_diff_v4",
        importpath = "sigs.k8s.io/structured-merge-diff/v4",
        version = "v4.0.1",
        sum = "h1:YXTMot5Qz/X1iBRJhAt+vI+HVttY0WkSqqhKxQ0xVbA=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "io_k8s_sigs_yaml",
        importpath = "sigs.k8s.io/yaml",
        version = "v1.2.0",
        sum = "h1:kr/MCeFWJWTwyaHoR9c8EjH9OumOmoF9YGiZd7lFm/Q=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
    go_repository(
        name = "ml_vbom_util",
        importpath = "vbom.ml/util",
        version = "v0.0.0-20160121211510-db5cfe13f5cc",
        sum = "h1:MksmcCZQWAQJCTA5T0jgI/0sJ51AVm4Z41MrmfczEoc=",
        build_extra_args = [
            "-go_naming_convention=go_default_library",
            "-go_naming_convention_external=go_default_library",
        ],
    )
