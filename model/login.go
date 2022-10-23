package model

type LoginValidator struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required,min=6,max=6"`
}

type SendCodeValidator struct {
	Email string `json:"email" binding:"required,email"`
}
