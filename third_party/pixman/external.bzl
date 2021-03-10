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

def pixman_external(name, version):
    sums = {
        "0.40.0": "6d200dec3740d9ec4ec8d1180e25779c00bc749f94278c8b9021f5534db223fc",
    }

    http_archive(
        name = name,
        sha256 = sums[version],
        build_file = "@//third_party/pixman:pixman.bzl",
        strip_prefix = "pixman-" + version + "/pixman",
        urls = ["https://www.cairographics.org/releases/pixman-%s.tar.gz" % version],
    )
