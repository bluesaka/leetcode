package main

import "fmt"

/**
方法一：先比较再遍历，时间O(mn)，空间O(mn)
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, v := range matrix {
		if len(v) == 0 {
			break
		}
		if target < v[0] {
			break
		}
		if target > v[len(v)-1] {
			continue
		}
		for _, vv := range v {
			if vv == target {
				return true
			}
		}
	}
	return false
}

/**
（推荐）方法二：Binary Search Tree 二叉搜索树，时间O(m+n)，空间O(m+n)
从右上角看，这个矩阵就像是一个二叉搜索树

  15
 /   \
11	  19
...

所以从右上角开始找：
- 如果target比当前位置小，则col--
- 如果target比当前位置大，则row++
- 如果相等，则返回true
*/
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	row := 0
	col := n - 1
	for row < m && col >= 0 {
		if target < matrix[row][col] {
			col--
		} else if target > matrix[row][col] {
			row++
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray2(matrix, 5))
}
