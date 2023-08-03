package main

import "fmt"

const engleshPrefix = "Hello, "
const spanishPrefix = "Hola, "
const deafultName = "world"
const spanish = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = deafultName
	}

	if language == spanish {
		return spanishPrefix + name
	}

	return engleshPrefix + name
}

func main() {
	fmt.Println(Hello("Nathan", "Spanish"))
}
