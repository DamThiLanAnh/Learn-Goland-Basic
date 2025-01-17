package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) doSomthing() {
	p.Name = "Lan Anh- doSomthing"
	fmt.Println("doSomthing method")
}

// Xu ly request or API thi sai poiter
func (p *Person) doSomthing1() {
	p.Name = "Lan Anh- doSomthing1"
	fmt.Println("doSomthing method poiter")
}
func main() {
	var p Person
	p.doSomthing()
	//p.Name = "hihi"
	fmt.Println(p.Name)
	p.doSomthing1()
	fmt.Println(p.Name)
}

// Trung hop cai struct qua lon thi minh sai poiter
// Hoac can sua truc tiep gi do o struct
