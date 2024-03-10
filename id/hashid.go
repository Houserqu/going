package id

import "github.com/speps/go-hashids/v2"

/**
 * 生成 hash id
 */
func NumberToHashID(id int64, salt string) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 10

	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64([]int64{id})

	return e
}

/**
 * hash id 转数字
 */
func HashIDToNumber(hash string, salt string) int64 {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 10

	h, _ := hashids.NewWithData(hd)
	d := h.DecodeInt64(hash)

	return d[0]
}
