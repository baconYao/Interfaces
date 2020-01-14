package main

import (
	"fmt"
)

// 宣告一個 bot interface
type bot interface {
	// 如果你是一個擁有名 getGreeting function 且 return type 為 string的type，
	// 那麼就有資格使用bot interface。
	// 所以由於 englishBot&spanishBot 都有getGreeting method，因此可以使用bot interface
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	fmt.Println("==========================================")

	getGooglePage()
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (en englishBot) getGreeting() string {
	return "Hi There!"
}

func (sp spanishBot) getGreeting() string {
	return "Hola!"
}