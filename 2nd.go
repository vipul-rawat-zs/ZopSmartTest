package main

import "fmt"

func x(s string) interface{} {
	switch s {
	case "BOB":
		return "bob"
	default:
		return "value not found"
	}
}

func main() {
	var s, ss string = "BOB", "BOB"
	if s == ss {
		fmt.Println("Hello ")
	}

}
