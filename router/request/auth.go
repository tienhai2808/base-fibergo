package request

type TestRequest struct {
	Request string `json:"request"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Email string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3"`
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName string `json:"last_name" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=6"`
}