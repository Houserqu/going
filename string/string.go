package string

import "strings"

/**
 * 替换字符串中的最后一个匹配项
 */
func ReplaceLast(str, old, new string) string {
	i := strings.LastIndex(str, old)
	return str[:i] + strings.Replace(str[i:], old, new, 1)
}
