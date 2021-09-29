package paymentMethod

import (
	"belee/business"
	"context"

	"time"
)

type PaymentUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewPaymentUsecase(repo Repository, timeout time.Duration) Usecase {
	return &PaymentUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (pu *PaymentUsecase) Add(ctx context.Context, data *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	result, err := pu.Repo.Add(ctx, data)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (pu *PaymentUsecase) FindAll(ctx context.Context) ([]Domain, error) {
	resp, err := pu.Repo.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, err
}

func (pu *PaymentUsecase) FindById(ctx context.Context, id int) (Domain, error) {
	if id <= 0 {
		return Domain{}, business.ErrIDNotFound
	}

	resp, err := pu.Repo.FindById(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}
