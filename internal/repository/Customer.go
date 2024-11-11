package repository

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/wignn/api-with-fiber/domain"
)

type CustomerRepository struct {
	db *goqu.Database
}



func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &CustomerRepository{
		db: goqu.New("default", con),
	}
}

func (cr *CustomerRepository) Save(ctx context.Context, c *domain.Customer) error {
	executor := cr.db.Insert("customers").Rows(c).Executor()
	_,err := executor.ExecContext(ctx)
	return err
}

func (cr *CustomerRepository) FindAll(ctx context.Context) (result []domain.Customer, err error) {
	dataset := cr.db.From("customers").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (cr *CustomerRepository) FindById(ctx context.Context, id string) (result domain.Customer,err error) {
	dataset := cr.db.From("customers").Where(goqu.C("delete_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (cr *CustomerRepository) Create(ctx context.Context, c *domain.Customer) error {
	panic("not implemented")
}

func (cr *CustomerRepository) Update(ctx context.Context, c *domain.Customer) error {
	panic("not implemented")
}

func (cr *CustomerRepository) Delete(ctx context.Context, id string) error {
	panic("not implemented")
}
