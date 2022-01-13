package concurrency

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		input1   []int
		input2   chan int
		input3   string
		expected int
	}{
		{[]int{1, 2, 3, 4}, make(chan int), "2", 10},
		{[]int{-11, 2, 3, 4}, make(chan int), "1", -1},
		{[]int{1, -2, 3, 4}, make(chan int), "1", 8},
	}

	for _, v := range testCases {
		t.Run("Running test cases on wait calls", func(t *testing.T) {
			Sum(v.input1, v.input2, v.input3)
			out := <-v.input2
			fmt.Println(out)
			if out != v.expected {
				t.Errorf("expected %v but got %v", v.expected, out)
			}
			close(v.input2)
		})
	}

}
