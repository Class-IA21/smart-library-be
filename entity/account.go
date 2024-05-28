package entity

type AccountRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	NPM      string `json:"npm" validate:"required_if=Level student"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Level    string `json:"level" validate:"required"`
}

type AccountChangePasswordRequest struct {
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type Account struct {
	ID    int    `json:"id"`
	Email string `json:"email" validate:"required"`
	Level string `json:"level" validate:"required"`
}

type AccountResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email" validate:"required"`
	Level string `json:"level" validate:"required"`
}
