package repo

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUserById(uid string) string {
	return "pong" + uid
}