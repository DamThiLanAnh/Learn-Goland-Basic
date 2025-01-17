package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Coffee struct {
	Name   string
	Price  float64
	Recipe map[string]int
}

type CoffeeMachine struct {
	Menu      []Coffee
	Inventory map[string]int
	Mutex     sync.Mutex
}

// Hiển thị menu
func (cm *CoffeeMachine) DisplayMenu() {
	fmt.Println("=== Menu ===")
	for _, coffee := range cm.Menu {
		fmt.Printf("%s: $%.2f\n", coffee.Name, coffee.Price)
	}
	fmt.Println("====================")
}

// Chọn cà phê
func (cm *CoffeeMachine) SelectCoffee(name string, payment float64) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()

	// Tìm loại cà phê
	var selectedCoffee *Coffee
	for _, coffee := range cm.Menu {
		if coffee.Name == name {
			selectedCoffee = &coffee
			break
		}
	}
	if selectedCoffee == nil {
		fmt.Println("Loại cà phê không có trong menu.")
		return
	}

	// Kiểm tra tiền
	if payment < selectedCoffee.Price {
		fmt.Printf("Số tiền không đủ. Giá %s là $%.2f.\n", selectedCoffee.Name, selectedCoffee.Price)
		return
	}

	// Kiểm tra nguyên liệu
	for ingredient, amount := range selectedCoffee.Recipe {
		if cm.Inventory[ingredient] < amount {
			fmt.Printf("Không đủ %s để pha %s.\n", ingredient, selectedCoffee.Name)
			return
		}
	}

	// Giảm nguyên liệu và trả lại tiền thừa
	for ingredient, amount := range selectedCoffee.Recipe {
		cm.Inventory[ingredient] -= amount
	}
	change := payment - selectedCoffee.Price
	fmt.Printf("Đang pha chế %s. Tiền thừa của bạn: $%.2f\n", selectedCoffee.Name, change)
}

// Hàm chính
func main() {
	// Khởi tạo máy cà phê
	machine := CoffeeMachine{
		Menu: []Coffee{
			{Name: "Espresso", Price: 2.5, Recipe: map[string]int{"coffee": 2, "water": 1}},
			{Name: "Cappuccino", Price: 3.0, Recipe: map[string]int{"coffee": 2, "water": 1, "milk": 1}},
			{Name: "Latte", Price: 3.5, Recipe: map[string]int{"coffee": 2, "water": 1, "milk": 2}},
		},
		Inventory: map[string]int{"coffee": 10, "water": 10, "milk": 10},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		machine.DisplayMenu()

		// Nhập loại cà phê
		fmt.Print("Nhập loại cà phê bạn muốn (hoặc nhập 'exit' để thoát): ")
		coffeeName, _ := reader.ReadString('\n')
		coffeeName = strings.TrimSpace(coffeeName)
		if coffeeName == "exit" {
			fmt.Println("Cảm ơn bạn đã sử dụng máy cà phê!")
			break
		}

		// Nhập số tiền
		fmt.Print("Nhập số tiền bạn đưa: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		payment, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Số tiền không hợp lệ. Vui lòng thử lại.")
			continue
		}

		// Xử lý chọn cà phê
		machine.SelectCoffee(coffeeName, payment)
		fmt.Println()
	}
}
