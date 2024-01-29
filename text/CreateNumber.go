package text

import (
	"strings"
	"go-diversify-captcha/constants"
	"go-diversify-captcha/utils"
)

/*
 *@author: 随风飘的云
 *@description: 数字验证码生成器
 *@date: 2023-10-02 21:33
 */
type textNumber struct{}

func (t *textNumber) GetTextContent(contextType string, length int) string {
	if strings.EqualFold(contextType, constants.NUMBER) {
		return utils.GetRandomNumber(length)
	}else if strings.EqualFold(contextType, constants.NUMBER_ZH) {
		return utils.GetRandomComplexNumber(length)
	}
	return ""
}