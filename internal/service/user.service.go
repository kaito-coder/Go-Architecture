package service

import "github.com/kaito-coder/go-ecommerce-architecture/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserById(uid string) string {
	return us.userRepo.GetUserById(uid)
}

