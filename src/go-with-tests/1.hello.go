package main

import "fmt"

func main() {
	fmt.Println(Hello("Vedant"))
}

func Hello(name string) string {
	return "Hello, " + name + "!"
}
