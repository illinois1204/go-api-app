package authDto

type UserCreate struct {
	Name     string `json:"name" validate:"omitempty"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}
