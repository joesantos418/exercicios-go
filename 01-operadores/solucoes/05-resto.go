package main

import "fmt"

func main() {
	var a, b int

	fmt.Printf("Digite um numero: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Digite outro numero: ")
	fmt.Scanf("%d", &b)
	fmt.Printf("O resto da divisão do primeiro pelo segundo é: %d\n", a%b)
}
