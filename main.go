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

func describe(m map[string]string) {
	fmt.Println("ALL COLORS:")
	for key, value := range m {
		fmt.Printf("%v: %v\n", key, value)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"black": "#000",
		"white": "#fff",
	}
	colors["blue"] = "#0000ff"

	describe(colors)
	fmt.Println("FOUND COLOR:", getColor(colors, "blue"))

	delete(colors, "blue")

	describe(colors)
	fmt.Println("FOUND COLOR:", getColor(colors, "blue"))
}
