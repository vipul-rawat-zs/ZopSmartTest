package template

import (
	"os"
	"text/template"
)

type User struct {
	Id    int
	Name  string
	Phone string
}

func CreateUser(id int, name, phone string) *User {
	user := User{Id: id, Name: name, Phone: phone}
	return &user
}

func Example(user *User) {
	greet := `Hello {{.Name}}, Please confirm your Id {{.Id}} and Phone number {{.Phone}}`
	t := template.Must(template.New("Greet").Parse(greet))
	err := t.Execute(os.Stdout, user)
	if err != nil {
	}
}
