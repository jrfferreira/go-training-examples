package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout")
	}
}
