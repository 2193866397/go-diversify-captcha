package text

import (
	"bytes"
	"go-diversify-captcha/constants"
	"go-diversify-captcha/utils"
	"math"
	"strconv"
	"strings"
)

/*
 *@author: 随风飘的云
 *@description: 生成算术验证码
 *@date: 2023-12-11 20:48
 */
type textArithmetic struct{}

/**
 * @description: 获取算术验证码生成器
 * @param {string} contextType
 * @param {int} length
 * @return {*}
 */
func (t *textArithmetic) GetTextContent(contextType string, length int) string {
	if strings.EqualFold(contextType, constants.ARITHMETIC) {
		return getArithmetic(length)
	} else if strings.EqualFold(contextType, constants.ARITHMETIC_ZH) {
		return getArithmeticChinese(length)
	}
	return ""
}

// 阿拉伯数字算术
func getArithmetic(length int) string {
	len1 := utils.GetRandBetween(1, length/2+1)
	len2 := length - len1
	str1 := utils.GetRandomNumber(len1)
	str2 := utils.GetRandomNumber(len2)
	x, _ := strconv.Atoi(str1)
	y, _ := strconv.Atoi(str2)
	operation := int(math.Ceil(utils.GetRandomFloat64() * 2))
	return splicing(operation, x, y, 0, str1, str2)
}

// 简体中文数字算术
func getArithmeticChinese(length int) string {
	len1 := utils.GetRandBetween(1, length/2+1)
	len2 := length - len1
	x := 0
	y := 0
	operation := int(math.Ceil(utils.GetRandomFloat64() * 2))
	var buffer1, buffer2 bytes.Buffer
	for i := 0; i < len1; i++ {
		temp := utils.GetRandomInt(len(constants.BASE_COMPLEX_NUMBER))
		if temp == 0 && i == 0 && len1 > 1 {
			continue
		}
		x = x*10 + temp
		buffer1.WriteString(constants.BASE_COMPLEX_NUMBER[temp])
	}
	for i := 0; i < len2; i++ {
		temp := utils.GetRandomInt(len(constants.BASE_COMPLEX_NUMBER))
		if temp == 0 && i == 0 && len1 > 1 {
			continue
		}
		y = y*10 + temp
		buffer2.WriteString(constants.BASE_COMPLEX_NUMBER[temp])
	}
	return splicing(operation, x, y, 1, strings.TrimLeft(buffer1.String(), "零"), strings.TrimLeft(buffer2.String(), "零"))
}

// 分隔获取算术内容
func splicing(operation, x, y, textType int, str1, str2 string) string {
	var buff bytes.Buffer
	result := 0
	if operation == 0 {
		result = x * y
		buff.WriteString(str1)
		if textType == 0 {
			buff.WriteString("*")
		} else {
			buff.WriteString("乘")
		}
		buff.WriteString(str2)
	} else if operation == 1 {
		if x != 0 && y%x == 0 {
			result = y / x
			buff.WriteString(str2)
			if textType == 0 {
				buff.WriteString("/")
			} else {
				buff.WriteString("除以")
			}
			buff.WriteString(str1)
		} else {
			result = x + y
			buff.WriteString(str1)
			if textType == 0 {
				buff.WriteString("+")
			} else {
				buff.WriteString("加")
			}
			buff.WriteString(str2)
		}
	} else if operation == 2 {
		if x >= y {
			result = x - y
			buff.WriteString(str1)
			if textType == 0 {
				buff.WriteString("-")
			} else {
				buff.WriteString("减")
			}
			buff.WriteString(str2)
		} else {
			result = y - x
			buff.WriteString(str2)
			if textType == 0 {
				buff.WriteString("-")
			} else {
				buff.WriteString("减")
			}
			buff.WriteString(str1)
		}
	} else {
		result = x + y
		buff.WriteString(str1)
		if textType == 0 {
			buff.WriteString("+")
		} else {
			buff.WriteString("加")
		}
		buff.WriteString(str2)
	}
	buff.WriteString("=?@" + strconv.Itoa(result))
	return buff.String()
}
