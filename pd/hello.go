package main

import "fmt"

const hello_eng = "Hello, "
const hello_spanish = "Hola, "

func hello(name string, language string) string {
	if language == "Spanish" {
		return hello_spanish + name
	}

	if name == "" {
		name = "world"
	}
	return hello_eng + name
}

func main() {
	fmt.Println(hello("emmmma", "Spanish"))
}
