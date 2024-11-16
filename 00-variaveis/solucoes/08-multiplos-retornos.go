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
	nome := "Pedro Joaquim"

	pedro, joaquim := divideNome(nome)

	fmt.Println(pedro)
	fmt.Println(joaquim)
}
