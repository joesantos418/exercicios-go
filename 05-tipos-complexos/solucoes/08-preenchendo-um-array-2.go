package main

import "fmt"

func main() {
	var a [15]int

	for i := 0; i < 15; i++ {
		a[i] = i + 1
	}

	for _, valor := range a {
		fmt.Println(valor)
	}
}
