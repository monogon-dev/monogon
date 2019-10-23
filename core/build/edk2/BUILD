genrule(
    name = "firmware",
    srcs = glob(["**"]),
    outs = [
        "OVMF_CODE.fd",
        "OVMF_VARS.fd",
    ],
    cmd = """
    (
        # The edk2 build does not like Bazel's default genrule environment.
        set +u

        cd external/edk2
        . edksetup.sh
        make -C BaseTools/Source/C
        build -DTPM2_ENABLE -DSECURE_BOOT_ENABLE -t GCC5 -a X64 -b RELEASE -p $$PWD/OvmfPkg/OvmfPkgX64.dsc
    ) > /dev/null

    cp external/edk2/Build/OvmfX64/RELEASE_GCC5/FV/{OVMF_CODE.fd,OVMF_VARS.fd} $(RULEDIR)
    """,
    visibility = ["//visibility:public"],
)
