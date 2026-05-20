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

	orderedColors := orderedMapStringString{
		keySlice:   []string{},
		orderedMap: map[string]string{},
	}
	orderedColors.set("red", "#ff0000")
	orderedColors.set("green", "#00ff00")
	orderedColors.set("blue", "#0000ff")
	orderedColors.set("black", "#000")
	orderedColors.set("white", "#fff")

	describe(colors)
	describe(colors)
	describe(colors)
	describe(colors)

	fmt.Println("\nALL COLORS:")
	orderedColors.describe()
	fmt.Println("\nALL COLORS:")
	orderedColors.describe()
	fmt.Println("\nALL COLORS:")
	orderedColors.describe()
	fmt.Println("\nALL COLORS:")
	orderedColors.describe()
}
