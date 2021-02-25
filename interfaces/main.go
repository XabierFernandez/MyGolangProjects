package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

func main() {

	fmt.Println(" ")

}
func (eb englishBot) getGreeting() string {
	return "Hi there!"
}
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
