package main

import "fmt"

type orderedMapStringString struct {
	keySlice   []string
	orderedMap map[string]string
}

func (o *orderedMapStringString) set(key, value string) {
	// 1. Check if key already exists
	// 2. If yes, no need to append key cause it's an update
	exists := false
	for _, k := range o.keySlice {
		if k == key {
			exists = true
		}
	}

	if !exists {
		o.keySlice = append(o.keySlice, key)
	}

	o.orderedMap[key] = value
}

func (o *orderedMapStringString) remove(key string) {
	// Delete from map if key exists
	_, exists := o.orderedMap[key]
	if exists {
		delete(o.orderedMap, key)
	}

	// Get key index
	idx := -1
	for i, k := range o.keySlice {
		if k == key {
			idx = i
		}
	}

	// Delete from slice if key exists
	if idx != -1 {
		o.keySlice = append(o.keySlice[:idx], o.keySlice[idx+1:]...)
	}
}

func (o *orderedMapStringString) describe() {
	for _, k := range o.keySlice {
		val, ok := o.orderedMap[k]
		if ok {
			fmt.Printf("%v: %v\n", k, val)
		}
	}
}
