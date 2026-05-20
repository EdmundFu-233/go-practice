package datastructures

import (
	"hash"
	"hash/fnv"
	"math"
)

// BloomFilter implements a space-efficient probabilistic data structure
// for testing set membership with possible false positives.
type BloomFilter struct {
	bits    []bool
	size    uint
	hashers []hash.Hash64
}

// NewBloomFilter creates a Bloom filter optimized for n elements with p false positive rate.
func NewBloomFilter(n uint, p float64) *BloomFilter {
	m := uint(math.Ceil(-float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
	k := uint(math.Ceil(float64(m) / float64(n) * math.Log(2)))

	bf := &BloomFilter{
		bits: make([]bool, m),
		size: m,
	}

	for i := uint(0); i < k; i++ {
		h := fnv.New64a()
		h.Write([]byte{byte(i)})
		bf.hashers = append(bf.hashers, h)
	}

	return bf
}

// Add inserts an element into the Bloom filter.
func (bf *BloomFilter) Add(data []byte) {
	for _, h := range bf.hashers {
		h.Reset()
		h.Write(data)
		idx := h.Sum64() % uint64(bf.size)
		bf.bits[idx] = true
	}
}

// Contains checks if an element may be in the set (false positives possible).
func (bf *BloomFilter) Contains(data []byte) bool {
	for _, h := range bf.hashers {
		h.Reset()
		h.Write(data)
		idx := h.Sum64() % uint64(bf.size)
		if !bf.bits[idx] {
			return false
		}
	}
	return true
}

