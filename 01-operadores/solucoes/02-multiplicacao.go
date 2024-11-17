package main

import "fmt"

func main() {
	var a, b int

	fmt.Printf("Digite um numero: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Digite outro numero: ")
	fmt.Scanf("%d", &b)
	fmt.Printf("O primeiro multiplicado pelo segundo resulta em: %d\n", a*b)
}
