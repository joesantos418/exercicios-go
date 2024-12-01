package main

import "fmt"

func main() {
	var op string
	var num int

	fmt.Printf(`Opções:

a: Soma o número dado com 3
s: Subtrai 5 do número dado
m: Multiplica o número dado por 7
d: Divide o número dado por 2 (divisão inteira)

Digite sua opção: `)

	fmt.Scanf("%s", &op)

	switch op {
	case "a":
		fmt.Printf("Digite seu número: ")
		fmt.Scanf("%d", &num)
		fmt.Println(num + 3)
	case "s":
		fmt.Printf("Digite seu número: ")
		fmt.Scanf("%d", &num)
		fmt.Println(num - 5)
	case "m":
		fmt.Printf("Digite seu número: ")
		fmt.Scanf("%d", &num)
		fmt.Println(num * 7)
	case "d":
		fmt.Printf("Digite seu número: ")
		fmt.Scanf("%d", &num)
		fmt.Println(num / 2)
	default:
		fmt.Println("Opção inválida")
	}
}
