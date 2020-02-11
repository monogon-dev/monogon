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

def kubernetes_external(name, version):
    sums = {
        "1.16.4": "3a49373ba56c73c282deb0cfa2ec7bfcc6bf46acb6992f01319eb703cbf68996",
    }
    http_archive(
        name = name,
        patch_args = ["-p1"],
        patches = [
            "//third_party/kubernetes/external:0001-avoid-unexpected-keyword-error-by-using-positional-p.patch",
        ],
        sha256 = sums[version],
        urls = ["https://dl.k8s.io/v%s/kubernetes-src.tar.gz" % version],
    )
