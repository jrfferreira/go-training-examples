package main

import "fmt"

func me() (string, int) {
	return "joao", 27
}

func main() {
	_, idade := me()
	fmt.Println(idade)
}
