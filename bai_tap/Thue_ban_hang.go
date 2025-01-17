package main

import "fmt"

func main() {
	//	Giá trị của giỏ hàng
	cartAmount := 95.0

	//	Thuế bán hàng của tiều bang và quận
	stateTaxRate := 0.04
	countryTaxRate := 0.02

	//	Tính thuế bán hàng của tiểu bang và quận
	stateTax := cartAmount * stateTaxRate
	countyTax := cartAmount * countryTaxRate

	//	Tính tổng thuế
	totalTax := stateTax + countyTax

	fmt.Printf("Tổnh thuế bán hàng: %.2f đô", totalTax)
}
