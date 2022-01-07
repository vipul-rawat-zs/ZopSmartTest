package main

import "testing"

func TestFib(t *testing.T) {
	fibTests := []struct {
		description string
		n           int
		expect      int
	}{
		{"case1", 1, 1},
		{"case2", 2, 1},
		{"case3", 3, 2},
		{"case4", 4, 3},
		{"case5", 5, 5},
		{"case6", 6, 8},
		{"case7", 7, 13},
		{"case8", -1, 0},
	}

	for _, test := range fibTests {
		t.Run(test.description, func(t *testing.T) {
			if actual := fib(test.n); actual != test.expect {
				t.Errorf("%v , expected %v but got %v ", test.description, test.expect, actual)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(i)
	}
}
