package main

import (
	"fmt"

	"workshop/auth"
)

func finish() {
	fmt.Println("fim da autenticacao")
	if err := recover(); err != nil {
		fmt.Println("RECOVERED!", err)
	}
}

func main() {
	defer finish()
	u := auth.User{Name: "Lorena"}
	fmt.Println(auth.Authenticate(u))
}
