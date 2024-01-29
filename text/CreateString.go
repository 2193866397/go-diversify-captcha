package text

import (
	"strings"
	"go-diversify-captcha/constants"
	"go-diversify-captcha/utils"
)

/*
 *@author: 随风飘的云
 *@description: 混合验证码生成器
 *@date: 2023-10-02 21:34
 */
type textString struct{}

func (t *textString) GetTextContent(contextType string, length int) string {
	if strings.EqualFold(contextType, constants.STRING) {
		return utils.GetRandomString(length)
	}else if strings.EqualFold(contextType, constants.UPPER_STRING) {
		return utils.GetRandomStringUpper(length)
	}
	return ""
}
