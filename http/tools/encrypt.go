package tools

import (
	"crypto/md5"
	"fmt"
)

//MD5 MD5 加密
func MD5(val string) string {
	data := []byte(val)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
