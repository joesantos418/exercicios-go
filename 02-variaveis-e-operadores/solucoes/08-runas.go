package main

import "fmt"

func main() {
	nome := "João da Silva Málaga"

	runas := []rune(nome)
	for i := 0; i < len(runas); i++ {
		fmt.Printf("%s ", string(runas[i]))
	}

	fmt.Println("")
}
