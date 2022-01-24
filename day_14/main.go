package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type user struct {
	ID   int
	Name string
}

func (u *user) GreetName() string {
	if len(u.Name) > 0 {
		return u.Name
	} else {
		return "World"
	}
}

func greet(u user) {
	greetString := `Hello, {{ .GreetName }}`
	t, err := template.New("greet").Parse(greetString)
	if err != nil {
		log.Printf("error parsing, %v", err)
	}
	t.Execute(os.Stdout, u)
}

func main() {
	u1 := user{1, "surya"}
	u2 := user{2, ""}

	greet(u1)
	fmt.Println()
	greet(u2)
}
