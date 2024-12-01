package main

import (
	"fmt"
	"math"
)

func main() {
	var vi, a, dx float64
	fmt.Printf("Digite a velocidade inicial (vi): ")
	fmt.Scanf("%f", &vi)
	fmt.Printf("Digite a aceleração (a): ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Digite a variação de espaço (dx): ")
	fmt.Scanf("%f", &dx)

	vf := math.Sqrt((vi * vi) + (2 * a * dx))

	fmt.Printf("A velocidade final é: %f\n", vf)
}
