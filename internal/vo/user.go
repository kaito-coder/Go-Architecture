package vo

type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required,email"` 
	Purpose string `json:"purpose" binding:"required"`
}