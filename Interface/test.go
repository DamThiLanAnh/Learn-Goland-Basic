package main

import "fmt"

func main() {
	var w Writer = &ConsoleWriter{}
	w.Write([]byte("Hello World"))
}

type Writer interface {
	Write([]byte) (int, error) //Khai báo tên hàm, tham số, giá trị trả về, ko có thân hàm
}
type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}