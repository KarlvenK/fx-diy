package server

import (
	"github.com/gofiber/fiber/v2"
)

func AddHandler(
	srv *Server,
) {
	v1 := srv.HttpS.App.Group("/v1/")
	v1.Get("/hello", Hello)
}

func Hello(c *fiber.Ctx) error {
	type resStruct struct {
		Message string
		Number  int
	}

	data := resStruct{
		Message: "hello from server!",
		Number:  10086,
	}
	return c.JSON(data)
}
