package main

import "fmt"

//Variadic function
//1. Khai bao variadic function
//2. pass 1 slice vao 1 variadic function
//3. change slice
//4. Ung dung cu variadic function: append(slice []Type, elems ...Type) []Type
//									append([]Type, args1, args2, argsN)
func addItem(item int, list ...int) {
	//100, 200, 300, 400 -> int[] {100, 200, 300, 400}
	// []int {1,2,3,4} -> int[] {[]int {1,2,3,4}}
	list = append(list, item, 11, 22, 33, 44, 55, 66)
	fmt.Println(list)
}

func change(list ...int) {
	list[0] = 999
}

func main() {
	addItem(1, 100, 200, 300, 400)

	var list = []int{1, 2, 3, 4}
	addItem(100, list...)

	change(list...)
	fmt.Println(list)
}
