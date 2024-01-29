package constants

/*
 *@author: 随风飘的云
 *@description: 定义验证码图片常用的常量
 *@date: 2023-09-30 16:56
 */
const (
	// 边框
	CAPTCHA_BORDER string = "captcha.border"

	// 边框颜色
	CAPTCHA_BORDER_COLOR string = "captcha.border.color"

	// 边框厚度
	CAPTCHA_BORDER_THICKNESS string = "captcha.border.thickness"

	// 干扰圆数量
	CAPTCHA_NOISE_NUMBER string = "captcha.noise.number"

	// 干扰线数量
	CAPTCHA_SHADE_NUMBER string = "captcha.shade.number"

	// 干扰线厚度
	CAPTCHA_SHADE_THICKNESS string = "captcha.shade.thickness"

	// 生成验证码内容类型
	CAPTCHA_TEXT_TYPE string = "captcha.text.type"

	// 验证码长度
	CAPTCHA_TEXT_LENGTH string = "captcha.text.length"

	// 验证码字体名称
	CAPTCHA_TEXT_FONT_NAMES string = "captcha.text.font.names"

	// 验证码字体大小
	CAPTCHA_TEXT_FONT_SIZE string = "captcha.text.font.size"

	//验证码字体厚度
	CAPTCHA_TEXT_FONT_THICKNESS string = "captcha.text.font.thickness"

	// 设置字体颜色
	CAPTCHA_TEXT_FONT_COLOR string = "captcha.text.font.color"

	// 设置背景颜色
	CAPTCHA_BACKGROUND_COLOR string = "captcha.background.color"

	// 添加扭曲效果
	CAPTCHA_IMAGE_SHEAR_ADD string = "captcha.image.shear.add"

	// 设置图片类型
	CAPTCHA_IMAGE_TYPE string = "captcha.image.type"

	// 图片宽度
	CAPTCHA_IMAGE_WIDTH string = "captcha.image.width"

	// 图片高度
	CAPTCHA_IMAGE_HEIGHT string = "captcha.image.height"
)
const (
	// 无压缩质量,原图
	QualityNone = iota
	// 质量压缩程度 1-5 级别，压缩级别越低图像越清晰
	QualityLevel1 = 100
	QualityLevel2 = 80
	QualityLevel3 = 60
	QualityLevel4 = 40
	QualityLevel5 = 20
)
const (
	ImageStringDpi float64 = 72.0
)