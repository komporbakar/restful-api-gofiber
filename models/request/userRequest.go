package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Phone    string `json:"phone" `
}

type UserUpdateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type EmailUpdateRequest struct {
	Email string `json:"email" validate:"required,email"`
}
