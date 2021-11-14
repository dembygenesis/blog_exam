package api

import "github.com/gofiber/fiber/v2"

func (s *Server) AddArticle(c *fiber.Ctx) error {
	return c.SendString("hello")
}

func (s *Server) ReadArticles(c *fiber.Ctx) error {
	return c.SendString("hello")
}