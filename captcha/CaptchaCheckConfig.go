package captcha

import (
	"errors"
	"image/color"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
 *@author: 随风飘的云
 *@description: 验证码检验配置
 *@date: 2023-10-29 16:11
 */

type CaptchaCheckConfig struct{}

// 检查用户设置的类型
func CaptchaCheckType(parName, parValue, defaultType, checkType string) (string, error) {
	var isType string
	var isCheck bool
	if parValue == "" || len(parValue) == 0 {
		isType = defaultType
	} else {
		parValue = strings.ToUpper(parValue)
		val := strings.Split(checkType, ",")
		length := len(val)
		for i := 0; i < length; i++ {
			if strings.EqualFold(parValue, val[i]) {
				isType = val[i]
				isCheck = true
				break
			}
		}
		if !isCheck {
			return "", errors.New("The content format set by the user does not meet the requirements, please reset" + parName + " and " + parValue)
		}
	}
	return isType, nil
}

/**
 * 创建颜色
 */
func CaptchaGetColor(parName, parValue string, c color.Color) color.Color {
	var co color.Color
	if parValue == "" || len(parValue) == 0 {
		co = c
	} else {
		co, _ = CaptchaCreateColor(parName, parValue)
	}
	return co
}

/**
 * 创建颜色
 */
func CaptchaCreateColor(parName, parValue string) (color.Color, error) {
	var co color.Color
	if strings.IndexAny(parValue, ",") > 0 {
		par := strings.ReplaceAll(parValue, " ", "")
		rgb := strings.Split(par, ",")
		r, _ := strconv.Atoi(rgb[0])
		g, _ := strconv.Atoi(rgb[1])
		b, _ := strconv.Atoi(rgb[2])
		if len(rgb) == 4 {
			a, _ := strconv.Atoi(rgb[3])
			co = &color.RGBA64{
				R: uint16(r),
				G: uint16(g),
				B: uint16(b),
				A: uint16(a),
			}
			return co, nil
		} else if len(rgb) == 3 {
			co = &color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
			}
			return co, nil
		} else {
			return co, errors.New(parName + ": The color need rgb which can only have 3(RGB) or 4 (RGB with Alpha) values")
		}
	}
	return nil, errors.New(parName + ": Only supports color conversion in 0-255 format, temporarily does not support hexadecimal color conversion")
}

// 设置长度，数字，比如说验证码的宽度，长度，字符长度，干扰线长度
func SetPositiveInt(parName, parValue string, defaultValue int) (int, error) {
	var intValue int
	if parValue == "" || len(parValue) == 0 {
		intValue = defaultValue
	} else {
		intValue, _ = strconv.Atoi(parValue)
		if intValue < 0 {
			return intValue, errors.New(parName + ": The metric you set is not appropriate, please set it again! ")
		}
	}
	return intValue, nil
}

// 设置检查类型
func SetType(parName, parValue, defaultType, checkType string) (string, error) {
	var typeValue string
	var isCheck bool
	if parValue == "" || len(parValue) == 0 {
		typeValue = defaultType
	} else {
		parValue := strings.ToUpper(parValue)
		value := strings.Split(checkType, ",")
		for _, val := range value {
			if strings.EqualFold(parValue, parValue) {
				typeValue = val
				isCheck = true
				break
			}
		}
		if !isCheck {
			return "", errors.New(parName + ": The content format set by the user does not meet the requirements, please reset again! ")
		}
	}
	return typeValue, nil
}

/**
 * @description: 设置检查是否需要
 * @param {*} parName 参数键
 * @param {string} parValue 参数值
 * @param {bool} defaultValue 默认值
 * @return {*}
 */
func SetCheckBoolean(parName, parValue string, defaultValue bool) (bool, error) {
	var value bool
	if parValue == "" || len(parValue) == 0 {
		value = defaultValue
	} else {
		parValue = strings.ToLower(parValue)
		if strings.EqualFold(parValue, "yes") {
			value = defaultValue
		} else if strings.EqualFold(parValue, "no") {
			value = false
		} else {
			return false, errors.New(parName + ": Must be set to YES or NO and not case sensitive")
		}
	}
	return value, nil
}

/**
 * @description: 中文长度
 * @param {string} str
 * @return int
 */
func LenChineseChar(str string) int {
	return utf8.RuneCountInString(str)
}

/**
 * @description: 把十六进制颜色转 color.RGBA
 * @param {string} co
 * @return {*}
 */
func ParseHexColor(co string) (c color.RGBA, err error) {
	c.A = 0xFF
	if co[0] != '#' {
		return c, errors.New("hex color must start with '#'")
	}
	if len(co) == 7 {
		c.R = HexToByte(co[1])<<4 + HexToByte(co[2])
		c.G = HexToByte(co[3])<<4 + HexToByte(co[4])
		c.B = HexToByte(co[5])<<4 + HexToByte(co[6])
	} else if len(co) == 4 {
		c.R = HexToByte(co[1]) * 17
		c.G = HexToByte(co[2]) * 17
		c.B = HexToByte(co[3]) * 17
	} else {
		err = errors.New("color value component is invalid")
	}
	return c, err
}

/**
 * @description: 转换byte
 * @param {byte} b
 * @return {*}
 */
func HexToByte(b byte) byte {
	switch {
	case b >= '0' && b <= '9':
		return b - '0'
	case b >= 'a' && b <= 'f':
		return b - 'a' + 10
	case b >= 'A' && b <= 'F':
		return b - 'A' + 10
	}
	return 0
}
