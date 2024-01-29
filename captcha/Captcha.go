package captcha

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"go-diversify-captcha/constants"
	"go-diversify-captcha/text"
	"go-diversify-captcha/utils"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"log"
	"math"
	"sync"
	"unicode/utf8"
)

/*
 *@author: 随风飘的云
 *@description: 验证码图片汇总生成
 *@date: 2023-10-01 01:14
 */
type Captcha struct {
	context *text.FactoryContext
	config  *CaptchaConfig
	draw    *CaptchaDraw
}

var _instance *Captcha
var _once sync.Once

/**
 * @description: 创建验证码
 * @return {*}
 */
func initCaptCha() *Captcha {
	return &Captcha{
		context: &text.FactoryContext{},
		config:  GetCaptchaConfig(),
		draw: &CaptchaDraw{},
	}
}

/**
 * @description: 初始化验证码
 * @return {*}
 */
func NewCaptCha() *Captcha {
	_once.Do(func() {
		_instance = initCaptCha()
	})
	return _instance
}

/**
 * @description: 生成验证码方法汇总
 * @param {string} text
 * @return {*}
 */
func (c *Captcha) GenerateCaptcha() func(context string) (string, string, string, string, error) {
	cType := c.config.imageType
	p := &PngCaptcha{Captcha: c}
	j := &JpegCaptcha{Captcha: c}
	g := &GifCaptcha{Captcha: c}
	s := &SlideCaptcha{Captcha: c}
	factory := map[string]func(context string) (string, string, string, string, error){
		constants.CAPTCHA_TYPE_PNG:  p.CreateCaptcha,
		constants.CAPTCHA_TYPE_JPEG: j.CreateCaptcha,
		constants.CAPTCHA_TYPE_GIF:  g.CreateCaptcha,
		constants.CAPTCHA_TYPE_SLIDE : s.CreateCaptcha,
	}
	return factory[cType]
}

/**
 * @description: 获取验证码的文本信息
 * @param {string} contextType
 * @param {int} length
 * @return {*}
 */
func (c *Captcha) GetTextContext() string {
	return c.context.GetContextText(c.config.textType, c.config.textLength)
}

/**
 * @description: 获取底层绘制验证码的逻辑
 * @return {*}
 */
func (c *Captcha) GetDraw() *CaptchaDraw {
	return c.draw
}

/**
 * @description: 获取底层绘制验证码的配置
 * @return {*}
 */
func (c *Captcha) GetCaptchaConfig() *CaptchaConfig {
	return c.config
}

func (c *Captcha) CreateCaptchaToPng(context string) (*image.NRGBA, error)  {
	// 获取验证码配置
	width := c.config.imageWidth
	height := c.config.imageHeight
	bg, _:= ParseHexColor(c.config.backColor)

	// 创建验证码画板
	canvas := c.draw.CreateCanvasWithBg(width, height, bg)

	// 读取字体
	fontConfig := c.config.LoadFontByName(c.config.fontName)
	
	// 画干扰圆
	if c.config.noiseNumber > 0 {
		for i := 0; i < c.config.noiseNumber; i++ {
			x := utils.GetRandomInt(width)
			y := utils.GetRandomInt(height)
			radius := utils.GetRandomFloat64() * 10 + 5
			c.draw.drawCircle(canvas, x, y, int(radius), utils.GetRandomColor())
		}
	}
	//
	// 画干扰曲线
	if c.config.shadeNumber > 0 {
		for i := 0; i < c.config.shadeNumber; i++ {
			c.draw.drawRandomCurve(canvas, utils.GetRandomColor(), c.config.shadeThickNess)
		}
	}

	// 画随机直线
	//if c.config.lineNumber > 0 {
	//	c.draw.drawSlimLine(canvas, width, height, c.config.lineNumber)
	//}

	// 获取内容字体
	//dots := c.genDots(context, c.config.padding)
	// 画图片内容
	//for _, dot := range dots {
		err := c.DrawText(canvas, context, fontConfig, utils.RandDeepColor())
		if err != nil {
			log.Fatal("can not write context value: ", err)
		}
	//}

	//// 画扭曲效果
	if c.config.shearAdd {
		c.draw.warpX(canvas, utils.GetRandomFloat64() * 20.0)
		c.draw.warpY(canvas, utils.GetRandomFloat64() * 20.0)
		c.draw.distortImage(canvas, utils.GetRandomFloat64() * 20.0, utils.GetRandomFloat64() * 20.0)
	}


	// 画图片边框
	if c.config.isBorder {
		c.draw.drawBorder(canvas, utils.GetRandomColor(), c.config.borderThickness)
	}
	return canvas, nil
}

