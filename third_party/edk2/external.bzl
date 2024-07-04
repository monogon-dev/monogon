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

load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")

def edk2_external(name):
    new_git_repository(
        name = name,
        build_file = "//third_party/edk2/external:BUILD.repo",
        commit = "3e722403cd16388a0e4044e705a2b34c841d76ca",  # stable202405
        recursive_init_submodules = True,
        remote = "https://github.com/tianocore/edk2",
        patches = ["//third_party/edk2/patches:disable-werror.patch", "//third_party/edk2/patches:remove-brotli-build.patch"],
        patch_args = ["-p1"],
    )
