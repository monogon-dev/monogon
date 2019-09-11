#!/usr/bin/env bash
set -eo pipefail

if [ ! -d "$root/linux" ] ; then
    echo "Please first call scripts/fetch_third_party.sh"
fi

root=$(git rev-parse --show-toplevel)/third_party

# nasm + Python 3.7 + iasl
if [ ! -d "$root/edk2" ] ; then
    git clone --recurse-submodules https://github.com/tianocore/edk2 $root/edk2
fi
cd $root/edk2
git checkout --recurse-submodules edk2-stable201908
. edksetup.sh
make -C $root/edk2/BaseTools/Source/C
build -DTPM2_ENABLE -DSECURE_BOOT_ENABLE -t GCC5 -a X64 -b RELEASE -p $PWD/OvmfPkg/OvmfPkgX64.dsc

musl_prefix=$root/musl-prefix

cd $root/linux
make headers_install ARCH=x86_64 INSTALL_HDR_PATH=$musl_prefix

mkdir -p $root/musl
curl -L https://www.musl-libc.org/releases/musl-1.1.23.tar.gz | tar -xzf - -C $root/musl --strip-components 1
cd $root/musl

./configure --prefix=$musl_prefix --syslibdir=$musl_prefix/lib
make -j8
make install

mkdir -p $root/util-linux
curl -L https://git.kernel.org/pub/scm/utils/util-linux/util-linux.git/snapshot/util-linux-2.34.tar.gz | tar -xzf - -C $root/util-linux --strip-components 1
cd $root/util-linux
./autogen.sh
./configure CC=$musl_prefix/bin/musl-gcc --without-systemd --without-udev --without-btrfs --disable-pylibmount --without-tinfo --prefix=$musl_prefix --disable-makeinstall-chown --disable-makeinstall-setuid --with-bashcompletiondir=$musl_prefix/usr/share/bash-completion
make -j8
make install

mkdir -p $root/xfsprogs-dev
curl -L https://git.kernel.org/pub/scm/fs/xfs/xfsprogs-dev.git/snapshot/xfsprogs-dev-5.2.1.tar.gz | tar -xzf - -C $root/xfsprogs-dev --strip-components 1
cd $root/xfsprogs-dev
patch -p1 < ../../patches/xfsprogs-dev/*.patch
./configure CC=$musl_prefix/bin/musl-gcc "CFLAGS=-static -I$musl_prefix/include -L$musl_prefix/lib" "LDFLAGS=-L$musl_prefix/lib"
make -j8 mkfs
cp $root/xfsprogs-dev/mkfs/mkfs.xfs