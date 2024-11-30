# Equação de Torricelli para o movimento uniformemente variado

Em física, é chamado movimento uniformemente variado o movimento realizado com
aceleração constante. A fórmula de Torricelli nos permite entender esse
movimento quando não temos informação sobre o tempo decorrido. A fórmula é dada
por:

```
vf^2 = vi^2 + 2 * a * dx
```

Onde:

- vf é a valocidade final
- vi é a velocidade inicial
- a é a aceleração
- dx é o deslocamento

Para sabermos `vf` tendo `vi`, `a` e `dx`, podemos usar a fómula:

```
vf = sqrt(vi^2 + 2 * a * dx)
```

Lembrando que `sqrt` é como notamos a raíz quadrada.

Escreva um programa que implementa a fórmula acima e recebe do usuário os
valores do tipo `float64` para `vi`, `a` e `dx`.

Veja a [solução](./solucoes/04-torricelli-muv.go)