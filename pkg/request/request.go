package request

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Create contextwrapper interface and struct
type (
	contextwrapperService interface {
		Bind(data any) error
	}

	contextwrapper struct {
		Context   echo.Context
		validator *validator.Validate
	}
)

// function instance contextWrapper
func ContextWrapper(ctx echo.Context) contextwrapperService {
	return &contextwrapper{
		Context:   ctx,
		validator: validator.New(),
	}
}

// function validate
func (c *contextwrapper) Bind(data any) error {
	if err := c.Context.Bind(data); err != nil {
		log.Printf("Error: Bind data faild: %s", err.Error())
	}

	if err := c.validator.Struct(data); err != nil {
		log.Printf("Error: validate data faild: %s", err.Error())
	}

	return nil
}
