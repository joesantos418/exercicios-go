package main

import "fmt"

func main() {
	var nome string
	fmt.Printf("Digite seu nome: ")
	fmt.Scanf("%s", &nome)
	fmt.Printf("Olá, %s!\n", nome)
}
