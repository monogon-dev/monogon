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
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        importpath = "github.com/Azure/go-ansiterm",
        version = "v0.0.0-20170929234023-d6e3b3328b78",
        sum = "h1:w+iIsaOQNcT7OZ575w+acHgRric5iCyQh+xv+KJ4HB8=",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        version = "v0.3.1",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    )
    go_repository(
        name = "com_github_makenowjust_heredoc",
        importpath = "github.com/MakeNowJust/heredoc",
        version = "v0.0.0-20170808103936-bb23615498cd",
        sum = "h1:sjQovDkwrZp8u+gxLtPgKGjk5hCxuy2hrRejBTA9xFU=",
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        importpath = "github.com/Microsoft/go-winio",
        version = "v0.4.14",
        sum = "h1:+hMXMk01us9KgxGb7ftKQt2Xpf5hH/yky+TDA+qxleU=",
    )
    go_repository(
        name = "com_github_microsoft_hcsshim",
        importpath = "github.com/Microsoft/hcsshim",
        version = "v0.8.9",
        sum = "h1:VrfodqvztU8YSOvygU+DN1BGaSGxmrNfqOv5oOuX2Bk=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "com_github_nytimes_gziphandler",
        importpath = "github.com/NYTimes/gziphandler",
        version = "v0.0.0-20170623195520-56545f4a5d46",
        sum = "h1:lsxEuwrXEAokXB9qhlbKWPpo3KMLZQ5WB5WLQRW1uq0=",
    )
    go_repository(
        name = "com_github_puerkitobio_purell",
        importpath = "github.com/PuerkitoBio/purell",
        version = "v1.1.1",
        sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
    )
    go_repository(
        name = "com_github_puerkitobio_urlesc",
        importpath = "github.com/PuerkitoBio/urlesc",
        version = "v0.0.0-20170810143723-de5bf2ad4578",
        sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
    )
    go_repository(
        name = "com_github_alexflint_go_filemutex",
        importpath = "github.com/alexflint/go-filemutex",
        version = "v0.0.0-20171022225611-72bdc8eae2ae",
        sum = "h1:AMzIhMUqU3jMrZiTuW0zkYeKlKDAFD+DG20IoO421/Y=",
    )
    go_repository(
        name = "com_github_armon_circbuf",
        importpath = "github.com/armon/circbuf",
        version = "v0.0.0-20150827004946-bbbad097214e",
        sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
    )
    go_repository(
        name = "com_github_armon_go_metrics",
        importpath = "github.com/armon/go-metrics",
        version = "v0.0.0-20180917152333-f0300d1749da",
        sum = "h1:8GUt8eRujhVEGZFFEjBj46YV4rDjvGrNxb0KMWYkL2I=",
    )
    go_repository(
        name = "com_github_armon_go_radix",
        importpath = "github.com/armon/go-radix",
        version = "v0.0.0-20180808171621-7fddfc383310",
        sum = "h1:BUAU3CGlLvorLI26FmByPp2eC2qla6E1Tw+scpcg/to=",
    )
    go_repository(
        name = "com_github_asaskevich_govalidator",
        importpath = "github.com/asaskevich/govalidator",
        version = "v0.0.0-20190424111038-f61b66f89f4a",
        sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        version = "v1.0.1",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    )
    go_repository(
        name = "com_github_bgentry_speakeasy",
        importpath = "github.com/bgentry/speakeasy",
        version = "v0.1.0",
        sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
    )
    go_repository(
        name = "com_github_blang_semver",
        importpath = "github.com/blang/semver",
        version = "v3.5.0+incompatible",
        sum = "h1:CGxCgetQ64DKk7rdZ++Vfnb1+ogGNnB17OJKJXD2Cfs=",
    )
    go_repository(
        name = "com_github_c9s_goprocinfo",
        importpath = "github.com/c9s/goprocinfo",
        version = "v0.0.0-20190309065803-0b2ad9ac246b",
        sum = "h1:4yfM1Zm+7U+m0inJ0g6JvdqGePXD8eG4nXUTbcLT6gk=",
    )
    go_repository(
        name = "com_github_cenkalti_backoff",
        importpath = "github.com/cenkalti/backoff",
        version = "v0.0.0-20190506075156-2146c9339422",
        sum = "h1:+FKjzBIdfBHYDvxCv+djmDJdes/AoDtg8gpcxowBlF8=",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        version = "v4.0.2",
        sum = "h1:JIufpQLbh4DkbQoii76ItQIUFzevQSqOLZca4eamEDs=",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        version = "v0.2.1",
        sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
        build_file_proto_mode = "disable",
        build_extra_args = [
            "-exclude=src",
        ],
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        version = "v2.1.1",
        sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
    )
    go_repository(
        name = "com_github_chai2010_gettext_go",
        importpath = "github.com/chai2010/gettext-go",
        version = "v0.0.0-20160711120539-c6fed771bfd5",
        sum = "h1:7aWHqerlJ41y6FOsEUvknqgXnGmJyJSbjhAWq5pO4F8=",
    )
    go_repository(
        name = "com_github_checkpoint_restore_go_criu_v4",
        importpath = "github.com/checkpoint-restore/go-criu/v4",
        version = "v4.0.2",
        sum = "h1:jt+rnBIhFtPw0fhtpYGcUOilh4aO9Hj7r+YLEtf30uA=",
    )
    go_repository(
        name = "com_github_cilium_arping",
        importpath = "github.com/cilium/arping",
        version = "v1.0.1-0.20190728065459-c5eaf8d7a710",
        sum = "h1:htVjkajqUYy6JmLMGlZYxfZ4urQq7rDvgUfmSJX7fSg=",
    )
    go_repository(
        name = "com_github_cilium_cilium",
        importpath = "github.com/cilium/cilium",
        version = "v1.8.0-rc1",
        sum = "h1:tbMNmz8RjjnZ1LHJ8D88mHeQcwEr0aW6eqaratxspu8=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "com_github_cilium_ebpf",
        importpath = "github.com/cilium/ebpf",
        version = "v0.0.0-20200702112145-1c8d4c9ef775",
        sum = "h1:cHzBGGVew0ezFsq2grfy2RsB8hO/eNyBgOLHBCqfR1U=",
    )
    go_repository(
        name = "com_github_cilium_ipam",
        importpath = "github.com/cilium/ipam",
        version = "v0.0.0-20200420133938-2f672ef3ad54",
        sum = "h1:YOrdErbkc+X+6wflk5idOHZ1IJtLNr3Vnz8JlznG0VI=",
    )
    go_repository(
        name = "com_github_cilium_proxy",
        importpath = "github.com/cilium/proxy",
        version = "v0.0.0-20200309181938-3cf80fe45d03",
        sum = "h1:vkRt49aGUyDbrmR8lVXWUPhS9uYvUZB+jwXyer9aq0w=",
        build_file_proto_mode = "disable",
        build_file_generation = "on",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        version = "v0.0.0-20191230090109-edbea6a78f6d",
        sum = "h1:F6x9XOn7D+HmM4z8vuG/vvlE53rWPWebGLdIy3Nh+XM=",
    )
    go_repository(
        name = "com_github_container_storage_interface_spec",
        importpath = "github.com/container-storage-interface/spec",
        version = "v1.2.0",
        sum = "h1:bD9KIVgaVKKkQ/UbVUY9kCaH/CJbhNxe0eeB4JeJV2s=",
    )
    go_repository(
        name = "com_github_containerd_btrfs",
        importpath = "github.com/containerd/btrfs",
        version = "v0.0.0-20200117014249-153935315f4a",
        sum = "h1:u5X1yvVEsXLcuTWYsFSpTgQKRvo2VTB5gOHcERpF9ZI=",
    )
    go_repository(
        name = "com_github_containerd_cgroups",
        importpath = "github.com/containerd/cgroups",
        version = "v0.0.0-20200710171044-318312a37340",
        sum = "h1:9atoWyI9RtXFwf7UDbme/6M8Ud0rFrx+Q3ZWgSnsxtw=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "com_github_containerd_console",
        importpath = "github.com/containerd/console",
        version = "v1.0.0",
        sum = "h1:fU3UuQapBs+zLJu82NhR11Rif1ny2zfMMAyPJzSN5tQ=",
    )
    go_repository(
        name = "com_github_containerd_containerd",
        importpath = "github.com/containerd/containerd",
        version = "v1.4.0-beta.2",
        sum = "h1:qZelipNh4yeTHIyzcNteRPoo/Mb9sFCrDtCNWWSXJHQ=",
        build_file_proto_mode = "disable",
        build_tags = [
            "no_zfs",
            "no_aufs",
            "no_devicemapper",
            "no_btrfs",
        ],
    )
    go_repository(
        name = "com_github_containerd_continuity",
        importpath = "github.com/containerd/continuity",
        version = "v0.0.0-20200413184840-d3ef23f19fbb",
        sum = "h1:nXPkFq8X1a9ycY3GYQpFNxHh3j2JgY7zDZfq2EXMIzk=",
    )
    go_repository(
        name = "com_github_containerd_cri",
        importpath = "github.com/containerd/cri",
        version = "v1.11.1-0.20200705100038-8fb244a65baa",
        sum = "h1:qqB+Jjek9F6LdsEzQwYWu3PmKkWvFyPr8eCUZPUfCoU=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "com_github_containerd_fifo",
        importpath = "github.com/containerd/fifo",
        version = "v0.0.0-20200410184934-f15a3290365b",
        sum = "h1:qUtCegLdOUVfVJOw+KDg6eJyE1TGvLlkGEd1091kSSQ=",
    )
    go_repository(
        name = "com_github_containerd_go_cni",
        importpath = "github.com/containerd/go-cni",
        version = "v1.0.0",
        sum = "h1:A681A9YQ5Du9V2/gZGk/pTm6g69wF0aGd9qFN9syB1E=",
    )
    go_repository(
        name = "com_github_containerd_go_runc",
        importpath = "github.com/containerd/go-runc",
        version = "v0.0.0-20200220073739-7016d3ce2328",
        sum = "h1:PRTagVMbJcCezLcHXe8UJvR1oBzp2lG3CEumeFOLOds=",
    )
    go_repository(
        name = "com_github_containerd_imgcrypt",
        importpath = "github.com/containerd/imgcrypt",
        version = "v1.0.1",
        sum = "h1:IyI3IIP4m6zrNFuNFT7HizGVcuD6BYJFpdM1JvPKCbQ=",
    )
    go_repository(
        name = "com_github_containerd_ttrpc",
        importpath = "github.com/containerd/ttrpc",
        version = "v1.0.1",
        sum = "h1:IfVOxKbjyBn9maoye2JN95pgGYOmPkQVqxtOu7rtNIc=",
    )
    go_repository(
        name = "com_github_containerd_typeurl",
        importpath = "github.com/containerd/typeurl",
        version = "v1.0.1",
        sum = "h1:PvuK4E3D5S5q6IqsPDCy928FhP0LUIGcmZ/Yhgp5Djw=",
    )
    go_repository(
        name = "com_github_containernetworking_cni",
        importpath = "github.com/containernetworking/cni",
        version = "v0.7.1",
        sum = "h1:fE3r16wpSEyaqY4Z4oFrLMmIGfBYIKpPrHK31EJ9FzE=",
    )
    go_repository(
        name = "com_github_containernetworking_plugins",
        importpath = "github.com/containernetworking/plugins",
        version = "v0.8.2",
        sum = "h1:5lnwfsAYO+V7yXhysJKy3E1A2Gy9oVut031zfdOzI9w=",
    )
    go_repository(
        name = "com_github_containers_ocicrypt",
        importpath = "github.com/containers/ocicrypt",
        version = "v1.0.1",
        sum = "h1:EToign46OSLTFWnb2oNj9RG3XDnkOX8r28ZIXUuk5Pc=",
    )
    go_repository(
        name = "com_github_coreos_go_iptables",
        importpath = "github.com/coreos/go-iptables",
        version = "v0.4.2",
        sum = "h1:KH0EwId05JwWIfb96gWvkiT2cbuOu8ygqUaB+yPAwIg=",
    )
    go_repository(
        name = "com_github_coreos_go_oidc",
        importpath = "github.com/coreos/go-oidc",
        version = "v2.1.0+incompatible",
        sum = "h1:sdJrfw8akMnCuUlaZU3tE/uYXFgfqom8DBE9so9EBsM=",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        version = "v0.3.0",
        sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
    )
    go_repository(
        name = "com_github_coreos_go_systemd",
        importpath = "github.com/coreos/go-systemd",
        version = "v0.0.0-20190321100706-95778dfbb74e",
        sum = "h1:Wf6HqHfScWJN9/ZjdUKyjop4mf3Qdd+1TvvltAvM3m8=",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        importpath = "github.com/coreos/go-systemd/v22",
        version = "v22.0.0",
        sum = "h1:XJIw/+VlJ+87J+doOxznsAWIdmWuViOVhkQamW5YV28=",
    )
    go_repository(
        name = "com_github_coreos_pkg",
        importpath = "github.com/coreos/pkg",
        version = "v0.0.0-20180928190104-399ea9e2e55f",
        sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
    )
    go_repository(
        name = "com_github_cosiner_argv",
        importpath = "github.com/cosiner/argv",
        version = "v0.0.0-20170225145430-13bacc38a0a5",
        sum = "h1:rIXlvz2IWiupMFlC45cZCXZFvKX/ExBcSLrDy2G0Lp8=",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        version = "v2.0.0",
        sum = "h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=",
    )
    go_repository(
        name = "com_github_cyphar_filepath_securejoin",
        importpath = "github.com/cyphar/filepath-securejoin",
        version = "v0.2.2",
        sum = "h1:jCwT2GTP+PY5nBz3c/YL5PAIbusElVrPujOBSCj8xRg=",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        version = "v1.1.1",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    )
    go_repository(
        name = "com_github_daviddengcn_go_colortext",
        importpath = "github.com/daviddengcn/go-colortext",
        version = "v0.0.0-20160507010035-511bcaf42ccd",
        sum = "h1:uVsMphB1eRx7xB1njzL3fuMdWRN8HtVzoUOItHMwv5c=",
    )
    go_repository(
        name = "com_github_denisenkom_go_mssqldb",
        importpath = "github.com/denisenkom/go-mssqldb",
        version = "v0.0.0-20200206145737-bbfc9a55622e",
        sum = "h1:LzwWXEScfcTu7vUZNlDDWDARoSGEtvlDKK2BYHowNeE=",
    )
    go_repository(
        name = "com_github_dgrijalva_jwt_go",
        importpath = "github.com/dgrijalva/jwt-go",
        version = "v3.2.0+incompatible",
        sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
    )
    go_repository(
        name = "com_github_diskfs_go_diskfs",
        importpath = "github.com/diskfs/go-diskfs",
        version = "v1.0.0",
        sum = "h1:sLQnXItICiYgiHcYNNujKT9kOKnk7diOvZGEKvxrwpc=",
    )
    go_repository(
        name = "com_github_docker_distribution",
        importpath = "github.com/docker/distribution",
        version = "v2.7.1+incompatible",
        sum = "h1:a5mlkVzth6W5A4fOsS3D2EO5BUmsJpcB+cRlLU7cSug=",
    )
    go_repository(
        name = "com_github_docker_docker",
        importpath = "github.com/docker/docker",
        version = "v17.12.0-ce-rc1.0.20200310163718-4634ce647cf2+incompatible",
        sum = "h1:ax4NateCD5bjRTqLvQBlFrSUPOoZRgEXWpJ6Bmu6OO0=",
    )
    go_repository(
        name = "com_github_docker_go_connections",
        importpath = "github.com/docker/go-connections",
        version = "v0.4.0",
        sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
    )
    go_repository(
        name = "com_github_docker_go_events",
        importpath = "github.com/docker/go-events",
        version = "v0.0.0-20190806004212-e31b211e4f1c",
        sum = "h1:+pKlWGMw7gf6bQ+oDZB4KHQFypsfjYlq/C4rfL7D3g8=",
    )
    go_repository(
        name = "com_github_docker_go_metrics",
        importpath = "github.com/docker/go-metrics",
        version = "v0.0.1",
        sum = "h1:AgB/0SvBxihN0X8OR4SjsblXkbMvalQ8cjmtKQ2rQV8=",
    )
    go_repository(
        name = "com_github_docker_go_units",
        importpath = "github.com/docker/go-units",
        version = "v0.4.0",
        sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
    )
    go_repository(
        name = "com_github_docker_spdystream",
        importpath = "github.com/docker/spdystream",
        version = "v0.0.0-20160310174837-449fdfce4d96",
        sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        importpath = "github.com/dustin/go-humanize",
        version = "v1.0.0",
        sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
    )
    go_repository(
        name = "com_github_emicklei_go_restful",
        importpath = "github.com/emicklei/go-restful",
        version = "v2.9.5+incompatible",
        sum = "h1:spTtZBk5DYEvbxMVutUuTyh1Ao2r4iyvLdACqsl/Ljk=",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        version = "v0.3.0-java",
        sum = "h1:bV5JGEB1ouEzZa0hgVDFFiClrUEuGWRaAc/3mxR2QK0=",
    )
    go_repository(
        name = "com_github_ericlagergren_decimal",
        importpath = "github.com/ericlagergren/decimal",
        version = "v0.0.0-20181231230500-73749d4874d5",
        sum = "h1:HQGCJNlqt1dUs/BhtEKmqWd6LWS+DWYVxi9+Jo4r0jE=",
    )
    go_repository(
        name = "com_github_euank_go_kmsg_parser",
        importpath = "github.com/euank/go-kmsg-parser",
        version = "v2.0.0+incompatible",
        sum = "h1:cHD53+PLQuuQyLZeriD1V/esuG4MuU0Pjs5y6iknohY=",
    )
    go_repository(
        name = "com_github_evanphx_json_patch",
        importpath = "github.com/evanphx/json-patch",
        version = "v0.0.0-20190815234213-e83c0a1c26c8",
        sum = "h1:DM7gHzQfHwIj+St8zaPOI6iQEPAxOwIkskvw6s9rDaM=",
    )
    go_repository(
        name = "com_github_exponent_io_jsonpath",
        importpath = "github.com/exponent-io/jsonpath",
        version = "v0.0.0-20151013193312-d6023ce2651d",
        sum = "h1:105gxyaGwCFad8crR9dcMQWvV9Hvulu6hwUh4tWPJnM=",
    )
    go_repository(
        name = "com_github_fatih_camelcase",
        importpath = "github.com/fatih/camelcase",
        version = "v1.0.0",
        sum = "h1:hxNvNX/xYBp0ovncs8WyWZrOrpBNub/JfaMvbURyft8=",
    )
    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        version = "v1.7.0",
        sum = "h1:DkWD4oS2D8LGGgTQ6IvwJJXSL5Vp2ffcQg58nFV38Ys=",
    )
    go_repository(
        name = "com_github_friendsofgo_errors",
        importpath = "github.com/friendsofgo/errors",
        version = "v0.9.2",
        sum = "h1:X6NYxef4efCBdwI7BgS820zFaN7Cphrmb+Pljdzjtgk=",
    )
    go_repository(
        name = "com_github_fullsailor_pkcs7",
        importpath = "github.com/fullsailor/pkcs7",
        version = "v0.0.0-20180613152042-8306686428a5",
        sum = "h1:v+vxrd9XS8uWIXG2RK0BHCnXc30qLVQXVqbK+IOmpXk=",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        version = "v1.0.0",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    )
    go_repository(
        name = "com_github_glerchundi_sqlboiler_crdb_v4",
        importpath = "github.com/glerchundi/sqlboiler-crdb/v4",
        version = "v4.0.0-20200507103349-d540ee52783e",
        sum = "h1:p1FS4Qf4pgi/ntt4XI1n1rM7EAmOQLRmvPYNXtNlxWA=",
    )
    go_repository(
        name = "com_github_go_delve_delve",
        importpath = "github.com/go-delve/delve",
        version = "v1.4.1",
        sum = "h1:kZs0umEv+VKnK84kY9/ZXWrakdLTeRTyYjFdgLelZCQ=",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        version = "v0.2.0",
        sum = "h1:QvGt2nLcHH0WK9orKa+ppBPAxREcH364nPUedEpK0TY=",
    )
    go_repository(
        name = "com_github_go_openapi_analysis",
        importpath = "github.com/go-openapi/analysis",
        version = "v0.19.5",
        sum = "h1:8b2ZgKfKIUTVQpTb77MoRDIMEIwvDVw40o3aOXdfYzI=",
    )
    go_repository(
        name = "com_github_go_openapi_errors",
        importpath = "github.com/go-openapi/errors",
        version = "v0.19.2",
        sum = "h1:a2kIyV3w+OS3S97zxUndRVD46+FhGOUBDFY7nmu4CsY=",
    )
    go_repository(
        name = "com_github_go_openapi_jsonpointer",
        importpath = "github.com/go-openapi/jsonpointer",
        version = "v0.19.3",
        sum = "h1:gihV7YNZK1iK6Tgwwsxo2rJbD1GTbdm72325Bq8FI3w=",
    )
    go_repository(
        name = "com_github_go_openapi_jsonreference",
        importpath = "github.com/go-openapi/jsonreference",
        version = "v0.19.3",
        sum = "h1:5cxNfTy0UVC3X8JL5ymxzyoUZmo8iZb+jeTWn7tUa8o=",
    )
    go_repository(
        name = "com_github_go_openapi_loads",
        importpath = "github.com/go-openapi/loads",
        version = "v0.19.4",
        sum = "h1:5I4CCSqoWzT+82bBkNIvmLc0UOsoKKQ4Fz+3VxOB7SY=",
    )
    go_repository(
        name = "com_github_go_openapi_runtime",
        importpath = "github.com/go-openapi/runtime",
        version = "v0.19.4",
        sum = "h1:csnOgcgAiuGoM/Po7PEpKDoNulCcF3FGbSnbHfxgjMI=",
    )
    go_repository(
        name = "com_github_go_openapi_spec",
        importpath = "github.com/go-openapi/spec",
        version = "v0.19.3",
        sum = "h1:0XRyw8kguri6Yw4SxhsQA/atC88yqrk0+G4YhI2wabc=",
    )
    go_repository(
        name = "com_github_go_openapi_strfmt",
        importpath = "github.com/go-openapi/strfmt",
        version = "v0.19.3",
        sum = "h1:eRfyY5SkaNJCAwmmMcADjY31ow9+N7MCLW7oRkbsINA=",
    )
    go_repository(
        name = "com_github_go_openapi_swag",
        importpath = "github.com/go-openapi/swag",
        version = "v0.19.5",
        sum = "h1:lTz6Ys4CmqqCQmZPBlbQENR1/GucA2bzYTE12Pw4tFY=",
    )
    go_repository(
        name = "com_github_go_openapi_validate",
        importpath = "github.com/go-openapi/validate",
        version = "v0.19.5",
        sum = "h1:QhCBKRYqZR+SKo4gl1lPhPahope8/RLt6EVgY8X80w0=",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        importpath = "github.com/go-sql-driver/mysql",
        version = "v1.5.0",
        sum = "h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        version = "v1.8.0",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        importpath = "github.com/godbus/dbus/v5",
        version = "v5.0.3",
        sum = "h1:ZqHaoEF7TBzh4jzPmqVhE/5A1z9of6orkAe5uHoAeME=",
    )
    go_repository(
        name = "com_github_gofrs_flock",
        importpath = "github.com/gofrs/flock",
        version = "v0.6.1-0.20180915234121-886344bea079",
        sum = "h1:JFTFz3HZTGmgMz4E1TabNBNJljROSYgja1b4l50FNVs=",
    )
    go_repository(
        name = "com_github_gofrs_uuid",
        importpath = "github.com/gofrs/uuid",
        version = "v3.2.0+incompatible",
        sum = "h1:y12jRkkFxsd7GpqdSZ+/KCs/fJbqpEXSGd4+jfEaewE=",
    )
    go_repository(
        name = "com_github_gogo_googleapis",
        importpath = "github.com/gogo/googleapis",
        version = "v1.3.2",
        sum = "h1:kX1es4djPJrsDhY7aZKJy7aZasdcB5oSOEphMjSB53c=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        version = "v1.3.1",
        sum = "h1:DqDEcV5aeaTmdFBePNpYsp3FlcVH/2ISVVM9Qf8PSls=",
    )
    go_repository(
        name = "com_github_golang_sql_civil",
        importpath = "github.com/golang-sql/civil",
        version = "v0.0.0-20190719163853-cb61b32ac6fe",
        sum = "h1:lXe2qZdvpiX5WZkZR4hgp4KJVfY3nMkvmwbVkpv1rVY=",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        version = "v0.0.0-20191227052852-215e87163ea7",
        sum = "h1:5ZkaAPbicIKTF2I64qf5Fh8Aa83Q/dnOafMYV0OMwjA=",
    )
    go_repository(
        name = "com_github_golang_snappy",
        importpath = "github.com/golang/snappy",
        version = "v0.0.1",
        sum = "h1:Qgr9rKW7uDUkrbSmQeiDsGa8SjGyCOGtuasMWwvp2P4=",
    )
    go_repository(
        name = "com_github_google_btree",
        importpath = "github.com/google/btree",
        version = "v1.0.0",
        sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
    )
    go_repository(
        name = "com_github_google_cadvisor",
        importpath = "github.com/google/cadvisor",
        version = "v0.36.1-0.20200623171404-8450c56c21bc",
        sum = "h1:il4pi2iOP5NRkBgnZH3n0GDqSCNEJ/QIRJrCAfU5h38=",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        version = "v0.4.0",
        sum = "h1:xsAVV57WRhGj6kEIi8ReJzQlHHqcBYCElAvkovg3B/4=",
    )
    go_repository(
        name = "com_github_google_go_dap",
        importpath = "github.com/google/go-dap",
        version = "v0.2.0",
        sum = "h1:whjIGQRumwbR40qRU7CEKuFLmePUUc2s4Nt9DoXXxWk=",
    )
    go_repository(
        name = "com_github_google_go_tpm",
        importpath = "github.com/google/go-tpm",
        version = "v0.1.2-0.20190725015402-ae6dd98980d4",
        sum = "h1:GNNkIb6NSjYfw+KvgUFW590mcgsSFihocSrbXct1sEw=",
    )
    go_repository(
        name = "com_github_google_go_tpm_tools",
        importpath = "github.com/google/go-tpm-tools",
        version = "v0.0.0-20190731025042-f8c04ff88181",
        sum = "h1:1Y5W2uh6E7I6hhI6c0WVSbV+Ae15uhemqi3RvSgtZpk=",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        version = "v1.1.0",
        sum = "h1:Hsa8mG0dQ46ij8Sl2AYJDUv1oA9/d6Vk+3LG99Oe02g=",
    )
    go_repository(
        name = "com_github_google_gopacket",
        importpath = "github.com/google/gopacket",
        version = "v1.1.17",
        sum = "h1:rMrlX2ZY2UbvT+sdz3+6J+pp2z+msCq9MxTU6ymxbBY=",
    )
    go_repository(
        name = "com_github_google_gops",
        importpath = "github.com/google/gops",
        version = "v0.3.6",
        sum = "h1:6akvbMlpZrEYOuoebn2kR+ZJekbZqJ28fJXTs84+8to=",
    )
    go_repository(
        name = "com_github_google_gvisor",
        importpath = "github.com/google/gvisor",
        version = "v0.0.0-20200511005220-c52195d25825",
        sum = "h1:Ryt0ml851mYbHu2ibbtjOCyJCDYdqdhEv5INoPR6Ovs=",
        patches = [
            "//third_party/go/patches:gvisor.patch",
        ],
        patch_args = ["-p1"],
    )
    go_repository(
        name = "com_github_google_gvisor_containerd_shim",
        importpath = "github.com/google/gvisor-containerd-shim",
        version = "v0.0.4",
        sum = "h1:RdBNQHpoQ3ekzfXYIV4+nQJ3a2xLnIHuZJkM40OEtyA=",
        patches = [
            "//third_party/go/patches:gvisor-containerd-shim.patch",
            "//third_party/go/patches:gvisor-containerd-shim-build.patch",
            "//third_party/go/patches:gvisor-containerd-shim-nogo.patch",
            "//third_party/go/patches:gvisor-shim-root.patch",
        ],
        patch_args = ["-p1"],
    )
    go_repository(
        name = "com_github_google_nftables",
        importpath = "github.com/google/nftables",
        version = "v0.0.0-20200316075819-7127d9d22474",
        sum = "h1:D6bN82zzK92ywYsE+Zjca7EHZCRZbcNTU3At7WdxQ+c=",
    )
    go_repository(
        name = "com_github_google_subcommands",
        importpath = "github.com/google/subcommands",
        version = "v0.0.0-20190508160503-636abe8753b8",
        sum = "h1:GZGUPQiZfYrd9uOqyqwbQcHPkz/EZJVkZB1MkaO9UBI=",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        version = "v1.1.1",
        sum = "h1:Gkbcsh/GbpXz7lPftLA3P6TYMwjCLYm83jiFQZF/3gY=",
    )
    go_repository(
        name = "com_github_googleapis_gnostic",
        importpath = "github.com/googleapis/gnostic",
        version = "v0.4.1",
        sum = "h1:DLJCy1n/vrD4HPjOvYcT8aYQXpPIzoRZONaYwyycI+I=",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        importpath = "github.com/gorilla/websocket",
        version = "v1.4.0",
        sum = "h1:WDFjx/TMzVgy9VdMMQi2K2Emtwi2QcUQsztZ/zLaH/Q=",
    )
    go_repository(
        name = "com_github_gregjones_httpcache",
        importpath = "github.com/gregjones/httpcache",
        version = "v0.0.0-20180305231024-9cad4c3443a7",
        sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        version = "v1.0.1-0.20190118093823-f849b5445de4",
        sum = "h1:z53tR0945TRRQO/fLEVPI6SMv7ZflF0TEaTAoU7tOzg=",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        version = "v1.2.0",
        sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        version = "v1.9.5",
        sum = "h1:UImYN5qQ8tuGpGE16ZmjvcTtTw24zw1QAp/SlnNrZhI=",
    )
    go_repository(
        name = "com_github_grpc_grpc",
        importpath = "github.com/grpc/grpc",
        version = "v1.26.0",
        sum = "h1:0/fjvIF5JHJdr34/JPEk1DJFFonjW37pDLvuAy9YieQ=",
    )
    go_repository(
        name = "com_github_hashicorp_consul_api",
        importpath = "github.com/hashicorp/consul/api",
        version = "v1.2.0",
        sum = "h1:oPsuzLp2uk7I7rojPKuncWbZ+m5TMoD4Ivs+2Rkeh4Y=",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        importpath = "github.com/hashicorp/errwrap",
        version = "v1.0.0",
        sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
    )
    go_repository(
        name = "com_github_hashicorp_go_cleanhttp",
        importpath = "github.com/hashicorp/go-cleanhttp",
        version = "v0.5.1",
        sum = "h1:dH3aiDG9Jvb5r5+bYHsikaOUIpcM0xvgMXVoDkXMzJM=",
    )
    go_repository(
        name = "com_github_hashicorp_go_immutable_radix",
        importpath = "github.com/hashicorp/go-immutable-radix",
        version = "v1.1.0",
        sum = "h1:vN9wG1D6KG6YHRTWr8512cxGOVgTMEfgEdSj/hr8MPc=",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        importpath = "github.com/hashicorp/go-multierror",
        version = "v1.0.0",
        sum = "h1:iVjPR7a6H0tWELX5NxNe7bYopibicUzc7uPribsnS6o=",
    )
    go_repository(
        name = "com_github_hashicorp_go_rootcerts",
        importpath = "github.com/hashicorp/go-rootcerts",
        version = "v1.0.0",
        sum = "h1:Rqb66Oo1X/eSV1x66xbDccZjhJigjg0+e82kpwzSwCI=",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        version = "v0.5.3",
        sum = "h1:YPkqC67at8FYaadspW/6uE0COsBxS2656RLEr8Bppgk=",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        importpath = "github.com/hashicorp/hcl",
        version = "v1.0.0",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
    )
    go_repository(
        name = "com_github_hashicorp_serf",
        importpath = "github.com/hashicorp/serf",
        version = "v0.8.2",
        sum = "h1:YZ7UKsJv+hKjqGVUUbtE3HNj79Eln2oQ75tniF6iPt0=",
    )
    go_repository(
        name = "com_github_imdario_mergo",
        importpath = "github.com/imdario/mergo",
        version = "v0.3.7",
        sum = "h1:Y+UAYTZ7gDEuOfhxKWy+dvb5dRQ6rJjFSdX2HZY1/gI=",
    )
    go_repository(
        name = "com_github_insomniacslk_dhcp",
        importpath = "github.com/insomniacslk/dhcp",
        version = "v0.0.0-20200402185128-5dd7202f1971",
        sum = "h1:P1pxzF2xvdnSY12ODSSwjxA4tyEjDEJNn829OXKnqks=",
    )
    go_repository(
        name = "com_github_j_keck_arping",
        importpath = "github.com/j-keck/arping",
        version = "v0.0.0-20160618110441-2cf9dc699c56",
        sum = "h1:742eGXur0715JMq73aD95/FU0XpVKXqNuTnEfXsLOYQ=",
    )
    go_repository(
        name = "com_github_joho_godotenv",
        importpath = "github.com/joho/godotenv",
        version = "v1.3.0",
        sum = "h1:Zjp+RcGpHhGlrMbJzXTrZZPrWj+1vfm90La1wgB6Bhc=",
    )
    go_repository(
        name = "com_github_jonboulle_clockwork",
        importpath = "github.com/jonboulle/clockwork",
        version = "v0.1.0",
        sum = "h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        version = "v1.1.9",
        sum = "h1:9yzud/Ht36ygwatGx56VwCZtlI/2AD15T1X2sjSuGns=",
    )
    go_repository(
        name = "com_github_kardianos_osext",
        importpath = "github.com/kardianos/osext",
        version = "v0.0.0-20170510131534-ae77be60afb1",
        sum = "h1:PJPDf8OUfOK1bb/NeTKd4f1QXZItOX389VN3B6qC8ro=",
    )
    go_repository(
        name = "com_github_karrick_godirwalk",
        importpath = "github.com/karrick/godirwalk",
        version = "v1.7.5",
        sum = "h1:VbzFqwXwNbAZoA6W5odrLr+hKK197CcENcPh6E/gJ0M=",
    )
    go_repository(
        name = "com_github_kevinburke_go_bindata",
        importpath = "github.com/kevinburke/go-bindata",
        version = "v3.16.0+incompatible",
        sum = "h1:TFzFZop2KxGhqNwsyjgmIh5JOrpG940MZlm5gNbxr8g=",
    )
    go_repository(
        name = "com_github_koneu_natend",
        importpath = "github.com/koneu/natend",
        version = "v0.0.0-20150829182554-ec0926ea948d",
        sum = "h1:MFX8DxRnKMY/2M3H61iSsVbo/n3h0MWGmWNN1UViOU0=",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        version = "v1.0.3",
        sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        version = "v0.2.0",
        sum = "h1:s5hAObm+yFO5uHYt5dYjxi2rXrsnmRpJx4OYvIWUaQs=",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        version = "v1.1.1",
        sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        version = "v0.1.0",
        sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
    )
    go_repository(
        name = "com_github_lib_pq",
        importpath = "github.com/lib/pq",
        version = "v1.2.1-0.20191011153232-f91d3411e481",
        sum = "h1:r9fnMM01mkhtfe6QfLrr/90mBVLnJHge2jGeBvApOjk=",
    )
    go_repository(
        name = "com_github_liggitt_tabwriter",
        importpath = "github.com/liggitt/tabwriter",
        version = "v0.0.0-20181228230101-89fcab3d43de",
        sum = "h1:9TO3cAIGXtEhnIaL+V+BEER86oLrvS+kWobKpbJuye0=",
    )
    go_repository(
        name = "com_github_lithammer_dedent",
        importpath = "github.com/lithammer/dedent",
        version = "v1.1.0",
        sum = "h1:VNzHMVCBNG1j0fh3OrsFRkVUwStdDArbgBWoPAffktY=",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star",
        importpath = "github.com/lyft/protoc-gen-star",
        version = "v0.4.14",
        sum = "h1:HUkD4H4dYFIgu3Bns/3N6J5GmKHCEGnhYBwNu3fvXgA=",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        version = "v1.8.0",
        sum = "h1:LLgXmsheXeRoUOBOjtwPQCWIYqM/LU1ayDtDePerRcY=",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        importpath = "github.com/mailru/easyjson",
        version = "v0.7.0",
        sum = "h1:aizVhC/NAAcKWb+5QsU1iNOZb4Yws5UO2I+aIprQITM=",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        version = "v0.0.9",
        sum = "h1:UVL0vNpWh04HeJXV0KLcaT7r06gOH2l4OW6ddYRUIY4=",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        version = "v0.0.4",
        sum = "h1:bnP0vzxcAdeI1zdubAl5PjU6zsERjGZb7raWodagDYs=",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        version = "v0.0.2",
        sum = "h1:UnlwIPBGaTZfPQ6T1IGzPI0EkYAQmT9fAEJ/poFC63o=",
    )
    go_repository(
        name = "com_github_mattn_go_shellwords",
        importpath = "github.com/mattn/go-shellwords",
        version = "v1.0.5",
        sum = "h1:JhhFTIOslh5ZsPrpa3Wdg8bF0WI3b44EMblmU9wIsXc=",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        importpath = "github.com/mattn/go-sqlite3",
        version = "v1.12.0",
        sum = "h1:u/x3mp++qUxvYfulZ4HKOvVO0JWhk7HtE8lWhbGz/Do=",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        version = "v1.0.1",
        sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
    )
    go_repository(
        name = "com_github_mdlayher_ethernet",
        importpath = "github.com/mdlayher/ethernet",
        version = "v0.0.0-20190606142754-0394541c37b7",
        sum = "h1:lez6TS6aAau+8wXUP3G9I3TGlmPFEq2CTxBaRqY6AGE=",
    )
    go_repository(
        name = "com_github_mdlayher_genetlink",
        importpath = "github.com/mdlayher/genetlink",
        version = "v1.0.0",
        sum = "h1:OoHN1OdyEIkScEmRgxLEe2M9U8ClMytqA5niynLtfj0=",
    )
    go_repository(
        name = "com_github_mdlayher_netlink",
        importpath = "github.com/mdlayher/netlink",
        version = "v1.1.0",
        sum = "h1:mpdLgm+brq10nI9zM1BpX1kpDbh3NLl3RSnVq6ZSkfg=",
    )
    go_repository(
        name = "com_github_mdlayher_raw",
        importpath = "github.com/mdlayher/raw",
        version = "v0.0.0-20190606142536-fef19f00fc18",
        sum = "h1:zwOa3e/13D6veNIz6zzuqrd3eZEMF0dzD0AQWKcYSs4=",
    )
    go_repository(
        name = "com_github_miekg_dns",
        importpath = "github.com/miekg/dns",
        version = "v1.1.4-0.20190417235132-8e25ec9a0ff3",
        sum = "h1:wenYMyWJ08dgEUUj0Ija8qdK/V9vL3ThAD5sjOYlFlg=",
        replace = "github.com/cilium/dns",
    )
    go_repository(
        name = "com_github_mindprince_gonvml",
        importpath = "github.com/mindprince/gonvml",
        version = "v0.0.0-20190828220739-9ebdce4bb989",
        sum = "h1:PS1dLCGtD8bb9RPKJrc8bS7qHL6JnW1CZvwzH9dPoUs=",
    )
    go_repository(
        name = "com_github_mistifyio_go_zfs",
        importpath = "github.com/mistifyio/go-zfs",
        version = "v2.1.2-0.20190413222219-f784269be439+incompatible",
        sum = "h1:aKW/4cBs+yK6gpqU3K/oIwk9Q/XICqd3zOX/UFuvqmk=",
    )
    go_repository(
        name = "com_github_mitchellh_cli",
        importpath = "github.com/mitchellh/cli",
        version = "v1.0.0",
        sum = "h1:iGBIsUe3+HZ/AD/Vd7DErOt5sU9fa8Uj7A2s1aggv1Y=",
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        importpath = "github.com/mitchellh/go-wordwrap",
        version = "v1.0.0",
        sum = "h1:6GlHJ/LTGMrIJbwgdqdl2eEH8o+Exx/0m8ir9Gns0u4=",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        version = "v1.1.2",
        sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
    )
    go_repository(
        name = "com_github_moby_sys_mountinfo",
        importpath = "github.com/moby/sys/mountinfo",
        version = "v0.1.3",
        sum = "h1:KIrhRO14+AkwKvG/g2yIpNMOUVZ02xNhOw8KY1WsLOI=",
    )
    go_repository(
        name = "com_github_moby_term",
        importpath = "github.com/moby/term",
        version = "v0.0.0-20200312100748-672ec06f55cd",
        sum = "h1:aY7OQNf2XqY/JQ6qREWamhI/81os/agb2BAGpcx5yWI=",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        version = "v1.0.1",
        sum = "h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
    )
    go_repository(
        name = "com_github_morikuni_aec",
        importpath = "github.com/morikuni/aec",
        version = "v1.0.0",
        sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
    )
    go_repository(
        name = "com_github_mrunalp_fileutils",
        importpath = "github.com/mrunalp/fileutils",
        version = "v0.0.0-20200520151820-abd8a0e76976",
        sum = "h1:aZQToFSLH8ejFeSkTc3r3L4dPImcj7Ib/KgmkQqbGGg=",
    )
    go_repository(
        name = "com_github_munnerz_goautoneg",
        importpath = "github.com/munnerz/goautoneg",
        version = "v0.0.0-20191010083416-a7dc8b61c822",
        sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
    )
    go_repository(
        name = "com_github_mxk_go_flowrate",
        importpath = "github.com/mxk/go-flowrate",
        version = "v0.0.0-20140419014527-cca7078d478f",
        sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
    )
    go_repository(
        name = "com_github_olekukonko_tablewriter",
        importpath = "github.com/olekukonko/tablewriter",
        version = "v0.0.0-20170122224234-a0225b3f23b5",
        sum = "h1:58+kh9C6jJVXYjt8IE48G2eWl6BjwU5Gj0gqY84fy78=",
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        importpath = "github.com/opencontainers/go-digest",
        version = "v1.0.0",
        sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        importpath = "github.com/opencontainers/image-spec",
        version = "v1.0.1",
        sum = "h1:JMemWkRwHx4Zj+fVxWoMCFm/8sYGGrUVojFA6h/TRcI=",
    )
    go_repository(
        name = "com_github_opencontainers_runc",
        importpath = "github.com/opencontainers/runc",
        version = "v1.0.0-rc91",
        sum = "h1:Tp8LWs5G8rFpzTsbRjAtQkPVexhCu0bnANE5IfIhJ6g=",
    )
    go_repository(
        name = "com_github_opencontainers_runtime-spec",
        importpath = "github.com/opencontainers/runtime-spec",
        version = "v1.0.3-0.20200520003142-237cc4f519e2",
        sum = "h1:9mv9SC7GWmRWE0J/+oD8w3GsN2KYGKtg6uwLN7hfP5E=",
    )
    go_repository(
        name = "com_github_opencontainers_selinux",
        importpath = "github.com/opencontainers/selinux",
        version = "v1.5.1",
        sum = "h1:jskKwSMFYqyTrHEuJgQoUlTcId0av64S6EWObrIfn5Y=",
        build_tags = [
            "selinux",
        ],
    )
    go_repository(
        name = "com_github_optiopay_kafka",
        importpath = "github.com/optiopay/kafka",
        version = "v0.0.0-20180809090225-01ce283b732b",
        sum = "h1:+bsFX/WOMIoaayXVyRem1awcpz3icz/HoL8Dxg/m6a4=",
        replace = "github.com/cilium/kafka",
    )
    go_repository(
        name = "com_github_pborman_uuid",
        importpath = "github.com/pborman/uuid",
        version = "v1.2.0",
        sum = "h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
    )
    go_repository(
        name = "com_github_peterbourgon_diskv",
        importpath = "github.com/peterbourgon/diskv",
        version = "v2.0.1+incompatible",
        sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
    )
    go_repository(
        name = "com_github_peterh_liner",
        importpath = "github.com/peterh/liner",
        version = "v0.0.0-20170317030525-88609521dc4b",
        sum = "h1:8uaXtUkxiy+T/zdLWuxa/PG4so0TPZDZfafFNNSaptE=",
    )
    go_repository(
        name = "com_github_petermattis_goid",
        importpath = "github.com/petermattis/goid",
        version = "v0.0.0-20180202154549-b0b1615b78e5",
        sum = "h1:q2e307iGHPdTGp0hoxKjt1H5pDo6utceo3dQVK3I5XQ=",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        version = "v0.9.1",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    )
    go_repository(
        name = "com_github_posener_complete",
        importpath = "github.com/posener/complete",
        version = "v1.1.1",
        sum = "h1:ccV59UEOTzVDnDUEFdT95ZzHVZ+5+158q8+SJb2QV5w=",
    )
    go_repository(
        name = "com_github_pquerna_cachecontrol",
        importpath = "github.com/pquerna/cachecontrol",
        version = "v0.0.0-20171018203845-0dec1b30a021",
        sum = "h1:0XM1XL/OFFJjXsYXlG30spTkV/E9+gmd5GD1w2HE8xM=",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        version = "v1.6.0",
        sum = "h1:YVPodQOcK15POxhgARIvnDRVpLcuK8mglnMrWfyrw6A=",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        version = "v0.2.0",
        sum = "h1:uq5h0d+GuxiXLJLNABMgp2qUWDPiLvgCzz2dUR+/W/M=",
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        version = "v0.9.1",
        sum = "h1:KOMtN28tlbam3/7ZKEYKHhKoJZYYj3gMH4uc62x7X7U=",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        version = "v0.0.11",
        sum = "h1:DhHlBtkHWPYi8O2y31JkK0TF+DGM+51OopZjH/Ia5qI=",
    )
    go_repository(
        name = "com_github_rekby_gpt",
        importpath = "github.com/rekby/gpt",
        version = "v0.0.0-20200219180433-a930afbc6edc",
        sum = "h1:goZGTwEEn8mWLcY012VouWZWkJ8GrXm9tS3VORMxT90=",
    )
    go_repository(
        name = "com_github_robfig_cron",
        importpath = "github.com/robfig/cron",
        version = "v1.1.0",
        sum = "h1:jk4/Hud3TTdcrJgUOBgsqrZBarcxl6ADIjSC2iniwLY=",
    )
    go_repository(
        name = "com_github_rubenv_sql_migrate",
        importpath = "github.com/rubenv/sql-migrate",
        version = "v0.0.0-20200429072036-ae26b214fa43",
        sum = "h1:0i6uTtxUGc/jpK/CngM4T2S2NFnqYUUxH+lKDgBLw8U=",
    )
    go_repository(
        name = "com_github_russross_blackfriday",
        importpath = "github.com/russross/blackfriday",
        version = "v1.5.2",
        sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
    )
    go_repository(
        name = "com_github_russross_blackfriday_v2",
        importpath = "github.com/russross/blackfriday/v2",
        version = "v2.0.1",
        sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
    )
    go_repository(
        name = "com_github_safchain_ethtool",
        importpath = "github.com/safchain/ethtool",
        version = "v0.0.0-20190326074333-42ed695e3de8",
        sum = "h1:2c1EFnZHIPCW8qKWgHMH/fX2PkSabFc5mrVzfUNdg5U=",
    )
    go_repository(
        name = "com_github_sasha_s_go_deadlock",
        importpath = "github.com/sasha-s/go-deadlock",
        version = "v0.2.1-0.20190427202633-1595213edefa",
        sum = "h1:0U2s5loxrTy6/VgfVoLuVLFJcURKLH49ie0zSch7gh4=",
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
    )
    go_repository(
        name = "com_github_sbezverk_nftableslib",
        importpath = "github.com/sbezverk/nftableslib",
        version = "v0.0.0-20200402150358-c20bed91f482",
        sum = "h1:k7gEZ/EwJhHDTRXFUZQlE4/p1cmoha7zL7PWCDG3ZHQ=",
    )
    go_repository(
        name = "com_github_seccomp_libseccomp_golang",
        importpath = "github.com/seccomp/libseccomp-golang",
        version = "v0.9.1",
        sum = "h1:NJjM5DNFOs0s3kYE1WUOr6G8V97sdt46rlXTMfXGWBo=",
    )
    go_repository(
        name = "com_github_servak_go_fastping",
        importpath = "github.com/servak/go-fastping",
        version = "v0.0.0-20160802140958-5718d12e20a0",
        sum = "h1:FFgMDF0otYdRIy7stdzyE6l1mbyw16XtOWXn6NJ8bEU=",
    )
    go_repository(
        name = "com_github_shirou_gopsutil",
        importpath = "github.com/shirou/gopsutil",
        version = "v0.0.0-20180427012116-c95755e4bcd7",
        sum = "h1:80VN+vGkqM773Br/uNNTSheo3KatTgV8IpjIKjvVLng=",
    )
    go_repository(
        name = "com_github_shurcool_sanitized_anchor_name",
        importpath = "github.com/shurcooL/sanitized_anchor_name",
        version = "v1.0.0",
        sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        version = "v1.6.0",
        sum = "h1:UBcNElsrwanuuMsnGSlYmtmgbb23qDR5dG+6X6Oo89I=",
    )
    go_repository(
        name = "com_github_soheilhy_cmux",
        importpath = "github.com/soheilhy/cmux",
        version = "v0.1.4",
        sum = "h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        version = "v1.2.2",
        sum = "h1:5jhuqJyZCZf2JRofRvN/nIFgIWNzPa3/Vz8mYylgbWc=",
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        version = "v1.3.1",
        sum = "h1:nFm6S0SMdyzrzcmThSipiEubIDy8WEXKNZ0UOgiRpng=",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        version = "v1.0.0",
        sum = "h1:6m/oheQuQ13N9ks4hubMG6BnvwOeaJrqSPLahSnczz8=",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        importpath = "github.com/spf13/jwalterweatherman",
        version = "v1.0.0",
        sum = "h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        version = "v1.0.5",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
    )
    go_repository(
        name = "com_github_spf13_viper",
        importpath = "github.com/spf13/viper",
        version = "v1.6.3",
        sum = "h1:pDDu1OyEDTKzpJwdq4TiuLyMsUgRa/BT5cn5O62NoHs=",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        version = "v1.4.0",
        sum = "h1:2E4SXV/wtOkTonXsotYi4li6zVWxYlZuYNCXe9XRJyk=",
    )
    go_repository(
        name = "com_github_subosito_gotenv",
        importpath = "github.com/subosito/gotenv",
        version = "v1.2.0",
        sum = "h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
    )
    go_repository(
        name = "com_github_syndtr_gocapability",
        importpath = "github.com/syndtr/gocapability",
        version = "v0.0.0-20180916011248-d98352740cb2",
        sum = "h1:b6uOv7YOFK0TYG7HtkIgExQo+2RdLuwRft63jn2HWj8=",
    )
    go_repository(
        name = "com_github_tchap_go_patricia",
        importpath = "github.com/tchap/go-patricia",
        version = "v2.2.6+incompatible",
        sum = "h1:JvoDL7JSoIP2HDE8AbDH3zC8QBPxmzYe32HHy5yQ+Ck=",
    )
    go_repository(
        name = "com_github_tmc_grpc_websocket_proxy",
        importpath = "github.com/tmc/grpc-websocket-proxy",
        version = "v0.0.0-20190109142713-0ad062ec5ee5",
        sum = "h1:LnC5Kc/wtumK+WB441p7ynQJzVuNRJiqddSIE3IlSEQ=",
    )
    go_repository(
        name = "com_github_u_root_u_root",
        importpath = "github.com/u-root/u-root",
        version = "v6.0.0+incompatible",
        sum = "h1:YqPGmRoRyYmeg17KIWFRSyVq6LX5T6GSzawyA6wG6EE=",
    )
    go_repository(
        name = "com_github_urfave_cli",
        importpath = "github.com/urfave/cli",
        version = "v1.22.1",
        sum = "h1:+mkCCcOFKPnCmVYVcURKps1Xe+3zP90gSYGNfRkjoIY=",
    )
    go_repository(
        name = "com_github_vishvananda_netlink",
        importpath = "github.com/vishvananda/netlink",
        version = "v1.1.0",
        sum = "h1:1iyaYNBLmP6L0220aDnYQpo1QEV4t4hJ+xEEhhJH8j0=",
    )
    go_repository(
        name = "com_github_vishvananda_netns",
        importpath = "github.com/vishvananda/netns",
        version = "v0.0.0-20200520041808-52d707b772fe",
        sum = "h1:mjAZxE1nh8yvuwhGHpdDqdhtNu2dgbpk93TwoXuk5so=",
    )
    go_repository(
        name = "com_github_volatiletech_inflect",
        importpath = "github.com/volatiletech/inflect",
        version = "v0.0.1",
        sum = "h1:2a6FcMQyhmPZcLa+uet3VJ8gLn/9svWhJxJYwvE8KsU=",
    )
    go_repository(
        name = "com_github_volatiletech_null_v8",
        importpath = "github.com/volatiletech/null/v8",
        version = "v8.1.0",
        sum = "h1:eAO3I31A5R04usY5SKMMfDcOCnEGyT/T4wRI0JVGp4U=",
    )
    go_repository(
        name = "com_github_volatiletech_randomize",
        importpath = "github.com/volatiletech/randomize",
        version = "v0.0.1",
        sum = "h1:eE5yajattWqTB2/eN8df4dw+8jwAzBtbdo5sbWC4nMk=",
    )
    go_repository(
        name = "com_github_volatiletech_sqlboiler_v4",
        importpath = "github.com/volatiletech/sqlboiler/v4",
        version = "v4.1.1",
        sum = "h1:cmpaEri8whb5lRv6q2ycWtmiWd42llsrDaERk2BkWbE=",
    )
    go_repository(
        name = "com_github_volatiletech_strmangle",
        importpath = "github.com/volatiletech/strmangle",
        version = "v0.0.1",
        sum = "h1:UKQoHmY6be/R3tSvD2nQYrH41k43OJkidwEiC74KIzk=",
    )
    go_repository(
        name = "com_github_xiang90_probing",
        importpath = "github.com/xiang90/probing",
        version = "v0.0.0-20190116061207-43a291ad63a2",
        sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
    )
    go_repository(
        name = "com_github_yalue_native_endian",
        importpath = "github.com/yalue/native_endian",
        version = "v0.0.0-20180607135909-51013b03be4f",
        sum = "h1:nsQCScpQ8RRf+wIooqfyyEUINV2cAPuo2uVtHSBbA4M=",
    )
    go_repository(
        name = "io_etcd_go_bbolt",
        importpath = "go.etcd.io/bbolt",
        version = "v1.3.5",
        sum = "h1:XAzx9gjCb0Rxj7EoqcClPD1d5ZBxZJk0jbuoPHenBt0=",
    )
    go_repository(
        name = "io_etcd_go_etcd",
        importpath = "go.etcd.io/etcd",
        version = "v0.5.0-alpha.5.0.20200520232829-54ba9589114f",
        sum = "h1:pBCD+Z7cy5WPTq+R6MmJJvDRpn88cp7bmTypBsn91g4=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        importpath = "go.mongodb.org/mongo-driver",
        version = "v1.1.2",
        sum = "h1:jxcFYjlkl8xaERsgLo+RNquI0epW6zuy/ZRQs6jnrFA=",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        version = "v0.22.0",
        sum = "h1:C9hSCOW830chIVkdja34wa6Ky+IzWllkUinR+BtRZd4=",
    )
    go_repository(
        name = "net_starlark_go",
        importpath = "go.starlark.net",
        version = "v0.0.0-20190702223751-32f345186213",
        sum = "h1:lkYv5AKwvvduv5XWP6szk/bvvgO6aDeUujhZQXIFTes=",
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        version = "v1.4.0",
        sum = "h1:cxzIVoETapQEqDhQu3QfnvXAV4AlzcvUCxkVUFw3+EU=",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        version = "v1.1.0",
        sum = "h1:HoEmRHQPVSqub6w2z2d2EOVs2fjyFRGyofhKuyDq0QI=",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        version = "v1.15.0",
        sum = "h1:ZZCA22JRF2gQE5FoNmhmrf7jeJJ2uhqDUNRYKm8dvmM=",
    )
    go_repository(
        name = "org_golang_x_arch",
        importpath = "golang.org/x/arch",
        version = "v0.0.0-20190927153633-4e8777c89be4",
        sum = "h1:QlVATYS7JBoZMVaf+cNjb90WD/beKVHnIxFKT4QaHVI=",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        version = "v0.0.0-20200220183623-bac4c82f6975",
        sum = "h1:/Tl7pH94bvbAAHBdZJT947M/+gp0+CqQXDtMRC0fseo=",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        version = "v0.3.0",
        sum = "h1:RM4zey1++hCTbCVQfnWeKs9/IEsaBLA8vTkd0WVtmH4=",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        version = "v0.0.0-20190311183353-d8887717615a",
        sum = "h1:oWX7TPOiFAMXLq8o0ikBYfCJVlRHBcsciT5bXOrH628=",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        version = "v0.0.0-20191202225959-858c2ad4c8b6",
        sum = "h1:pE8b58s1HRDMi8RDc79m0HISf9D4TzseP40cEA6IGfs=",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        version = "v0.0.0-20181108010431-42b317875d0f",
        sum = "h1:Bl/8QSvNqXvPGPGXa2z5xUTmV7VDcZyvRZ+QQXkXTZQ=",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        version = "v0.0.0-20200327173247-9dae0f8f5775",
        sum = "h1:TC0v2RSO1u2kn1ZugjrFXkRZAEaqMN/RW+OTZkBzmLE=",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        version = "v0.3.0",
        sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        version = "v0.0.0-20191024005414-555d28b269f0",
        sum = "h1:/5xXl8Y5W96D+TtHSlonuFqGHIWVuyCkGJLwGh9JJFs=",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        version = "v0.0.0-20191204190536-9bdfabe68543",
        sum = "h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
    )
    go_repository(
        name = "com_zx2c4_golang_wireguard_wgctrl",
        importpath = "golang.zx2c4.com/wireguard/wgctrl",
        version = "v0.0.0-20200515170644-ec7f26be9d9e",
        sum = "h1:fqDhK9OlzaaiFjnyaAfR9Q1RPKCK7OCTLlHGP9f74Nk=",
    )
    go_repository(
        name = "org_gonum_v1_gonum",
        importpath = "gonum.org/v1/gonum",
        version = "v0.6.2",
        sum = "h1:4r+yNT0+8SWcOkXP+63H2zQbN+USnC73cjGUxnDF94Q=",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        version = "v0.0.0-20200224152610-e50cd9704f63",
        sum = "h1:YzfoEYWbODU5Fbt37+h7X16BWQbad7Q4S6gclTKFXM8=",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        version = "v1.26.0",
        sum = "h1:2dTRdpdFEEhJYQD8EMLB61nnrzSCTbG38PhqdhvOltg=",
    )
    go_repository(
        name = "in_gopkg_djherbis_times_v1",
        importpath = "gopkg.in/djherbis/times.v1",
        version = "v1.2.0",
        sum = "h1:UCvDKl1L/fmBygl2Y7hubXCnY7t4Yj46ZrBFNUipFbM=",
    )
    go_repository(
        name = "in_gopkg_gorp_v1",
        importpath = "gopkg.in/gorp.v1",
        version = "v1.7.2",
        sum = "h1:j3DWlAyGVv8whO7AcIWznQ2Yj7yJkn34B8s63GViAAw=",
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        importpath = "gopkg.in/inf.v0",
        version = "v0.9.1",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    )
    go_repository(
        name = "in_gopkg_ini_v1",
        importpath = "gopkg.in/ini.v1",
        version = "v1.51.0",
        sum = "h1:AQvPpx3LzTDM0AjnIRlVFwFFGC+npRopjZxLJj6gdno=",
    )
    go_repository(
        name = "in_gopkg_natefinch_lumberjack_v2",
        importpath = "gopkg.in/natefinch/lumberjack.v2",
        version = "v2.0.0",
        sum = "h1:1Lc07Kr7qY4U2YPouBjpCLxpiyxIVoxqXgkXLknAOE8=",
    )
    go_repository(
        name = "in_gopkg_square_go_jose_v2",
        importpath = "gopkg.in/square/go-jose.v2",
        version = "v2.2.2",
        sum = "h1:orlkJ3myw8CN1nVQHBFfloD+L3egixIa4FvUP6RosSA=",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        version = "v2.2.8",
        sum = "h1:obN1ZagJSUGI0Ek/LBmuj4SNLPfIny3KsKFopxRdj10=",
    )
    go_repository(
        name = "io_k8s_api",
        importpath = "k8s.io/api",
        version = "v0.19.0-rc.0",
        sum = "h1:K+xi+F3RNAxpFyS1f7uHekMNprjFX7WVZDx2lJE+A3A=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "io_k8s_apiextensions_apiserver",
        importpath = "k8s.io/apiextensions-apiserver",
        version = "v0.19.0-rc.0",
        sum = "h1:XGNmUwNvh5gt6sYwCzaxLU6Dr461DVKWlGiaCSKZzyw=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "io_k8s_apimachinery",
        importpath = "k8s.io/apimachinery",
        version = "v0.20.0-alpha.0",
        sum = "h1:XCZhrYfFYSC8GBpI4OUJFTH1s5euLMYdoIDQ7u2aDPM=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "io_k8s_apiserver",
        importpath = "k8s.io/apiserver",
        version = "v0.19.0-rc.0",
        sum = "h1:SaF/gMgUeDPbQDKHTMvB2yynBUZpp6s4HYQIOx/LdDQ=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "io_k8s_cli_runtime",
        importpath = "k8s.io/cli-runtime",
        version = "v0.19.0-rc.0",
        sum = "h1:amuzfqubksp5ooo99cpiu6hYe6ua1bGEqw59vZKyRqA=",
    )
    go_repository(
        name = "io_k8s_client_go",
        importpath = "k8s.io/client-go",
        version = "v0.19.0-rc.0",
        sum = "h1:6WW8MElhoLeYcLiN4ky1159XG5E39KYdmLCrV/6lNiE=",
        patches = [
            "//third_party/go/patches:k8s-client-go.patch",
            "//third_party/go/patches:k8s-client-go-build.patch",
        ],
        patch_args = ["-p1"],
    )
    go_repository(
        name = "io_k8s_cloud_provider",
        importpath = "k8s.io/cloud-provider",
        version = "v0.19.0-rc.0",
        sum = "h1:W1YV1XhdklzoGFZcYmzJnm3D4O6uWaoEAFRF1X4h7uw=",
    )
    go_repository(
        name = "io_k8s_cluster_bootstrap",
        importpath = "k8s.io/cluster-bootstrap",
        version = "v0.19.0-rc.0",
        sum = "h1:2OCD/1YLoWlBisd7MPfPM35ZXFct/eA94TkRs/uAuhg=",
    )
    go_repository(
        name = "io_k8s_component_base",
        importpath = "k8s.io/component-base",
        version = "v0.19.0-rc.0",
        sum = "h1:S/jt6xey1Wg5i5A9/BCkPYekpjJ5zlfuSCCVlNSJ/Yc=",
    )
    go_repository(
        name = "io_k8s_cri_api",
        importpath = "k8s.io/cri-api",
        version = "v0.19.0-rc.0",
        sum = "h1:vXd1YUBZcQkkDb2jYdtaCm+XFA2euMVGVU08EKsN40k=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "io_k8s_csi_translation_lib",
        importpath = "k8s.io/csi-translation-lib",
        version = "v0.19.0-rc.0",
        sum = "h1:2xvrVxnNKtbhilsj/gcD60P9r2PGT+zAEhBWNynySgk=",
    )
    go_repository(
        name = "io_k8s_gengo",
        importpath = "k8s.io/gengo",
        version = "v0.0.0-20200428234225-8167cfdcfc14",
        sum = "h1:t4L10Qfx/p7ASH3gXCdIUtPbbIuegCoUJf3TMSFekjw=",
    )
    go_repository(
        name = "io_k8s_heapster",
        importpath = "k8s.io/heapster",
        version = "v1.2.0-beta.1",
        sum = "h1:lUsE/AHOMHpi3MLlBEkaU8Esxm5QhdyCrv1o7ot0s84=",
    )
    go_repository(
        name = "io_k8s_klog_v2",
        importpath = "k8s.io/klog/v2",
        version = "v2.2.0",
        sum = "h1:XRvcwJozkgZ1UQJmfMGpvRthQHOvihEhYtDfAaxMz/A=",
    )
    go_repository(
        name = "io_k8s_kube_aggregator",
        importpath = "k8s.io/kube-aggregator",
        version = "v0.19.0-rc.0",
        sum = "h1:+u9y1c0R2GF8fuaEnlJrdUtxoEmQOON98oatycSquOA=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "io_k8s_kube_controller_manager",
        importpath = "k8s.io/kube-controller-manager",
        version = "v0.19.0-rc.0",
        sum = "h1:b78T0fHLtRqOEe/70UzdTI0mN2hOph/krz9B5yI/DN4=",
    )
    go_repository(
        name = "io_k8s_kube_openapi",
        importpath = "k8s.io/kube-openapi",
        version = "v0.0.0-20200427153329-656914f816f9",
        sum = "h1:5NC2ITmvg8RoxoH0wgmL4zn4VZqXGsKbxrikjaQx6s4=",
    )
    go_repository(
        name = "io_k8s_kube_proxy",
        importpath = "k8s.io/kube-proxy",
        version = "v0.19.0-rc.0",
        sum = "h1:eYzuS4rtUGH8Nglk40WIWSNQyMSTj8pKcGB14BKVhHg=",
    )
    go_repository(
        name = "io_k8s_kube_scheduler",
        importpath = "k8s.io/kube-scheduler",
        version = "v0.19.0-rc.0",
        sum = "h1:KiKDepusDaex8fJj2R0F1y2zNj/oPaCzziC7JiuU09o=",
    )
    go_repository(
        name = "io_k8s_kubectl",
        importpath = "k8s.io/kubectl",
        version = "v0.19.0-rc.0",
        sum = "h1:JcCGByIwsglw1eQKUpTfYuxSjvQ5NUQTyxoGp1P/Bx4=",
    )
    go_repository(
        name = "io_k8s_kubelet",
        importpath = "k8s.io/kubelet",
        version = "v0.19.0-rc.0",
        sum = "h1:Eii9aWFKr4MtrRSlhxnaLkGZ0WkSb2p6sPyDuMul/Tc=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "io_k8s_kubernetes",
        importpath = "k8s.io/kubernetes",
        version = "v1.19.0-rc.0",
        sum = "h1:vKA6/0biZ/LJUPuWWzn1lfqIQrjfuJBVAtHn7AYScTs=",
        build_file_proto_mode = "disable",
        build_tags = [
            "providerless",
        ],
        patches = [
            "//third_party/go/patches:k8s-kubernetes.patch",
            "//third_party/go/patches:k8s-kubernetes-build.patch",
            "//third_party/go/patches:k8s-native-metrics.patch",
            "//third_party/go/patches:k8s-use-native.patch",
        ],
        patch_args = ["-p1"],
    )
    go_repository(
        name = "io_k8s_legacy_cloud_providers",
        importpath = "k8s.io/legacy-cloud-providers",
        version = "v0.19.0-rc.0",
        sum = "h1:cyf6e9AnQL/ATzZHXDqdwlD+lmRhtKCYPcfeFqb8wn0=",
    )
    go_repository(
        name = "io_k8s_metrics",
        importpath = "k8s.io/metrics",
        version = "v0.19.0-rc.0",
        sum = "h1:hPBuMVgXakpnLBLe0K9SZxF8T7mH9VaNTY/pKsU/958=",
        build_file_proto_mode = "disable",
    )
    go_repository(
        name = "io_k8s_repo_infra",
        importpath = "k8s.io/repo-infra",
        version = "v0.0.0-20190329054012-df02ded38f95",
        sum = "h1:PQyAIB6SRdV0a3Vj/VA39L1uANW36k/zg3tOk/Ffh3U=",
    )
    go_repository(
        name = "io_k8s_sample_apiserver",
        importpath = "k8s.io/sample-apiserver",
        version = "v0.19.0-rc.0",
        sum = "h1:ZsO1AWW9k79zA+tU1nu7nGMGT7XidiA1jDrfBvMZmzg=",
    )
    go_repository(
        name = "io_k8s_utils",
        importpath = "k8s.io/utils",
        version = "v0.0.0-20200619165400-6e3d28b6ed19",
        sum = "h1:7Nu2dTj82c6IaWvL7hImJzcXoTPz1MsSCH7r+0m6rfo=",
        patches = [
            "//third_party/go/patches:k8s-native-mounter.patch",
        ],
        patch_args = ["-p1"],
    )
    go_repository(
        name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
        importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
        version = "v0.0.9",
        sum = "h1:rusRLrDhjBp6aYtl9sGEvQJr6faoHoDLd0YcUBTZguI=",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize",
        importpath = "sigs.k8s.io/kustomize",
        version = "v2.0.3+incompatible",
        sum = "h1:JUufWFNlI44MdtnjUqVnvh29rR37PQFzPbLXqhyOyX0=",
    )
    go_repository(
        name = "io_k8s_sigs_structured_merge_diff_v3",
        importpath = "sigs.k8s.io/structured-merge-diff/v3",
        version = "v3.0.0",
        sum = "h1:dOmIZBMfhcHS09XZkMyUgkq5trg3/jRyJYFZUiaOp8E=",
    )
    go_repository(
        name = "io_k8s_sigs_yaml",
        importpath = "sigs.k8s.io/yaml",
        version = "v1.2.0",
        sum = "h1:kr/MCeFWJWTwyaHoR9c8EjH9OumOmoF9YGiZd7lFm/Q=",
    )
    go_repository(
        name = "ml_vbom_util",
        importpath = "vbom.ml/util",
        version = "v0.0.0-20160121211510-db5cfe13f5cc",
        sum = "h1:MksmcCZQWAQJCTA5T0jgI/0sJ51AVm4Z41MrmfczEoc=",
    )
