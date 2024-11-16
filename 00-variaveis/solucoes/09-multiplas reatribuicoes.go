package main

import (
	"fmt"
	"strings"
)

func divideNome(nomeCompleto string) (string, string) {
	nomes := strings.Fields(nomeCompleto)

	return nomes[0], nomes[1]
}

func main() {
	nome := "Chanandler Bong"
	primeiro, segundo := divideNome(nome)
	fmt.Println(primeiro)
	fmt.Println(segundo)

	nome = "Joey Tribbiani"
	primeiro, segundo = divideNome(nome)
	fmt.Println(primeiro)
	fmt.Println(segundo)

	nome = "Monica Geller"
	primeiro, segundo = divideNome(nome)
	fmt.Println(primeiro)
	fmt.Println(segundo)

	nome = "Phoebe Buffay"
	primeiro, segundo = divideNome(nome)
	fmt.Println(primeiro)
	fmt.Println(segundo)

	nome = "Rachel Green"
	primeiro, segundo = divideNome(nome)
	fmt.Println(primeiro)
	fmt.Println(segundo)

	nome = "Ross Geller"
	primeiro, segundo = divideNome(nome)
	fmt.Println(primeiro)
	fmt.Println(segundo)
}
