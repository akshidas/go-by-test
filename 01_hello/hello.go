package hello

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const french = "french"
const spanish = "spanish"

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix

	case french:
		return frenchHelloPrefix

	default:
		return englishHelloPrefix
	}
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
