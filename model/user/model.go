package UserModels

type RegisterInput struct {
	Email           string `json:"email"`
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
