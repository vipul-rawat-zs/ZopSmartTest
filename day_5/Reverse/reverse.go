package reverse

import "errors"

func Reverse(s string) (string, error) {
	var x string
	if len(s) == 0 {
		return "", errors.New("empty string")
	}
	for _, v := range s {
		x = string(v) + x
	}
	return x, nil
}
