package id

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

/**
 * 生成时间戳+随机数 ID
 */
func GenUnixID() int64 {
	t := time.Now().UnixMilli()
	n, _ := rand.Int(rand.Reader, big.NewInt(9999))
	return t*10000 + n.Int64()
}

/**
 * 生成前缀+日期+随机数 ID
 */
func GenDateID(prefix string) string {
	n, _ := rand.Int(rand.Reader, big.NewInt(9999))
	t := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s%s%d", prefix, t, n)
}
