package main

import "fmt"

/**
方法一：迭代法，时间O(n)，空间O(1)
*/
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	cur := head
	for cur.Next != nil && cur.Next.Val != val {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	dumpListNode(deleteNode(sliceToListNode(nums), 2))
}

func sliceToListNode(nums []int) *ListNode {
	var head *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		head = &ListNode{Val: nums[i], Next: head}
	}
	return head
}

func dumpListNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
