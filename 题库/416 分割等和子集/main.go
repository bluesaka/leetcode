package main

/**
方法一：01背包动态规划DP，时间O(NC)，空间O(NC)，其中N为数组元素的个数，C为数组元素和的一半

转化问题为0-1背包的可行性问题。

*/
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
		if dp[i][target] {
			return true
		}
	}

	return dp[n-1][target]
}

/**
方法二：01背包动态规划DP优化（滚动数组），时间O(NC)，空间O(C)，其中N为数组元素的个数，C为数组元素和的一半
*/
func canPartition2(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2

	dp := make([]bool, target+1)
	if nums[0] <= target {
		dp[nums[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			if dp[target] {
				return true
			}
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[target]
}

func main() {

}
