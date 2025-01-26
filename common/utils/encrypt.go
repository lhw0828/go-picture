package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const salt = "ning"

func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(salt + password))
	return hex.EncodeToString(h.Sum(nil))
}
