package main

import "fmt"

func main() {
	animais := []string{
		"Golfinho",
		"Macaco",
		"Leão",
		"Papagaio",
		"Girafa",
		"Baleia",
		"Polvo",
		"Cachorro",
	}

	for _, animal := range animais {
		if animal != "Golfinho" &&
			animal != "Baleia" &&
			animal != "Polvo" {
			continue
		}

		fmt.Println(animal)
	}
}
