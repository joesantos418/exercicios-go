package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &num)

	if num%15 == 0 {
		fmt.Println("baz")
	} else if num%5 == 0 {
		fmt.Println("bar")
	} else if num%3 == 0 {
		fmt.Println("foo")
	} else {
		fmt.Println(num)
	}
}
