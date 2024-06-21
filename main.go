package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create an instance of Echo
	e := echo.New()

	// Define the database connection parameters
	db, err := sql.Open("mysql", "MarttiCheng:fivestar@tcp(localhost:3306)/edutech")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL database")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS middleware
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", helloHandler)

	//Auth
	e.POST("/login", onLoginHandler(db))
	e.POST("/signup", onSignupHandler(db))

	//Course
	e.POST("/course", onCourseCreate(db))
	e.GET("/course/:id", onGetCourse(db))
	e.GET("/course/all", onGetAllCourses(db))

	//Enrollment
	e.POST("/enrollment", onEnrollmentCreate(db))
	e.GET("/enrollment/:id", onGetEnrollment(db))
	e.GET("/enrollment/all", onGetAllEnrollments(db))

	//Resource
	e.POST("/resource", onResourceCreate(db))
	e.GET("/resource/:id", onGetResource(db))
	e.GET("/resource/all", onGetAllResources(db))

	//Assessment
	e.POST("/assessment", onAssessmentCreate(db))
	e.GET("/assessment/:id", onGetAssessment(db))
	e.GET("/assessment/all", onGetAllAssessments(db))

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler function for the route "/"
func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
