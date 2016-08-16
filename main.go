package main

import "fmt"

type city struct {
	name string
	DDD  int
}

func (c *city) show() {
	fmt.Println("WIP ", c)
}

func main() {
	sp := city{name: "sp", DDD: 11}
	sp.show()
}
