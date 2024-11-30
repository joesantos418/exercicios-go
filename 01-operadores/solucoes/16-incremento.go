package main

import "fmt"

func main() {
	var num int

	fmt.Printf("Digite um numero inteiro: ")
	fmt.Scanf("%d", &num)
	num++
	fmt.Printf("%d\n", num)
}
