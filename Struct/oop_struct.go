package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type AccStatus string

const (
	Normal AccStatus = "normal"
	Locked AccStatus = "locked"
)

// Zero value
type Teacher struct {
	Name   string
	Email  string
	Age    int
	Gender string
	Status AccStatus
}

func Sum(numbers ...int) int {
	return 10
}

// Nil trong golang la khong co gia tri
// private vs public se tuong ung voi chu hoa va chu thuong
// []byte
func main() {
	st1 := Teacher{
		Name:   "Lan Anh",
		Email:  "lananh@gmail.com",
		Age:    18,
		Gender: "nữ",
		Status: "normal",
	}

	//Khai báo struct kiểu anonymus struct (struct ko ten)

	var demoStr = struct {
		Name string
		Age  int
	}{
		Name: "Lan Anh",
		Age:  18,
	}

	spew.Dump(demoStr)
	fmt.Println("==================")

	spew.Dump(st1)
	fmt.Println("==============")
	fmt.Println(st1)
}
