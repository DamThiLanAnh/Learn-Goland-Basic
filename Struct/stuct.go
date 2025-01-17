package main

import "fmt"

func main() {
	type Student struct {
		Name  string
		Msv   int
		Email string
	}
	var student Student = Student{
		Name:  "Lan anh",
		Msv:   42908,
		Email: "damthilananh8@gmail.com",
	}
	fmt.Println(student)
	fmt.Println(student.Msv)
	fmt.Println(student.Name)
	fmt.Println(student.Email)

	//Khai bao 1 struct ko ten (vo danh)
	var anonymousStruct = struct {
		Name  string
		Msv   int
		Email string
	}{
		Name:  "Lan anh",
		Msv:   42908,
		Email: "damthilananh8@gmail.com",
	}
	fmt.Println("======================")
	fmt.Println(anonymousStruct)

	if student == anonymousStruct {
		fmt.Println("Bang nhau nhe")
	} else {
		fmt.Println("Ko bang nhau nhe")
	}

	// Ke thua
	type Person struct {
		Student
		ID   string
		Khoa string
	}
	var person = Person{
		Student: Student{
			Name:  "Lan anh",
			Msv:   42908,
			Email: "damthilananh8@gmail.com",
		},
		ID:   "hihi",
		Khoa: "Toan Tin",
	}
	fmt.Println(person)
}

// Stuct giống class ở ngôn ngữ khác
// Muốn struct private thì viết thường:
/*
type Student struct {
		name  string
		msv   int
		email string
	}
*/
