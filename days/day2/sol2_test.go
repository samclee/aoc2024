package day2

import "testing"

func TestSol2(t *testing.T) {
	cases := []struct {
		Input    []int
		Low      int
		High     int
		Expected bool
	}{
		{
			Input:    []int{1, 2, 7, 8, 9},
			Low:      1,
			High:     3,
			Expected: false,
		},
		{
			Input:    []int{1, 2, 7, 8, 9},
			Low:      -3,
			High:     -1,
			Expected: false,
		},
		{
			Input:    []int{9, 7, 6, 2, 1},
			Low:      1,
			High:     3,
			Expected: false,
		},
		{
			Input:    []int{9, 7, 6, 2, 1},
			Low:      -3,
			High:     -1,
			Expected: false,
		},
	}

	for i, c := range cases {
		got := diffsInRangeWithIgnore(c.Input, c.Low, c.High)
		if got != c.Expected {
			t.Errorf("[%d] expected %v, got %v", i, c.Expected, got)
		}
	}
}
