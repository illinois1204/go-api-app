package dto

type UserCreate struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}
