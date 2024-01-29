package main

import (
	"fmt"
	"go-diversify-captcha/captcha"
)

func main() {
	c := captcha.NewCaptCha()
	text := c.GetTextContext()
	captchaFunc := c.GenerateCaptcha()
	fmt.Println(text)
	captchaFunc(text)
}
