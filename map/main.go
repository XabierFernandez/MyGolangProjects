package main

import "fmt"

func main() {
	// var colors map[string]string
	/* colors := make(map[string]string)
	colors["white"] = "#ffffff"
	delete(colors, "white") */
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	printMap(colors)
}

func printMap(c map[string]string){
	for color,hex := range c{
		fmt.Println("Hex code for",color, "is", hex)
	}
}