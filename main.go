package main

import (
	"fmt"
	"workshop/auth"
)

func main() {
	u := auth.User{Name: "Lorena"}
	_, err := auth.Authenticate(u)

	if err != nil {
		fmt.Println(err)
	}
}
