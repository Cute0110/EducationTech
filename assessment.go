package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Assessment struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Max_score int    `json:"max_score"`
	Course_id int    `json:"course_id"`
}

func onAssessmentCreate(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		assessment := new(Assessment)
		if err := c.Bind(assessment); err != nil {
			return err
		}

		fmt.Println(assessment)

		query := fmt.Sprintf("INSERT INTO assessments (title, type, max_score, course_id) VALUES ('%s', '%s', '%d', '%d')", assessment.Title, assessment.Type, assessment.Max_score, assessment.Course_id)
		_, err := db.Query(query)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"result":  true,
			"message": "Assessment Created Successfuly",
		})
	}
}

func onGetAllAssessments(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT * from assessments")
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Assessment

		for rows.Next() {
			var value Assessment

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Title, &value.Type, &value.Max_score, &value.Course_id)
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

func onGetAssessment(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		rows, err := db.Query("SELECT * FROM assessments WHERE id = ?", id)
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Assessment

		for rows.Next() {
			var value Assessment

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Title, &value.Type, &value.Max_score, &value.Course_id)
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
