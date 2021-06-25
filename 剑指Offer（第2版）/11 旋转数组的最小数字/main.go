package main

import "fmt"

/**
方法一：迭代法，时间O(n)，空间O(1)
*/
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	min := numbers[0]
	for _, v := range numbers {
		if min > v {
			return v
		}
	}
	return min
}

/**
方法二：二分查找，分而治之
因为是半有序数组，所以最小值肯定在中间位置
*/
func minArray2(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right) / 2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] == numbers[right] {
			right = right - 1
		} else {
			right = mid
		}
	}
	return numbers[left]
}

func main() {
	fmt.Println(minArray([]int{33, 44, 55, 66, 77, 2, 3, 4, 5}))
	fmt.Println(minArray2([]int{2, 2, 2, 2, 2, 0, 2}))
}
