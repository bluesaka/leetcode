package main

import "fmt"

/**
方法一：完全背包动态规划DP，时间O(n根号n)，空间O(n)
*/
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func main() {
	fmt.Println(numSquares(6))
}
