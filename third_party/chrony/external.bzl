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

def chrony_external(name):
    # ONCHANGE(//third_party/chrony:chrony.bzl): version needs to be kept in sync
    version = "4.1"

    http_archive(
        name = name,
        sha256 = "ed76f2d3f9347ac6221a91ad4bd553dd0565ac188cd7490d0801d08f7171164c",
        build_file = "@//third_party/chrony:chrony.bzl",
        strip_prefix = "chrony-" + version,
        patch_args = ["-p1"],
        patches = [
            "//third_party/chrony/patches:disable_defaults.patch",
            "//third_party/chrony/patches:support_fixed_uids.patch",
        ],
        urls = ["https://download.tuxfamily.org/chrony/chrony-%s.tar.gz" % version],
    )
