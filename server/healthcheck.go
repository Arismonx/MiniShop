package server

import (
	"net/http"

	"github.com/Arismonx/MiniShop/pkg/response"
	"github.com/labstack/echo/v4"
)

// struct for check name
type healthCheck struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

// function healthCheckService
func (s *server) healthCheck(ctx echo.Context) error {
	return response.SuccessResponse(ctx, http.StatusOK, &healthCheck{
		App:    s.cfg.App.Name,
		Status: "OK",
	})
}
