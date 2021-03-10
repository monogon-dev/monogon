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

def qemu_external(name, version):
    sums = {
        "5.2.0": "cb18d889b628fbe637672b0326789d9b0e3b8027e0445b936537c78549df17bc",
    }

    http_archive(
        name = name,
        patch_args = ["-p1"],
        patches = [
            "//third_party/qemu/patches:fix_code_issues.patch",
            "//third_party/qemu/patches:bazel_support.patch",
            "//third_party/qemu/patches:pregenerated_config_files.patch",
            "//third_party/qemu/patches:headers_fix.patch",
        ],
        sha256 = sums[version],
        strip_prefix = "qemu-" + version,
        urls = ["https://download.qemu.org/qemu-%s.tar.xz" % version],
    )
