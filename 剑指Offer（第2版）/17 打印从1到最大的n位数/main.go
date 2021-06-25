package main

import "fmt"

/**
方法一：迭代法，时间O(10^n)，空间O(1)，列表作为返回结果不计入额外的空间
*/
func printNumbers(n int) []int {
	//var nums []int
	//end := int(math.Pow10(n))
	//for i := 1; i < end; i++ {
	//	nums = append(nums, i)
	//}
	//return nums

	end := 1
	for ; n > 0; n-- {
		end *= 10
	}
	nums := make([]int, end-1, end-1)
	for i := 0; i < end-1; i++ {
		nums[i] = i + 1
	}
	return nums
}

/**
大数越界问题，使用字符串存储
*/
func printNumbers2(n int) []string {
	nums := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	loop := make([]byte, n)
	var result []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			// 去除前导0
			k := 0
			for ; k < len(loop); k++ {
				if loop[k] != '0' {
					break
				}
			}
			result = append(result, string(loop[k:]))
			return
		}
		for _, v := range nums {
			loop[i] = v
			dfs(i + 1)
		}
	}
	dfs(0)
	return result
}

func main() {
	//fmt.Println(printNumbers(3))
	fmt.Println(printNumbers2(2))
}
