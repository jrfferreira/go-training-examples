package main

import (
	"fmt"
	"time"
)

var cities = map[string]int{}

func writeCities() {
	cities["curitiba"] = 41
	cities["rio"] = 21
	cities["imperatriz"] = 99
}

func main() {
	go writeCities()
	cities["sp"] = 11
	time.Sleep(10 * time.Second)

	fmt.Println(cities)
}
