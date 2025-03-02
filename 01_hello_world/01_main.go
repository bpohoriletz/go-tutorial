package helloworld

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func Hello(name, lang string) string {
	if "" == name {
		name = "World"
	}

	return greetingPrefix(lang) + name + "!"
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
