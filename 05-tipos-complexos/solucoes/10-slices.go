package main

import "fmt"

func main() {
	a := []int{1, 3, 7, 4, 93}
	imprimeSlice(a)

	fmt.Println("---")

	b := []int{3, 21, 89}
	imprimeSlice(b)

	fmt.Println("---")

	c := []int{1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 10, 100, 1000}
	imprimeSlice(c)
}

func imprimeSlice(s []int) {
	for _, valor := range s {
		fmt.Println(valor)
	}
}
