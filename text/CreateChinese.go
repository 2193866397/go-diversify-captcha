package text

import (
	"strings"
	"go-diversify-captcha/constants"
	"go-diversify-captcha/utils"
)

/*
 *@author: 随风飘的云
 *@description: 中文验证码生成器
 *@date: 2023-10-02 21:33
 */
type textChinese struct{}

func (t *textChinese) GetTextContent(contextType string, length int) string {
	if strings.EqualFold(contextType, constants.CHINESE) {
		return utils.GetRandomChinese(length)
	}else if strings.EqualFold(contextType, constants.CHINESE_IDIOM) {
		str := utils.GetRandomIdiom()
		num := utils.GetRandomIdiomSort()
		b := make([]byte, 0, 2*len(str) + 1)
		b = append(b, str...)
		// 作为成语区分
		b = append(b, ',')
		for i := 0; i < len(str); i++ {
			b = append(b, str[num[i]])
		}
		return string(b)
	}
	return ""
}