package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"
	korean = "Korean"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	koreanHelloPrefix = "Annyeonghaseyo, "
)

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix

func main() {
	fmt.Println(hello("Elodie", "Korean"))
}
