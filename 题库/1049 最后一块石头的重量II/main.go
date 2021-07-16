package main

import "fmt"

/**
方法一：01背包动态规划DP

转化为0-1背包的最小值问题。
*/
func lastStoneWeightII(stones []int) int {
	n := len(stones)
	sum := 0
	for _, num := range stones {
		sum += num
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sum/2+1)
	}

	var max func(x, y int) int
	max = func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= sum/2; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= stones[i-1] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i-1]]+stones[i-1])
			}
		}
	}
	return sum - dp[n][sum/2]*2
}

/**
方法二：01背包动态规划DP优化，滚动数组
*/
func lastStoneWeightII2(stones []int) int {
	n := len(stones)
	sum := 0
	for _, num := range stones {
		sum += num
	}
	dp := make([]int, sum/2+1)

	var max func(x, y int) int
	max = func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}

	for i := 0; i < n; i++ {
		for j := sum / 2; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - dp[sum/2]*2
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeightII2([]int{2, 7, 4, 1, 8, 1}))
}
