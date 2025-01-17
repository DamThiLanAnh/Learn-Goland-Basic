package main

import "fmt"

func main() {
	//Khai bao array
	var myArray [4]int
	fmt.Println(myArray)
	//Gan gia tri cho array vua khai bao
	myArray[0] = 123
	fmt.Println(myArray)
	myArray[1] = 234
	fmt.Println(myArray)

	//Khai bao 1 array va khoi tao gia tri luon cho no
	a := [3]int{2, 4, 7}
	fmt.Println(a)

	//Khai bao mang ko can set size
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(len(arr))
	fmt.Println(arr[3])

	//array la value type khong phai reference type
	countries := [...]string{"VN", "US", "France"}
	copyCountries := countries
	fmt.Println(countries)
	fmt.Println(copyCountries)

	fmt.Println("=====================")
	//loop array
	//cach1
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}
	fmt.Println("======================")
	//cach2
	for index, value := range countries {
		fmt.Printf("i = %d, value = %s\n", index, value)
	}

	//Khai bao mang 2 chieu [hang][cot] [row][col]
	matrix := [4][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{6, 7},
	}
	fmt.Println(matrix)
	fmt.Println("==================")
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
