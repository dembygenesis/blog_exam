package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func (s *Server) AddArticle(c *fiber.Ctx) error {
	return c.SendString("AddArticle")
}

func (s *Server) ReadArticles(c *fiber.Ctx) error {
	var id int
	var err error
	status := http.StatusOK

	// Parse query
	queryId := c.Params("article_id")
	if queryId != "" {
		id, err = strconv.Atoi(queryId)
		if err != nil {
			status = http.StatusInternalServerError
			return c.Status(status).JSON(Response{
				Status:  status,
				Message: fmt.Sprintf("error trying to parse the query_id into int: %v", err.Error()),
				Data:    nil,
			})
		}
	}

	// Fetch articles
	data, err := s.logic.Read(id)
	if err != nil {
		status = http.StatusInternalServerError
		return c.Status(status).JSON(Response{
			Status:  status,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.Status(status).JSON(Response{
		Status:  status,
		Message: "Success",
		Data:    data,
	})
}