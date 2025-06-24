package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type CreateBuyerRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

type UpdateBuyerRequest struct {
	ID        string `json:"id" binding:"required,uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type SettingsRequest struct {
	Currency string `json:"currency" binding:"required"`
	Language string `json:"language" binding:"required"`
}
