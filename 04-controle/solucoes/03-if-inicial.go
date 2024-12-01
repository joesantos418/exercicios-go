package main

import (
	"errors"
	"fmt"
)

func validaParEDobra(num int) (int, error) {
	if num%2 == 0 {
		return num * 2, nil
	}

	return 0, errors.New("Numero dado nao pode ser impar")
}

func main() {
	var num int
	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &num)

	if dobro, err := validaParEDobra(num); err == nil {
		fmt.Println(dobro)
	} else {
		fmt.Println(err)
	}
}
