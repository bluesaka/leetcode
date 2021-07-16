package main

import "fmt"

/**
贪心算法，时间O(n)，空间O(1)
不推荐使用动态规划取余，效果不好
*/
func cuttingRope(n int) int {
	if n < 3 {
		return 1
	}
	if n == 3 {
		return 2
	}
	res := 1
	for n > 4 {
		res *= 3
		res %= 1000000007
		n -= 3
	}
	return res * n % 1000000007
}

func main() {
	fmt.Println(cuttingRope(10))
}
