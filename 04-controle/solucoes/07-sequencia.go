package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Digite um numero: ")
	fmt.Scanf("%d", &num)

	if num < 0 || num > 20 {
		return
	}

	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}
