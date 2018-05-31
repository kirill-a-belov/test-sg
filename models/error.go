package models

const (
	ErrUnknownRequestCode = iota
	ErrUnknownErrorCode
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}
