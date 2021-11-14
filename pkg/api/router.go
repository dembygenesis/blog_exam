package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func getBadRequestHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		status := http.StatusBadRequest
		return c.Status(status).JSON(Response{
			Status:  status,
			Message: "Invalid request",
			Data:    nil,
		})
	}
}

// SetRoutes configures the multiplexer
func (s *Server) SetRoutes() {

	// =======================
	// Articles
	// =======================
	api := s.app.Group("/articles")

	api.Get("/:article_id", s.ReadArticles)
	api.Get("/", s.ReadArticles)
	api.Post("/", s.AddArticle)

	// Unidentified handler
	s.app.Get("*", getBadRequestHandler())
}
