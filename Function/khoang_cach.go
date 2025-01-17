package main

import (
	"fmt"
	"math"
)

func kc(xA, yA, xB, yB float64) float64 {
	a := math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
	return a
}

func main() {
	result := kc(3.2, -1.4, -5.7, 6.1)
	fmt.Println(result)
}
