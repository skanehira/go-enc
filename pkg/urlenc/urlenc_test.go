package urlenc

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{`こんにちわ 世界`, `%E3%81%93%E3%82%93%E3%81%AB%E3%81%A1%E3%82%8F%20%E4%B8%96%E7%95%8C`},
		{`hello world`, `hello%20world`},
		{`"\><[]|`, `%22%5C%3E%3C%5B%5D%7C`},
	}

	for _, test := range tests {
		out, err := Encode(test.in)
		if err != nil {
			t.Errorf("in: %s, error: %s", test.in, err)
			continue
		}

		if test.out != out {
			t.Errorf("want: %s, got: %s", test.out, out)
		}
	}
}
