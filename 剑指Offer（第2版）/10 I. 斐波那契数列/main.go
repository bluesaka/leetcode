package main

import "fmt"

/**
方法一：记忆化递归，注意递归容易引起超时，可采用记忆化递归来避免超时

递归法：
f(n) = f(n-1) + f(n-2) 会重复计算f(n-2)，容易造成超时
时间O(2^n)，空间O(n)

记忆递归法：
使用长度为n的数组/slice来存储已经计算的值，时间O(n)，空间O(n)
需要额外使用O(n)的空间
*/
func fib(n int) int {
	// 递归容易引起超时
	//if n == 0 || n == 1 {
	//	return n
	//}
	//return (fib(n-1) + fib(n-2)) % 1000000007

	// 记忆化递归
	m := make([]int, n+1)
	return doFib(n, m)
}

func doFib(n int, m []int) int {
	if n == 0 || n == 1 {
		m[n] = n
		return n
	}
	if m[n] != 0 {
		return m[n]
	}
	m[n] = (doFib(n-1, m) + doFib(n-2, m)) % 1000000007
	return m[n]
}

/**
（推荐）方法二：迭代法，动态规划，时间O(n)，空间O(1)
*/
func fib2(n int) int {
	//nums := []int{0, 1}
	//for i := 2; i <= n; i++ {
	//	nums = append(nums, (nums[i-1]+nums[i-2])%1000000007)
	//}
	//return nums[n]

	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}

func main() {
	fmt.Println(fib(7))
	fmt.Println(fib2(7))
}
