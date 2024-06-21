package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Course struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Creater_id  int    `json:"creater_id"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

func onCourseCreate(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		course := new(Course)
		if err := c.Bind(course); err != nil {
			return err
		}

		query := fmt.Sprintf("INSERT INTO courses (title, description, creater_id, created_at, updated_at) VALUES ('%s', '%s', '%d', '%s', '%s')", course.Title, course.Description, course.Creater_id, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		_, err := db.Query(query)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"result":  true,
			"message": "Course Created Successfuly",
		})
	}
}

func onGetAllCourses(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT * from courses")
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Course

		for rows.Next() {
			var value Course

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Title, &value.Description, &value.Creater_id, &value.Created_at, &value.Updated_at)
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

func onGetCourse(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		rows, err := db.Query("SELECT * FROM courses WHERE id = ?", id)
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Course

		for rows.Next() {
			var value Course

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Title, &value.Description, &value.Creater_id, &value.Created_at, &value.Updated_at)
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
