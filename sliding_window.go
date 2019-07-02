package main

// n, k >= 0; len(arr) >= n
func MaxSum(arr []int, n int, k int) (sum int) {
	if k == 0 {
		sum = 0
		return
	}
	if n <= k {
		sum = 0
		for i := 0; i < n; i++ {
			sum += arr[i]
		}
		return
	}

	sum = 0
	// Calculate window sum
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	// Slide window and check max sum
	for right := k; right < n; right++ {
		left := right - k // k = window size
		if arr[right] > arr[left] {
			sum = sum - arr[left] + arr[right]
		}
	}
	return
}
