package reverse

import (
	"errors"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
		err      error
	}{
		{"Hello, world", "dlrow ,olleH", nil},
		{"Hello, 世界", "界世 ,olleH", nil},
		{"", "", errors.New("empty string")},
	}
	for _, c := range cases {
		got, err := Reverse(c.in)
		if err != nil && err.Error() != c.err.Error() {
			t.Errorf("got %v, want %v", err, c.err)
		}
		if got != c.want {
			t.Errorf("got %q, want %q", got, c.want)
		}
	}
}
