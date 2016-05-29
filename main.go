// Package choose provides functions for making choices.
package choose

import (
	"math"
	"sort"
)

// UniformFloat64 returns the indices of n elements of a sequence such that the
// first and last elements are included, and the rest are chosen to have a more
// or less uniform distance between two successive elements.
func UniformFloat64(sequence []float64, n uint) []uint {
	m := uint(len(sequence))

	if m <= n {
		indices := make([]uint, m)
		for i := uint(0); i < m; i++ {
			indices[i] = i
		}
		return indices
	} else if n == 0 {
		return []uint{}
	} else if n == 1 {
		return []uint{0}
	}

	indices := make([]uint, n)
	indices[0], indices[n-1] = 0, m-1

	pool := make(map[uint]bool)
	for i := uint(1); i+1 < m; i++ {
		pool[i] = true
	}

	min, max := float64(sequence[0]), float64(sequence[m-1])
	Δ := (max - min) / (float64(n) - 1)

	for i := uint(1); i+1 < n; i++ {
		k, x := ^uint(0), min+float64(i)*Δ
		for j := range pool {
			if k == ^uint(0) {
				k = j
				continue
			}
			Δ1 := math.Abs(x - sequence[j])
			Δ2 := math.Abs(x - sequence[k])
			if Δ1 < Δ2 || Δ1 == Δ2 && j < k {
				k = j
			}
		}
		indices[i] = k
		delete(pool, k)
	}

	sort.Sort(uints(indices))

	return indices
}

// UniformUint is an analog of UniformFloat64 for uint.
func UniformUint(sequence []uint, n uint) []uint {
	clone := make([]float64, len(sequence))
	for i := range sequence {
		clone[i] = float64(sequence[i])
	}
	return UniformFloat64(clone, n)
}
