package main

import "fmt"

const (
	Quinta = iota
	Sexta
	Sabado
	Domingo
	Segunda
)

func main() {
	fmt.Println(Quinta)
	fmt.Println(Sexta)
	fmt.Println(Sabado)
	fmt.Println(Domingo)
	fmt.Println(Segunda)
}
