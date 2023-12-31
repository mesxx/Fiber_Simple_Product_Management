package main

import (
	"fiber_simple_product_management_api/servers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	port := os.Getenv("PORT")

	server := servers.Server()
	server.Listen(":" + port)
}
