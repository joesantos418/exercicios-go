# Fórmula de Bhaskara

Uma equação de segundo grau é uma equação na forma:

```
ax^2 + bx + c = 0
```

onde `a`, `b` e `c` são números reais. Note que `x^2` significa "x ao quadrado".
É Comum utilizar o acento circumflexo para denotar potências quando escrevemos
fórmulas em texto;

Equações dessa forma possuem 0, 1 ou 2 soluções reais, dadas pela fórmula
conhecida como fórmula de Bhaskara:

```
x1 = (-b + sqrt(b^2 - 4ac)) / 2a

e

x2 = (-b - sqrt(b^2 - 4ac)) / 2a

```

Note que `sqrt()` significa raíz quadrada, uma abreviação do inglês
<em>square root<em>.

Quando o termo `b^2 - 4ac`, chamado de discriminante, é negativo, a raíz
quadrada real não existe e a equação possui zero soluções reais. Quando o
discriminante é igual a zero, `x1` é igual a `x2` e a equação possui apenas uma
raíz real.

Escreva um programa que implementa a fórmula de Bhaskara para resolver a
seguinte equação de segundo grau:

```
x^2 + 12x - 13 = 0
```

Dica use a função [`math.Sqrt`](https://pkg.go.dev/math#Sqrt) e [conversão de inteiro para float](https://www.digitalocean.com/community/tutorials/how-to-convert-data-types-in-go-pt#convertendo-os-numeros-inteiros-em-floats)

Veja a [solução](./solucoes/00-bhaskara.go)