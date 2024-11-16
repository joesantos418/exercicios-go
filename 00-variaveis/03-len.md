# Comprimento (length)

Escreva um programa que recebe o nome do usuário e retorna a quantidade de
letras do nome, supondo apenas caracteres alfanuméricos sem acentuação.

Escreva também uma saudação e imprima na tela a saudação e a quantidade de
caracteres da saudação.

Exemplo:

```
Digite seu nome:
Joe
Seu nome tem 3 caracteres
Ola, Joe
Essa saudação tem 8 caracteres
```

Note que `len` dá a quantidade de caracteres apenas para strings não acentuadas.

Dica: considere as seguintes documentações:

- [Função `fmt.Sprintf`](https://pkg.go.dev/fmt#Sprintf)
- [Função `len`](https://pkg.go.dev/builtin#len)

Veja a [solução](./solucoes/03-len.go)