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

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def xfsprogs_external(name, version):
    sums = {
        "5.2.1": "6187f25f1744d1ecbb028b0ea210ad586d0f2dae24e258e4688c67740cc861ef",
        "5.10.0": "e807ca9fd8f01e45c9ec8ffb3c123bdb7dfcfd8e05340520d2ff1ddbc3bd7c88",
    }

    http_archive(
        name = name,
        patch_args = ["-p1"],
        patches = ["//third_party/xfsprogs/patches:bazel_cc_fix.patch"],
        sha256 = sums[version],
        build_file = "@//third_party/xfsprogs:xfsprogs.bzl",
        strip_prefix = "xfsprogs-dev-" + version,
        urls = ["https://git.kernel.org/pub/scm/fs/xfs/xfsprogs-dev.git/snapshot/xfsprogs-dev-%s.tar.gz" % version],
    )
