package main

import (
	"fmt"
	"math"
)

/**
方法一：动态规划，时间O(n^2)，空间O(n)
*/
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

/**
方法二：贪心算法，时间O(n)，空间O(1)

尽可能把绳子分成长度为3的小段，这样乘积最大
*/
func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}
	return res * n
}

/**
方法三：需要用到数论推导，时间O(1)，空间O(1)
*/
func cuttingRope3(n int) int {
	if n <= 3 {
		return n - 1
	}
	return int(math.Pow(3, float64(n/3))) * 4 / (4 - n%3)
}

func main() {
	fmt.Println(cuttingRope(10))
	fmt.Println(cuttingRope2(10))
	fmt.Println(cuttingRope3(10))
}
