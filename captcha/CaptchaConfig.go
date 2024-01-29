package captcha

import (
	"fmt"
	const_captcha "go-diversify-captcha/constants"
	_interface "go-diversify-captcha/interface"
	"image/color"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
)

/*
 *@author: 随风飘的云
 *@description: 验证码配置
 *@date: 2023-10-02 21:29
 */
// 长度大小
type CodeLen struct {
	Min, Max int
}

// 字体大小
type FontSizes struct {
	Width, Height int
}

// 定义点击图片的位置
type AreaPoint struct {
	MinX, MaxX, MinY, MaxY int
}

// 验证码的配置
type CaptchaConfig struct {
	// 是否设置边框
	isBorder bool
	// 边框颜色
	borderColor color.Color
	// 边框厚度
	borderThickness int
	// 干扰圆数量
	noiseNumber int
	// 干扰线数量
	shadeNumber int
	// 干扰线厚度
	shadeThickNess int
	// 随机直线数量
	lineNumber int
	// 是否设置字体
	isFont bool
	// 默认字体名称
	fontName     string
	fontsArray   []*truetype.Font
	fontsStorage _interface.FontsStorage
	// 字体大小
	fontSize FontSizes
	// 字体大小范围
	fontCodeLen CodeLen
	// 字体厚度
	fontThickness int
	// 字体颜色
	fontColor string
	// 文字间的间距
	padding int
	// 背景颜色
	backColor string
	// 添加扭曲效果
	shearAdd bool
	// 验证码文本扭曲程度
	imageFontDistort int
	// 验证码文本透明度 0-1
	imageFontAlpha float64
	// 图片宽度
	imageWidth int
	// 图片高度
	imageHeight int
	// 内容类型
	textType string
	// 内容长度
	textLength int
	// 图片类型
	imageType string
	// 随机缩略背景图		格式：图片绝对路径字符串
	rangBackground []string
	// 缩略图文字扭曲程度，值为 Distort...,
	fontDistort int
}

var captchaConfig CaptchaConfig

func init() {
	captchaConfig.isBorder = false
}

func GetCaptchaConfig() *CaptchaConfig {
	return &CaptchaConfig{
		isBorder: true,
		// 边框厚度
		borderThickness: 5,
		// 干扰圆数量
		noiseNumber: 15,
		// 干扰线数量
		shadeNumber: 4,
		// 干扰线厚度
		shadeThickNess: 1,
		lineNumber: 5,
		// 设置自定义字体
		isFont:   false,
		fontName: "./assets/fonts/actionj.ttf",
		// 背景颜色
		backColor: "0xffff",
		// 添加扭曲效果
		shearAdd: false,
		// 验证码文本扭曲程度
		imageFontDistort: 0,
		// 验证码文本透明度 0-1
		imageFontAlpha: 0,
		// 图片宽度
		imageWidth: 200,
		// 图片高度
		imageHeight: 100,
		textType:    const_captcha.STRING,
		textLength:  7,
		imageType:   const_captcha.CAPTCHA_TYPE_PNG,
		// 字的间距
		padding: 10,
		fontCodeLen: CodeLen{25, 30},
	}
}

/**
 * @description: 设置内容类型
 * @param {string} textType
 * @return {*}
 */
func (c *CaptchaConfig) SetContentType(textType string) {
	c.textType, _ = CaptchaCheckType(const_captcha.CAPTCHA_TEXT_TYPE, textType, const_captcha.STRING, const_captcha.TEXT_VALUE)
}

/**
 * @description: 设置图片类型
 * @param {string} imageType
 * @return {*}
 */
func (c *CaptchaConfig) SetImageType(imageType string) {
	c.imageType, _ = CaptchaCheckType(const_captcha.CAPTCHA_IMAGE_TYPE, imageType, const_captcha.CAPTCHA_TYPE_PNG, const_captcha.TYPE_VALUE)
}

