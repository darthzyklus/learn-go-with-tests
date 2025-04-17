package main

import "fmt"

const Spanish = "Spanish"
const French = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == Spanish {
		return spanishHelloPrefix + name
	} else if language == French {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Andres", ""))
}
