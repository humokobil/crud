package web

type ModelError struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
}
