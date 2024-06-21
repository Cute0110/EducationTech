package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Resource struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Url       string `json:"url"`
	Course_id int    `json:"course_id"`
}

func onResourceCreate(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		resource := new(Resource)
		if err := c.Bind(resource); err != nil {
			return err
		}

		query := fmt.Sprintf("INSERT INTO resources (title, type, url, course_id) VALUES ('%s', '%s', '%s', '%d')", resource.Title, resource.Type, resource.Url, resource.Course_id)
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

func onGetAllResources(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT * from resources")
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Resource

		for rows.Next() {
			var value Resource

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Title, &value.Type, &value.Url, &value.Course_id)
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

func onGetResource(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		rows, err := db.Query("SELECT * FROM resources WHERE id = ?", id)
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Resource

		for rows.Next() {
			var value Resource

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Title, &value.Type, &value.Url, &value.Course_id)
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
