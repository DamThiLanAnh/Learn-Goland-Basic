package main

import (
	"fmt"
	"time"
)

func dowloadFile() {
	fmt.Println("file downloading...")
}
func main() {
	go dowloadFile()
	go func() {
		fmt.Println("start")
	}()
	var f = func() {
		fmt.Println("end")
	}
	go f()
	time.Sleep(1 * time.Second)
}
