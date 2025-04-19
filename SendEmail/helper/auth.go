package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"
)

func ComputeContentHash(content []byte) string {
	hash := sha256.New()
	hash.Write(content)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func ComputeSignature(stringToSign, secret string) string {
	decodedSecret, _ := base64.StdEncoding.DecodeString(secret)
	h := hmac.New(sha256.New, decodedSecret)
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func FormatDate(dt time.Time) string {
	return dt.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
}
