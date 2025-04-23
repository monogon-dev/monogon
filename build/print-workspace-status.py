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
import re
import subprocess

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

# Git commit date.
git_commit_date: str = (
    subprocess.check_output(["git", "show", "--pretty=format:%cI", "--no-patch", "HEAD"]).decode().strip()
)

# Git tags pointing at this commit.
git_tags_b: [bytes] = subprocess.check_output(
    ["git", "tag", "--sort=-version:refname", "--points-at", "HEAD"]
).split(b"\n")
git_tags: [str] = [t.decode().strip() for t in git_tags_b if t.decode().strip() != ""]

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


# Is this a release build of the given product?
is_release: dict[str, bool] = {}


for product in ["metropolis", "cloud"]:
    # Get exact version from tags.
    version = None
    for tag in git_tags:
        version = parse_tag(tag, product)
        if version is not None:
            break

    is_release[product] = version is not None and git_tree_state == "clean"

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
# //go.mod.
def parse_go_mod(path: str) -> dict[str, str]:
    """
    Shoddily parse a go.mod into a map of name->version.

    This relies heavily on go.mod being correctly formatted and
    sorted.

    If this breaks, it's probably best to try to port this to Go
    and parse it using golang.org/x/mod/modfile, shell out to
    "go mod edit -json", or similar.
    """

    # Just a copied together regex to find the url followed by a semver.
    NAME_VERSION_REGEX = r"([-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*) v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?)"

    res = {}
    for line in open(path):
        matches = re.findall(NAME_VERSION_REGEX, line)
        if not matches:
            continue

        [name, version] = matches[0][0].strip().split(" ")

        # If we already saw a package, skip it.
        if name in res:
            continue

        res[name] = version

    return res


# Parse go.mod.
go_versions = parse_go_mod("go.mod")

# Find Kubernetes version.
kubernetes_version: str = go_versions.get("k8s.io/kubernetes")
if kubernetes_version is None:
    raise Exception("could not figure out Kubernetes version")
kubernetes_version_parsed = re.match(
    r"^v([0-9]+)\.([0-9]+)\.[0-9]+$", kubernetes_version
)
if not kubernetes_version_parsed:
    raise Exception("invalid Kubernetes version: " + kubernetes_version)

variables["STABLE_KUBERNETES_gitMajor"] = kubernetes_version_parsed[1]
variables["STABLE_KUBERNETES_gitMinor"] = kubernetes_version_parsed[2]
variables["STABLE_KUBERNETES_gitVersion"] = kubernetes_version + "+mngn"

# Stamp commit info into Kubernetes only for release builds, to avoid
# unnecessary rebuilds of hyperkube during development.
if is_release["metropolis"]:
    variables["STABLE_KUBERNETES_gitCommit"] = git_commit
    variables["STABLE_KUBERNETES_gitTreeState"] = git_tree_state
    variables["STABLE_KUBERNETES_buildDate"] = git_commit_date

# Emit variables to stdout for consumption by Bazel and targets.
for key in sorted(variables.keys()):
    print("{} {}".format(key, variables[key]))
