# Múltiplas reatribuições

Considere a seguinte função:

```
func divideNome(nomeCompleto string) (string, string) {
    nomes := strings.Fields(nomeCompleto)

    return nomes[0], nomes[1]
}
```
Reutilizando as variáveis declaradas, escreva um programa que chama a função
dada acima para os seguintes nomes:

- Chanandler Bong
- Joey Tribbiani
- Monica Geller
- Phoebe Buffay
- Rachel Green
- Ross Geller

Veja a [solução](./solucoes/09-multiplas-reatribuicoes.go)
