# Múltiplos retornos

Considere a seguinte função:

```
func divideNome(nomeCompleto string) (string, string) {
    nomes := strings.Fields(nomeCompleto)

    return nomes[0], nomes[1]
}
```

Escreva um programa que recebe o primeiro e o último nome do usuário separados
por espaço e os imprime na tela usando a função dada acima.

Veja a [solução](./solucoes/08-multiplos-retornos.go)
