"""Rules for generating EFI unified kernel images. These are EFI-bootable PE/COFF files containing a stub loader,
a kernel, and optional commandline and initramfs in one file.
See https://systemd.io/BOOT_LOADER_SPECIFICATION/#type-2-efi-unified-kernel-images for more information.
"""

load("//build/toolchain/llvm-efi:transition.bzl", "build_efi_transition")

def _efi_unified_kernel_image_impl(ctx):
    out = ctx.actions.declare_file(ctx.attr.name + ".efi")

    toolchain_info = ctx.attr._toolchain[platform_common.ToolchainInfo]

    sections = [
        dict(name = ".linux", file = ctx.file.kernel, vma = 0x2000000),
    ]

    if ctx.attr.cmdline != "":
        cmdline_file = ctx.actions.declare_file("cmdline")
        ctx.actions.write(
            output = cmdline_file,
            content = ctx.attr.cmdline,
        )
        sections.append(dict(name = ".cmdline", file = cmdline_file, vma = 0x30000))

    if ctx.file.initramfs:
        sections += [dict(name = ".initrd", file = ctx.file.initramfs, vma = 0x5000000)]
    if ctx.file.os_release:
        sections += [dict(name = ".osrel", file = ctx.file.os_release, vma = 0x20000)]
    if ctx.file.splash:
        sections += [dict(name = ".splash", file = ctx.file.splash, vma = 0x40000)]

    args = []
    for sec in sections:
        args.append("--add-section")
        args.append("{}={}".format(sec["name"], sec["file"].path))
        args.append("--change-section-vma")
        args.append("{}={}".format(sec["name"], sec["vma"]))

    ctx.actions.run(
        mnemonic = "GenEFIKernelImage",
        progress_message = "Generating EFI unified kernel image",
        inputs = [ctx.file.stub] + [s["file"] for s in sections],
        outputs = [out],
        executable = toolchain_info.objcopy_executable,
        arguments = args + [
            ctx.file.stub.path,
            out.path,
        ],
    )
    return [DefaultInfo(files = depset([out]), runfiles = ctx.runfiles(files = [out]))]

efi_unified_kernel_image = rule(
    implementation = _efi_unified_kernel_image_impl,
    attrs = {
        "kernel": attr.label(
            doc = "The Linux kernel executable bzImage. Needs to have EFI handover and EFI stub enabled.",
            mandatory = True,
            allow_single_file = True,
        ),
        "cmdline": attr.string(
            doc = "The kernel commandline to be embedded.",
        ),
        "initramfs": attr.label(
            doc = "The initramfs to be embedded.",
            allow_single_file = True,
        ),
        "os_release": attr.label(
            doc = """
                The os-release file identifying the operating system.
                See https://www.freedesktop.org/software/systemd/man/os-release.html for format.
            """,
            allow_single_file = True,
        ),
        "splash": attr.label(
            doc = "An image in BMP format which will be displayed as a splash screen until the kernel takes over.",
            allow_single_file = True,
        ),
        "stub": attr.label(
            doc = "The stub executable itself as a PE/COFF executable.",
            default = "@efistub//:efistub",
            allow_single_file = True,
            executable = True,
            cfg = build_efi_transition,
        ),
        "_toolchain": attr.label(
            doc = "The toolchain used for objcopy.",
            default = "//build/toolchain/llvm-efi:efi_cc_suite",
            providers = [platform_common.ToolchainInfo],
        ),
        # Allow for transitions to be attached to this rule.
        "_whitelist_function_transition": attr.label(
            default = "@bazel_tools//tools/whitelists/function_transition_whitelist",
        ),
    },
)
