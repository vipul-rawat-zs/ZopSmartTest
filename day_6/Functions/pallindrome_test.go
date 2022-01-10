package functions

import (
	"testing"
)

func TestPallindrome(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{"case 1", "A man, A plan ,A canal :Panama", true},
		{"case 2", "Hello, Moto", false},
		{"case 3", "Ma'am", true},
		{"case 4", "ThIs ana Siht", true},
		{"case 6", "", true},
	}

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			out := CheckPallindrome(v.input)
			if out != v.expected {
				t.Errorf("Expected %v but got %v", v.expected, out)
			}
		})
	}
}
