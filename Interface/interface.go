package main

import (
	"errors"
	"fmt"
)

const Pi = 3.14

type Shape interface {
	Area(x, y float64) (float64, error) // Để nhận tham số x và y
	Perimeter() (float64, error)
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float32
}

// Diện tích hình chữ nhật sử dụng x, y như width, height
func (r Rectangle) Area(x, y float64) (float64, error) {
	if x <= 0 || y <= 0 {
		return 0.0, errors.New("width and height must be positive")
	}
	return x * y, nil
}

// Diện tích hình tròn, x và y không được sử dụng, nhưng cần để phù hợp với giao diện
func (c Circle) Area(x, y float64) (float64, error) {
	if c.radius <= 0 {
		return 0.0, errors.New("radius must be positive")
	}
	return float64(c.radius * c.radius * Pi), nil
}

func (r Rectangle) Perimeter() (float64, error) {
	if r.width <= 0 || r.height <= 0 {
		return 0.0, errors.New("width and height must be positive")
	}
	return 2 * (r.width + r.height), nil
}

func (c Circle) Perimeter() (float64, error) {
	if c.radius <= 0 {
		return 0.0, errors.New("radius must be positive")
	}
	return float64(c.radius * 2 * Pi), nil
}

func main() {
	var s Shape

	// Tạo 1 hình chữ nhật
	r := Rectangle{width: -2, height: 5}
	s = r
	area, err := s.Area(r.width, r.height) // Truyền width và height
	if err != nil {
		fmt.Println("Error calculating area:", err)
	} else {
		fmt.Println("Rectangle Area:", area)
	}

	perimeter, err := s.Perimeter()
	if err != nil {
		fmt.Println("Error calculating perimeter:", err)
	} else {
		fmt.Println("Rectangle Perimeter:", perimeter)
	}

	// Tạo 1 hình tròn
	c := Circle{radius: 7}
	s = c
	area, err = s.Area(0, 0) // Truyền 0, 0 vì Circle không sử dụng x, y
	if err != nil {
		fmt.Println("Error calculating area:", err)
	} else {
		fmt.Println("Circle Area:", area)
	}

	perimeter, err = s.Perimeter()
	if err != nil {
		fmt.Println("Error calculating perimeter:", err)
	} else {
		fmt.Println("Circle Perimeter:", perimeter)
	}
}
