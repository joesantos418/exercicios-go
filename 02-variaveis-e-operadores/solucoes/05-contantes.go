package main

import "fmt"

const (
	MinValue = 20
	MaxValue = 25
)

func main() {
	var num int
	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &num)

	fmt.Println(num > MinValue && num < MaxValue)
}
