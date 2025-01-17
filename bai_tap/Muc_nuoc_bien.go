package main

import "fmt"

func tang_len(muc_nuoc float64, nam int) float64 {
	return muc_nuoc * float64(nam)
}

func main() {
	var muc_nuoc float64 = 1.5

	// Tính mức nước tăng lên sau 5, 7, và 10 năm
	tang_5_nam := tang_len(muc_nuoc, 5)
	tang_7_nam := tang_len(muc_nuoc, 7)
	tang_10_nam := tang_len(muc_nuoc, 10)

	// Hiển thị kết quả
	fmt.Printf("Số mm sẽ tăng lên so với mức nước biển hiện tại sau 5 năm là: %.2f mm\n", tang_5_nam)
	fmt.Printf("Số mm sẽ tăng lên so với mức nước biển hiện tại sau 7 năm là: %.2f mm\n", tang_7_nam)
	fmt.Printf("Số mm sẽ tăng lên so với mức nước biển hiện tại sau 10 năm là: %.2f mm\n", tang_10_nam)
}
