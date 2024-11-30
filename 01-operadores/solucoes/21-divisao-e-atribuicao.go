package main

import "fmt"

func main() {
	num := 1024
	num /= 256
	num /= 2

	fmt.Println(num)
}
