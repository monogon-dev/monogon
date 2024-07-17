//go:build tools
// +build tools

package kube_code_generator

import (
	_ "k8s.io/code-generator/cmd/client-gen"
	_ "k8s.io/code-generator/cmd/deepcopy-gen"
	_ "k8s.io/code-generator/cmd/informer-gen"
	_ "k8s.io/code-generator/cmd/lister-gen"
)
