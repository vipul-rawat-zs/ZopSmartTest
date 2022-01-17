package main

import "fmt"

func main() {
	var input interface{}

	type z struct {
		Id   int
		Name string
	}
	user := z{Id: 1, Name: "test"}
	input = user
	out := input.(z)
	fmt.Println(out.Id, out.Name)
}
