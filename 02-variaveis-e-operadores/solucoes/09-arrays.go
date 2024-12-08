package main

import "fmt"

func main() {
	a := [5]int{1, 3, 7, 4, 93}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}
