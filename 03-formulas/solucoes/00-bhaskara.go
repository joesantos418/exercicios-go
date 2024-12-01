package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1
	b := 12
	c := -13

	x1 := (float64(-1*b) + math.Sqrt(float64((b*b)-(4*a*c)))) / (float64(2 * a))
	x2 := (float64(-1*b) - math.Sqrt(float64((b*b)-(4*a*c)))) / (float64(2 * a))

	fmt.Println(x1)
	fmt.Println(x2)
}
