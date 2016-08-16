package main

import "fmt"

type user struct {
	name string
}

func authenticate(u user) bool {
	return u.name == "joao"
}

func main() {
	u := user{"joao"}
	fmt.Println(authenticate(u))
}
