load(
    "@bazel_tools//tools/build_defs/repo:utils.bzl",
    "patch",
    "update_attrs",
    "workspace_and_buildfile",
)

def _build_archive_url(owner, repo, ref):
    return "https://github.com/{owner}/{repo}/archive/{ref}.tar.gz".format(
        owner = owner,
        repo = repo,
        ref = ref,
    ), "{repo}-{ref}".format(repo = repo, ref = ref)

def build_submodule_info_url(owner, repo, submodule, ref):
    return "https://api.github.com/repos/{owner}/{repo}/contents/{submodule}?ref={ref}".format(
        owner = owner,
        repo = repo,
        submodule = submodule,
        ref = ref,
    )

def parse_github_url(url):
    url = url.removeprefix("https://github.com/")
    url = url.removesuffix(".git")
    (owner, repo) = url.split("/")
    return owner, repo

def _github_repository(ctx):
    base_repo_archive_url, base_repo_archive_prefix = _build_archive_url(
        owner = ctx.attr.owner,
        repo = ctx.attr.repo,
        ref = ctx.attr.ref,
    )

    base_repo_download_info = ctx.download_and_extract(
        url = base_repo_archive_url,
        stripPrefix = base_repo_archive_prefix,
        integrity = ctx.attr.integrity,
        type = "tar.gz",
    )

    for submodule, integrity in ctx.attr.submodules.items():
        if submodule not in ctx.attr.submodule_info:
            url = build_submodule_info_url(
                owner = ctx.attr.owner,
                repo = ctx.attr.repo,
                ref = ctx.attr.ref,
                submodule = submodule,
            )

            submodule_info_path = submodule + ".submodule_info"
            ctx.download(
                url = url,
                headers = {
                    "Accept": "application/vnd.github+json",
                    "X-GitHub-Api-Version": "2022-11-28",
                },
                output = submodule_info_path,
            )

            submodule_info = json.decode(ctx.read(submodule_info_path))

            # buildifier: disable=print
            print("Missing submodule_info for submodule %s. Consider adding it: \n%s" % (submodule, submodule_info))
        else:
            submodule_info = json.decode(ctx.attr.submodule_info[submodule])

        if submodule_info["type"] != "submodule":
            fail("provided submodule path is not a submodule")

        submodule_owner, submodule_repo = parse_github_url(
            url = submodule_info["submodule_git_url"],
        )

        submodule_url, submodule_strip_prefix = _build_archive_url(
            owner = submodule_owner,
            repo = submodule_repo,
            ref = submodule_info["sha"],
        )

        download_info = ctx.download_and_extract(
            url = submodule_url,
            stripPrefix = submodule_strip_prefix,
            integrity = integrity,
            type = "tar.gz",
            output = submodule_info["path"],
        )
        if integrity == "":
            # buildifier: disable=print
            print("Missing integrity for submodule {submodule}. Consider adding it:\n\"{submodule}\": \"{integrity}\".".format(
                submodule = submodule,
                integrity = download_info.integrity,
            ))

    workspace_and_buildfile(ctx)

    patch(ctx)

    return update_attrs(ctx.attr, _github_repository_attrs.keys(), {"integrity": base_repo_download_info.integrity})

_github_repository_attrs = {
    "owner": attr.string(
        mandatory = True,
        doc = "The Owner of the Github repository",
    ),
    "repo": attr.string(
        mandatory = True,
        doc = "The Name of Github repository",
    ),
    "submodules": attr.string_dict(
        mandatory = False,
        default = {},
        doc = "The list of submodules with their integrity as value",
    ),
    "submodule_info": attr.string_dict(
        mandatory = False,
        default = {},
        doc = """The list of submodules with their GitHub API response as value.
        This is a workaround until https://github.com/bazelbuild/bazel/issues/24777
        is implemented to prevent hitting GitHub API limits.""",
    ),
    "ref": attr.string(
        default = "",
        doc =
            "The specific ref to be checked out.",
    ),
    "integrity": attr.string(
        doc = """Expected checksum in Subresource Integrity format of the file downloaded.

    This must match the checksum of the file downloaded. _It is a security risk
    to omit the checksum as remote files can change._ At best omitting this
    field will make your build non-hermetic. It is optional to make development
    easier but either this attribute or `sha256` should be set before shipping.""",
    ),
    "patches": attr.label_list(
        default = [],
        doc =
            "A list of files that are to be applied as patches after " +
            "extracting the archive. By default, it uses the Bazel-native patch implementation " +
            "which doesn't support fuzz match and binary patch, but Bazel will fall back to use " +
            "patch command line tool if `patch_tool` attribute is specified or there are " +
            "arguments other than `-p` in `patch_args` attribute.",
    ),
    "patch_tool": attr.string(
        default = "",
        doc = "The patch(1) utility to use. If this is specified, Bazel will use the specified " +
              "patch tool instead of the Bazel-native patch implementation.",
    ),
    "patch_args": attr.string_list(
        default = ["-p0"],
        doc =
            "The arguments given to the patch tool. Defaults to -p0, " +
            "however -p1 will usually be needed for patches generated by " +
            "git. If multiple -p arguments are specified, the last one will take effect." +
            "If arguments other than -p are specified, Bazel will fall back to use patch " +
            "command line tool instead of the Bazel-native patch implementation. When falling " +
            "back to patch command line tool and patch_tool attribute is not specified, " +
            "`patch` will be used. This only affects patch files in the `patches` attribute.",
    ),
    "patch_cmds": attr.string_list(
        default = [],
        doc = "Sequence of Bash commands to be applied on Linux/Macos after patches are applied.",
    ),
    "build_file": attr.label(
        allow_single_file = True,
        doc =
            "The file to use as the BUILD file for this repository." +
            "This attribute is an absolute label (use '@//' for the main " +
            "repo). The file does not need to be named BUILD, but can " +
            "be (something like BUILD.new-repo-name may work well for " +
            "distinguishing it from the repository's actual BUILD files. " +
            "Either build_file or build_file_content can be specified, but " +
            "not both.",
    ),
    "build_file_content": attr.string(
        doc =
            "The content for the BUILD file for this repository. " +
            "Either build_file or build_file_content can be specified, but " +
            "not both.",
    ),
    "workspace_file": attr.label(
        doc =
            "The file to use as the `WORKSPACE` file for this repository. " +
            "Either `workspace_file` or `workspace_file_content` can be " +
            "specified, or neither, but not both.",
    ),
    "workspace_file_content": attr.string(
        doc =
            "The content for the WORKSPACE file for this repository. " +
            "Either `workspace_file` or `workspace_file_content` can be " +
            "specified, or neither, but not both.",
    ),
}

github_repository = repository_rule(
    implementation = _github_repository,
    attrs = _github_repository_attrs,
)
