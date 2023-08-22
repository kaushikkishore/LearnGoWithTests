package main 

import "fmt"

const (
	spanish = "Spanish"
 	french = "French"
 	hindi = "Hindi"

 	frenchHelloPrefix = "Bonjour, "
 	spanishHelloPrefix = "Hola, "
 	hindiHelloPrefix = "Namaste, "
 	englishHelloPrefix = "Hello, " 
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("Kaushik", "English"))
}