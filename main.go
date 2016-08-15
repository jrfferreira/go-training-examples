package main

import "fmt"

func main() {
	nomeIdade := map[string]int{}
	nomeIdade["joao"] = 27
	nomeIdade["lorena"] = 27
	nomeIdade["babi"] = 1

	for nome, idade := range nomeIdade {
		fmt.Println(nome, idade)
	}
}
