FROM docker.io/fedora:37@sha256:3487c98481d1bba7e769cf7bcecd6343c2d383fdd6bed34ec541b6b23ef07664

RUN dnf -y upgrade && \
	dnf -y install --setopt=tsflags=nodocs \
	java-11-openjdk-headless \
	dotnet-runtime-6.0 \
	jq \
	tini \
    golang && \
    dnf clean all

# (java-11-openjdk-headless and tini are required for the Jenkins CI agent)
# (dotnet-runtime-5.0 and jq are required for the GitHub Actions runner)
# (golang is required for us to build bazelisk - see below)

# Create CI build user. It is used by CI infrastructure to run build agents as.
# The newly created user will have a UID of 500, and a corresponding CI group
# of GID 500 will be created as well. This UID:GID pair's numeric values are
# relied on by the CI infrastructure and must not change without coordination.
RUN set -e -x ;\
	useradd -u 500 -U -m -d /home/ci ci

# Use a shared Go module cache for gazelle
# https://github.com/bazelbuild/bazel-gazelle/pull/535
ENV GO_REPOSITORY_USE_HOST_CACHE=1

USER ci

# Install Bazelisk. We do not have to pin a checksum since Go automatically
# relies on its built-in transparency log for pinning.
RUN go install github.com/bazelbuild/bazelisk@v1.15.0
ENV PATH="/home/ci/go/bin:${PATH}"

WORKDIR /work
