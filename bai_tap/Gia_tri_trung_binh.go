package main

import "fmt"

func main() {
	var (
		a int = 28
		b int = 32
		c int = 37
		d int = 24
		e int = 33
	)
	var sum = a + b + c + d + e
	trung_binh := sum / 5
	fmt.Println(trung_binh)
}
