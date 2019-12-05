package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5String(strings string) string {

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strings))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func Md5ByByte(bytes []byte) string {

	md5Ctx := md5.New()
	md5Ctx.Write(bytes)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
