package text

import "strings"

/*
 *@author: 随风飘的云
 *@description: 创建文本
 *@date: 2023-10-02 17:00
 */
type FactoryContext struct {
	textFactory map[string]func(contextType string, length int) string
}

func (c *FactoryContext)initText(contextType string) func(contextType string, length int) string{
	number := textNumber{}
	chinese := textChinese{}
	chars := textChar{}
	text := textString{}
	arithmetic := textArithmetic{}
	// 优化switch case
	factory := &FactoryContext{
		textFactory: map[string]func(contextType string, length int) string{
			"NUMBER" : number.GetTextContent,
			"CHINESE" : chinese.GetTextContent,
			"CHAR" : chars.GetTextContent,
			"STRING" : text.GetTextContent,
			"ARITHMETIC": arithmetic.GetTextContent,
		},
	}
	// 切割内容类型分类
	contextType = strings.Split(contextType, "_")[0]
	return factory.textFactory[contextType]
}

func (c *FactoryContext)GetContextText(contextType string, length int) string {
	textFunc := c.initText(contextType)
	return textFunc(contextType, length)
}