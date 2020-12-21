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

def ktest(deps, tester, initramfs_extra, cmdline):
    native.genrule(
        name = "test_initramfs",
        srcs = [
            "//metropolis/test/ktest/init",
        ] + deps + [tester],
        outs = [
            "initramfs.cpio.lz4",
        ],
        testonly = True,
        cmd = """
        $(location @linux//:gen_init_cpio) - <<- 'EOF' | lz4 -l > \"$@\" 
dir /dev 0755 0 0
nod /dev/console 0600 0 0 c 5 1
nod /dev/null 0644 0 0 c 1 3
file /init $(location //metropolis/test/ktest/init) 0755 0 0
file /tester $(location """ + tester + """) 0755 0 0
""" + initramfs_extra + """
EOF
        """,
        tools = [
            "@linux//:gen_init_cpio",
        ],
    )

    native.sh_test(
        name = "ktest",
        args = [
            "$(location //metropolis/test/ktest)",
            "$(location :test_initramfs)",
            "$(location //metropolis/test/ktest:linux-testing)",
            cmdline,
        ],
        size = "small",
        srcs = ["//metropolis/test/ktest:test-script"],
        data = [
            "//metropolis/test/ktest",
            ":test_initramfs",
            "//metropolis/test/ktest:linux-testing",
            "@com_github_bonzini_qboot//:qboot-bin",
        ],
    )
