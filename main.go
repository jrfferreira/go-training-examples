package main

import (
	"fmt"
	"workshop/auth"
)

func authenticate(u auth.User) bool {
	return u.Name == "joao"
}

func main() {
	u := auth.User{"joao"}
	fmt.Println(authenticate(u))
}
