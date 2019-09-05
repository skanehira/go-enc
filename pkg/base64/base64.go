package base64

import (
	"encoding/base64"
)

type Base64 struct {
}

func New() *Base64 {
	return &Base64{}
}

func (b *Base64) Encode(str string) string {
	if str == "" {
		return ""
	}

	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (b *Base64) Decode(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	result, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(result), err
}
