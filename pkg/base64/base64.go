package base64

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
)

var NoStringError = errors.New("there are no specified strings")

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
		return "", NoStringError
	}

	result, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(result), err
}

func (b *Base64) EncodeFile(fileName string) (string, error) {
	if fileName == "" {
		return "", NoStringError
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