/**
 * @description: 设置是否需要边框
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetIsBorder(parValue string) {
	c.isBorder, _ = SetCheckBoolean(const_captcha.CAPTCHA_BORDER, parValue, true)
}

/**
 * @description: 设置边框颜色（未测试，需要定位接口和结构体的区别）
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetBorderColor(parValue string) {
	c.borderColor = CaptchaGetColor(const_captcha.CAPTCHA_BORDER_COLOR, parValue, color.Gray{})
}

/**
 * @description: 边框厚度
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetThickness(parValue string) {
	c.borderThickness, _ = SetPositiveInt(const_captcha.CAPTCHA_BORDER_THICKNESS, parValue, 1)
}

/**
 * @description: 干扰圆数量
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetNoiseNumber(parValue string) {
	c.noiseNumber, _ = SetPositiveInt(const_captcha.CAPTCHA_NOISE_NUMBER, parValue, 2)
}

/**
 * @description: 干扰线数量
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetShadeNumber(parValue string) {
	c.shadeNumber, _ = SetPositiveInt(const_captcha.CAPTCHA_SHADE_NUMBER, parValue, 1)
}

/**
 * @description: 干扰线厚度
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetShadeThickness(parValue string) {
	c.shadeThickNess, _ = SetPositiveInt(const_captcha.CAPTCHA_SHADE_THICKNESS, parValue, 1)
}

/**
 * @description: 设置扭曲效果
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetAddShear(parValue string) {
	c.shearAdd, _ = SetCheckBoolean(const_captcha.CAPTCHA_IMAGE_SHEAR_ADD, parValue, true)
}

/**
 * @description: 设置内容长度
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetTextLength(parValue string) {

}

func (c *CaptchaConfig) SetFontNames(fons []string) {

}

/**
 * @description: 根据路径设置字体
 * @param {string} path
 * @return {*}
 */
func (c *CaptchaConfig) SetFontByPath(path string) {
	fs := c.LoadFontByPath(path)
	tfs := []*truetype.Font{}
	for _, fff := range fs {
		tfs = append(tfs, fff)
	}
	if len(tfs) != 0 {
		c.fontsArray = tfs
	}
}

func (c *CaptchaConfig) SetFontSize(size int) {

}

func (c *CaptchaConfig) SetFontThickness(parValue string) {

}

func (c *CaptchaConfig) SetFontColor() {

}

func (c *CaptchaConfig) SetBackColor() {

}

func (c *CaptchaConfig) SetImageFontDistort() {

}

func (c *CaptchaConfig) SetImageFontAlpha() {

}

/**
 * @description: 设置图片宽度
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetImageWidth(parValue string) {
	c.imageWidth, _ = SetPositiveInt(const_captcha.CAPTCHA_IMAGE_WIDTH, parValue, 160)
}

/**
 * @description: 设置图片宽度
 * @param {string} parValue
 * @return {*}
 */
func (c *CaptchaConfig) SetImageHeight(parValue string) {
	c.imageHeight, _ = SetPositiveInt(const_captcha.CAPTCHA_IMAGE_HEIGHT, parValue, 60)
}

/**
 * @description: 根据名称读取字体
 * @param {string} name
 * @return {*}
 */
func (c *CaptchaConfig) LoadFontByName(name string) *truetype.Font {
	fontFile, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("failed to open font file: %v", err)
	}
	trueTypeFont, err := truetype.Parse(fontFile)
	if err != nil {
		log.Fatal("failed to Parse font file: %v", err)
	}
	return trueTypeFont
}

/**
 * @description: 根据路径读取所有的字体
 * @param {string} path
 * @return {*}
 */
func (c *CaptchaConfig) LoadFontByPath(path string) []*truetype.Font {
	fonts := make([]*truetype.Font, 0)
	fontPath, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("can not read font path: %v", err)
	}
	for _, fontEntiy := range fontPath {
		fontStr := fmt.Sprintf("%v", fontEntiy)
		f := c.LoadFontByName(fontStr)
		fonts = append(fonts, f)
	}
	return fonts
}

/**
 * @description: 根据字符串地址数组读取字体
 * @param {[]string} assetFontNames
 * @return {*}
 */
func (c *CaptchaConfig) LoadFontsByNames(assetFontNames []string) []*truetype.Font {
	fonts := make([]*truetype.Font, 0)
	for _, assetName := range assetFontNames {
		f := c.LoadFontByName(assetName)
		fonts = append(fonts, f)
	}
	return fonts
}
