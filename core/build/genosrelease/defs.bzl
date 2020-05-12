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

def _os_release_impl(ctx):
    ctx.actions.run(
        mnemonic = "GenOSRelease",
        progress_message = "Generating os-release",
        inputs = [ctx.info_file],
        outputs = [ctx.outputs.out],
        executable = ctx.executable._genosrelease,
        arguments = [
            "-status_file",
            ctx.info_file.path,
            "-out_file",
            ctx.outputs.out.path,
            "-stamp_var",
            ctx.attr.stamp_var,
            "-name",
            ctx.attr.os_name,
            "-id",
            ctx.attr.os_id,
        ],
    )

os_release = rule(
    implementation = _os_release_impl,
    attrs = {
        "os_name": attr.string(mandatory = True),
        "os_id": attr.string(mandatory = True),
        "stamp_var": attr.string(mandatory = True),
        "_genosrelease": attr.label(
            default = Label("//core/build/genosrelease"),
            cfg = "host",
            executable = True,
            allow_files = True,
        ),
    },
    outputs = {
        "out": "os-release",
    },
)
