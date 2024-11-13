package models

import "errors"

var (
	ErrLinkNotFound = errors.New("link not found")
)

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
