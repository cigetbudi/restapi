package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type Users struct {
	Email       string `json:"email" form:"email"`
	Nama        string `json:"nama" form:"nama"`
	NoHandphone string `json:"no_handphone" form:"no_handphone"`
	ALamat      string `json:"alamat" form:"alamat"`
	Ktp         string `json:"ktp" form:"ktp"`
}

func main() {
	route := echo.New()
	route.POST("user/create_user", func(c echo.Context) error {
		user := new(Users)
		c.Bind(user)
		// var user Users
		// c.Bind(&user)
		contentType := c.Request().Header.Get("Content-Type")
		if contentType == "application/json" {
			fmt.Println("Request dari json")
		} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
			file, err := c.FormFile("ktp")
			if err != nil {
				fmt.Println("ktp kosong")
			} else {
				src, err := file.Open()
				if err != nil {
					return err
				}
				defer src.Close()
				dst, err := os.Create(file.Filename)
				if err != nil {
					return err
				}
				defer dst.Close()
				if _, err = io.Copy(dst, src); err != nil {
					return err
				}

				user.Ktp = file.Filename
				fmt.Println("fileada, akan disimpan")
			}
		}
		response := struct {
			Message string
			Data    Users
		}{
			Message: "Sukses menambahkan data",
			Data:    *user,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PUT("user/update_user/:email", func(c echo.Context) error {
		user := new(Users)
		c.Bind(user)
		user.Email = c.Param("email")
		//ngapain ya
		response := struct {
			Message string
			Data    Users
		}{
			Message: "Sukses mengubah data",
			Data:    *user,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("user/delete_user/:email", func(c echo.Context) error {
		user := new(Users)
		user.Email = c.Param("email")
		//ngapain ya
		response := struct {
			Message string
			ID      string
		}{
			Message: "Sukses menghapus data",
			ID:      user.Email,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("user/search_user", func(c echo.Context) error {
		user := new(Users)
		user.Email = c.QueryParam("keywords")
		user.Nama = "Sigit Budi"
		user.ALamat = "Jalan apa ya"
		user.Ktp = "file.jpg"
		//ngapain ya
		response := struct {
			Message string
			Data    Users
		}{
			Message: "Berhasil cek data",
			Data:    *user,
		}
		return c.JSON(http.StatusOK, response)

	})
	route.Start(":9000")
}
