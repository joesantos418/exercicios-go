# Números

Nosso sistema de numeração é chamado de decimal, pois ele tem base 10, mas
existem outras bases e algumas delas são muito utilizadas em sistemas
computacionais:

- Base 2: também chamado de sistema binário. Os zeros e uns que um computador usa internamente
- Base 8: sistema octal. 8 é a potência de 2 mais próxima de dez.
- Base 16: sistema hexadecimal. Permite escrever numéros maiores usando menos caracteres.

Para diferenciar números escritos em octal, usamos um zero à esquerda do número.
Por exemplo:

- 06 é o número seis em octal
- 011 é o número treze em octal
- 02331136 é o número seicentos e trinta e cinco mil, quatrocentos e oitenta e seis (635.486) em octal

Para diferenciar números em hexadecimal, usamos os prefixos: 0x, 0X ou x. Note
que o sistema hexadecimal usa as letras A, B, C, D, E e F para representar os
números dez, onze, doze, treze, catorze e quinze, respectivamente Por exemplo:

- 0x5 é o número cinco em hexadecimal
- 0XA é o número dez em hexadecimal
- x9B25E é o número seicentos e trinta e cinco mil, quatrocentos e oitenta e seis (635.486) em hexadecimal

Escreva um programa que recebe um número do usuário e o imprime em notação
decimal, octal, hexadecimal e binária.

Dica: considere as seguintes documentações:

- [Pacote `fmt`](https://pkg.go.dev/fmt)
- [Função `fmt.Printf`](https://pkg.go.dev/fmt#Printf)
- [Documentação adicional dos verbos do pacote `fmt`](https://www.w3schools.com/go/go_formatting_verbs.php)

Veja a [solução](./solucoes/02-numeros.go)