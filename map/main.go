package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#b5d68",
		"white": "#fffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
