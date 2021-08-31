#  Copyright 2021 The Monogon Project Authors.
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

def cap_external(name, version):
    sums = {
        "1.2.55": "e29322032ea94e90696ae2d17530edc914c71765232ee8fd4fde38ba639cb256",
    }

    http_archive(
        name = name,
        sha256 = sums[version],
        build_file = "@//third_party/cap:cap.bzl",
        strip_prefix = "libcap-cap/v%s/libcap" % version,
        patch_args = ["-p1"],
        patches = [
            "//third_party/cap/patches:add_go_codegen.patch",
        ],
        urls = ["https://git.kernel.org/pub/scm/libs/libcap/libcap.git/snapshot/libcap-cap/v%s.tar.gz" % version],
    )
