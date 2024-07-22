package models

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	Name     string `json:"name" required:"true"`
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" required:"true"`
	Email    string `json:"email" required:"true" gorm:"unique"`
	Password string `json:"password" required:"true"`
}

type UserRegisterResponse struct {
	Token string `json:"token"`
}