func (c *Captcha) CreateCaptchaToJpeg()  {
	//// 获取验证码配置
	//width := c.config.imageWidth
	//height := c.config.imageHeight
	//bg, _:= ParseHexColor(c.config.backColor)
	//
	//// 创建验证码画板
	//canvas := c.draw.CreateCanvasWithBg(width, height, bg)
	//// 获取内容字体
	//dots := c.genDots(context, c.config.padding)
	//
	//// 读取字体
	//fontConfig := c.config.LoadFontByName(c.config.fontName)
	//
	//for _, dot := range dots {
	//
	//}
}

func (c *Captcha) CreateCaptchaToSlide()  {

}

func (c *Captcha) DrawText(img *image.NRGBA, text string, fontConfig *truetype.Font, fc color.RGBA) (err error) {
	dc := freetype.NewContext()
	dc.SetDPI(float64(72))
	dc.SetClip(img.Bounds())
	dc.SetDst(img)

	// 文字大小
	//dc.SetFontSize(float64(25))
	dc.SetHinting(font.HintingFull)

	fontWidth := c.config.imageWidth / len(text)

	for i, s := range text {
		fontSize := c.config.imageHeight * (utils.GetRandomInt(7) + 7) / 16
		dc.SetSrc(image.NewUniform(utils.RandDeepColor()))
		dc.SetFontSize(float64(utils.GetRandBetween(c.config.fontCodeLen.Min, c.config.fontCodeLen.Max)))
		dc.SetFont(fontConfig)
		x := fontWidth * i + fontWidth / fontSize + 5
		y := c.config.imageHeight / 2 + fontSize / 2 - utils.GetRandomInt(c.config.imageHeight / 16 * 3)
		pt := freetype.Pt(x, y)
		if _, err := dc.DrawString(string(s), pt); err != nil {
			return err
		}
	}
	return nil
}

/**
 * @description: 配置验证码字出现的位置
 * @param {CharDot} dot
 * @param {[]color.RGBA} colorArr
 * @param {color.Color} fc
 * @return {*}
 */
func (c *Captcha) genDots(context string, padding int) map[int]CharDot {
	dots := make(map[int]CharDot) // 各个文字点位置
	width := c.config.imageWidth
	height := c.config.imageHeight
	if padding > 0 {
		width -= padding
		height -= padding
	}
	strs := []rune(context)
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		// 随机角度
		randAngle := c.getRandAngle()
		// 随机颜色
		randColor := utils.GetRandomColor()
		// 随机文字大小
		randFontSize := utils.GetRandBetween(c.config.fontCodeLen.Min, c.config.fontCodeLen.Max)
		fontHeight := randFontSize
		fontWidth := randFontSize
		if utf8.RuneCountInString(string(str)) > 1 {
			fontWidth = randFontSize * LenChineseChar(string(str))
			if randAngle > 0 {
				surplus := fontWidth - randFontSize
				ra := randAngle % 90
				pr := float64(surplus) / 90
				h := math.Max(float64(ra)*pr, 1)
				fontHeight = fontHeight + int(h)
			}
		}
		_w := width / len(strs)
		rd := math.Abs(float64(_w) - float64(fontWidth))
		x := (i * _w) + utils.GetRandBetween(0, int(math.Max(rd, 1)))
		x = int(math.Min(math.Max(float64(x), 10), float64(width-10-(padding*2))))
		y := utils.GetRandBetween(10, height+fontHeight)
		y = int(math.Min(math.Max(float64(y), float64(fontHeight+10)), float64(height+(fontHeight/2)-(padding*2))))
		text := fmt.Sprintf("%s", str)

		dot := CharDot{i, x, y, randFontSize, fontWidth, fontHeight, text, randAngle, randColor}
		dots[i] = dot
	}
	return dots
}

/**
 * @Description: 获取随机角度
 * @receiver cc
 * @return int
 */
func (c *Captcha) getRandAngle() int {
	angles := []CodeLen{
		{20, 35},
		{35, 45},
		{45, 60},
		{1, 15},
		{15, 30},
		{30, 45},
		{315, 330},
		{330, 345},
		{345, 359},
		{290, 305},
		{305, 325},
		{325, 330},
	}
	anglesLen := len(angles)
	index := utils.GetRandBetween(0, anglesLen)
	if index >= anglesLen {
		index = anglesLen - 1
	}
	angle := angles[index]
	res := utils.GetRandBetween(angle.Min, angle.Max)
	return res
}
