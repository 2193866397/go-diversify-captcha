package captcha

type SlideCaptcha struct {
	*Captcha
}

func (s *SlideCaptcha) CreateCaptcha(context string) (string, string, string, string, error) {
	return "", "", "", "", nil
}
