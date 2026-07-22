package main

import (
	"fmt"
)

func getPoint() (x int, y int) {
	return 3, 4
}

func main() {
	// ignore y value
	x, _ := getPoint()
	fmt.Println(x)
}
