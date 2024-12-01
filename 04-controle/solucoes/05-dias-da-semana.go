package main

import "fmt"

func main() {
	var dia int
	fmt.Printf("Digite o código do dia: ")
	fmt.Scanf("%d", &dia)

	switch dia {
	case 21, 82, 43, 64, 55:
		fmt.Println("Dia de trabalho")
	case 30, 86:
		fmt.Println("Dia de descanso")
	}
}
