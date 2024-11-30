package main

import "fmt"

func main() {
	var pointer = new(string)
	*pointer = "hello"

	fmt.Printf("O tipo é: %T, o valor referenciado, %s\n", pointer, *pointer)
}
