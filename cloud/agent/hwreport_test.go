package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"source.monogon.dev/cloud/agent/api"
)

func TestParseCpuinfoAMD64(t *testing.T) {
	var cases = []struct {
		name     string
		data     string
		expected *api.CPU
	}{
		{
			"QEMUSingleCore",
			"cpuinfo_qemu_virtual.txt",
			&api.CPU{
				Vendor:          "GenuineIntel",
				Model:           "QEMU Virtual CPU version 2.1.0",
				Cores:           1,
				HardwareThreads: 1,
				Architecture: &api.CPU_X86_64_{X86_64: &api.CPU_X86_64{
					Family:   6,
					Model:    6,
					Stepping: 3,
				}},
			},
		},
		{
			"AMDEpyc7402P",
			"cpuinfo_amd_7402p.txt",
			&api.CPU{
				Vendor:          "AuthenticAMD",
				Model:           "AMD EPYC 7402P 24-Core Processor",
				Cores:           24,
				HardwareThreads: 48,
				Architecture: &api.CPU_X86_64_{X86_64: &api.CPU_X86_64{
					Family:   23,
					Model:    49,
					Stepping: 0,
				}},
			},
		},
		{
			"Intel12900K",
			"cpuinfo_intel_12900k.txt",
			&api.CPU{
				Vendor:          "GenuineIntel",
				Model:           "12th Gen Intel(R) Core(TM) i9-12900K",
				Cores:           16,
				HardwareThreads: 24,
				Architecture: &api.CPU_X86_64_{X86_64: &api.CPU_X86_64{
					Family:   6,
					Model:    151,
					Stepping: 2,
				}},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rawData, err := os.ReadFile("testdata/" + c.data)
			if err != nil {
				t.Fatalf("unable to read testdata file: %v", err)
			}
			res, errs := parseCpuinfoAMD64(rawData)
			if len(errs) > 0 {
				t.Fatal(errs[0])
			}
			assert.Equal(t, c.expected.Vendor, res.Vendor, "vendor mismatch")
			assert.Equal(t, c.expected.Model, res.Model, "model mismatch")
			assert.Equal(t, c.expected.Cores, res.Cores, "cores mismatch")
			assert.Equal(t, c.expected.HardwareThreads, res.HardwareThreads, "hardware threads mismatch")
			x86_64, ok := res.Architecture.(*api.CPU_X86_64_)
			if !ok {
				t.Fatal("CPU architecture not X86_64")
			}
			expectedX86_64 := c.expected.Architecture.(*api.CPU_X86_64_)
			assert.Equal(t, expectedX86_64.X86_64.Family, x86_64.X86_64.Family, "x86_64 family mismatch")
			assert.Equal(t, expectedX86_64.X86_64.Model, x86_64.X86_64.Model, "x86_64 model mismatch")
			assert.Equal(t, expectedX86_64.X86_64.Stepping, x86_64.X86_64.Stepping, "x86_64 stepping mismatch")
		})
	}
}
