# Múltiplos retornos

Considere a seguinte função:

```
func divideNome(nomeCompleto string) (string, string) {
    nomes := strings.Fields(nomeCompleto)

    return nomes[0], nomes[1]
}
```

Escreva um programa que usa a função dada acima para seprar o nome
"Pedro Joaquim" e imprimir cada um em uma linha na tela.

Veja a [solução](./solucoes/08-multiplos-retornos.go)
