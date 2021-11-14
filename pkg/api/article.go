package api

import "github.com/gofiber/fiber/v2"

func (s *Server) AddArticle(c *fiber.Ctx) error {
	return c.SendString("AddArticle")
}

func (s *Server) ReadArticles(c *fiber.Ctx) error {
	s.logic.Read(3)
	return c.SendString("ReadArticles")
}