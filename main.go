package main

import "fmt"

func fibo(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		c <- x
	}
	close(c)
}

func main() {
	c := make(chan int)

	go fibo(10, c)

	for n := range c {
		fmt.Println(n)
	}
}
