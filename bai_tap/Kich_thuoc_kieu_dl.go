package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var c byte
	var i int
	var f float32
	var d float64
	var s string

	fmt.Printf("Kich thuoc cua kieu char (byte): %d byte\n", unsafe.Sizeof(c))
	fmt.Printf("Kich thuoc cua kieu int: %d byte\n", unsafe.Sizeof(i))
	fmt.Printf("Kich thuoc cua kieu float32: %d byte\n", unsafe.Sizeof(f))
	fmt.Printf("Kich thuoc cua kieu float64: %d byte\n", unsafe.Sizeof(d))
	fmt.Printf("Kich thuoc cua kieu string: %d byte\n", unsafe.Sizeof(s))

}
