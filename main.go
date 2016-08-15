package main

import "fmt"

func main() {
	var name []byte
	name = []byte("joao")

	fmt.Println(name, string(name))
}
