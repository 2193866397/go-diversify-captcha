package _interface

type TextFactory interface {
	GetTextContent(contextType string, length int) string
}

type CaptchaFactory interface {
	CreateCaptcha(context string) (string, string, string, string, error)
}