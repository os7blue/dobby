package model

type EmailChannel struct {
	Host           string `json:"host" validate:"required"`
	Port           string `json:"port" validate:"required,min=2"`
	Username       string `json:"username" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
	ToEmailListStr string `json:"toEmailListStr" validate:"required"`
}
