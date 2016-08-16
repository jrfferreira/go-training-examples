package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(1 * time.Second)
	var x chan int

	for {
		select {
		case m := <-x:
			fmt.Println(m)
		case <-c:
			fmt.Println("Please wait")
		}
	}
}
