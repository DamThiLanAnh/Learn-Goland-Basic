package main

import (
	"errors"
	"fmt"
)

/*
Viet 1 ham chia 2 so tra ve ket qua va 1 loi
Goi ham errors.New de tao 1 doi tuong loi voi thong bao "divide by zero"
*/

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	//tao 2 bien de hung 2 gia tri tra ve cua ham divide
	result, err := divide(3.4, 1.2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	result, err = divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
