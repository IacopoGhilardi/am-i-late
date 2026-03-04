package commons

import "time"

type BaseResponse[T any] struct {
	Status  string `json:"status"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func Success[T any](data T) BaseResponse[T] {
	return BaseResponse[T]{Status: "OK", Data: data}
}

func Fail(err string) BaseResponse[any] {
	return BaseResponse[any]{Status: "KO", Error: err}
}

type AuthResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
