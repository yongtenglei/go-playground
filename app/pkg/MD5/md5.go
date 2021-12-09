package MD5

import (
	"crypto/md5"
	"encoding/hex"
)

const SECRET = "yongtenglei.com"

func EncryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(SECRET))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
