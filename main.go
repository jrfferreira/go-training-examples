package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4}
	fmt.Println(num, len(num), cap(num))
	num = append(num, 1, 2, 3, 4)
	fmt.Println(num, len(num), cap(num))
	num = append(num, 1)
	fmt.Println(num, len(num), cap(num))
}
