package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	// 2nd way of creating a map
	// colors := make(map[string]string)
	// colors["white"] = "ffffff"

	// delete(colors, "white")

	// 3rd way
	// var colorsMap map[string]string

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
