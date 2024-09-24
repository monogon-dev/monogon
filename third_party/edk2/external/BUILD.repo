filegroup(
    name = "all",
    srcs = glob(
        ["**"],
        exclude = [
            "CryptoPkg/Library/OpensslLib/openssl/boringssl/fuzz/*_corpus/**",
            "CryptoPkg/Library/OpensslLib/openssl/fuzz/corpora/**",
        ],
    ),
)

genrule(
    name = "firmware",
    srcs = [":all"],
    outs = [
        "OVMF_CODE.fd",
        "OVMF_VARS.fd",
    ],
    cmd = """
    (
        # The edk2 build does not like Bazel's default genrule environment.
        set +u

        cd {path}
        . edksetup.sh
        make -C BaseTools/Source/C
        build -DTPM2_ENABLE -DSECURE_BOOT_ENABLE -t GCC5 -a X64 -b RELEASE -p $$PWD/OvmfPkg/OvmfPkgX64.dsc
    ) > /dev/null

    cp {path}/Build/OvmfX64/RELEASE_GCC5/FV/OVMF_CODE.fd $(RULEDIR)
    cp {path}/Build/OvmfX64/RELEASE_GCC5/FV/OVMF_VARS.fd $(RULEDIR)
    """.format(path = package_relative_label(":all").workspace_root),
    visibility = ["//visibility:public"],
)
