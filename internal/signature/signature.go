package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"time"
)

// GenerateSignature creates a Smile ID HMAC-SHA256 signature
func GenerateSignature(apiKey, partnerID string) (signature, timestamp string, err error) {
	timestamp = time.Now().UTC().Format(time.RFC3339Nano)
	h := hmac.New(sha256.New, []byte(apiKey))
	h.Write([]byte(timestamp))
	h.Write([]byte(partnerID))
	h.Write([]byte("sid_request"))
	sig := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return sig, timestamp, nil
}
