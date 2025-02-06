// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// localconfig modifies the project's IntelliJ config to include project-specific
// settings. This is usually done by checking in the .idea directory, but we do not
// want to do this: it conflicts with the Bazel plugin's way of conducting its
// workspace business, lacks backwards compatibility, and is a common source of
// spurious Git diffs, especially when the IDE/JDK/random plugins are updated and
// team members run different versions.
//
// Instead, we use the officially supported way of shipping IntelliJ Bazel project
// configs - a .bazelproject file that can be imported using the Bazel project
// import wizard, with local configs. We then use this tool to mangle the local
// configs to add additional custom configuration beyond run configurations. This
// avoids merge conflicts and allows us to intelligently handle schema changes.
package main

import (
	"log"
	"os"
	"path"

	"source.monogon.dev/intellij/localconfig/watchers"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: localconfig <project dir>")
	}

	projectDir := os.Args[1]
	if _, err := os.Stat(path.Join(projectDir, ".ijwb")); err != nil {
		log.Fatalf("invalid project dir: %v", err)
	}

	if err := watchers.RewriteConfig(projectDir); err != nil {
		log.Fatal(err)
	}
}
