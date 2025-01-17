package main

import (
	"errors"
	"fmt"
)

// Ham ko co gia tri tra ve
func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Ham tra ve 1 gia tri
func sum(num int) int {
	var result = 0
	for i := 1; i <= num; i++ {
		result += i
	}
	return result
}

func sumV2(nums ...int) int { //Ham nhieu bien
	var result = 0
	for i := 0; i <= len(nums); i++ {
		result += nums[i]
	}
	return result
}

// Ham co gia tri tra ve
func greeting(name string) string {
	result := fmt.Sprintf("Hello %s", name)
	return result
}
func main() {
	sayHello("Lan Anh")

	result := greeting("Bob")
	fmt.Println(result)

	var result1 = sum(9)
	fmt.Println(result1)

	result2 := sumV2(1, 2, 3, 4, 5, 6)
	fmt.Println(result2)

	var result3, err = division(1, 0)
	fmt.Println(result3)

	if err != nil {
		fmt.Println(err)
	}

	var result4, _ = division(10, 0)
	fmt.Println(result4)

	if err != nil {
		fmt.Println(err)
	}
}

func division(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("Division by zero")
	}
	var result = num1 / num2
	return result, nil
}

// Sd con tro
func divisionV2(num1, num2 float64) (*float64, error) {
	if num2 == 0 {
		return nil, errors.New("Division by zero")
	}
	var result = num1 / num2
	return &result, nil
}
