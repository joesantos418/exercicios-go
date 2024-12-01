package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}
}
