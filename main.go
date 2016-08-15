package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	nomeIdade := map[string]user{
		"joao":   user{"joao", 27},
		"lorena": user{"lorena", 27},
		"babi":   user{"babi", 1},
	}

	for nome, u := range nomeIdade {
		fmt.Println(nome, u)
	}

	chico, exists := nomeIdade["chico"]

	fmt.Printf("%#v, exists? %v", chico, exists)
}
