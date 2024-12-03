package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const exclamationMark = "!"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
		case "Spanish":
			return spanishHelloPrefix + name + exclamationMark
		case "French":
			return frenchHelloPrefix + name + exclamationMark
	}

	return englishHelloPrefix + name + exclamationMark
}

func main() {
	fmt.Println(Hello("Vish⚡", ""))
}