package utils

import (
	"strings"
	"unicode"
)

// 单词全部转化为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 单词全部转化为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线单词转为大写驼峰单词
func UderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线单词转为小写驼峰单词
func UderscoreToLowerCamelCase(s string) string {
	s = UderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
	return s
}

// 驼峰单词转下划线单词
func CamelCaseToUdnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}

// removeJSONTags 函数用于去掉JSON文本中的 ```json``` 标签
func RemoveJSONTags(jsonText string) string {
	// 去掉开头的 ```json```
	jsonText = strings.TrimPrefix(jsonText, "```json")
	// 去掉开头可能存在的换行符
	jsonText = strings.TrimLeft(jsonText, "\n")
	// 去掉结尾的 ```
	jsonText = strings.TrimSuffix(jsonText, "```")
	// 去掉结尾可能存在的换行符
	jsonText = strings.TrimRight(jsonText, "\n")
	return jsonText
}
