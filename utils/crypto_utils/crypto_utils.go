package crypto_utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// Hash a password with a secret key
func HashPassword(password string) string {
	hash := hmac.New(sha256.New, []byte(os.Getenv("SECRETKEY")))

	io.WriteString(hash, password)

	hashedValue := hash.Sum(nil)

	return hex.EncodeToString(hashedValue)
}
