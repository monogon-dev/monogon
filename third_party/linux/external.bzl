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

def linux_external(name, version):
    sums = {
        "5.4.7": "abc9b21d9146d95853dac35f4c4489a0199aff53ee6eee4b0563d1b37079fcc9",
        "5.6": "e342b04a2aa63808ea0ef1baab28fc520bd031ef8cf93d9ee4a31d4058fcb622",
    }
    http_archive(
        name = name,
        build_file = "//third_party/linux/external:BUILD.repo",
        patch_args = ["-p1"],
        patches = [
            # Enable built-in cmdline for efistub
            "//third_party/linux/external:0001-x86-Allow-built-in-command-line-to-work-in-early-ker.patch",
        ],
        sha256 = sums[version],
        strip_prefix = "linux-" + version,
        urls = ["https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-%s.tar.xz" % version],
    )

