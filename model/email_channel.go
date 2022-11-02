package model

type EmailChannel struct {
	Host           string `json:"host" binding:"not null,uri"`
	Port           string `json:"port" binding:"not null,min=2"`
	Username       string `json:"username" binding:"not null,email"`
	Password       string `json:"password" binding:"not null"`
	ToEmailListStr string `json:"toEmailListStr" binding:"not null"`
}
