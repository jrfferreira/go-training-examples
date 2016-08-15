package main

import "fmt"

func me() (string, int) {
	return "joao", 27
}

func main() {
	nome, idade := me()
	fmt.Println(nome, idade)
}
