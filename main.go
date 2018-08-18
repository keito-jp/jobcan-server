package main

import (
	"net/http"

	"github.com/keito-jp/jobcan-cli/jobcan"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Account is account information of Jobcan.
type Account struct {
	ClientID string `json:"client_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JobcanStatus is status of Jobcan account.
type JobcanStatus struct {
	Status string `json:"status"`
}

// readStatus is function get status of Jobcan account.
func readStatus(c echo.Context) (err error) {
	a := &Account{}
	if err = c.Bind(a); err != nil {
		return err
	}
	j, err := jobcan.NewJobcan(a.ClientID, a.Email, a.Password)
	if err != nil {
		return err
	}
	s, err := j.Status()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &JobcanStatus{Status: s})
}

// punch is function punch in.
func punch(c echo.Context) (err error) {
	a := &Account{}
	if err = c.Bind(a); err != nil {
		return err
	}
	j, err := jobcan.NewJobcan(a.ClientID, a.Email, a.Password)
	if err != nil {
		return err
	}
	err = j.Punch()
	if err != nil {
		return err
	}
	s, err := j.Status()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &JobcanStatus{Status: s})
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/status", readStatus)
	e.POST("/punch", punch)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
