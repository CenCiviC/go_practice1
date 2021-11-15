package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func mult(a, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	const name string = "nico"
	var name1 string = "nico"
	name2 := "nico" //func 안에서는 축약형을 하면 compiler가 자동으로 type binding

	repeatMe("hi", "hello")
	fmt.Println(mult(2, 2))
	fmt.Println(lenAndUpper("hello"))
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(name2)
}
