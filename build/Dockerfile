#
# The CI only rebuilds this Dockerfile if its hash changes.
# Do not reference any external files, since modifications to them won't trigger a rebuild.
#

FROM fedora:32

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
	expect \
	grpc-cli \
	nc \
	python-unversioned-command \
	openssl-devel

# Install Bazel binary
RUN curl -o /usr/local/bin/bazel \
	https://releases.bazel.build/3.7.2/release/bazel-3.7.2-linux-x86_64 && \
	echo '70dc0bee198a4c3d332925a32d464d9036a831977501f66d4996854ad4e4fc0d  /usr/local/bin/bazel' | sha256sum --check && \
	chmod +x /usr/local/bin/bazel

# Use a shared Go module cache for gazelle
# https://github.com/bazelbuild/bazel-gazelle/pull/535
ENV GO_REPOSITORY_USE_HOST_CACHE=1

# --userns=keep-id uses the workdir as $HOME otherwise
RUN mkdir /user
ENV HOME=/user

WORKDIR /work
