package web

type ModelError struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
