package main

import "fmt"

type Human struct{}
type Cat struct{}

func (Human) Speak() {
	fmt.Println("Oi!")
}

func (Cat) Speak() {
	fmt.Println("miau!")
}

type Speaker interface {
	Speak()
}

func speak(speaker Speaker) {
	speaker.Speak()
}

func main() {
	h := Human{}
	c := Cat{}
	speak(h)
	speak(c)
}
