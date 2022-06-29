package pkg

import (
	"crypto/md5"
	"encoding/hex"
)

func GetWebhookID(userID string) string {
	hash := md5.Sum([]byte(userID))
	return hex.EncodeToString(hash[:])
}
