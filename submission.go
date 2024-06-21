package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Submission struct {
	Id            int    `json:"id"`
	Score         int    `json:"score"`
	Feedback      string `json:"feedback"`
	Assessment_id int    `json:"assessment_id"`
	Creater_id    int    `json:"creater_id"`
}

func onSubmissionCreate(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		submission := new(Submission)
		if err := c.Bind(submission); err != nil {
			return err
		}

		query := fmt.Sprintf("INSERT INTO submissions (score, feedback, assessment_id, creater_id) VALUES ('%d', '%s', '%d', '%d')", submission.Score, submission.Feedback, submission.Assessment_id, submission.Creater_id)
		_, err := db.Query(query)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"result":  true,
			"message": "Submission Created Successfuly",
		})
	}
}

func onGetAllSubmissions(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT * from submissions")
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Submission

		for rows.Next() {
			var value Submission

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Score, &value.Feedback, &value.Assessment_id, &value.Creater_id)
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

func onGetSubmission(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		rows, err := db.Query("SELECT * FROM submissions WHERE id = ?", id)
		if err != nil {
			return err
		}
		defer rows.Close()

		var data []Submission

		for rows.Next() {
			var value Submission

			// Scan the column values into the variables
			err = rows.Scan(&value.Id, &value.Score, &value.Feedback, &value.Assessment_id, &value.Creater_id)
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
