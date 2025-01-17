package main

import "fmt"

func main() {
	//Tổng doanh thu
	totalRevenua := 8.6

	//	Tỷ lệ phần trăm doanh thu của Phòng sale The East Coast
	eastCoastPercentage := 0.58

	// Tính doanh thu của phòng sale The East Coast
	eastCoastRevenue := totalRevenua * eastCoastPercentage

	fmt.Printf("%.2f\n", eastCoastRevenue) //%.2f\n lấy giá trị thứ 2 sau dấu phẩy
}
