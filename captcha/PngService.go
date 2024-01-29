package captcha

import (
	"fmt"
	"image/png"
	"os"
)

/*
 *@author: 随风飘的云
 *@description: 创建Png格式验证码
 *@date: 2023-10-02 17:01
 */
type PngCaptcha struct {
	*Captcha
}

func (p *PngCaptcha) CreateCaptcha(context string) (string, string, string, string, error)  {
	pngFile, _ := p.CreateCaptchaToPng(context)
	file, err := os.Create("./text.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return "", "", "", "", nil
	}
	defer file.Close()

	err = png.Encode(file, pngFile)
	if err != nil {
		fmt.Println("Error encoding PNG:", err)
		return "", "", "", "", nil
	}

	fmt.Println("Image saved successfully!")
	return "", "", "", "", nil
}
