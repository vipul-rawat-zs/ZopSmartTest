package main

import "fmt"

func test() {
	s := "this"
	defer fmt.Println(s)
	s = "and"
	fmt.Println(s)
	s = "that"
	defer fmt.Println(s)
}

func main() {
	defer fmt.Println("hello")
	test()
	defer fmt.Println("world")
}
