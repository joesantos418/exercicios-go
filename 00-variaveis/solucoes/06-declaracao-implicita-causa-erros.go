package main

import (
	"fmt"
	"time"
)

func main() {
	a := int64(695998980)
	b := time.Unix(a, 0)
	fmt.Println(b.Format(time.RFC3339Nano))
}
