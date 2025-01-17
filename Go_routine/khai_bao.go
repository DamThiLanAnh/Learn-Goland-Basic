package main

import (
	"fmt"
	"time"
)

func main() {
	// Tạo 1 goroutine chạy 1 hàm ẩn danh (anonymous function)
	go func() {
		fmt.Println("inside Goroutin") //In ra chuỗi "inside Goroutinr"
	}()
	// Tạo một goroutine gọi hàm runTask
	go runTask()
	// Chờ một chút để các goroutine có thời gian chạy xong

	time.Sleep(1 * time.Second)
}
func runTask() {
	fmt.Println("inside runTask")
}
