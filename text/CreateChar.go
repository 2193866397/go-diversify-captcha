package text

import (
	"strings"
	"go-diversify-captcha/constants"
	"go-diversify-captcha/utils"
)

/*
 *@author: 随风飘的云
 *@description: 字符验证码生成器
 *@date: 2023-10-02 21:32
 */
type textChar struct{}

func (t *textChar) GetTextContent(contextType string, length int) string {
	if strings.EqualFold(contextType, constants.CHAR) {
		return utils.GetRandomChar(length)
	}else if strings.EqualFold(contextType, constants.UPPER_CHAR) {
		return utils.GetRandomUpperChar(length)
	}
	return ""
}