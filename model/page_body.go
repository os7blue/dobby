package model

type PageBody[T any] struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Param T   `json:"param"`
}
