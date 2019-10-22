#!/usr/bin/env bash
set -eo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$(realpath ${DIR}/../.vendor)

echo "Vendor build root: $ROOT"

if [ ! -d "$ROOT/linux" ] ; then
    echo "Please first call scripts/fetch_third_party.sh"
    exit 1
fi

if [ ! -d "$ROOT/edk2" ] ; then
    git clone --single-branch --branch edk2-stable201908 --depth=1 --recurse-submodules https://github.com/tianocore/edk2 $ROOT/edk2
fi

(
  cd $ROOT/edk2
  . edksetup.sh
  make -C BaseTools/Source/C
  build -DTPM2_ENABLE -DSECURE_BOOT_ENABLE -t GCC5 -a X64 -b RELEASE -p $PWD/OvmfPkg/OvmfPkgX64.dsc

  cp Build/OvmfX64/RELEASE_GCC5/FV/{OVMF_CODE.fd,OVMF_VARS.fd} $ROOT/../.artifacts
)

musl_prefix=$ROOT/musl-prefix

(
  cd $ROOT/linux
  make headers_install ARCH=x86_64 INSTALL_HDR_PATH=$musl_prefix
)

mkdir -p $ROOT/musl
curl -L https://www.musl-libc.org/releases/musl-1.1.23.tar.gz | tar -xzf - -C $ROOT/musl --strip-components 1

(
  cd $ROOT/musl

  ./configure --prefix=$musl_prefix --syslibdir=$musl_prefix/lib
  make -j8
  make install
)

mkdir -p $ROOT/util-linux
curl -L https://git.kernel.org/pub/scm/utils/util-linux/util-linux.git/snapshot/util-linux-2.34.tar.gz | tar -xzf - -C $ROOT/util-linux --strip-components 1

(
  cd $ROOT/util-linux
  ./autogen.sh
  ./configure \
    CC=$musl_prefix/bin/musl-gcc \
    --without-systemd \
    --without-udev \
    --without-btrfs \
    --disable-pylibmount \
    --without-tinfo \
    --prefix=$musl_prefix \
    --disable-makeinstall-chown \
    --disable-makeinstall-setuid \
    --with-bashcompletiondir=$musl_prefix/usr/share/bash-completion
  make -j8
  make install
)

mkdir -p $ROOT/xfsprogs-dev
curl -L https://git.kernel.org/pub/scm/fs/xfs/xfsprogs-dev.git/snapshot/xfsprogs-dev-5.2.1.tar.gz | tar -xzf - -C $ROOT/xfsprogs-dev --strip-components 1

(
  cd $ROOT/xfsprogs-dev
  patch -p1 < ../../patches/xfsprogs-dev/*.patch
  make configure
  ./configure CC=$musl_prefix/bin/musl-gcc "CFLAGS=-static -I$musl_prefix/include -L$musl_prefix/lib" "LDFLAGS=-L$musl_prefix/lib"
  make -j8 mkfs
  cp $ROOT/xfsprogs-dev/mkfs/mkfs.xfs $ROOT/../.artifacts
)
