package main

import "fmt"

const (
	CincoNonos  = float64(5) / float64(9)
	TrintaEDois = float64(32)
)

func main() {
	var f float64
	fmt.Printf("Digite uma temperatura em Farenheit: ")
	fmt.Scanf("%f", &f)

	c := (f - TrintaEDois) * CincoNonos

	fmt.Println(c)
}
