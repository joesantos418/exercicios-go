package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Printf("Digite um numero: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Digite outro numero: ")
	fmt.Scanf("%d", &b)
	fmt.Printf("Digite outro numero: ")
	fmt.Scanf("%d", &c)
	fmt.Printf("Digite outro numero: ")
	fmt.Scanf("%d", &d)
	fmt.Println(a == b && c == d)
}
