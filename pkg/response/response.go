package response

import "github.com/labstack/echo/v4"

//Create package response
type (
	MsgResponse struct {
		Message string `json:"message"`
	}
)

//function response successful
func SuccessResponse(ctx echo.Context, statusCode int, data any) error {
	return ctx.JSON(statusCode, data)
}

//function response error
func ErrResponse(ctx echo.Context, statusCode int, message string) error {
	return ctx.JSON(statusCode, &MsgResponse{Message: message})
}
