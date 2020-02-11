#
# The CI only rebuilds this Dockerfile if its hash changes.
# Do not reference any external files, since modifications to them won't trigger a rebuild.
#

FROM fedora:30

RUN dnf -y upgrade && \
    dnf -y install \
	"@Development Tools" \
	g++ \
	libuuid-devel \
	python3 \
	nasm \
	acpica-tools \
	gettext-devel \
	autoconf \
	bison \
	libtool \
	automake \
	flex \
	glibc-static \
	elfutils-libelf-devel \
	libblkid-devel \
	lz4 \
	bc \
	hostname \
	which \
	swtpm-tools \
	rsync \
	qemu-system-x86-core \
	postgresql \
	expect \
	grpc-cli \
	nc \
	python-unversioned-command

# Workaround for a binutils bugs in F30, which generates invalid ELF binaries
# when linking statically with musl.
RUN dnf -y install fedora-repos-rawhide && \
	dnf -y --disablerepo=* --enablerepo=rawhide --releasever=32 upgrade binutils

# Install Bazel binary
RUN curl -o /usr/local/bin/bazel \
    https://releases.bazel.build/2.0.0/release/bazel-2.0.0-linux-x86_64 && \
    echo '4df79462c6c3ecdeeee7af99fc269b52ab1aa4828ef3bc359c1837d3fafeeee7  /usr/local/bin/bazel' | sha256sum --check && \
    chmod +x /usr/local/bin/bazel

# Use a shared Go module cache for gazelle
# https://github.com/bazelbuild/bazel-gazelle/pull/535
ENV GO_REPOSITORY_USE_HOST_CACHE=1

# --userns=keep-id uses the workdir as $HOME otherwise
RUN mkdir /user
ENV HOME=/user

WORKDIR /work
