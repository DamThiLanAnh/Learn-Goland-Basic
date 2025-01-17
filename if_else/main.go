package main

import "fmt"

// Khai bao bien
// var n int = 10
// var a int = 20
/*
var (
	n = 10
	a = 20
)
var PI = 3.1415926

// Ham init he thong se goi cai ham nay
func init() {
	fmt.Println("init")
	fmt.Println(PI)
}
func main() {
	var email string = "damthilananh8@gmail.com"
	full_name := "Dam Thi Lan Anh"
	fmt.Println(n)
	fmt.Println(a)
	fmt.Println(PI)
	fmt.Println(email, "  ", full_name)
}
*/
func main() {
	//Cau lenh dieu khien
	var number = 10
	if number%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	//Cau lenh switch case
	switch number {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	}
}
