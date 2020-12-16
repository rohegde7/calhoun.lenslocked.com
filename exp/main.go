package main

import (
	"os"
	"text/template"
)

type User struct {
	Name string
	Age  int
	Dogs []Dog
}

type Dog struct {
	Name string
	Age  float32
}

func main() {

	templatePointer, err := template.ParseFiles("exp/hello.gohtml")

	if err != nil {
		panic(err)
	}

	dog1 := Dog{
		Name: "Ruby",
		Age:  1.22222227,
	}
	dog2 := Dog{
		Name: "YO Yo",
		Age:  1.0909,
	}

	user := User{
		Name: "Manna Bhaji",
		Age:  24,
		Dogs: []Dog{dog1, dog2},
	}

	var templateExecuteError = templatePointer.Execute(os.Stdout, user)

	if templateExecuteError != nil {
		panic(templateExecuteError)
	}
}
