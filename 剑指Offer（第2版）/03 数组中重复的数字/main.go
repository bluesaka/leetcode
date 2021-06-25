package main

import (
	"fmt"
	"sort"
)

/**
方法一：哈希表，时间O(n)，空间O(n)
*/
func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if m[num] > 0 {
			return num
		}
		m[num] = 1
	}
	return -1
}

/**
（推荐）方法二：原地置换，鸽巢原理，时间O(n)，空间O(1)
题目中给的元素值是小于len(nums)的，所以可以在位置i的地方放元素i
当位置i的元素不是i时，nums[i] 与 nums[nums[i]] 进行交换，如果此时 nums[i] = nums[nums[i]]，说明元素相等，则返回
 */
func findRepeatNumber2(nums []int) int {
	for i, v := range nums {
		for i != v {
			if v == nums[v] {
				return v
			}
			nums[i], nums[v] = nums[v], nums[i]
		}
	}
	return -1
}

/**
方法三：先排序，然后看相邻元素是否有相等的，有则返回，时间O(nlogn)，空间O(1)
 */
func findRepeatNumber3(nums []int) int {
	sort.Ints(nums)
	pre := nums[0]
	for _, num := range nums[1:] {
		if pre == num {
			return pre
		}
		pre = num
	}
	return -1
}



func main() {
	fmt.Println(findRepeatNumber2([]int{2, 3, 1, 0, 2, 5, 3}))
}
