package main

// Input: A array of int, N int as length of A, K int
// Output: Max summary of K elements
func FindMaxSum(arr []int, n int, k int) (max int) {
	// n, k >= 0; len(arr) >= n
	if k == 0 {
		max = 0
		return
	}
	if n <= k {
		max = 0
		for i := 0; i < n; i++ {
			max += arr[i]
		}
		return
	}

	max = 0
	// Calculate window sum
	for i := 0; i < k; i++ {
		max += arr[i]
	}
	// Slide window and check max sum
	sum := max
	for right := k; right < n; right++ {
		left := right - k // k = window size
		sum = sum - arr[left] + arr[right]
		if sum > max {
			max = sum
		}
	}
	return
}

// Input: A as array of characters, N int as length of array, K int
// Output: Longest substring with K distinct character
func FindLongestDistinctChars(a []byte, n int, k int) (chars []byte) {
	return
}
