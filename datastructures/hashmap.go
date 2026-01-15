package main

import (
	"fmt"
	"hash/fnv"
)

type HashMap struct {
	buckets [][]kv
	size    int
}

type kv struct {
	key   string
	value int
}

func NewHashMap(size int) *HashMap {
	return &HashMap{buckets: make([][]kv, size), size: size}
}

func (h *HashMap) hash(key string) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32()) % h.size
}

func (h *HashMap) Put(key string, value int) {
	idx := h.hash(key)
	for i, pair := range h.buckets[idx] {
		if pair.key == key {
			h.buckets[idx][i].value = value
			return
		}
	}
	h.buckets[idx] = append(h.buckets[idx], kv{key, value})
}

func (h *HashMap) Get(key string) (int, bool) {
	idx := h.hash(key)
	for _, pair := range h.buckets[idx] {
		if pair.key == key {
			return pair.value, true
		}
	}
	return 0, false
}

func main() {
	hm := NewHashMap(16)
	hm.Put("hello", 42)
	if val, ok := hm.Get("hello"); ok {
		fmt.Println("Found:", val)
	}
}
