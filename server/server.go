package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Arismonx/MiniShop/config"
	_middlewareHandler "github.com/Arismonx/MiniShop/modules/middleware/middlewareHandler"
	_middlewareRepository "github.com/Arismonx/MiniShop/modules/middleware/middlewareRepository"
	_middlewareUsecase "github.com/Arismonx/MiniShop/modules/middleware/middlewareUsecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Create struct server
type (
	server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware _middlewareHandler.MiddlewareHandlerService
	}
)

// initialize middleware
func newMiddleware(cfg *config.Config) _middlewareHandler.MiddlewareHandlerService {
	repo := _middlewareRepository.NewMiddlewareRepository()
	usecase := _middlewareUsecase.NewmiddlewareUsecase(repo)
	return _middlewareHandler.NewMiddlewareHandler(cfg, usecase)
}

// function concurrency
func (s *server) gracefulShutdown(pctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start service: %s", s.cfg.App.Name)
	<-quit
	log.Printf("shutting service: %s", s.cfg.App.Name)

	// Force shutdown after 10 seconds
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	//Shutdown stops the server gracefully
	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// function Listening
func (s *server) httpListening() {
	//Server is down normally and abnormally
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
}

// Create function start server
func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	//initialize struct server
	s := &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: newMiddleware(cfg),
	}

	//Basic Middleware
	//Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error: Requese Timeout",
		Timeout:      30 * time.Second,
	}))

	//CORS
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	//Body limit
	s.app.Use(middleware.BodyLimit("10M"))

	//choose start service
	switch s.cfg.App.Name {
	case "auth":
	case "item":
	case "player":
	case "inventory":
	case "payment":
	}

	//Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s.app.Use(middleware.Logger())

	go s.gracefulShutdown(pctx, quit)

	//listening
	s.httpListening()
}
