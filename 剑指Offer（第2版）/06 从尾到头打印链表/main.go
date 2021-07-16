package main

import "fmt"

/**
 * 方法一：迭代法
 */
func reversePrint(head *ListNode) []int {
	var nums []int
	for head != nil {
		nums = append([]int{head.Val}, nums...)
		head = head.Next
	}
	return nums
}

/**
 * 方法二：递归法
 */
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint2(head.Next), head.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := sliceToListNode([]int{1, 3, 2})
	fmt.Println(reversePrint(head))
}

func sliceToListNode(nums []int) *ListNode {
	// [1, 3, 2]
	//node3 := &ListNode{Val: 2}
	//node2 := &ListNode{Val: 3, Next: node3}
	//node1 := &ListNode{Val: 1, Next: node2}
	var head *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		head = &ListNode{Val: nums[i], Next: head}
	}
	return head
}
