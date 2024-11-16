package main

import "fmt"

func main() {
	var nome string
	fmt.Printf("Digite seu nome: ")
	fmt.Scanf("%s", &nome)
	fmt.Printf("Seu nome tem %d caracteres\n", len(nome))

	var saudacao = fmt.Sprintf("Ola, %s", nome)
	fmt.Println(saudacao)
	fmt.Printf("Essa saudação tem %d caracteres\n", len(saudacao))
}
