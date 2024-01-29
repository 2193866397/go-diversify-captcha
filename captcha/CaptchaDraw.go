package captcha

import (
	"go-diversify-captcha/utils"
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
)

/*
 *@author: 随风飘的云
 *@description: 创建验证码画板
 *@date: 2023-10-04 16:26
 */
type CaptchaDraw struct{}

type CharDot struct {
	// 顺序索引
	Index int
	// x,y位置
	Dx int
	Dy int
	// 字体大小
	Size int
	// 字体宽
	Width int
	// 字体高
	Height int
	// 字符文本
	Text string
	// 字体角度
	Angle int
	// 每个字的颜色
	Color color.RGBA
}
type Point struct {
	X, Y int
}

/**
 * @description: 创建验证码像素画板
 * @param {*} width
 * @param {int} height
 * @param {[]color.RGBA} bg
 * @return {*}
 */
func (cd *CaptchaDraw) CreateCaptchaDrawWithSize(width, height int, bg []color.RGBA) *image.Paletted {
	// 默认的颜色为白色
	p := []color.Color{
		color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00},
	}
	for _, co := range bg {
		p = append(p, co)
	}
	return image.NewPaletted(image.Rect(0, 0, width, height), p)
}

/**
 * @description: 创建验证码画板
 * @param {bool} isAlpha
 * @return {*}
 */
func (cd *CaptchaDraw) CreateCanvas(width, height int, isAlpha bool) (img *image.NRGBA) {
	img = image.NewNRGBA(image.Rect(0, 0, width, height))
	// 画背景
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if isAlpha {
				img.Set(x, y, color.Alpha{A: uint8(0)})
			} else {
				img.Set(x, y, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}
	return
}

/**
 * @description: 创建具有背景颜色的画板
 * @param {*} width
 * @param {int} height
 * @param {color.RGBA} bg
 * @return {*}
 */
func (cd *CaptchaDraw) CreateCanvasWithBg(width, height int, bg color.RGBA) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)
	return img
}

/**
 * @description: 设置透明度
 * @param {float64} val
 * @return {*}
 */
func (cd *CaptchaDraw) formatAlpha(val float64) uint8 {
	a := math.Min(val, 1)
	alpha := a * 255
	return uint8(alpha)
}

/**
 * @description: 画随机直线
 * @param {*image.NRGBA} p
 * @param {Point} point1
 * @param {Point} point2
 * @param {color.RGBA} lineColor
 * @return {*}
 */
func (cd *CaptchaDraw) drawBeeline(p *image.NRGBA, point1 Point, point2 Point, lineColor color.RGBA) {
	dx := math.Abs(float64(point1.X - point2.X))
	dy := math.Abs(float64(point2.Y - point1.Y))
	sx, sy := 1, 1
	if point1.X >= point2.X {
		sx = -1
	}
	if point1.Y >= point2.Y {
		sy = -1
	}
	err := dx - dy
	for {
		p.Set(point1.X, point1.Y, lineColor)
		p.Set(point1.X+1, point1.Y, lineColor)
		p.Set(point1.X-1, point1.Y, lineColor)
		p.Set(point1.X+2, point1.Y, lineColor)
		p.Set(point1.X-2, point1.Y, lineColor)
		if point1.X == point2.X && point1.Y == point2.Y {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			point1.X += sx
		}
		if e2 < dx {
			err += dx
			point1.Y += sy
		}
	}
}

/**
 * @description: 画随机直线
 * @param {*image.NRGBA} p
 * @param {*} width
 * @param {*} height
 * @param {int} num
 * @return {*}
 */
func (cd *CaptchaDraw) drawSlimLine(p *image.NRGBA, width, height, num int) {

	first := width / 10
	end := first * 9

	y := height / 3

	for i := 0; i < num; i++ {

		point1 := Point{X: rand.Intn(first), Y: rand.Intn(y)}
		point2 := Point{X: rand.Intn(first) + end, Y: rand.Intn(y)}

		if i%2 == 0 {
			point1.Y = rand.Intn(y) + y*2
			point2.Y = rand.Intn(y)
		} else {
			point1.Y = rand.Intn(y) + y*(i%2)
			point2.Y = rand.Intn(y) + y*2
		}
		cd.drawBeeline(p, point1, point2, utils.GetRandomColor())
	}
}

