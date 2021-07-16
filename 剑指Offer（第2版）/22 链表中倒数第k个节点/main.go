package main

import "fmt"

/**
方法一，迭代法，时间O(n+k)，空间O(1)
两次迭代，第一次计算链表长度，第二次定位
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	l := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		l++
	}

	if k > l {
		return nil
	}
	n := 0
	for n != l-k && head != nil {
		head = head.Next
		n++
	}

	return head
}

/**
方法二：快慢双指针，可以不用计算链表长度，时间O(n)，空间O(1)
前指针先走k步，然后指针和后指针共同移动，直到前指针走过链表，此时后指针指向倒数第k个节点，返回后指针
*/
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	former, later := head, head

	// former先走k步
	for i := 0; i < k; i++ {
		former = former.Next
	}

	// former和later共同走，直到former走过链表
	for former != nil {
		former = former.Next
		later = later.Next
	}

	return later
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	dumpListNode(getKthFromEnd(sliceToListNode([]int{1, 2, 3, 4, 5}), 3))
	dumpListNode(getKthFromEnd2(sliceToListNode([]int{1, 2, 3, 4, 5}), 3))
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
