package repo

import (
	"github.com/kaito-coder/go-ecommerce-architecture/global"
	"github.com/kaito-coder/go-ecommerce-architecture/internal/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}
type UserRepository struct {
}

// GetUserByEmail implements IUserRepository.
func (u *UserRepository) GetUserByEmail(email string) bool {
	row := global.Mdb.Table(TableNameCrmUser).Where("usr_email = ?", email).First(&model.CrmUser{}).RowsAffected
	return row != NumberNull
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}
