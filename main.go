package main

import (
	"fmt"
)

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(en englishBot) {
	fmt.Println(en.getGreeting())
}

func printGreeting(sp spanishBot) {
	fmt.Println(sp.getGreeting())
}

func (en englishBot) getGreeting() string {
	return "Hi There!"
}

func (sp spanishBot) getGreeting() string {
	return "Hola!"
}