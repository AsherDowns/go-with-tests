package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"
	korean  = "Korean"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	koreanHelloPrefix  = "Annyeonghaseyo, "
)

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case korean:
		prefix = koreanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(hello("Elodie", "Korean"))
}
