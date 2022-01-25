"""Rules for generating EFI unified kernel images. These are EFI-bootable PE/COFF files containing a stub loader,
a kernel, and optional commandline and initramfs in one file.
See https://systemd.io/BOOT_LOADER_SPECIFICATION/#type-2-efi-unified-kernel-images for more information.
"""

load("//build/toolchain/llvm-efi:transition.bzl", "build_efi_transition")
load("//metropolis/node/build:def.bzl", "VerityConfig")

def _efi_unified_kernel_image_impl(ctx):
    # Find the dependency paths to be passed to mkpayload.
    deps = {
      "linux": ctx.file.kernel,
      "initrd": ctx.file.initramfs,
      "osrel": ctx.file.os_release,
      "splash": ctx.file.splash,
      "stub": ctx.file.stub,
    }
    # Since cmdline is a string attribute, put it into a file, then append
    # that file to deps.
    if ctx.attr.cmdline and ctx.attr.cmdline != "":
        cmdline = ctx.actions.declare_file("cmdline")
        ctx.actions.write(
            output = cmdline,
            content = ctx.attr.cmdline,
        )
        deps["cmdline"] = cmdline
    # Get the dm-verity target table from VerityConfig provider.
    if ctx.attr.verity:
      deps["rootfs_dm_table"] = ctx.attr.verity[VerityConfig].table
    # Format deps into command line arguments while keeping track of mkpayload
    # runtime inputs.
    args = []
    inputs = []
    for name, file in deps.items():
      if file:
        args.append("-{}={}".format(name, file.path))
        inputs.append(file)
    # Append the output parameter separately, as it doesn't belong with the
    # runtime inputs.
    image = ctx.actions.declare_file(ctx.attr.name + ".efi")
    args.append("-output={}".format(image.path))
    # Append the objcopy parameter separately, as it's not of File type, and
    # it does not constitute an input, since it's part of the toolchain.
    objcopy = ctx.attr._toolchain[platform_common.ToolchainInfo].objcopy_executable
    args.append("-objcopy={}".format(objcopy))
    # Run mkpayload.
    ctx.actions.run(
      mnemonic = "GenEFIKernelImage",
      progress_message = "Generating EFI unified kernel image",
      inputs = inputs,
      outputs = [image],
      executable = ctx.file._mkpayload,
      arguments = args
    )
    # Return the unified kernel image file.
    return [DefaultInfo(files = depset([image]), runfiles = ctx.runfiles(files = [image]))]

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
        "verity": attr.label(
            doc = "The DeviceMapper Verity rootfs target table.",
            allow_single_file = True,
            providers = [DefaultInfo, VerityConfig],
        ),
        "_mkpayload": attr.label(
            doc = "The mkpayload executable.",
            default = "//metropolis/node/build/mkpayload",
            allow_single_file = True,
            executable = True,
            cfg = "exec",
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
