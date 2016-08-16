package main

import (
	"fmt"

	"workshop/auth"
)

func main() {
	u := auth.User{Name: "joao"}
	fmt.Println(auth.Authenticate(u))
}
