package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	u := user{name: "joao", age: 27}
	fmt.Printf("%#v", u)
}
