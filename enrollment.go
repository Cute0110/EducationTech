package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Enrollment struct {
	Id         int    `json:"id"`
	Creater_id int    `json:"creater_id"`
	Course_id  int    `json:"course_id"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func onEnrollmentCreate(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		enrollment := new(Enrollment)
		if err := c.Bind(enrollment); err != nil {
			return err
		}

		query := fmt.Sprintf("INSERT INTO enrollments (creater_id, course_id, created_at, updated_at) VALUES ('%d', '%d', '%s', '%s')", enrollment.Creater_id, enrollment.Course_id, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		_, err := db.Query(query)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"result":  true,
			"message": "Enrollment Created Successfuly",
		})
	}
}

func onGetAllEnrollments(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT * from enrollments")
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Enrollment

		for rows.Next() {
			var value Enrollment

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Creater_id, &value.Course_id, &value.Created_at, &value.Updated_at)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, value)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"result":  true,
			"courses": data,
		})
	}
}

func onGetEnrollment(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		rows, err := db.Query("SELECT * FROM enrollments WHERE id = ?", id)
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Enrollment

		for rows.Next() {
			var value Enrollment

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Creater_id, &value.Course_id, &value.Created_at, &value.Updated_at)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, value)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"result":  true,
			"courses": data[0],
		})
	}
}
