BuildKindProvider = provider(fields = ['type'])

def _impl(ctx):
    values = ['full', 'lite'] 
    value = ctx.build_setting_value
    if value not in values:
        fail(str(ctx.label) + " build setting allowed to take values {full, lite} but was set to " + value)

    return BuildKindProvider(type = value)

buildkind = rule(
    implementation = _impl,
    build_setting = config.string(flag = True),
    doc = """
        Build kind for userspace tools, either full (will have a direct
        dependency on data files) or lite (will not have a direct dependency on
        data files and will not attempt to load them).
    """,
)
