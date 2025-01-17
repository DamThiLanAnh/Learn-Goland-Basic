package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type CustomerAccounts struct {
	Name               string
	Address            string
	Company            string
	Phone              string
	Balance            string
	DateOfPaymentFinal time.Time
}

func NewCustomerAccounts(name string, address string, city string,
	phone string, balance string, dateOfPaymentFinal time.Time) CustomerAccounts {
	return CustomerAccounts{
		Name:               name,
		Address:            address,
		Company:            city,
		Phone:              phone,
		Balance:            balance,
		DateOfPaymentFinal: dateOfPaymentFinal,
	}
}

func (c *CustomerAccounts) Input() {
	var customer CustomerAccounts //doi tuong duoc khoi tao tu ham dung
	fmt.Println("Input your accounts information: ")

	fmt.Println("Name: ")
	fmt.Scan(&c.Name)

	fmt.Println("Address: ")
	fmt.Scan(&c.Address)

	fmt.Println("Company: ")
	fmt.Scan(&c.Company)

	fmt.Println("Phone: ")
	fmt.Scan(&c.Phone)
	for {
		fmt.Print("\nBalance: ")
		fmt.Scan(&customer.Balance)
		if customer.Balance < "0" {
			fmt.Println("Balance can't be less than 0, please input again")
		} else {
			break
		}
	}
	for {
		fmt.Print("Date of Payment Final (YYYY-MM-DD): ")
		var dateStr string
		fmt.Scan(&dateStr)
		var err error
		c.DateOfPaymentFinal, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("Invalid date format. Please try again.")
		} else {
			break
		}
	}
}

// func Update() CustomerAccounts {
//
// }
func (c CustomerAccounts) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent(" ", "  ") //Đặt định dạng cho JSON
	return encoder.Encode(c)
}

func List(filename string) ([]CustomerAccounts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var accounts []CustomerAccounts
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func main() {

	fmt.Println("************* MENU *************")
	fmt.Println("1. Nhap vao danh sach tai khoan")
	fmt.Println("2. Sua doi thong tin tai khoan")
	fmt.Println("3. Hien thi danh sach tai khoan")
	fmt.Println("4. Thoat khoi chuong trinh")
	fmt.Println("********************************")

	fmt.Println("Moi ban nhap: ")

	account := &CustomerAccounts{}
	account.Input()

	err := account.SaveToFile("accounts.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Account information saved to customer_accounts.json")
	}
	List()

}
