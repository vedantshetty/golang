package examples

import "fmt"

func main() {
	fmt.Println(Hello(""))
}

const prefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	fmt.Println(name)
	return prefix + name + "!"
}
