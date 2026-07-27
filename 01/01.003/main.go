package main

import "fmt"

func main() {
	type persona struct {
		nome  string
		eta   int
		email string
	}

	raffaele := persona{nome: "raffaele", eta: 15, email: "xxx@example.com"}

	fmt.Println(raffaele.nome)
}
