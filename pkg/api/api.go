package api

import (
	"fmt"
	"github.com/dembygenesis/blog_exam/pkg/logic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// Response holds our response structure
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Server holds our app variables
type Server struct {
	logic logic.Logic
	port  int
	app   *fiber.App
}

// NewServer starts a new HTTP server
func NewServer(logic logic.Logic, port int) *Server {
	s := &Server{
		logic: logic,
		port:  port,
		app: fiber.New(fiber.Config{
			BodyLimit: 20971520,
		}),
	}
	s.SetRoutes()
	return s
}

// Start initialize a new fiber instance
func (s *Server) Start() error {
	// ======================================================
	// Graceful shutdown & recover
	s.app.Use(recover.New())
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdown
		fmt.Println("Gracefully shutting down")
		err := s.app.Shutdown()
		if err != nil {
			fmt.Println("Shutting down error", err)
		}
	}()

	// Start listener
	return s.app.Listen(":" + strconv.Itoa(s.port))
}
