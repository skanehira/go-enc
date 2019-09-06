package urlenc

import (
	"net/url"

	"github.com/skanehira/go-enc/pkg/errors"
)

type URLEnc struct {
}

func New() *URLEnc {
	return &URLEnc{}
}

func (u *URLEnc) Encode(rawURL string) (string, error) {
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

func (u *URLEnc) Decode(rawURL string) (string, error) {
	if rawURL == "" {
		return "", errors.NoStringError
	}

	baseURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	if query, err := url.QueryUnescape(baseURL.RawQuery); err != nil {
		return "", err
	} else {
		baseURL.RawQuery = query
	}

	return baseURL.String(), nil
}
