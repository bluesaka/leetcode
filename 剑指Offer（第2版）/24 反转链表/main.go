package main

import "fmt"

/**
方法一：迭代法，时间O(n)，空间O(1)
*/
func reverseList(head *ListNode) *ListNode {
	// 定义一个值来存储前节点
	var prev *ListNode
	for head != nil {
		// 当前节点的下一个节点
		next := head.Next
		// 当前节点的下一个节点变为前一个节点
		head.Next = prev
		// 前一个节点变为当前节点，准备后面的迭代
		prev = head
		// 当前节点变为下一个节点，开始后面的迭代
		head = next
	}
	return prev
}

/**
方法一：递归法，时间O(n)，空间O(n)
*/
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := sliceToListNode([]int{1, 2, 3, 4, 5})
	dumpListNode(reverseList(head))

	head2 := sliceToListNode([]int{11, 22, 33, 44, 55})
	dumpListNode(reverseList2(head2))
}

func sliceToListNode(slice []int) *ListNode {
	var head *ListNode
	for i := len(slice) - 1; i >= 0; i-- {
		head = &ListNode{Val: slice[i], Next: head}
	}
	return head
}

func dumpListNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
