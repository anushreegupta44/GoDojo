package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "
const hindiHelloPrefix = "Namaste "
const spanish = "Spanish"
const french = "French"
const hindi = "Hindi"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return getPrefix(language) + name
}

func getPrefix(language string) string {
	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("anushree", "English"))
}
