package main

import "fmt"

// func Hello() string {
// 	return "Hello, world"
// }

// const englishHelloPrefix = "Hello, "

// func Hello(name string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	return englishHelloPrefix + name
// }
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "spanish"))
}
