package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	outrosNumeros := []int{10, 11, 12}

	numeros = append(numeros, outrosNumeros...)

	fmt.Println(numeros)
}
