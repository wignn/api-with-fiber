package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wignn/api-with-fiber/internal/config"
	"github.com/wignn/api-with-fiber/internal/connection"
	"github.com/wignn/api-with-fiber/internal/repository"
	"github.com/wignn/api-with-fiber/internal/service"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	CustomerRepository := repository.NewCustomer(dbConnection)
	
	customerService = service.NewCustomer(CustomerRepository)

	_= app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

func Work(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("api Working")
}

func developers(ctx *fiber.Ctx) error { 
	return ctx.Status(200).JSON("data")
}
