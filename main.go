package main

import (
	"fmt"
	"strconv"
)

func main() {
	numeroString := "10"

	numero, _ := strconv.Atoi(numeroString)
	fmt.Println(numero + 10)
}
