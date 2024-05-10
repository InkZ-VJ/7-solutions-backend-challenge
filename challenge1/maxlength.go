package main

func maxSum(input [][]int) int {
	// Start from the second last row
	for i := len(input) - 2; i >= 0; i-- {
		for j := 0; j < len(input[i]); j++ {
			// Choose the maximum value between current and next row
			input[i][j] += max(input[i+1][j], input[i+1][j+1])
		}
	}
	// The maximum sum will be at the top of the pyramid
	return input[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
