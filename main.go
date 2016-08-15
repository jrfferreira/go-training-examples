package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	outrosNumeros := []int{10, 11, 12}

	for _, numero := range outrosNumeros {
		numeros = append(numeros, numero)
	}

	fmt.Println(numeros)
}
