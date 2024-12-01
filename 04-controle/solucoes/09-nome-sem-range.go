package main

import "fmt"

func main() {
	nomes := []string{"Monica", "Chandler", "Ross", "Rachel", "Joey", "Phoebe"}

	for i := 0; i < len(nomes); i++ {
		fmt.Printf("Nome: %s, indice: %d\n", nomes[i], i)
	}
}
