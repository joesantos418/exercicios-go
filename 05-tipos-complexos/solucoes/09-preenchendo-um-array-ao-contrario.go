package main

import "fmt"

func main() {
	var a [15]int

	for i := 0; i < 15; i++ {
		a[i] = 14 - i
	}

	for _, valor := range a {
		fmt.Println(valor)
	}
}
