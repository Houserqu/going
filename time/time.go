package time

import "time"

/**
 * 获取当前时间的字符串
 */
func TimeNowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
