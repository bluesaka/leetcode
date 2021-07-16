package main

import "fmt"

/**
方法一：01背包动态规划DP，时间O(n*(sum-target)))，空间O(n*neg)，其中neg为负数的和的绝对值

转化问题为0-1背包的方案数问题。

不推荐使用回溯算法，时间O(2^n)，空间O(n)，指数级的时间复杂度
*/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	neg := diff / 2
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		num := nums[i-1]
		for j := 0; j <= neg; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= num {
				dp[i][j] += dp[i-1][j-num]
			}
		}
	}

	return dp[n][neg]
}

/**
方法一：01背包动态规划DP优化（滚动数组），时间O(n*(sum-target)))，空间O(sum-target)，其中neg为负数的和的绝对值
*/
func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1

	for _, num := range nums {
		for j := neg; j >= num; j++ {
			dp[j] += dp[j-num]
		}
	}

	return dp[neg]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
