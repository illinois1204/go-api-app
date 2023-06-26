package dto

type UserCreate struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}
