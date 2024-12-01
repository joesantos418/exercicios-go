package main

import "fmt"

const (
	HorasEmUmDia       = 24
	MinutosEmUmaHora   = 60
	SegundosEmUmMinuto = 60
)

func main() {
	var dias int
	fmt.Printf("Digite uma quantidade de dias: ")
	fmt.Scanf("%d", &dias)

	segundos := dias * HorasEmUmDia * MinutosEmUmaHora * SegundosEmUmMinuto

	fmt.Printf("%d dias tem %d segundos\n", dias, segundos)
}
