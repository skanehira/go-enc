package urlenc

import (
	"net/url"

	"github.com/skanehira/go-enc/pkg/errors"
)

func Encode(rawURL string) (string, error) {
	if rawURL == "" {
		return "", errors.NoStringError
	}

	baseURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	query := url.Values{}
	for key, value := range baseURL.Query() {
		for _, v := range value {
			query.Add(key, v)
		}
	}
	baseURL.RawQuery = query.Encode()

	return baseURL.String(), nil
}

func Decode(rawURL string) (string, error) {
	if rawURL == "" {
		return "", errors.NoStringError
	}

	return url.QueryUnescape(rawURL)
}
