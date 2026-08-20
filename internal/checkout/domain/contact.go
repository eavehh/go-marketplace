package domain

type Contact struct {
	First_name string `json:"first_name" validate:"required,max=100"`
	Last_name  string `json:"Last_name" validate:"required,max=100"`
	Email      string `json:"email" validate:"required, email, max=255"`
}
