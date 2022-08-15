package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(intup string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(intup))
	return hex.EncodeToString(hash.Sum(nil))
}
