package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(v string) string {
	if v != "" {
		d := []byte(v)
		m := md5.New()
		m.Write(d)
		return hex.EncodeToString(m.Sum(nil))
	}
	return ""
}