/**
 * @description: 旋转任意角度
 * @param {image.Paletted} p
 * @param {int} angle
 * @return {*}
 */
func (cd *CaptchaDraw) Rotate(p *image.Paletted, angle int) {
	tarImg := p
	width := tarImg.Bounds().Max.X
	height := tarImg.Bounds().Max.Y
	r := width / 2
	retImg := image.NewPaletted(image.Rect(0, 0, width, height), tarImg.Palette)
	for x := 0; x <= retImg.Bounds().Max.X; x++ {
		for y := 0; y <= retImg.Bounds().Max.Y; y++ {
			tx, ty := cd.angleSwapPoint(float64(x), float64(y), float64(r), float64(angle))
			retImg.SetColorIndex(x, y, tarImg.ColorIndexAt(int(tx), int(ty)))
		}
	}

	nW := retImg.Bounds().Max.X
	nH := retImg.Bounds().Max.Y
	for x := 0; x < nW; x++ {
		for y := 0; y < nH; y++ {
			p.SetColorIndex(x, y, retImg.ColorIndexAt(x, y))
		}
	}
}

/**
 * @description: 角度转换坐标
 * @param {*} x
 * @param {*} y
 * @param {*} r
 * @param {float64} angle
 * @return {*}
 */
func (cd *CaptchaDraw) angleSwapPoint(x, y, r, angle float64) (tarX, tarY float64) {
	x -= r
	y = r - y
	sinVal := math.Sin(angle * (math.Pi / 180))
	cosVal := math.Cos(angle * (math.Pi / 180))
	tarX = x*cosVal + y*sinVal
	tarY = -x*sinVal + y*cosVal
	tarX += r
	tarY = r - tarY
	return
}

/**
 * @description: 图片扭曲
 * @param {image.Paletted} p
 * @param {float64} amplude
 * @param {float64} period
 * @return {*}
 */
func (cd *CaptchaDraw) distort(p *image.Paletted, amplude float64, period float64) {
	w := p.Bounds().Max.X
	h := p.Bounds().Max.Y
	newP := image.NewPaletted(image.Rect(0, 0, w, h), p.Palette)
	dx := 2.0 * math.Pi / period
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			xo := amplude * math.Sin(float64(y)*dx)
			yo := amplude * math.Cos(float64(x)*dx)
			newP.SetColorIndex(x, y, p.ColorIndexAt(x+int(xo), y+int(yo)))
		}
	}

	nW := newP.Bounds().Max.X
	nH := newP.Bounds().Max.Y
	for x := 0; x < nW; x++ {
		for y := 0; y < nH; y++ {
			p.SetColorIndex(x, y, newP.ColorIndexAt(x, y))
		}
	}

	newP.Palette = nil
}

/**
 * @description: 绘制非实心小圆圈
 * @param {*image.RGBA} img
 * @param {*} x
 * @param {*} y
 * @param {int} radius
 * @param {color.RGBA} c
 * @return {*}
 */
func (cd *CaptchaDraw) drawCircle(img *image.NRGBA, x, y, radius int, c color.RGBA) {
	for angle := 0.0; angle < 360; angle += 0.1 {
		radians := angle * (math.Pi / 180.0)
		xPos := x + int(float64(radius)*math.Cos(radians))
		yPos := y + int(float64(radius)*math.Sin(radians))

		if xPos >= 0 && xPos < img.Bounds().Max.X && yPos >= 0 && yPos < img.Bounds().Max.Y {
			img.Set(xPos, yPos, c)
		}
	}
}

/**
 * @description: 画随机曲线
 * @param {*image.NRGBA} img
 * @param {color.RGBA} c
 * @return {*}
 */
