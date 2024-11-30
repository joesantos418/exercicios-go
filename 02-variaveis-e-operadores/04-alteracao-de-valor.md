# Alteração de valor

Considere o código abaixo:

```
package main

import "fmt"

func main() {
	a := 10
	minha_funcao(a)

	fmt.Println(a)
}

func minha_funcao(a int) {
	a = 5
}

```

Esse código deveria substituir o valor de `a`, mas isso não acontece. Utilizando
ponteiros, conserte esse código para que o valor impresso de `a` seja 5.

Veja a [solução](./solucoes/04-alteracao-de-valor.go)