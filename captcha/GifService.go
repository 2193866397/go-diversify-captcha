package captcha

/*
 *@author: 随风飘的云
 *@description: 创建Gif格式验证码
 *@date: 2023-10-02 17:02
 */
type GifCaptcha struct {
	*Captcha
}
func (g *GifCaptcha) CreateCaptcha(context string) (string, string, string, string, error)  {
	return "", "", "", "", nil
}