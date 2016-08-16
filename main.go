package main

func soma(x, y int) int {
	return x + y
}

func main() {
	go soma(10, 1)
}
