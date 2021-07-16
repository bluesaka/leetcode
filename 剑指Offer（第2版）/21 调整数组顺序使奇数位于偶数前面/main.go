package main

import "fmt"

/**
方法一：首尾指针，时间O(n)，空间O(1)
首指针向后面搜索偶数，尾指针向前面搜索奇数，
两者都找的话进行交换 nums[left] 和 nums[right]
重复以上操作，直到left == right
*/
func exchange(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]&1 == 1 {
			left++
			continue
		}
		if nums[right]&1 == 0 {
			right--
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

/**
方法二：快慢指针
定义快指针fast，慢指针slow，fast在前，slow在后，
fast向后搜索奇数，slow指向下一个奇数存放的位置，当fast搜索到奇数时， nums[fast]和nums[slow]进行交换
重复上述操作，直到fast指向数组末尾
*/
func exchange2(nums []int) []int {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast]&1 == 1 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 3, 4, 2, 5, 7, 6}))
	fmt.Println(exchange2([]int{1, 3, 4, 2, 5, 7, 6}))
}
