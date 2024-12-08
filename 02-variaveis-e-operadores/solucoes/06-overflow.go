package main

import "fmt"

func main() {
	var a uint8 = 255

	a += uint8(10)

	fmt.Println(a)
}
