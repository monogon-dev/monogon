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

def uring_external(name, version):
    sums = {
        "2.0": "ca069ecc4aa1baf1031bd772e4e97f7e26dfb6bb733d79f70159589b22ab4dc0",
    }

    http_archive(
        name = name,
        patch_args = ["-p1"],
        patches = [
            "//third_party/uring/patches:bazel_cc_fix.patch",
            "//third_party/uring/patches:include-compat-h.patch",
        ],
        sha256 = sums[version],
        build_file = "@//third_party/uring:uring.bzl",
        strip_prefix = "liburing-liburing-" + version,
        urls = ["https://github.com/axboe/liburing/archive/liburing-%s.tar.gz" % version],
    )
