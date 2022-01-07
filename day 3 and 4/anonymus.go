package main

import "fmt"

func Xyx() func(string) string {
	i := "hello"
	return func(a string) string {
		i = i + " " + a
		return i
	}
}

func main() {
	f := Xyx()
	g := Xyx()
	fmt.Println(f("You"))
	fmt.Println(g("are"))
	fmt.Println(f("my"))

}
