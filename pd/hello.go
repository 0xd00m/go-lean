package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const hello_eng = "Hello, "
const hello_spanish = "Hola, "
const hello_french = "Bonjour, "

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greeting(language) + name
}

func greeting(language string) (prefix string) {
	switch language {
	case french:
		prefix = hello_french
	case spanish:
		prefix = hello_spanish
	default:
		prefix = hello_eng
	}
	return
}

func main() {
	fmt.Println(hello("emmmma", "Spanish"))
}
