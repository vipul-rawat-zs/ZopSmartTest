package main

import (
	"fmt"
	"strings"
)

func getFruit(s string) string {
	s = strings.ToLower(s)
	switch s {
	case "red":
		return "apple"
	case "green":
		return "grapes"
	case "orange":
		return "orange"
	case "yellow":
		return "banana"
	default:
		return "unknown"
	}
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(getFruit(s))
}
