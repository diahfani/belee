//logic business
//usecase bisa mengakses ke database/drivers
//melalui domain yang dihubungkan oleh interface

package buyers

import (
	"context"
	"errors"
	"time"
)

type BuyerUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

//responsenya adalah interface usecase
//interface usecase akan dipasangkan dgn controllers
func NewBuyerUsecase(repo Repository, timeOut time.Duration) Usecase {
	return &BuyerUsecase{
		Repo:           repo,
		contextTimeout: timeOut,
	}
}

//core business login
//validasi ada di controllers (disarankan)
//untuk handlers (bagian depan) itu untuk mem-binding

func (uc *BuyerUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("Email Empty")
	}

	if password == "" {
		return Domain{}, errors.New("Password Empty")
	}

	buyer, err := uc.Login(ctx, email, password)

	if err != nil {
		return Domain{}, err
	}

	return buyer, nil
}
