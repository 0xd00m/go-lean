package main

import "fmt"

const helloprefix = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloprefix + name
}

func main() {
	fmt.Println(hello(""))
}
