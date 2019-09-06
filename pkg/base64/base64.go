package base64

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/skanehira/go-enc/pkg/errors"
)

type Base64 struct {
}

func New() *Base64 {
	return &Base64{}
}

func (b *Base64) Encode(bytes []byte) string {
	if len(bytes) == 0 {
		return ""
	}

	return base64.StdEncoding.EncodeToString(bytes)
}

func (b *Base64) Decode(str string) (string, error) {
	if str == "" {
		return "", errors.NoStringError
	}

	result, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(result), err
}

func (b *Base64) EncodeFile(fileName string) (string, error) {
	if fileName == "" {
		return "", errors.NoStringError
	}

	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return b.Encode(data), nil
}
