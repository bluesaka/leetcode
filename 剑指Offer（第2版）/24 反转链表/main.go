package main

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

}
