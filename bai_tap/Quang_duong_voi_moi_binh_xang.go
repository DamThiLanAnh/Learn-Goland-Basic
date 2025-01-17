package main

import "fmt"

//Ct tinh quang duong = dung tich binh xang * muc tieu thu nhien lieu
//
//func quang_duong(a, b float32) float32 {
//	return a * b
//}
//func main() {
//	var dung_tich_binh_xang float32 = 20
//	var thanh_pho float32 = 23.5
//	var cao_toc float32 = 28.9
//
//	Sthanhpho := quang_duong(dung_tich_binh_xang, thanh_pho)
//	Scaotoc := quang_duong(dung_tich_binh_xang, cao_toc)
//
//	fmt.Println("Quang duong di chuyen trong thanh pho la: ", Sthanhpho)
//	fmt.Println("Quang duong khi di chuyen tren duong cao toc la: ", Scaotoc)
//}

//Sử dụng hằng số thay vì dùng biến trong hàm main để cho thấy rằng
//những giá trị này là cố định và không thay đổi trong quá trình thực thi.

const dung_tich_binh_xang float32 = 20
const thanh_pho float32 = 23.5
const cao_toc float32 = 28.9

func quangDuong(a, b float32) float32 {
	return a * b
}

func main() {
	Sthanhpho := quangDuong(dung_tich_binh_xang, thanh_pho)
	Scaotoc := quangDuong(dung_tich_binh_xang, cao_toc)

	fmt.Println(Sthanhpho, Scaotoc)
}
