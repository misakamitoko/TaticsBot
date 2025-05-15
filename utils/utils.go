package utils

import "strings"

// 分割：平均名次3.07第一名23.45%前四名率78.35%挑选率2.74%
func SplitString(s string) []string {
	s = strings.Replace(s, "平均名次", "", 1)
	s = strings.Replace(s, "第一名", ";", 1)
	s = strings.Replace(s, "前四名率", ";", 1)
	s = strings.Replace(s, "挑选率", ";", 1)
	return strings.Split(s, ";")
}
