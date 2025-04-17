package main

import "fmt"

const (
	Spanish            = "Spanish"
	French             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func greetingPrefix(language string) string {
	if language == French {
		return frenchHelloPrefix
	}

	if language == Spanish {
		return spanishHelloPrefix
	}

	return englishHelloPrefix
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Andres", ""))
}
