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

"""
Ktest provides a macro to run tests under a normal Metropolis node kernel
"""

load("//metropolis/node/build:def.bzl", "node_initramfs")

def _dict_union(x, y):
    z = {}
    z.update(x)
    z.update(y)
    return z

def ktest(tester, cmdline = "", files = {}, fsspecs = [], files_cc = {}):
    node_initramfs(
        name = "test_initramfs",
        fsspecs = [
            "//metropolis/node/build:earlydev.fsspec",
        ] + fsspecs,
        files = _dict_union({
            "//osbase/test/ktest/init": "/init",
            tester: "/tester",
        }, files),
        files_cc = files_cc,
        testonly = True,
    )

    native.sh_test(
        name = "ktest",
        args = [
            "$(location //osbase/test/ktest)",
            "$(location :test_initramfs)",
            "$(location //osbase/test/ktest:linux-testing)",
            cmdline,
        ],
        size = "small",
        srcs = ["//osbase/test/ktest:test-script"],
        data = [
            "//osbase/test/ktest",
            ":test_initramfs",
            "//osbase/test/ktest:linux-testing",
        ],
    )
