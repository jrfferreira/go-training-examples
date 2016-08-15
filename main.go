package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	u := user{"joao", 27}
	fmt.Printf("%#v", u)
}
