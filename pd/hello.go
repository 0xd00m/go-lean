package main

import "fmt"

const helloprefix = "Hello, "

func hello(name string) string {
	return helloprefix + name
}

func main() {
	fmt.Println(hello("disk"))
}
