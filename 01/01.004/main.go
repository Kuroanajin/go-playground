package main

import (
	"errors"
	"fmt"
)

func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

func main() {
	mul, div, err := calculator(9, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Scanln()
		return
	}

	fmt.Printf("multiplication: %d, Division: %d", mul, div)
	fmt.Scanln()
}
