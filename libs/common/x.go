package common

import (
	"crypto/md5"
	"fmt"
)

func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", hash)
}
