package service

import (
	"github.com/kaito-coder/go-ecommerce-architecture/internal/repo"
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
}

// Register implements IUserService.
func (us *UserService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}
