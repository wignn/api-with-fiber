package service

import (
	"context"

	"github.com/wignn/api-with-fiber/domain"
	"github.com/wignn/api-with-fiber/dto"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (c customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData
	for _, v:= range customers {
		customerData =append(customerData, dto.CustomerData{
			ID: v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}
	return customerData, nil
}
