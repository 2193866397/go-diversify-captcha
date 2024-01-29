package utils

import (
	"go-diversify-captcha/constants"
	"image/color"
	"math"
	"strings"
)
/*
 *@author: 随风飘的云
 *@description: 验证码内容随机工具
 *@date: 2023-10-02 20:54
 */

// 获取len长度的中文汉字
func GetRandomChinese(len int) string {
	return RandomString(constants.MODERN_CHINESE, len)
}

// 随机获取一个成语
func GetRandomIdiom() string {
	return constants.CHINESE_SIMPLE_IDIOM[GetRandomInt(len(constants.CHINESE_SIMPLE_IDIOM))]
}

// 随机获取一个成语的排列顺序
func GetRandomIdiomSort() []int {
	return constants.IDIOM_ID[GetRandomInt(len(constants.IDIOM_ID))]
}

// 获取一个字符串，只包含数字。
func GetRandomNumber(len int) string {
	str := RandomByteToString(constants.BASE_NUMBER, len)
	return strings.TrimLeft(str, "0")
}

// 随机获取0~9的中文简体数字
func GetRandomComplexNumber(len int) string {
	return RandomString(constants.BASE_COMPLEX_NUMBER, len)
}

// 获取一个字符串，只包含字母。
func GetRandomChar(len int) string {
	return RandomByteToString(constants.BASE_CHAR, len)
}

// 获取一个字符串，只包含字母，大写
func GetRandomUpperChar(len int) string {
	return RandomByteToString(constants.BASE_UPPER_CHAR, len)
}

// 获取一个随机字符串（只包含数字和字符）
func GetRandomString(len int) string {
	return RandomByteToString(constants.BASE_NUMBER_CHAR, len)
}

// 获取一个随机字符串（只包含数字和大小写字符）
func GetRandomStringUpper(len int) string {
	return RandomByteToString(constants.UPPER_NUMBER_CHAR, len)
}

// 获取一个随机颜色
func GetRandomColor() color.RGBA {
	c := constants.COLORS_ID[GetRandomInt(len(constants.COLORS_ID))]
	red := c[0]
	green := c[1]
	blue := c[2]
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

// 随机生成深色系.
func RandDeepColor() color.RGBA {
	randColor := GetRandomColor()
	increase := float64(30 + GetRandomInt(255))
	red := math.Abs(math.Min(float64(randColor.R)-increase, 255))
	green := math.Abs(math.Min(float64(randColor.G)-increase, 255))
	blue := math.Abs(math.Min(float64(randColor.B)-increase, 255))

	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

// 随机生成浅色.
func RandLightColor() color.RGBA {
	red := GetRandomInt(55) + 200
	green := GetRandomInt(55) + 200
	blue := GetRandomInt(55) + 200
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

// 获取byte类型的随机内容
func RandomByteToString(base []byte, length int) string {
	baseLen := len(base)
	if baseLen == 0 {
		return ""
	}
	buffer := make([]byte, 0, length)
	if length < 1 {
		length = 1
	}
	for i := 0; i < length; i++ {
		index := GetRandomInt(baseLen)
		buffer = append(buffer, base[index])
	}
	return string(buffer)
}

// 获取string类型的随机内容
func RandomString(base []string, length int) string {
	baseLen := len(base)
	if baseLen == 0 {
		return ""
	}
	buffer := make([]byte, 0, length)
	if length < 1 {
		length = 1
	}
	for i := 0; i < length; i++ {
		index := GetRandomInt(baseLen)
		str := base[index]
		buffer = append(buffer, str...)
	}
	return string(buffer)
}
