package base64

import (
	"encoding/base64"

	"github.com/skanehira/go-enc/pkg/errors"
)

func Encode(bytes []byte) string {
	if len(bytes) == 0 {
		return ""
	}

	return base64.StdEncoding.EncodeToString(bytes)
}

func Decode(str string) (string, error) {
	if str == "" {
		return "", errors.NoStringError
	}

	result, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(result), err
}
