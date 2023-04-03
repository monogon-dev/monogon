package main

import (
	"fmt"

	"go.starlark.net/starlark"
)

// getBazelDNFFiles parses third_party/sandboxroot/repositories.bzl (at the given
// path) into a list of rpmDefs. It does so by loading the .bzl file into a
// minimal starlark interpreter that emulates enough of the Bazel internal API to
// get things going.
func getBazelDNFFiles(path string) ([]*rpmDef, error) {
	var res []*rpmDef

	// rpm will be called any time the Starlark code calls rpm() from
	// @bazeldnf//:deps.bzl.
	rpm := func(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
		var name, sha256 starlark.String
		var urls *starlark.List
		if err := starlark.UnpackArgs("rpm", args, kwargs, "name", &name, "sha256", &sha256, "urls", &urls); err != nil {
			return nil, err
		}
		it := urls.Iterate()
		defer it.Done()

		var urlsS []string
		var url starlark.Value
		for it.Next(&url) {
			if url.Type() != "string" {
				return nil, fmt.Errorf("urls must be a list of strings")
			}
			urlS := url.(starlark.String)
			urlsS = append(urlsS, urlS.GoString())
		}

		ext, err := newRPMDef(name.GoString(), sha256.GoString(), urlsS)
		if err != nil {
			return nil, fmt.Errorf("invalid rpm: %v", err)
		}
		res = append(res, ext)
		return starlark.None, nil
	}

	thread := &starlark.Thread{
		Name: "fakebazel",
		Load: func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
			switch module {
			case "@bazeldnf//:deps.bzl":
				return map[string]starlark.Value{
					"rpm": starlark.NewBuiltin("rpm", rpm),
				}, nil
			}
			return nil, fmt.Errorf("not implemented in fakebazel")
		},
	}
	globals, err := starlark.ExecFile(thread, path, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("executing failed: %w", err)
	}
	if !globals.Has("sandbox_dependencies") {
		return nil, fmt.Errorf("does not contain sandbox_dupendencies")
	}
	_, err = starlark.Call(thread, globals["sandbox_dependencies"], nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call sandbox_dependencies: %w", err)
	}
	return res, nil
}
