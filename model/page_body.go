package model

type PageBody[T any] struct {
	Page  int `json:"page" binding:"required"`
	Limit int `json:"limit" binding:"required"`
	Param T   `json:"params"`
}
