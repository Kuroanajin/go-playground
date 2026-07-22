package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	f := func(operation func(int, int) int, number int) int {
		return operation(number, 10)
	}

	result := f(sum, 5)

	fmt.Println("the result is: ", result)
}
