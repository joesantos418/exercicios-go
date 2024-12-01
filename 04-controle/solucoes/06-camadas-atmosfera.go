package main

import "fmt"

func main() {
	var altura int
	fmt.Printf("Digite uma altura: ")
	fmt.Scanf("%d", &altura)

	switch {
	case altura >= 0 && altura <= 11:
		fmt.Println("Troposfera")
	case altura >= 12 && altura <= 50:
		fmt.Println("Estratosfera")
	case altura >= 51 && altura <= 80:
		fmt.Println("Mesosfera")
	case altura >= 81 && altura <= 710:
		fmt.Println("Termosfera")
	case altura >= 711 && altura <= 9978:
		fmt.Println("Exosfera")
	}
}
