package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &num)

	fmt.Printf("Seu número em decimal: %d\n", num)
	fmt.Printf("Seu número em octal: %O\n", num)
	fmt.Printf("Seu número em hexadecimal: %X\n", num)
	fmt.Printf("Seu número em binário: %b\n", num)
}
