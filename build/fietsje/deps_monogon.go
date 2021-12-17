package fietsje

// deps_monogon.go contains the main entrypoint for Monogon-specific
// dependencies to be handled by fietsje, and calls out to all other functions
// in deps_*.go.

import (
	"bytes"
	"fmt"
	"os"
)

// Monogon runs fietsje for all Monogon transitive dependencies.
func Monogon(shelfPath, repositoriesBzlPath string) error {
	shelf, err := shelfLoad(shelfPath)
	if err != nil {
		return fmt.Errorf("could not load shelf: %w", err)
	}

	p := &planner{
		available: make(map[string]*dependency),
		enabled:   make(map[string]bool),
		seen:      make(map[string]string),

		shelf: shelf,
	}

	// Currently can't bump past v1.30.0, as that removes the old balancer.Picker API
	// that go-etcd depends upon. See https://github.com/etcd-io/etcd/pull/12398 .
	p.collect(
		"google.golang.org/grpc", "v1.29.1",
	).use(
		"golang.org/x/text",
	)

	depsKubernetes(p)
	depsContainerd(p)
	depsGVisor(p)

	// our own deps, common
	p.collectOverride("go.uber.org/zap", "v1.15.0")
	p.collectOverride("golang.org/x/mod", "v0.3.0", useImportAliasNaming)
	p.collectOverride("github.com/spf13/viper", "v1.9.0").use(
		"gopkg.in/ini.v1",
		"github.com/subosito/gotenv",
	)
	p.collectOverride("github.com/spf13/cobra", "v1.2.1")
	p.collect("github.com/cenkalti/backoff/v4", "v4.0.2")

	p.collect("github.com/google/go-tpm", "ae6dd98980d4")
	p.collect("github.com/google/go-tpm-tools", "f8c04ff88181")
	p.collect("github.com/google/certificate-transparency-go", "v1.1.0")
	p.collect("github.com/insomniacslk/dhcp", "67c425063dcad32c5d14ce9a520c8865240dc945").use(
		"github.com/mdlayher/ethernet",
		"github.com/mdlayher/raw",
		"github.com/u-root/u-root",
	)
	p.collect("github.com/rekby/gpt", "a930afbc6edcc89c83d39b79e52025698156178d")
	p.collect("github.com/yalue/native_endian", "51013b03be4fd97b0aabf29a6923e60359294186")

	// Used by //build/bazel_cc_fix, override to make sure we use the latest version
	p.collectOverride("github.com/mattn/go-shellwords", "v1.0.11")

	// Used by //metropolis/build/mkimage
	p.collect("github.com/diskfs/go-diskfs", "v1.2.0").use(
		"gopkg.in/djherbis/times.v1",
		"github.com/pkg/xattr",
		"github.com/pierrec/lz4",
		"github.com/ulikunitz/xz",
	)

	// Used by //metropolis/build/genosrelease
	p.collect("github.com/joho/godotenv", "v1.3.0")

	// used by //build/bindata
	p.collect("github.com/kevinburke/go-bindata", "v3.16.0")

	// for interactive debugging during development (//:dlv alias)
	depsDelve(p)

	// Used by //metropolis/test/nanoswitch
	p.collect("github.com/google/nftables", "7127d9d22474b437f0e8136ddb21855df29790bf").use(
		"github.com/koneu/natend",
	)

	// Used by //metropolis/node/core/network/dhcp4c
	p.collect("github.com/google/gopacket", "v1.1.17")

	// used by //core//kubernetes/clusternet
	p.collect("golang.zx2c4.com/wireguard/wgctrl", "ec7f26be9d9e47a32a2789f8c346031978485cbf").use(
		"github.com/mdlayher/netlink",
		"github.com/mdlayher/genetlink",
	)

	p.collect(
		"github.com/sbezverk/nfproxy", "7fac5f39824e7f34228b08ba8b7640770ca6a9f4",
		patches("nfproxy.patch"),
	).use(
		"github.com/sbezverk/nftableslib",
	)

	p.collect("github.com/coredns/coredns", "v1.7.0",
		prePatches("coredns-remove-unused-plugins.patch"),
	).use(
		"github.com/caddyserver/caddy",
		"github.com/dnstap/golang-dnstap",
		"github.com/farsightsec/golang-framestream",
		"github.com/flynn/go-shlex",
		"github.com/grpc-ecosystem/grpc-opentracing",
		"github.com/infobloxopen/go-trees",
		"github.com/miekg/dns",
		"github.com/opentracing/opentracing-go",
	)

	// goimports with import grouping patch
	// https://go-review.googlesource.com/c/tools/+/321409/
	p.collectOverride("golang.org/x/tools", "v0.1.2-0.20210518182153-17b346669257", patches("goimports-group-merging.patch"))

	// commentwrap is used as a nogo analyzer to stick to a maximum line
	// length for comments.
	p.collect("github.com/corverroos/commentwrap", "2926638be44ce0c6c0ee2471e9b5ad9473c984cd").use(
		"github.com/muesli/reflow",
	)

	// Used by metroctl to resolve XDG directories
	p.collect("github.com/adrg/xdg", "v0.4.0")

	// Used for generating proto docs in //build/proto_docs
	p.collect("github.com/pseudomuto/protoc-gen-doc", "v1.5.0", patches("protoc-gen-doc-no-gogo.patch")).use(
		"github.com/Masterminds/sprig",
		"github.com/Masterminds/semver",
		"github.com/aokoli/goutils",
		"github.com/huandu/xstrings",
	).with(
		disabledProtoBuild,
	).use(
		"github.com/pseudomuto/protokit",
	)

	// First generate the repositories starlark rule into memory. This is because
	// rendering will lock all unlocked dependencies, which might take a while. If a
	// use were to interrupt it now, they would end up with an incomplete
	// repositories.bzl and would have to restore from git.
	buf := bytes.NewBuffer(nil)
	err = p.render(buf)
	if err != nil {
		return fmt.Errorf("could not render deps: %w", err)
	}

	err = os.WriteFile(repositoriesBzlPath, buf.Bytes(), 0666)
	if err != nil {
		return fmt.Errorf("could not write deps: %w", err)
	}
	return nil
}
