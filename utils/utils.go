package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func PasswordEncode(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	herStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(herStr)
}
