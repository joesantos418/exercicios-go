package main

import "fmt"

func main() {
	var pointer = new(int)
	*pointer = 99

	fmt.Printf("O tipo é: %T, o valor, %v\n", pointer, pointer)
}
