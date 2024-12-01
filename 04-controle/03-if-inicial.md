# If inicial

Considere a função abaixo:

```
func validaParEDobra(num int) (int, error) {
    if num % 2 == 0 {
        return num * 2, nil
    }

    return 0, errors.New("Numero dado nao pode ser impar")
}
```

Escreva um programa que recebe um número do usuário, passa esse número para a
função acima e, caso o erro seja igual a `nil` (caso de sucesso), imprime o
número retornado na tela. Caso contrário, o programa deve imprimir o erra dado
na tela. A chamada à função acima deve estar dentro da condição do `if`.

Veja a [solução](./solucoes/03-if-inicial.go)