package main

import "fmt"

/**
方法一：迭代法，类似于斐波那契数列
 */
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}

func main() {
	fmt.Println(numWays(7))
}
