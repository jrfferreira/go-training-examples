package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}

	for i, n := range num {
		fmt.Println(i, n)
	}

	for i := 0; i < len(num); i++ {
		fmt.Println(i, num[i])
	}
}
