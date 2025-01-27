package utils

import (
	"crypto/md5"
	"encoding/hex"
	"picture/common/errorx"
)

const salt = "ning"

func EncryptPassword(password string) (string, error) {
	if password == "" {
		return "", errorx.NewCodeError(errorx.PasswordIsNull, "密码不能为空")
	}
	h := md5.New()
	h.Write([]byte(salt + password))
	return hex.EncodeToString(h.Sum(nil)), nil
}
