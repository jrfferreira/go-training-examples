package main

import "fmt"

func soma(x, y int, c chan int) {
	c <- x + y

}

func main() {
	c := make(chan int)

	go soma(10, 1, c)

	go soma(10, 2, c)

	go soma(10, 3, c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
