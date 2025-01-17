package main

import "fmt"

func So_dam_tren_moi_gallon(gallon_xang int, di_chuyen int) int {
	return gallon_xang / di_chuyen
}
func main() {
	var gallon_xang int = 15
	var di_chuyen int = 375

	result := So_dam_tren_moi_gallon(gallon_xang, di_chuyen)
	fmt.Println(result)
}
