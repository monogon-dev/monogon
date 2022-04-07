#!/usr/bin/env bash
set -euo pipefail
shopt -s nullglob

main() {
    # Our local user needs write access to /dev/kvm (best accomplished by
    # adding your user to the kvm group).
    if ! touch /dev/kvm; then
      echo "Cannot write to /dev/kvm - please verify permissions." >&2
      exit 1
    fi
    
    # The KVM module needs to be loaded, since our container is unprivileged
    # and won't be able to do it itself.
    if ! [[ -d /sys/module/kvm ]]; then
      echo "kvm module not loaded - please modprobe kvm" >&2
      exit 1
    fi

    local dockerfile="build/ci/Dockerfile"
    if ! [[ -f "${dockerfile}" ]]; then
        echo "Dockerfile not found at path ${dockerfile}. Make sure to run this script from the root of the Monogon checkout." >&2
        exit 1
    fi
    
    # Rebuild base image purely with no build context (-) ensuring that the
    # builder image does not contain any other data from the repository.
    podman build -t gcr.io/monogon-infra/monogon-builder - < "${dockerfile}"

    # TODO(serge): stop using pods for the builder, this is a historical artifact.
    podman pod create --name monogon
    
    # Mount bazel root to identical paths inside and outside the container.
    # This caches build state even if the container is destroyed.
    #
    # TODO(serge): do not hardcode this path? This breaks if attempting to use
    # the build container from multiple Monogon checkouts on disk.
    local bazel_root="${HOME}/.cache/bazel-monogon"
    mkdir -p "${bazel_root}"
    
    # When IntelliJ's Bazel plugin uses //scripts/bin/bazel to either build targets
    # or run syncs, it adds a --override_repository flag to the bazel command
    # line that points @intellij_aspect into a path on the filesystem. This
    # external repository contains a Bazel Aspect definition which Bazel
    # executes to provide the IntelliJ Bazel plugin with information about the
    # workspace / build targets / etc...
    #
    # We need to make this path available within the build container. However,
    # instead of directly pointing into that path on the host, we redirect through
    # a patched copy of this repository. That patch applies fixes related to the
    # Monogon codebase in general, not specific to the fact that we're running in a
    # container.
    #
    # What this ends up doing is that the path mounted within the container
    # looks as if it's a path straight from the host IntelliJ config directory,
    # but in fact points into a patched copy. It looks weird, but this setup
    # allows us to let IntelliJ's Bazel integration to trigger Bazel without us
    # having to replace the override_repository flag that it passes to Bazel.
    # We at some point used to do that, but parsing and replacing Bazel flags
    # in the scripts/bin/bazel wrapper script is error prone and fragile.

    # Find all IntelliJ installation/config directories.
    local ij_home_paths=("${HOME}/.local/share/JetBrains/IntelliJIdea"*)
    # Get the newest one, if any.
    local ij_home=""
    if ! [[  ${#ij_home_paths[@]} -eq 0 ]]; then
        # Reverse sort paths by name, with the first being the newest IntelliJ
        # installation.
        IFS=$'\n'
        local sorted=($(sort -r <<<"${ij_home_paths[*]}"))
        unset IFS
        ij_home="${sorted[0]}"
    fi

    # If we don't have or can't find ij_home, don't bother with attempting to patch anything.
    # If we do, podman_flags will get populated with extra flags that it will
    # run with to support IntelliJ/Bazel integration.
    declare -a podman_flags
    if [[ -d "${ij_home}" ]]; then
        echo "IntelliJ found at ${ij_home}, patching and mounting aspect repository."
        # aspect_orig is the path to the aspect external repository that IntelliJ will
        # inject into bazel via --override_repository. It is the path that it expects
        # to be available to Bazel within the container, and also the path that is
        # directly visible in the host.
        local aspect_orig="${ij_home}/ijwb/aspect"
        # aspect_patched is the path to our patched copy of the aspect
        # repository. We keep this in bazel_root for convenience, as that's a
        # path that we control on the host anyway, so we can be sure we're not
        # trampling some other process.
        local aspect_patched="${bazel_root}/ijwb_aspect"

        # If we already have a patched version of the aspect, remove it.
        if [[ -d "${aspect_patched}" ]]; then
            rm -rf "${aspect_patched}"
        fi

        # Copy and patch the aspect.
        cp -r "${aspect_orig}" "${aspect_patched}"
        patch -d "${aspect_patched}" -p1 < scripts/patches/bazel_intellij_aspect_filter.patch

        # Make podman mount the patched aspect into the original aspect path in the build container.
        podman_flags+=(-v "${aspect_patched}:${aspect_orig}")
    else
        echo "No IntelliJ found, not patching/mounting aspect repository."
    fi

    podman run -it -d \
        -v $(pwd):$(pwd):z \
        -w $(pwd) \
        --tmpfs=/tmp \
        --volume="${bazel_root}:${bazel_root}" \
        --device /dev/kvm \
        --privileged \
        --pod monogon \
        --name=monogon-dev \
        --net=host \
        "${podman_flags[@]}" \
        gcr.io/monogon-infra/monogon-builder
}

main
