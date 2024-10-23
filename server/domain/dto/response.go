package dto

type ApiResponse[T any] struct {
 StatusCode     int `json:"statusCode"`
 Message 				string `json:"message"`
 Data           T      `json:"data"`
}