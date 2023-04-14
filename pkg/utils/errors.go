package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

// add switch other variant
func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	default:
		// e.Errors["body"] = v.Error()
		e.Errors["status"] = "fail"
		e.Errors["message"] = v.Error()
	}
	return e
}

func NewValidatorError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.Errors[v.Field()] = fmt.Sprintf("%v", v.Tag())
	}
	return e
}

func AccessForbidden() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	// e.Errors["body"] = "access forbidden"
	e.Errors["status"] = "fail"
	e.Errors["message"] = "access forbidden"
	return e
}

func NotFound(obj string) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	// e.Errors["body"] = "resource not found"
	e.Errors["status"] = "fail"
	e.Errors["message"] = fmt.Sprintf("%s not found", obj)
	return e
}

func BadRequest() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["status"] = "fail"
	e.Errors["message"] = "bad  request"
	return e
}
func BadGateway() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["status"] = "error"
	e.Errors["message"] = "Something bad happened"
	return e
}
func InvalidCredentials() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["status"] = "fail"
	e.Errors["message"] = "Invalid email or password"
	return e
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
