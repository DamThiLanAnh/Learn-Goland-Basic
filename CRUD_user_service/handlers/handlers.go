package handlers

import (
	"github.com/labstack/echo/v4"
	"golang-basic/CRUD_user_service/entity"
	"net/http"
	"strconv"
)

var (
	users = map[int]*entity.User{}
	seq   = 1
)

// Khai báo một hàm xử lý yêu cầu tạo người dùng, trả về error nếu có lỗi.
func CreateUser(c echo.Context) error {
	// Tạo 1 con trỏ u đến 1 đối tượng User moi, khoi tao ID voi gia tri tu seq
	u := &entity.User{
		ID: seq,
		//Name: "Lan Anh",
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u //doan nay trong kieu du lieu map
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
