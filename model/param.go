package model

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     int    `json:"code"`
}
