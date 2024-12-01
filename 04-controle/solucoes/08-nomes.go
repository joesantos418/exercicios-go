package main

import "fmt"

func main() {
	nomes := []string{"Monica", "Chandler", "Ross", "Rachel", "Joey", "Phoebe"}

	for indice, nome := range nomes {
		fmt.Printf("Nome: %s, indice: %d\n", nome, indice)
	}
}
