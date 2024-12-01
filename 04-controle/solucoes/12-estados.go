package main

import "fmt"

func main() {
	estados := []string{
		"São Paulo",
		"Sergipe",
		"Alagoas",
		"Mato Grosso do Sul",
		"Paraná",
		"Minas Gerais",
		"Pará",
		"Roraima",
	}

	for _, estado := range estados {
		if estado == "Mato Grosso" ||
			estado == "Goiás" ||
			estado == "Mato Grosso do Sul" {
			break
		}

		fmt.Println(estado)
	}
}
