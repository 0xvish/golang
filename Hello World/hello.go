package main

import "fmt"

const ( // This is a way to group constants together. It's a good practice to group them together like this as it makes it easier to see what values are related to each other. It also makes it easier to add or remove values from the group.
	spanish = "Spanish"
	french = "French"
	hindi = "Hindi"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	hindiHelloPrefix = "Namaste, "
	exclamationMark = "!"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + exclamationMark
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
		case hindi:
			prefix = hindiHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
		return // This is called a naked return. We have named the return value (prefix) in the function signature (prefix string) so we can just return it like this. This is not very common in longer functions as it can harm readability.
}

func main() {
	fmt.Println(Hello("Vish⚡", ""))
}