# Bytes

Considere o código a seguir:

```
	type pessoa struct {
		Nome  string
		Idade int
	}

	p := pessoa{
		Nome:  "John Doe",
		Idade: 25,
	}

	pessoaBytes, _ := json.Marshal(p)
```

Escreva um programa que imprime a respresentação em texto da variável `p`

Veja a [solução](./solucoes/07-bytes.go)