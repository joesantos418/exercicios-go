package main

import "fmt"

var nome string

func main() {
	eletrodomesticos := []string{
		"Televisão",
		"Máquina de lavar",
		"Microondas",
		"Air Fryer",
		"Geladeira",
		"Aspirador de pó",
	}

	i := 0

PRINT:
	if i >= len(eletrodomesticos) {
		return
	}
	fmt.Println(eletrodomesticos[i])
	i++
	goto PRINT
}
