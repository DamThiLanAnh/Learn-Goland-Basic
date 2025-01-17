package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}
func main() {
	var a int = 50
	var b int = 100

	var total int = sum(a, b)

	fmt.Printf("Tong cua 50 va 100 la: %d\n", total)
}
