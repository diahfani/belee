package paymentMethod

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Add(ctx context.Context, data *Domain) (Domain, error)
	FindAll(ctx context.Context) ([]Domain, error)
	FindById(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Find(ctx context.Context) ([]Domain, error)
	Add(ctx context.Context, data *Domain) (Domain, error)
	FindById(id int) (Domain, error)
}
