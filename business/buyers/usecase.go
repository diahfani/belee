//logic business
//usecase bisa mengakses ke database/drivers
//melalui domain yang dihubungkan oleh interface

package buyers

import (
	"belee/app/middleware"
	"context"
	"final_project/belee/business"
	"strings"
	"time"
)

type BuyerUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJwt
}

//responsenya adalah interface usecase
//interface usecase akan dipasangkan dgn controllers
func NewBuyerUsecase(repo Repository, jwtAuth *middleware.ConfigJwt, timeout time.Duration) Usecase {
	return &BuyerUsecase{
		Repo:           repo,
		jwtAuth:        jwtAuth,
		contextTimeout: timeout,
	}
}

//core business login
//validasi ada di controllers (disarankan)
//untuk handlers (bagian depan) itu untuk mem-binding

func (uc *BuyerUsecase) Login(ctx context.Context, buyerDomain Domain) (Domain, string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if buyerDomain.Email == "" {
		return Domain{}, "", business.ErrPhonePasswordNotFound
	}

	if buyerDomain.Password == "" {
		return Domain{}, "", business.ErrPhonePasswordNotFound
	}

	buyer, err := uc.Repo.Login(ctx, buyerDomain.Email, buyerDomain.Password)

	if err != nil {
		return Domain{}, "", business.ErrInvalidCredential
	}

	token, err := uc.jwtAuth.GenerateToken(buyer.Id)
	if err != nil {
		return Domain{}, "", err
	}
	return buyer, token, nil

}

func (uc *BuyerUsecase) Register(ctx context.Context, data Domain) (Domain, string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedBuyer, err := uc.Repo.GetByEmail(ctx, data.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, "", err
		}
	}
	if existedBuyer != (Domain{}) {
		return Domain{}, "", business.ErrDuplicateData
	}

	buyer, err := uc.Repo.Store(ctx, data)
	if err != nil {
		return Domain{}, "", err
	}

	// token, err := uc.jwtAuth.GenerateToken(buyer.Id)
	// if err != nil {
	// 	return Domain{}, err
	// }

	return buyer, "", nil

	// err = uc.Repo.Register(ctx, buyerDomain)
	// if err != nil {
	// 	return err
	// }

	// return nil
}
