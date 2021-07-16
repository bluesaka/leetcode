package main

import "fmt"

/**
方法一：完全背包动态规划DP，时间(SN)，空间O(S)
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(5, []int{1, 2, 5})
}
