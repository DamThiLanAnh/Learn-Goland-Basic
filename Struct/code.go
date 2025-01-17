package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	age  int
}

func main() {
	st1 := Student{
		id:   123,
		name: "Lan Anh",
		age:  20,
	}
	fmt.Println(st1)
	fmt.Println(st1.id)

	st2 := Student{345, "Gia Huy", 11}
	fmt.Println("Ten: ", st2.name)

	//Kieu khai bao ko can sd type struct
	var st3 Students = struct {
		id   int
		name string
	}{
		555,
		"hihi",
	}
	fmt.Println("Student3: ", st3)
}
