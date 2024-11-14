package models

import "errors"

var (
	ErrLinkNotFound = errors.New("link not found")
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
