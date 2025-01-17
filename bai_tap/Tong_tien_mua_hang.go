package main

import "fmt"

func main() {
	var (
		do_1 = 15.59
		do_2 = 24.59
		do_3 = 6.59
		do_4 = 12.59
		do_5 = 3.59
	)
	sum := do_1 + do_2 + do_3 + do_4 + do_5
	tong_phai_tra := sum * 7 / 100
	fmt.Println("tong tam tinh: ", sum)
	fmt.Println("tong phai tra", tong_phai_tra)

}
