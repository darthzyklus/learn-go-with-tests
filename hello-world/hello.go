package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	if language == spanish {
		return spanishHelloPrefix
	}

	if language == french {
		return frenchHelloPrefix
	}

	return englishHelloPrefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
