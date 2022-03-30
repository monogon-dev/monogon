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

def seccomp_external(name, version):
    # NOTE: Remember to update seccomp.bzl's seccomp.h template rule
    # with the correct version.
    sums = {
        "2.5.1": "76ad54e31d143b39a99083564045212a965e026a1010a742edd793d26d699829",
    }

    http_archive(
        name = name,
        patch_args = ["-p1"],
        patches = [
            "//third_party/seccomp/patches:bazel_cc_fix.patch",
            "//third_party/seccomp/patches:fix_generated_includes.patch",
        ],
        sha256 = sums[version],
        build_file = "@//third_party/seccomp:seccomp.bzl",
        strip_prefix = "libseccomp-" + version,
        # We cannot use the actual release tarball as it contains files generated incorrectly for our environment
        urls = ["https://github.com/seccomp/libseccomp/archive/v%s.tar.gz" % version],
    )
