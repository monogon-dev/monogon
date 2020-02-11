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

def musl_external(name, version):
    sums = {
        "1.1.24": "1370c9a812b2cf2a7d92802510cca0058cc37e66a7bedd70051f0a34015022a3"
    }
    all_content = """filegroup(name = "all", srcs = glob(["**"]), visibility = ["//visibility:public"])"""

    http_archive(
        name = name,
        build_file_content = all_content,
        sha256 = sums[version],
        strip_prefix = "musl-" + version,
        urls = ["https://www.musl-libc.org/releases/musl-%s.tar.gz" % version],
    )

