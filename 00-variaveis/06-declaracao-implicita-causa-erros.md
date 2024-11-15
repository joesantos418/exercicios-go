# Cuidado com a declaração implícita

Uma data em Linux timestamp é dada por um número que representa a quantidade de
segundos passados desde 01/01/1970 00:00:00.

Considere o seguinte código:

```

func main() {
	a := 695998980
	b := time.Unix(a, 0)
	fmt.Println(b.Format(time.RFC3339Nano))
}

```

Note que ele falha devido ao fato da declaração implícita dar à variável `a` o
tipo `Int` enquanto a função `time.Unix` espera o tipo `Int64`. Corrija esse
código para que ele funcione.

Veja a [solução](./solucoes/06-declaracao-implicita-causa-erros.go)
