package main

import "fmt"

const (
	NoveQuintos = float64(9) / float64(5)
	TrintaEDois = float64(32)
)

func main() {
	var c float64
	fmt.Printf("Digite uma temperatura em Celsius: ")
	fmt.Scanf("%f", &c)

	f := (c * NoveQuintos) + TrintaEDois

	fmt.Println(f)
}
