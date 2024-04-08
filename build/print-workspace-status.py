#!/usr/bin/env python3
"""Workspace status script used for build stamping."""

# Treat this script as shell code, but with Python syntax. We want to remain as
# simple as possible, and absolutely never use any non-standard Python library.
# This script should be able to run on any 'modern' Linux distribution with
# Python 3.8 or newer.

# The following versioning concepts apply:
# 1. Version numbers follow the Semantic Versioning 2.0 spec.
# 2. Git tags in the form `<product>-vX.Y.Z` will be used as a basis for
#    versioning a build. If the currently built release is exactly the same as
#    such a tag, it will be versioned at vX.Y.Z. Otherwise, a devNNN prerelease
#    identifier will be appended to signify the amount of commits since the
#    release.
# 3. Product git tags are only made up of a major/minor/patch version.
#    Prerelease and build tags are assigned by the build system and this
#    script, Git tags have no influence on them.
# 4. 'Products' are release numbering trains within the Monogon monorepo. This
#    means there is no such thing as a 'version' for the monorepo by itself,
#    only within the context of some product.

from dataclasses import dataclass
from datetime import datetime, timezone
import os
import re
import subprocess
import time

from typing import Optional


# Variables to output. These will be printed to stdout at the end of the script
# runtime, sorted by key.
variables: dict[str, str] = {}

# Git build tree status: clean or dirty.
git_tree_state: str = "clean"
git_status = subprocess.check_output(["git", "status", "--porcelain"])
if git_status.decode().strip() != "":
    git_tree_state = "dirty"

# Git commit hash.
git_commit: str = (
    subprocess.check_output(["git", "rev-parse", "HEAD^{commit}"]).decode().strip()
)

# Git tags pointing at this commit.
git_tags_b: [bytes] = subprocess.check_output(
    ["git", "tag", "--sort=-version:refname", "--points-at", "HEAD"]
).split(b"\n")
git_tags: [str] = [t.decode().strip() for t in git_tags_b if t.decode().strip() != ""]

# Build timestamp, respecting SOURCE_DATE_EPOCH for reproducible builds.
build_timestamp = int(time.time())
sde = os.environ.get("SOURCE_DATE_EPOCH")
if sde is not None:
    build_timestamp = int(sde)

variables["STABLE_MONOGON_gitCommit"] = git_commit
variables["STABLE_MONOGON_gitTreeState"] = git_tree_state

# Per product. Each product has it's own semver-style version number, which is
# deduced from git tags.
#
# For example: metropolis v. 1.2.3 would be tagged 'metropolis-v1.2.3'.
@dataclass
class Version:
    """Describes a semver version for a given product."""

    product: str
    version: str
    prerelease: [str]

    def __str__(self) -> str:
        ver = self.version
        if self.prerelease:
            ver += "-" + ".".join(self.prerelease)
        return ver


def parse_tag(tag: str, product: str) -> Optional[Version]:
    prefix = product + "-"
    if not tag.startswith(prefix):
        return None
    version = tag[len(prefix) :]
    # The first release of Metropolis was v0.1, which we extend to v0.1.0.
    if product == "metropolis" and version == "v0.1":
        version = "v0.1.0"
    # Only care about the limited major/minor/patch subset of semver from git
    # tags. All prerelease identifies will be appended by this code.
    if not re.match(r"^v[0-9]+\.[0-9]+\.[0-9]+$", version):
        return None
    return Version(product, version, [])


for product in ["metropolis", "cloud"]:
    # Get exact version from tags.
    version = None
    for tag in git_tags:
        version = parse_tag(tag, product)
        if version is not None:
            break

    if version is None:
        # No exact version found. Use latest tag for the given product and
        # append a 'devXXX' identifier based on number of commits since that
        # tag.
        for tag in (
            subprocess.check_output(
                ["git", "tag", "--sort=-version:refname", "--merged", "HEAD"]
            )
            .decode()
            .strip()
            .split("\n")
        ):
            version = parse_tag(tag, product)
            if version is None:
                continue
            # Found the latest tag for this product. Augment it with the
            # devXXX identifier and add it to our versions.
            count = (
                subprocess.check_output(["git", "rev-list", tag + "..HEAD", "--count"])
                .decode()
                .strip()
            )
            version.prerelease.append(f"dev{count}")
            break

    if version is None:
        # This product never had a release! Use v0.0.0 as a fallback.
        version = Version(product, "v0.0.0", [])
        # ... and count the number of all commits ever to use as the devXXX
        # prerelease identifier.
        count = (
            subprocess.check_output(["git", "rev-list", "HEAD", "--count"])
            .decode()
            .strip()
        )
        version.prerelease.append(f"dev{count}")

    version.prerelease.append(f"g{git_commit[:8]}")
    if git_tree_state == "dirty":
        version.prerelease.append("dirty")
    variables[f"STABLE_MONOGON_{product}_version"] = str(version)


# Special treatment for Kubernetes, which uses these stamp values in its build
# system. We populate the Kubernetes version from whatever is in
# //third_party/go/repositories.bzl.
def parse_repositories_bzl(path: str) -> dict[str, str]:
    """
    Shoddily parse a Gazelle-created repositories.bzl into a map of
    name->version.

    This relies heavily on repositories.bzl being correctly formatted and
    sorted.

    If this breaks, it's probably best to try to use the actual Python parser
    to deal with this, eg. by creating a fake environment for the .bzl file to
    be parsed.
    """

    # Main parser state: None where we don't expect a version line, set to some
    # value otherwise.
    name: Optional[str] = None

    res = {}
    for line in open(path):
        line = line.strip()
        if line == "go_repository(":
            name = None
            continue
        if line.startswith("name ="):
            if name is not None:
                raise Exception("parse error in repositories.bzl: repeated name?")
            if line.count('"') != 2:
                raise Exception(
                    "parse error in repositories.bzl: invalid name line: " + name
                )
            name = line.split('"')[1]
            continue
        if line.startswith("version ="):
            if name is None:
                raise Exception("parse error in repositories.bzl: version before name")
            if line.count('"') != 2:
                raise Exception(
                    "parse error in repositories.bzl: invalid name line: " + name
                )
            version = line.split('"')[1]
            res[name] = version
            name = None
    return res


# Parse repositories.bzl.
go_versions = parse_repositories_bzl("third_party/go/repositories.bzl")

# Find Kubernetes version.
kubernetes_version: str = go_versions.get("io_k8s_kubernetes")
if kubernetes_version is None:
    raise Exception("could not figure out Kubernetes version")
kubernetes_version_parsed = re.match(
    r"^v([0-9]+)\.([0-9]+)\.[0-9]+$", kubernetes_version
)
if not kubernetes_version_parsed:
    raise Exception("invalid Kubernetes version: " + kubernetes_version)

# The Kubernetes build tree is considered clean iff the monorepo build tree is
# considered clean.
variables["KUBERNETES_gitTreeState"] = git_tree_state
variables["KUBERNETES_buildDate"] = datetime.fromtimestamp(
    build_timestamp, timezone.utc
).strftime("%Y-%m-%dT%H:%M:%SZ")
variables["STABLE_KUBERNETES_gitMajor"] = kubernetes_version_parsed[1]
variables["STABLE_KUBERNETES_gitMinor"] = kubernetes_version_parsed[2]
variables["STABLE_KUBERNETES_gitVersion"] = kubernetes_version + "+mngn"

# Emit variables to stdout for consumption by Bazel and targets.
for key in sorted(variables.keys()):
    print("{} {}".format(key, variables[key]))
