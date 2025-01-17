package main

import (
	"fmt"
)

func main() {
	// Khai báo các biến
	mealCost := 88.67
	taxRate := 6.75 / 100
	tipRate := 20.0 / 100

	// Tính tiền thuế
	taxAmount := mealCost * taxRate

	// Tính tổng tiền sau khi cộng thuế
	totalWithTax := mealCost + taxAmount

	// Tính tiền tip dựa trên tổng tiền sau thuế
	tipAmount := totalWithTax * tipRate

	// Tính tổng hóa đơn
	totalBill := totalWithTax + tipAmount

	// Hiển thị kết quả
	fmt.Printf("Chi phí bữa ăn: $%.2f\n", mealCost)
	fmt.Printf("Tiền thuế: $%.2f\n", taxAmount)
	fmt.Printf("Tiền tip: $%.2f\n", tipAmount)
	fmt.Printf("Tổng hóa đơn: $%.2f\n", totalBill)
}
