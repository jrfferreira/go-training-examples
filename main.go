package main

import "fmt"

func main() {
	num := make([]int, 2, 10)
	fmt.Println(num, len(num), cap(num))
	num = append(num, 1, 2, 3, 4)
	fmt.Println(num, len(num), cap(num))
	num = append(num, 1, 2, 3, 4, 5)
	fmt.Println(num, len(num), cap(num))
}
