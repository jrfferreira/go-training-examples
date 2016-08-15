package main

import "fmt"

type user struct {
	name string
	age int
}

func main() {
	nomeIdade := map[string]user{}
	nomeIdade["joao"] = {name: "joao", idade:27}
	nomeIdade["lorena"] = {name: "lorena", idade:27}
	nomeIdade["babi"] = {name: "babi", idade:1}

	for nome, idade := range nomeIdade {
		fmt.Println(nome, idade)
	}
}
