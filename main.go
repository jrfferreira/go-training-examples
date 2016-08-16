package main

import "fmt"

type Human struct{}
type Cat struct{}

func (Human) Speak() {
	fmt.Println("Oi!")
}

func (Human) Walk() {
	fmt.Println("going")
}

func (Cat) Speak() {
	fmt.Println("miau!")
}

type Speaker interface {
	Speak()
}

type Walker interface {
	Walk()
}

type WalkerAndSpeaker interface {
	Walker
	Speaker
}

func speak(speaker Speaker) {
	speaker.Speak()
}

func walk(w Walker) {
	w.Walk()
}

func walkAndSpeak(w WalkerAndSpeaker) {
	w.Walk()
	w.Speak()
}

func main() {
	h := Human{}
	c := Cat{}
	walkAndSpeak(h)
	speak(c)
}
