package main

import "fmt"

func main() {
	a := [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 20, 30, 31}
	b := somandoAnterior(a)

	for _, valor := range b {
		fmt.Println(valor)
	}
}

func somandoAnterior(a [14]int) [14]int {
	var b [14]int
	anterior := 21

	for chave, valor := range a {
		b[chave] = valor + anterior
		anterior = valor
	}

	return b
}
