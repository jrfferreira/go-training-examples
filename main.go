package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	cities = map[string]int{}
	m      = sync.Mutex{}
)

func writeCities() {
	defer m.Unlock()
	m.Lock()
	cities["curitiba"] = 41
	cities["rio"] = 21
	cities["imperatriz"] = 99
}

func main() {
	go writeCities()

	defer m.Unlock()
	m.Lock()
	cities["sp"] = 11
	time.Sleep(10 * time.Second)

	fmt.Println(cities)
}
