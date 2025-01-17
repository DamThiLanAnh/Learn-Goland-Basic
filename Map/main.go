package main

import (
	"fmt"
)

func main() {
	//Khai bao 1 map trong go land
	var myMap = make(map[string]int)
	fmt.Println(myMap)
	if myMap == nil {
		fmt.Println("myMap = nil")
	} else {
		fmt.Println("myMap != nil")
	}

	fmt.Println()
	fmt.Println()

	var myMap1 map[string]int
	fmt.Println(myMap1)
	if myMap1 == nil {
		fmt.Println("myMap1 = nil")
	} else {
		fmt.Println("myMap1 != nil")
	}
	fmt.Println()
	//Khai bao voi gia tri khoi tao
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	fmt.Println(myMap2)

	//Them 1 phan tu vao map
	myMap2["key4"] = 4
	myMap2["key5"] = 5
	fmt.Println(myMap2)

	//delete 1 phan tu trong map = delete(map, key)
	delete(myMap2, "key1")
	fmt.Println(myMap2)

	//length cua map
	fmt.Println(len(myMap2))

	//Map la 1 reference type
	myMap3 := myMap2
	fmt.Println(myMap3)
	myMap3["key5"] = 1000
	fmt.Println(myMap3)
	fmt.Println(myMap2)

	//Truy cap 1 phan tu trong map
	value, found := myMap2["key3"]
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("Khong tim thay gia tri vs key 1000")
	}

	//Trong map khong co toan tu ==
}
