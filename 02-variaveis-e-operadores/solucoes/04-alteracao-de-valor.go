package main

import "fmt"

func main() {
	a := 10
	minha_funcao(&a)

	fmt.Println(a)
}

func minha_funcao(a *int) {
	*a = 5
}
