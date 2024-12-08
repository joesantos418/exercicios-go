# Runas

Considere o código a seguir:

```
nome := "João da Silva Málaga"

for i := 0; i < len(nome); i++ {
    fmt.Printf("%s ", string(nome[i]))
}
```

Esse código deveria imprimir o nome João da Silva Málaga com um espaço entre
cada letra, porém não é isso que está acontecendo. Corrija o erro para que ele
imprima corretamente.

Dica: Utilize o tipo [`rune`](https://go.dev/blog/strings#code-points-characters-and-runes)

Veja a [solução](./solucoes/07-bytes.go)
