package constants

/*
 *@author: 随风飘的云
 *@description: 验证码的内容类型
 *@date: 2023-09-30 17:32
 */
const (
	// 数字
	NUMBER string = "NUMBER"
	// 中文数字
	NUMBER_ZH string = "NUMBER_ZH"
	// 中文
	CHINESE string = "CHINESE"
	// 中文成语
	CHINESE_IDIOM string = "CHINESE_IDIOM"
	// 小写英文
	CHAR string = "CHAR"
	// 大写
	UPPER_CHAR string = "UPPER_CHAR"
	// 英文和中文混合
	STRING string = "STRING"
	// 数字和大小写字符验证码
	UPPER_STRING string = "UPPER_STRING"
	// 算术
	ARITHMETIC string = "ARITHMETIC"
	// 中文算术
	ARITHMETIC_ZH string = "ARITHMETIC_ZH"
	// 常量汇总（取巧方法）
	TEXT_VALUE string = "NUMBER,NUMBER_ZH,CHINESE,CHINESE_IDIOM,CHAR,UPPER_CHAR,STRING,UPPER_STRING,ARITHMETIC,ARITHMETIC_ZH"
)
