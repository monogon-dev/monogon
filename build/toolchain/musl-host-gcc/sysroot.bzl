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

load("//build/toolchain/musl-host-gcc:sysroot_repository.bzl", "musl_sysroot_rule")

def musl_sysroot_repositories():
    """
    Provides an external repository that contains the extracted musl/linux sysroot.
    """
    musl_sysroot_rule(
        name = "musl_sysroot",
        snapshot = "//build/toolchain/musl-host-gcc:sysroot.tar.xz",
    )
