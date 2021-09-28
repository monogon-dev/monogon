FROM fedora:32

RUN dnf -y upgrade && \
	dnf -y install \
	"@Development Tools" \
	g++ \
	llvm \
	lld \
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
	openssl-devel \
	java-11-openjdk

# Create CI build user. This is not used by scripts/bin/bazel, but instead only
# used by CI infrastructure to run build agents as.
# The newly created user will have a UID of 500, and a corresponding CI group
# of GID 500 will be created as well. This UID:GID pair's numeric values are
# relied on by the CI infrastructure and must not change without coordination.
RUN set -e -x ;\
	useradd -u 500 -U -m -d /home/ci ci

# Install Bazel binary
RUN curl -o /usr/local/bin/bazel \
	https://releases.bazel.build/4.0.0/release/bazel-4.0.0-linux-x86_64 && \
	echo '7bee349a626281fc8b8d04a7a0b0358492712377400ab12533aeb39c2eb2b901  /usr/local/bin/bazel' | sha256sum --check && \
	chmod +x /usr/local/bin/bazel

# Use a shared Go module cache for gazelle
# https://github.com/bazelbuild/bazel-gazelle/pull/535
ENV GO_REPOSITORY_USE_HOST_CACHE=1

# Install ibazel (bazel-watcher)
RUN set -e -x ;\
    cd /tmp ;\
    git clone -b v0.15.10 https://github.com/bazelbuild/bazel-watcher ;\
    cd bazel-watcher ;\
    [ $(git rev-parse HEAD) == "84cab6f15f64850fb972ea88701e634c8b611301" ] ;\
    bazel --output_user_root /tmp/bazel-watcher-cache build //ibazel ;\
    cp bazel-bin/ibazel/linux_amd64_stripped/ibazel /usr/local/bin/ibazel ;\
    cd /tmp ;\
    rm -rf bazel-watcher bazel-watcher-cache

# --userns=keep-id uses the workdir as $HOME otherwise
RUN mkdir /user
ENV HOME=/user

WORKDIR /work
