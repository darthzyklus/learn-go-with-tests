package main

import "fmt"

const (
	SPANISH        = "Spanish"
	FRANCE         = "France"
	ENGLISH_PREFIX = "Hello, "
	SPANISH_PREFIX = "Hola, "
	FRANCE_PREFIX  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	if language == SPANISH {
		return SPANISH_PREFIX
	}

	if language == FRANCE {
		return FRANCE_PREFIX
	}

	return ENGLISH_PREFIX
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
