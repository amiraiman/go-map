package main

import (
	"iter"
	"slices"
)

type orderedMap[K comparable, V any] struct {
	keys []K
	m    map[K]V
}

func (o *orderedMap[K, V]) set(key K, value V) {
	// 1. Check if key already exists
	// 2. If no, append key cause it's a new key
	exists := slices.Contains(o.keys, key)
	if !exists {
		o.keys = append(o.keys, key)
	}

	o.m[key] = value
}

func (o *orderedMap[K, V]) remove(key K) {
	// Delete from map if key exists
	_, exists := o.m[key]
	if exists {
		delete(o.m, key)
	}

	// Delete from slice if key exists
	idx := slices.Index(o.keys, key)
	if idx != -1 {
		o.keys = append(o.keys[:idx], o.keys[idx+1:]...)
	}
}

func (o *orderedMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range o.keys {
			if !yield(k, o.m[k]) {
				return
			}
		}
	}
}
