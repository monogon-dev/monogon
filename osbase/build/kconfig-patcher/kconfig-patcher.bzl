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

"""Override configs in a Linux kernel Kconfig
"""

def kconfig_patch(name, src, out, override_configs, **kwargs):
    native.genrule(
        name = name,
        srcs = [src],
        outs = [out],
        tools = [
            "//osbase/build/kconfig-patcher",
        ],
        cmd = """
        $(location //osbase/build/kconfig-patcher) \
            -in $< -out $@ '%s'
        """ % json.encode(struct(overrides = override_configs)),
        **kwargs
    )
