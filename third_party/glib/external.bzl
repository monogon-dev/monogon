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

def glib_external(name, version):
    sums = {
        "2.67.5": "410966db712638dc749054c0a3c3087545d5106643139c25806399a51a8d4ab1",
    }

    http_archive(
        name = name,
        patch_args = ["-p1", "-u"],
        patches = [
            "//third_party/glib/patches:bazel_cc_fix.patch",
            "//third_party/glib/patches:bazel_support.patch",
        ],
        sha256 = sums[version],
        strip_prefix = "glib-" + version,
        # We cannot use the actual release tarball as it contains files generated incorrectly for our environment
        urls = ["https://gitlab.gnome.org/GNOME/glib/-/archive/%s/glib-%s.tar.gz" % (version, version)],
    )
