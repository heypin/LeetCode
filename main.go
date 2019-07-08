package main

import (
	"fmt"
)

func test() *int {
	a := 5
	return &a
}
func main() {
	a := test()
	b := test()
	fmt.Println(*a, *b)
	*a = 10
	fmt.Println(*a, *b)
}
