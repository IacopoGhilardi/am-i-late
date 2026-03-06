package _interface

type EmailServiceInterface interface {
	SendResetPasswordEmail(to, token string) error
}
