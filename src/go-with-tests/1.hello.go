package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Vedant", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return prefix(language) + name + "!"
}

func prefix(language string) (prefix string) {

	switch language {
	case "Hindi":
		prefix = "नमस्ते, "
	case "French":
		prefix = "Bonjour, "
	case "German":
		prefix = "Hallo, "
	case "Spanish":
		prefix = "Hola, "
	case "Italian":
		prefix = "Ciao, "
	case "Russian":
		prefix = "Здравствуйте, "
	case "Japanese":
		prefix = "こんにちは, "
	case "Chinese":
		prefix = "你好, "
	default:
		prefix = "Hello, "

	}
	return
}
