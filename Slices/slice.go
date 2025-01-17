package main

import "fmt"

func main() {
	//Khai bao 1 slice ko can cung cap size cu the
	var slice []int
	fmt.Println(slice)

	//Khai bao va khoi tao slice
	var slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	//Tao mot slice tu 1 array
	var array = [4]int{1, 2, 3, 4}
	slice2 := array[1:3] //array[1] -> array[3-1=2]
	fmt.Println(slice2)

	slice3 := array[2:]
	fmt.Println(slice3)

	slice4 := array[:] //Lay het array
	fmt.Println(slice4)

	slice5 := array[:3] // 0, 1, 2
	fmt.Println(slice5)

	//Tao 1 slice tu 1 slice khac
	var slice6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[1:]
	fmt.Println(slice8)

	//Slice la 1 reference type
	//Array la 1 value type
	fmt.Println("====================")
	var array1 = [5]int{1, 2, 3, 4, 5}
	slice9 := array1[:]
	slice9[0] = 23
	fmt.Println(slice9)
	fmt.Println(array1)

	//length va capacity cua slice
	countries := [...]string{"VN", "USA", "Canada", "France", "China"}
	slice10 := countries[3:4]
	fmt.Println(slice10)

	fmt.Println(len(slice10)) //len la so luong phan tu cua slice
	fmt.Println(cap(slice10)) //cap la so luong phan tu cua underlying array bat dau tu vi tri start khi ma slice dc tao
	fmt.Println("============================")
	// make, copy, append
	//1. Ham make de khai bao 1 slice cung cap luon len va cap
	slice11 := make([]int, 2, 5)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	fmt.Println("============================")
	slice12 := make([]int, 2)
	fmt.Println(slice12)
	fmt.Println(len(slice12))
	fmt.Println(cap(slice12))

	fmt.Println("============================")
	//2.Ham append tham gia tri vao cai slice cua minh
	var slice13 []int
	slice13 = append(slice13, 100)
	fmt.Println(slice13)

	//3.Ham copy
	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 2)

	number := copy(dest, src)
	fmt.Println(dest)
	fmt.Println(number)

	//delete item with index = 1
	src = append(src[:1], src[2:]...) //noi 2 slice vs nhau = append (slice1, slice2...)
	fmt.Println(src)
}
