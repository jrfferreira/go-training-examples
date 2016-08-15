package main

import "fmt"

func exec(cmd string, args ...string) {
	fmt.Println(cmd, args)
}

func main() {
	exec("ps", "-e", "-f")
}
