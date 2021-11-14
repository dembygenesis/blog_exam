package api

import (
	"fmt"
	"github.com/dembygenesis/blog_exam/pkg/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func (s *Server) AddArticle(c *fiber.Ctx) error {
	var body models.Article

	// Parse body
	err := c.BodyParser(&body)
	if err != nil {
		status := http.StatusInternalServerError
		return c.Status(status).JSON(Response{
			Status:  status,
			Message: fmt.Sprintf("error trying to parse the body: %v", err.Error()),
			Data:    nil,
		})
	}

	// Validate inputs
	errs := body.Validate()
	if len(errs) > 0 {
		status := http.StatusBadRequest
		return c.Status(status).JSON(Response{
			Status:  status,
			Message: "errors on the body's inputs",
			Data:    errs,
		})
	}



	return c.JSON(body)
}

// ReadArticles api handler to for article data
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
