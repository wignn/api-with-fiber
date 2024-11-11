package api

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wignn/api-with-fiber/domain"
)

type customersAPi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService){
	ca := customersAPi{
		customerService: customerService,
	}
}


func (ca customersAPi) Index(ctx *fiber.Ctx)error{
	c, cencel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cencel()
	res, err := ca.customerService.Index(c)
	if err != nil{
		return ctx.Status(500).JSON("internal server eoror")
	}
}