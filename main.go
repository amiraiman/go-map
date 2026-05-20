package main

import (
	"fmt"
)

func getColor(m map[string]string, k string) string {
	val, ok := m[k]
	if !ok {
		return "No colors found"
	}

	return val
}

func describe[K comparable, V any](m map[K]V) {
	fmt.Println("\nALL COLORS:")
	for key, value := range m {
		fmt.Printf("%v: %v\n", key, value)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"black": "#000",
		"white": "#fff",
	}

	describe(colors)
	describe(colors)
	describe(colors)

	oc := orderedMap[string, string]{
		keys: []string{},
		m:    map[string]string{},
	}
	oc.set("red", "#ff0000")
	oc.set("green", "#00ff00")
	oc.set("blue", "#0000ff")
	oc.set("black", "#000")
	oc.set("white", "#fff")

	describe(oc.m)
	describe(oc.m)
	describe(oc.m)
}
