package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	FName     string `json:"fname"`
	LName     string `json:"lname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	CPassword string `json:"cpassword"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func onSignupHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}

		rows, err := db.Query("SELECT * FROM users WHERE Email = ?", user.Email)
		if err != nil {
			return err
		}
		defer rows.Close()

		// Check if rows are empty or not
		if rows.Next() {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result":  false,
				"message": "Email address already exists",
			})
		}

		// Create an MD5 hasher object
		hasher := md5.New()

		// Write data to the hasher
		hasher.Write([]byte(user.Password))

		// Calculate the MD5 hash
		hashBytes := hasher.Sum(nil)

		// Convert the hash bytes to a hexadecimal string
		hashString := hex.EncodeToString(hashBytes)

		query := fmt.Sprintf("INSERT INTO users (name, Email, Role, Password) VALUES ('%s', '%s', '%s', '%s')", user.FName+" "+user.LName, user.Email, user.Role, hashString)
		_, err = db.Query(query)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"result":  true,
			"message": "Register Success",
		})
	}
}

func onLoginHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}

		// Create an MD5 hasher object
		hasher := md5.New()

		// Write data to the hasher
		hasher.Write([]byte(user.Password))

		// Calculate the MD5 hash
		hashBytes := hasher.Sum(nil)

		// Convert the hash bytes to a hexadecimal string
		hashString := hex.EncodeToString(hashBytes)

		rows, err := db.Query("SELECT * FROM users WHERE Email = ?", user.Email)
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []UserInfo

		for rows.Next() {
			var value UserInfo

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Name, &value.Email, &value.Role, &value.Password)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, value)
		}
		if len(data) == 0 {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"result":  false,
				"message": "Login Failed",
			})
		} else {
			if hashString != data[0].Password {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"result":  false,
					"message": "Password is incorrect",
				})
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"result":  true,
			"message": "Login Success",
			"user":    data[0],
		})
	}
}
