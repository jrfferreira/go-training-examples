package main

import "fmt"

func inc(x *int) {
	*x++
	fmt.Println(x, &x) // &x = endereco de memoria
}

func main() {
	n := 10
	fmt.Println("antes", n, &n)
	inc(&n)
	fmt.Println("depois", n, &n)
}
