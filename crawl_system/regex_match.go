package main

import "regexp"

/**
	正则匹配
 */
func regexMatch(target string, regex string) (string, bool) {
	// 解析正则表达式
	reg, err := regexp.Compile(regex)
	if err != nil {
		return "", false
	}

	// 检测是否匹配
	if ok := reg.MatchString(target); !ok {
		return "", false
	}

	// 返回匹配内容
	return reg.FindString(target), true
}