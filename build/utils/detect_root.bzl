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

# Copyright Google Inc.
# This file's contents are licensed under the Apache License, Version 2.0.
# See third_party/licenses/LICENSE.APACHE20 file in this repository for a copy
# of the License.

# This file contains code adapted from github.com/bazelbuiled/rules_foreign_cc:
# Files:
#  - tools/build_defs/detect_root.bzl

def detect_root(source):
    """Detects the path to the topmost directory of the 'source' outputs.
    To be used with external build systems to point to the source code/tools directories.
"""

    root = ""
    sources = source.files.to_list()
    if (root and len(root) > 0) or len(sources) == 0:
        return root

    root = ""
    level = -1
    num_at_level = 0

    # find topmost directory
    for file in sources:
        file_level = _get_level(file.path)
        if level == -1 or level > file_level:
            root = file.path
            level = file_level
            num_at_level = 1
        elif level == file_level:
            num_at_level += 1

    if num_at_level == 1:
        return root

    (before, sep, after) = root.rpartition("/")
    if before and sep and after:
        return before
    return root

def _get_level(path):
    normalized = path
    for _ in range(len(path)):
        new_normalized = normalized.replace("//", "/")
        if len(new_normalized) == len(normalized):
            break
        normalized = new_normalized

    return normalized.count("/")
