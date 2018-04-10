package util

import (
	"crypto/md5"
	"fmt"
)

// MD5加密
func MD5(str string) (md5str string) {
	data := []byte(str)
	has := md5.Sum(data)
	md5str = fmt.Sprintf("%x", has)
	return
}
