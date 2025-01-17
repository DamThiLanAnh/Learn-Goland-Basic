package main

import "fmt"

func main() {
	var counter = 100
	var p *int = &counter

	fmt.Println(*p)       //Thong qua dia chi lay kq
	fmt.Println(p)        //Tro ve dia chi cua counter
	fmt.Println(&counter) //Lay dia chi

	var name1 = "Lan"
	var name2 = "Lan"

	var p1, p2 *string
	p1 = &name1
	p2 = &name2
	fmt.Println(p1 == p2)   //false vi 2 dia chi la khac nhau
	fmt.Println(*p1 == *p2) // gia tri cua bien
}
