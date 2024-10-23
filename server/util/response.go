package util

import (
	"server/domain/dto"
)

func BuildResponse[T any](status int, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		StatusCode: status,
		Message: message,
		Data: data,
	}
}