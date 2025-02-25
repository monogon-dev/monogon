# Flattens the previously generated entries together
def build_nogo_config(v):
    out = {}
    for exp in v:
        for check, cfg in exp.items():
            if check not in out:
                out[check] = {}

            for k, v in cfg.items():
                if k not in out[check]:
                    out[check][k] = {}

                out[check][k] = out[check][k] | v

    return out

def exclude_from_checks(path, *checks):
    return {
        check: {
            "exclude_files": {
                "external/.+%s/" % path: "",
            },
        }
        for check in checks
    }

def exclude_from_external(checks):
    return {
        check: {
            "exclude_files": {
                # Don't run linters on external dependencies
                "external/": "third_party",
                "bazel-out/": "generated_output",
                "cgo/": "cgo",
            },
        }
        for check in checks
    }
