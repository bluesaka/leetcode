package main

import "fmt"

// D、L、R 分别代表遍历 根节点、左子数、右子树
// 前序遍历（DLR）：[根、[左]、[右]]，根在左子数前面，左子数在右子数前面
// 中序遍历（LDR）：[[左]、根、[右]]，左子数在根前面，根在右子数前面
// 后序遍历（LRD）：[[左]、[右]、根]，左子数在右子数前面，右子数在根前面

/**
方法一：递归法

复杂度：
- 时间复杂度：O(n)，其中 n 是树中的节点个数。
- 空间复杂度：O(n)，除去返回的答案需要的 O(n) 空间之外，我们还需要使用 O(n) 的空间存储哈希映射，
  以及 O(h)（其中 h 是树的高度）的空间表示递归时栈空间。这里 h < n，所以总空间复杂度为 O(n)。
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		// 找到根节点，因为题目描述所有数字不一样，所以相等则为根节点
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 左子树为前序遍历与中序遍历的[左]交集
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	// 右子树为前序遍历与中序遍历的[右]交集
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	preorder := []int{3, 9, 8, 5, 4, 10, 20, 15, 7}
	inorder := []int{4, 5, 8, 10, 9, 3, 15, 20, 7}
	dumpTreeWithPreorder(buildTree(preorder, inorder))
}

func dumpTreeWithPreorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	dumpTreeWithPreorder(root.Left)
	dumpTreeWithPreorder(root.Right)
}
