package main

import "fmt"

func main() {
	num := 34
	pointer := &num

	fmt.Printf("O tipo é: %T, o valor, %v\n", pointer, pointer)
}
