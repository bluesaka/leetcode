package main

import (
	"fmt"
)

/**
方法一：循环检查二进制位，时间O(k)，k为int型的二进制位数32，空间O(1)，只需常数空间保存若干变量

2^i和n进行与运算，当n的第i位为1时，与运算结果部位0
*/
func hammingWeight(num uint32) int {
	var count int
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			count++
		}
	}
	return count
}

/**
方法二：位运算优化，时间O(logn)，循环次数为n的二进制中1的个数，空间O(1)

利用 num & (num-1)，把num的最低位的1变为0
 */
func hammingWeight2(num uint32) int {
	var count int
	for ; num > 0; num &= num - 1 {
		count++
	}
	return count
}

func main() {
	fmt.Println(hammingWeight(15))
	fmt.Println(hammingWeight2(15))
}
