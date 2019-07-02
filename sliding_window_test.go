package main

import "testing"

func TestDoSlidingWindow(t *testing.T) {
	var testCases = []struct {
		arr         []int
		n           int
		k           int
		expectedSum int
	}{
		{
			[]int{}, 0, 0, 0,
		},
		{
			[]int{1}, 1, 0, 0,
		},
		{
			[]int{1}, 1, 1, 1,
		},
		{
			[]int{1, 2}, 2, 1, 2,
		},
		{
			[]int{1, 2, 3}, 3, 2, 5,
		},
		{
			[]int{1, 2, -1}, 3, 2, 3,
		},
		{
			[]int{1, 2, 2, 2, 2, 1}, 6, 6, 10,
		},
		{
			[]int{1, 3, 2, 4, 1, 1, 1, 5, 1, 9, 2, 5, 8, 3, 1, 1, 1}, 17, 4, 24,
		},
	}

	for _, tc := range testCases {
		sum := MaxSum(tc.arr, tc.n, tc.k)
		if sum != tc.expectedSum {
			t.Errorf("Expect %v but actualy %v", tc.expectedSum, sum)
		}
	}
}
