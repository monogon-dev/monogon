// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package nvme

import (
	"math"
	"math/big"
)

// uint128 little endian composed of two uint64s, readable by binary.Read.
// Auxiliary type to simplify structures with uint128s (of which NVMe has
// quite a few).
type uint128le struct {
	Lo, Hi uint64
}

// BigInt returns u as a bigint
func (u uint128le) BigInt() *big.Int {
	v := new(big.Int).SetUint64(u.Hi)
	v = v.Lsh(v, 64)
	v = v.Or(v, new(big.Int).SetUint64(u.Lo))
	return v
}

// Uint64 returns u as a clamped uint64
func (u uint128le) Uint64() uint64 {
	if u.Hi > 0 {
		return math.MaxUint64
	}
	return u.Lo
}
