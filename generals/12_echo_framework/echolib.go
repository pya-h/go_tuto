package main

import (
	"errors"
	"github.com/labstack/echo"
	"net/http"
)

type APIError struct {
	Status  int
	Message string
}
type User struct {
	Name string `json:"name"`
	Id   uint8  `json:"id"`
}

func (self APIError) Error() string {
	return self.Message
}

func main() {

	e := echo.New()
	e.GET("/user", echoGetUser)
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]any{"message": "fuck world!"})
	})
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if apiError, ok := err.(APIError); ok {
			ctx.JSON(apiError.Status, apiError)
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]any{"error": "internal fucked up situation!"})
	}
	e.Logger.Fatal(e.Start(":8000"))
}

func UserNotFoundError() APIError {
	return APIError{http.StatusNotFound, "User not found!"}
}

func echoGetUser(ctx echo.Context) error {
	if user, err := getTestUser(); err != nil {
		return UserNotFoundError()
	} else {
		return ctx.JSON(http.StatusOK, user)
	}
}

func getTestUser() (*User, error) {
	return &User{"whatever", 1}, errors.New("not Found")
}
