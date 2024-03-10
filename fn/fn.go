package fn

import "errors"

/**
 * 重试函数，只有失败次数达到指定次数才会返回错误
 */
func RetryFun(fn func() error, times int) (err error) {
	if times <= 0 {
		return errors.New("times must be greater than 0")
	}

	for i := 0; i < times; i++ {
		err = fn()
		if err == nil {
			return
		}
	}

	return
}
