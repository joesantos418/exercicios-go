package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type pessoa struct {
		Nome  string
		Idade int
	}

	p := pessoa{
		Nome:  "John Doe",
		Idade: 25,
	}

	pessoaBytes, _ := json.Marshal(p)

	fmt.Println(string(pessoaBytes))
}
