package model

type EmailChannel struct {
	Host           string `json:"host" validate:"not null"`
	Port           string `json:"port" validate:"not null,min=2"`
	Username       string `json:"username" validate:"not null,email"`
	Password       string `json:"password" validate:"not null"`
	ToEmailListStr string `json:"toEmailListStr" validate:"not null"`
}
