package main

import "fmt"

func main() {
	a := [14]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 20, 30, 31}
	b := dobraImparesDividePares(a)

	for _, valor := range b {
		fmt.Println(valor)
	}
}

func dobraImparesDividePares(a [14]int) [14]int {
	var b [14]int
	for chave, valor := range a {
		if valor%2 == 0 {
			b[chave] = valor / 2
		} else {
			b[chave] = valor * 2
		}
	}

	return b
}
