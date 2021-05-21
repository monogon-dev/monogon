// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package watchers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	fileWatchersPath = ".ijwb/.idea/watcherTasks.xml"
	templatePath     = "intellij/localconfig/data/watcherTasks.xml"
)

type config struct {
	XMLName   xml.Name  `xml:"project"`
	Version   string    `xml:"version,attr"`
	Component component `xml:"component"`
}

type component struct {
	Name        string       `xml:"name,attr"`
	TaskOptions []taskOption `xml:"TaskOptions"`
}

type taskOption struct {
	IsEnabled string `xml:"isEnabled,attr"`
	Option    []struct {
		Name  string `xml:"name,attr"`
		Value string `xml:"value,attr,omitempty"`
		Data  string `xml:",innerxml"`
	} `xml:"option"`
	Envs struct {
		Env []struct {
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"env"`
	} `xml:"envs"`
}

func buildConfig(options []taskOption) *config {
	return &config{
		XMLName: xml.Name{Local: "project"},
		Version: "4",
		Component: component{
			Name:        "ProjectTasksOptions",
			TaskOptions: options,
		},
	}
}

func readConfig(filename string) (cfg *config, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed reading file: %w", err)
	}

	err = xml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed deserializing XML: %w", err)
	}

	return
}

func (cfg *config) atomicWriteFile(filename string) error {
	b, err := xml.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize: %w", err)
	}

	// Atomic write is needed, IntelliJ has inotify watches on its config and reloads
	// (but not applies) instantly.
	tmpPath := filename + ".tmp"
	defer os.Remove(tmpPath)
	if err := ioutil.WriteFile(tmpPath, []byte(xml.Header+string(b)), 0664); err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}
	if err := os.Rename(tmpPath, filename); err != nil {
		return fmt.Errorf("failed to rename: %w", err)
	}

	return nil
}

// RewriteConfig adds our watchers to projectDir's watchers config, overwriting
// existing entries with the same name.
func RewriteConfig(projectDir string) error {
	template, err := readConfig(path.Join(projectDir, templatePath))
	if err != nil {
		return fmt.Errorf("failed reading template config: %w", err)
	}

	if template.Version != "4" {
		return fmt.Errorf("unknown template config version: %s", template.Version)
	}

	// Read existing tasks, if any.
	tasks := make(map[string]taskOption)
	cfg, err := readConfig(path.Join(projectDir, fileWatchersPath))

	switch {
	case err == nil:
		// existing config, read tasks
		if cfg.Version != "4" {
			return fmt.Errorf("unknown watchers config version: %s", cfg.Version)
		}
		for _, v := range cfg.Component.TaskOptions {
			for _, o := range v.Option {
				if o.Name == "name" {
					tasks[o.Value] = v
				}
			}
		}
	case os.IsNotExist(err):
		// no existing config - continue with empty tasks
	default:
		// error is non-nil and not an ENOENT
		return fmt.Errorf("failed reading existing config: %w", err)
	}

	// Overwrite "our" entries, identified by name.
	for _, v := range template.Component.TaskOptions {
		for _, o := range v.Option {
			if o.Name == "name" {
				tasks[o.Value] = v
			}
		}
	}

	// Build new configuration
	options := make([]taskOption, 0, len(tasks))
	for _, t := range tasks {
		options = append(options, t)
	}

	out := buildConfig(options)

	err = out.atomicWriteFile(path.Join(projectDir, fileWatchersPath))
	if err != nil {
		return fmt.Errorf("failed writing to output file: %w", err)
	}

	return nil
}
