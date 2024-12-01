# Nomes e idades

Considere o codigo abaixo:

```
idadesENomes := map[int]string{
    26: "Ross",
    24: "Monica",
    24: "Rachel",
    26: "Chandler",
    25: "Joey",
    27: "Phoebe",
}
```

O código acima declara um [mapa](https://go.dev/blog/maps) de idades para nomes
das personagens da série [Friends](https://www.imdb.com/title/tt0108778/).

Escreva um programa que imprime uma linha para cada personagem no formato:

```
Personagem tem X anos.
```

Veja a [solução](./solucoes/10-nomes-e-idades.go)