func (cd *CaptchaDraw) drawRandomCurve(img *image.NRGBA, c color.RGBA, shadeThickNess int) {
	// 曲线颜色，这里用黑色
	curveColor := c

	// 曲线粗细
	curveWidth := shadeThickNess

	// 随机生成曲线的起始点和终止点
	startX := rand.Float64() * float64(img.Bounds().Max.X)
	startY := rand.Float64() * float64(img.Bounds().Max.Y)
	endX := rand.Float64() * float64(img.Bounds().Max.X)
	endY := rand.Float64() * float64(img.Bounds().Max.Y)

	// 随机生成曲线的控制点
	controlX := rand.Float64() * float64(img.Bounds().Max.X)
	controlY := rand.Float64() * float64(img.Bounds().Max.Y)

	// 遍历每个像素点，根据贝塞尔曲线的公式计算曲线上的点并画到图片上
	for t := 0.0; t <= 1.0; t += 0.001 {
		x := math.Pow(1-t, 2)*startX + 2*(1-t)*t*controlX + math.Pow(t, 2)*endX
		y := math.Pow(1-t, 2)*startY + 2*(1-t)*t*controlY + math.Pow(t, 2)*endY

		// 画点到图片上
		for i := -int(curveWidth) / 2; i <= int(curveWidth)/2; i++ {
			for j := -int(curveWidth) / 2; j <= int(curveWidth)/2; j++ {
				img.Set(int(x)+i, int(y)+j, curveColor)
			}
		}
	}
}

/**
 * @description: 画图片边框
 * @param {*image.RGBA} img
 * @param {color.Color} borderColor
 * @param {int} thickness
 * @return {*}
 */
func (cd *CaptchaDraw) drawBorder(img *image.NRGBA, borderColor color.Color, thickness int) {
	rect := img.Bounds()

	// 画上边框
	draw.Draw(img, image.Rect(0, 0, rect.Max.X, thickness), &image.Uniform{borderColor}, image.Point{}, draw.Src)

	// 画下边框
	draw.Draw(img, image.Rect(0, rect.Max.Y-thickness, rect.Max.X, rect.Max.Y), &image.Uniform{borderColor}, image.Point{}, draw.Src)

	// 画左边框
	draw.Draw(img, image.Rect(0, 0, thickness, rect.Max.Y), &image.Uniform{borderColor}, image.Point{}, draw.Src)

	// 画右边框
	draw.Draw(img, image.Rect(rect.Max.X-thickness, 0, rect.Max.X, rect.Max.Y), &image.Uniform{borderColor}, image.Point{}, draw.Src)
}

/**
 * @description: 在x轴上扭曲图片
 * @param {*image.NRGBA} img
 * @param {float64} factor
 * @return {*}
 */
func (cd *CaptchaDraw) warpX(img *image.NRGBA, factor float64) {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			newX := int(float64(x) + factor*math.Sin(2*math.Pi*float64(y)/float64(bounds.Max.Y)))
			if newX >= 0 && newX < bounds.Max.X {
				newImg.Set(x, y, img.At(newX, y))
			}
		}
	}

	draw.Draw(img, bounds, newImg, image.Point{}, draw.Over)
}

/**
 * @description: 在y轴上扭曲图片
 * @param {*image.NRGBA} img
 * @param {float64} factor
 * @return {*}
 */
func (cd *CaptchaDraw) warpY(img *image.NRGBA, factor float64) {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			newY := int(float64(y) + factor*math.Sin(2*math.Pi*float64(x)/float64(bounds.Max.X)))
			if newY >= 0 && newY < bounds.Max.Y {
				newImg.Set(x, y, img.At(x, newY))
			}
		}
	}

	draw.Draw(img, bounds, newImg, image.Point{}, draw.Over)
}

func (cd *CaptchaDraw) distortImage(img *image.NRGBA, xOffset, yOffset float64) *image.RGBA {
	imageWidth := 200.0
	imageHeight := 100.0
	bounds := img.Bounds()
	distortedImg := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			srcX := int(float64(x) + xOffset*math.Sin(2*math.Pi*float64(y)/imageHeight))
			srcY := int(float64(y) + yOffset*math.Sin(2*math.Pi*float64(x)/imageWidth))

			srcPoint := image.Point{X: srcX, Y: srcY}
			srcPoint = srcPoint.Add(bounds.Min)
			srcPoint = srcPoint.Mod(bounds)

			distortedImg.Set(x, y, img.At(srcPoint.X, srcPoint.Y))
		}
	}

	return distortedImg
}
