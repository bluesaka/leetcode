package main

import "fmt"

/**
方法一：DFS回溯，时间O(3^k mn)，空间O(k)，k为字符串长度
*/
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	visited := getVisited(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, visited, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, visited [][]bool, i, j, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] || visited[i][j] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	visited[i][j] = true
	res := dfs(board, word, visited, i+1, j, k+1) || dfs(board, word, visited, i-1, j, k+1) ||
		dfs(board, word, visited, i, j+1, k+1) || dfs(board, word, visited, i, j-1, k+1)
	if !res {
		visited[i][j] = false
	}
	return res
}

func getVisited(board [][]byte) [][]bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	return visited
}

func main() {
	test1()
}

func test1() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(word[0], len(word))
	fmt.Println(exist(board, word))
}

func test2() {
	board := [][]byte{
		{'a', 'a'},
	}
	word := "aa"
	fmt.Println(exist(board, word))
}

func test3() {
	board := [][]byte{
		{'a'},
	}
	word := "ab"
	fmt.Println(exist(board, word))
}

func test4() {
	board := [][]byte{
		{'a', 'a'},
	}
	word := "aaa"
	fmt.Println(exist(board, word))
}
