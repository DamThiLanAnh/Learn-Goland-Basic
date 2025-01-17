package main

import "fmt"

func main() {
	var height int
	fmt.Print("Nhập chều cao của kim tự tháp: ")
	fmt.Scan(&height)

	for i := 1; i <= height; i++ {
		//Vòng for đầu tien
		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
