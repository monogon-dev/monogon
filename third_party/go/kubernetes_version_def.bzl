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

# This reimplements k8s.io/kubernetes-compatible version_x_defs, while namespacing
# stamp variables with KUBERNETES_*.
# The generated build defs then are defines by the workspace status script, see
# //build/print-workspace-status.sh.

def version_x_defs():
    stamp_pkgs = [
        "k8s.io/component-base/version",
        "k8s.io/client-go/pkg/version",
    ]

    stamp_vars = [
        "buildDate",
        "gitCommit",
        "gitMajor",
        "gitMinor",
        "gitTreeState",
        "gitVersion",
    ]

    # Generate the cross-product.
    x_defs = {}
    for pkg in stamp_pkgs:
        for var in stamp_vars:
            x_defs["%s.%s" % (pkg, var)] = "{KUBERNETES_%s}" % var
    return x_defs
