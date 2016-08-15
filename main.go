package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	nomeIdade := map[string]user{}
	nomeIdade["joao"] = user{"joao", 27}
	nomeIdade["lorena"] = user{"lorena", 27}
	nomeIdade["babi"] = user{"babi", 1}

	for nome, idade := range nomeIdade {
		fmt.Println(nome, idade)
	}
}
