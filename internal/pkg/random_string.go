package pkg

import (
	"crypto/rand"
	"encoding/base64"
)

func Random32ByteString() string {
	randomKey := make([]byte, 32)

	_, _ = rand.Read(randomKey)

	return base64.RawURLEncoding.EncodeToString(randomKey)
}
