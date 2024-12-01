package main

import "fmt"

func main() {
	idadesENomes := map[int]string{
		26: "Ross",
		24: "Monica",
		25: "Rachel",
		27: "Chandler",
		28: "Joey",
		30: "Phoebe",
	}

	for idade, nome := range idadesENomes {
		fmt.Printf("%s tem %d anos.\n", nome, idade)
	}
}
