package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kaito-coder/go-ecommerce-architecture/internal/repo"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/utils"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/utils/crypto"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/utils/sendto"
	"github.com/kaito-coder/go-ecommerce-architecture/pkg/response"
)

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetUserById(uid string) string {
// 	return us.userRepo.GetUserById(uid)
// }

type IUserService interface {
	Register(email string, purpose string) int
}
type UserService struct {
	userRepo repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository	
}

// Register implements IUserService.
func (us *UserService) Register(email string, purpose string) int {
	// TODO: Implement the logic of user registration.
	// 1. hash email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail: %s", hashEmail)
	// 2. check if email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	// 3. new OTP -> ..
	otp := utils.GenerateSixDigitOtp()
	if purpose == "TEST"{
		otp = 123456
	}
	// 4. save OTP to redis
	err := us.userAuthRepo.AddOTP(email, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrOTPInvalid
	}
	// 5. send OTP to email
	err = sendto.SendEmailOTP([]string{email}, "kaitok57a01@gmail.com",strconv.Itoa(otp))
	if err != nil {
		return response.ErrorSendEmailFail
	}
	// 6. check OTP is available
	// 7. user spam ...
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
		userAuthRepo: userAuthRepo,
	}
}
