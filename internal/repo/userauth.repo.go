package repo

import (
	"fmt"
	"time"

	"github.com/kaito-coder/go-ecommerce-architecture/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type UserAuthRepository struct {
}

// AddOTP implements IUserAuthRepository.
func (u *UserAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("otp:%s", email)
	return global.Rdb.Set(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &UserAuthRepository{}
}
