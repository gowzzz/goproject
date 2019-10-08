package handler

import (
	"testing"
)

func TestFunc1(t *testing.T) {
	var tests = []struct {
		inputA int
		inputB int
		want   int
	}{
		{1, 2, 3},
		{2, 2, 4},
	}
	for _, test := range tests {
		res := Add(test.inputA, test.inputB)
		if res != test.want {
			t.Errorf("Add(%d,%d) want %d,but get %d", test.inputA, test.inputB, test.want, res)
		}
	}
}
