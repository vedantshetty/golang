package main

import (
	"fmt"
)

var j int = 32

func main() {
	fmt.Println(j)
	var (
		j         int    = 5
		companion string = "Angelina Jolie"
	)
	actorName := "Brad pitt"

	fmt.Printf("%v+ %v = %v\n", actorName, companion, j)

	var (
		i  int = 42
		j2 float32
	)
	j2 = float32(i)
	fmt.Printf("%v , %v", i, j2)
}
