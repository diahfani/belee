// struct or model business
package buyers

import (
	"context"
	"time"
)

//for business not related with database, response, request, etc
type Domain struct {
	Id        int
	Name      string
	Age       string
	NoHp      string
	Dob       string
	Address   string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//interface adalah list2 function yang menempel di sebuah struct yang akan digunakan usecase untuk domain

type Usecase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
}

//interface yang digunakan untuk mengakses database/drivers
type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
}
