FSSpecInfo = provider(
    "Provides parts of an FSSpec used to assemble filesystem images",
    fields = {
        "spec": "File containing the partial FSSpec as prototext",
        "referenced": "Files (potentially) referenced by the spec",
    },
)

def fsspec_core_impl(ctx, tool, output_file, extra_files = [], extra_fsspecs = []):
    """
    fsspec_core_impl implements the core of an fsspec-based rule. It takes
    input from the `files`,`files_cc`, `symlinks` and `fsspecs` attributes
    and calls `tool` with the `-out` parameter pointing to `output_file`
    and paths to all fsspecs as positional arguments.
    """
    fs_spec_name = ctx.label.name + ".prototxt"
    fs_spec = ctx.actions.declare_file(fs_spec_name)

    fs_files = []
    inputs = []
    for p, label in ctx.attr.files.items() + ctx.attr.files_cc.items() + extra_files:
        if not p.startswith("/"):
            fail("file {} invalid: must begin with /".format(p))

        # Figure out if this is an executable.
        is_executable = True

        di = label[DefaultInfo]
        if di.files_to_run.executable == None:
            # Generated non-executable files will have DefaultInfo.files_to_run.executable == None
            is_executable = False
        elif di.files_to_run.executable.is_source:
            # Source files will have executable.is_source == True
            is_executable = False

        # Ensure only single output is declared.
        # If you hit this error, figure out a better logic to find what file you need, maybe looking at providers other
        # than DefaultInfo.
        files = di.files.to_list()
        if len(files) > 1:
            fail("file {} has more than one output: {}", p, files)
        src = files[0]
        inputs.append(src)

        mode = 0o555 if is_executable else 0o444
        fs_files.append(struct(path = p, source_path = src.path, mode = mode, uid = 0, gid = 0))

    fs_symlinks = []
    for target, p in ctx.attr.symlinks.items():
        fs_symlinks.append(struct(path = p, target_path = target))

    fs_spec_content = struct(file = fs_files, directory = [], symbolic_link = fs_symlinks)
    ctx.actions.write(fs_spec, proto.encode_text(fs_spec_content))

    extra_specs = []

    for fsspec in ctx.attr.fsspecs + extra_fsspecs:
        if FSSpecInfo in fsspec:
            fsspec_info = fsspec[FSSpecInfo]
            extra_specs.append(fsspec_info.spec)
            for f in fsspec_info.referenced:
                inputs.append(f)
        else:
            # Raw .fsspec prototext. No referenced data allowed.
            di = fsspec[DefaultInfo]
            extra_specs += di.files.to_list()

    ctx.actions.run(
        mnemonic = "GenFSSpecImage",
        progress_message = "Generating a fsspec based image: {}".format(output_file.short_path),
        outputs = [output_file],
        inputs = [fs_spec] + inputs + extra_specs,
        tools = [tool],
        executable = tool,
        arguments = ["-out", output_file.path, fs_spec.path] + [s.path for s in extra_specs],
    )
    return
