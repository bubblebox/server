package base62_test

import (
	"testing"

	"github.com/bubblebox/server/base62"
)

var base62Tests = []struct {
	n uint64 // input
	c string // code
}{
	{1, "1"},
	{10, "a"},
	{61, "Z"},
	{62, "10"},
	{63, "11"},
	{124, "20"},
	{200000000, "dxb8s"},
}

func TestBase62Encoding(t *testing.T) {
	for _, tt := range base62Tests {
		actual := base62.Encode(tt.n)
		if actual != tt.c {
			t.Errorf("Base62 encoding of `%d` to '%s' failed, got '%s'", tt.n, tt.c, actual)
		}
	}
}

func TestBase62Decoding(t *testing.T) {
	for _, tt := range base62Tests {
		actual := base62.Decode(tt.c)
		if actual != tt.n {
			t.Errorf("Base62 decoding of '%s' to `%d` failed, got `%d`", tt.c, tt.n, actual)
		}
	}
}
