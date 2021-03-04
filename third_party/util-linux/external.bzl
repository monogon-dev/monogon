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

def util_linux_external(name, version):
    sums = {
        "2.34": "1d0c1a38f8c14a2c251681907203cccc78704f5702f2ef4b438bed08344242f7",
        "2.36.2": "084bc6f9bd479bfd78e8924e63f85bb17953681bf1dad240e4d8601a21294d0b",
    }

    http_archive(
        name = name,
        sha256 = sums[version],
        strip_prefix = "util-linux-" + version,
        build_file = "@//third_party/util-linux:util-linux.bzl",
        patch_args = ["-p1"],
        patches = ["//third_party/util-linux/patches:bazel_cc_fix.patch"],
        urls = ["https://git.kernel.org/pub/scm/utils/util-linux/util-linux.git/snapshot/util-linux-%s.tar.gz" % version],
    )
