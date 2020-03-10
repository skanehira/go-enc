package base64

import (
	"testing"

	"github.com/skanehira/go-enc/pkg/errors"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		in  []byte
		out string
	}{
		{[]byte("hello"), "aGVsbG8="},
		{[]byte{}, ""},
	}

	for _, test := range tests {
		out := Encode(test.in)
		if test.out != out {
			t.Errorf("want: %s, got: %s", test.out, out)
		}
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"aGVsbG8=", "hello"},
	}

	for _, test := range tests {
		out, err := Decode(test.in)
		if err != nil {
			t.Errorf("cannot decode ':%s', reason: %s", test.in, err)
		}
		if test.out != out {
			t.Errorf("want: %s, got: %s", test.out, out)
		}
	}
}

func TestDecodeError(t *testing.T) {
	tests := []struct {
		in  string
		err error
	}{
		{"", errors.NoStringError},
	}

	for _, test := range tests {
		_, err := Decode(test.in)
		if err == nil {
			t.Errorf("want error: %s, got nil", test.err)
		}

		if err != test.err {
			t.Errorf("want error: %s, got error: %s", test.err, err)
		}
	}
}
