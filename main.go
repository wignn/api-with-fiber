package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wignn/api-with-fiber/internal/config"
	"github.com/wignn/api-with-fiber/internal/connection"
	"github.com/wignn/api-with-fiber/internal/repository"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	CustomerRepository := repository.NewCustomer(dbConnection)
	
	app.Get("/developers", developers) 
	app.Get("/", Work)
	fmt.Println("Server running at  http://127.0.0.1:3000") 
	app.Listen(":3000")
}

func Work(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("api Working")
}

func developers(ctx *fiber.Ctx) error { 
	return ctx.Status(200).JSON("data")
}
