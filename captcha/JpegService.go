package captcha

/*
 *@author: 随风飘的云
 *@description: 创建位点格式验证码
 *@date: 2023-10-02 17:01
 */
type JpegCaptcha struct {
	*Captcha
}

func (j *JpegCaptcha) CreateCaptcha(context string)(string, string, string, string, error)  {
	return "", "", "", "", nil
}
