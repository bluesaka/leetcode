package main

import (
	"fmt"
)

/**
方法一：迭代累乘，时间O(n)，空间O(1)
注意此方法会超出时间限制
*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.00
	for i := 1; i <= n; i++ {
		res *= x
	}
	return res
}

/**
（推荐）方法二：快速幂，时间O(log2n)，空间O(1)
*/
func myPow2(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.00
	for n > 0 {
		if n&1 > 0 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

func main() {
	fmt.Println(myPow2(21.0012, 3))
	fmt.Println(myPow2(-21.0012, 3))
	fmt.Println(myPow2(21.0012, -3))
	fmt.Println(myPow2(-21.0012, -3))
}
