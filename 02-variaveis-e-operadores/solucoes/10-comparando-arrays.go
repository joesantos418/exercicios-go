package main

import "fmt"

func main() {
	a := [5]int{1, 3, 7, 4, 93}
	b := [5]int{1, 6, 89, 3, 5}

	for i := 0; i < 5; i++ {
		if a[i] > b[i] {
			fmt.Println(a[i])
		}
	}
}
