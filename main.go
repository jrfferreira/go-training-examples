package main

import "fmt"

func inc(x *int) {
	*x++
	fmt.Println("dentro", *x, x) // &x = endereco de memoria
}

/*
 * O ponteiro (*) traduz um endereço de memoria (&)
 * para o valor definido
 */

func main() {
	n := 10
	fmt.Println("antes", n, &n)
	inc(&n)
	fmt.Println("depois", n, &n)
}
