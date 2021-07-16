package main

import "fmt"

/**
方法一：完全背包动态规划DP，时间O(SN)，空间(S)，S为金额，N为面额数
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
