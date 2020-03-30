package main

import "fmt"

type bot interface {
	getGreeting() string
}

type engishBot struct{}
type spanishBot struct{}

func main() {
	eb := engishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (engishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Holal"
}
