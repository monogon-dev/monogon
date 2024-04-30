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
        "6.6.13": "88b89e7dd41ead4e3ab1e411c8bb8d592575acf815cf1df3c0dc57e2e882c0bc",
    }
    http_archive(
        name = name,
        build_file = "//third_party/linux/external:BUILD.repo",
        patch_args = ["-p1"],
        patches = [
            "//third_party/linux/external:0001-block-partition-expose-PARTUUID-through-uevent.patch",
            "//third_party/linux/external:disable-static-ifs.patch",
            "//third_party/linux/external:enable-pmsg.patch",
            "//third_party/linux/external:lacp_fix.patch",
        ],
        sha256 = sums[version],
        strip_prefix = "linux-" + version,
        urls = ["https://cdn.kernel.org/pub/linux/kernel/v6.x/linux-%s.tar.xz" % version],
    )
