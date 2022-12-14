package model

type Result struct {
	Code  int    `json:"code" default:"0"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Count int64  `json:"count" default:"0"`
}
