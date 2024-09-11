package repo

type IUserRepository interface {
	GetUserByEmail(email string) bool
}
type UserRepository struct {
}

// GetUserByEmail implements IUserRepository.
func (u *UserRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}
