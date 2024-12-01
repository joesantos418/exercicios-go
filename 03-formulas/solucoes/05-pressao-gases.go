package main

import "fmt"

func main() {
	const R = 8.31
	p := float64(10)
	n := float64(12)
	t := float64(308)

	v := (n * R * t) / p

	fmt.Printf("O volume é %f\n", v)
}
