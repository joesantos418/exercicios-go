package main

import "fmt"

func main() {

	var a [5]int

	a[0] = 1
	a[1] = 3
	a[2] = 7
	a[3] = 4
	a[4] = 93

	for valor, _ := range a {
		fmt.Println(valor)
	}
}
