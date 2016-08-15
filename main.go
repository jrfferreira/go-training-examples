package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	var u user
	u.name = "joao"
	u.age = 27
	fmt.Println(u, u.age, u.name)
}
