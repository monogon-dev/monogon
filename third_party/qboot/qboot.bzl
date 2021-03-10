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

cc_binary(
    name = "qboot-elf",
    srcs = [
        "code16.c",
        "code32seg.c",
        "cstart.S",
        "entry.S",
        "fw_cfg.c",
        "hwsetup.c",
        "linuxboot.c",
        "main.c",
        "malloc.c",
        "mptable.c",
        "pci.c",
        "printf.c",
        "string.c",
        "smbios.c",
        "tables.c",
        "benchmark.h",
    ] + glob(["include/*.h"]),
    copts = [
        "-m32",
        "-march=i386",
        "-mregparm=3",
        "-fno-stack-protector",
        "-fno-delete-null-pointer-checks",
        "-ffreestanding",
        "-mstringop-strategy=rep_byte",
        "-minline-all-stringops",
        "-fno-pic",
    ],
    features = ["-link_full_libc", "-cpp"],
    includes = [
        "include",
    ],
    additional_linker_inputs = [
        "flat.lds",
    ],
    linkopts = [
        "-nostdlib",
        "-m32",
        "-Wl,--build-id=none",
        "-Wl,-T$(location flat.lds)",
        "-no-pie",
    ],
)

# TODO(q3k): move to starlark rule for hermeticity, use toolchain objcopy
genrule(
    name = "qboot-bin",
    srcs = [":qboot-elf"],
    outs = ["bios.bin"],
    cmd = "objcopy -O binary $< $@",
    visibility = ["//visibility:public"],
)
