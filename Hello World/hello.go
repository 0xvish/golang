package main

import "fmt"

const englishHelloPrefix = "Hello, "
const exclamationMark = "!"

func Hello(name string) string {
	return englishHelloPrefix + name + exclamationMark
}

func main() {
	fmt.Println(Hello("Vish⚡"))
}