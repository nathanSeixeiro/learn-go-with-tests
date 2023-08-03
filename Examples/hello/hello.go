package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
	deafultName   = "world"
)

func Hello(name, language string) string {
	if name == "" {
		name = deafultName
	}
	prefix := gettingPrefixLang(language)
	return prefix + name
}

func gettingPrefixLang(lang string) (prefix string){
	switch lang {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Nathan", "Spanish"))
}
