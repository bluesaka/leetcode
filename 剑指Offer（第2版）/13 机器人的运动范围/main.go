package main

import (
	"container/list"
	"fmt"
)

/**
方法一：深度优先遍历DFS，时间O(mn)，空间O(mn)
*/
func movingCount(m int, n int, k int) int {
	visited := initVisited(m, n)
	return dfs(m, n, k, 0, 0, visited)
}

/**
方法二：广度优先搜索BFS，时间O(mn)，空间O(mn)
*/
func movingCount2(m int, n int, k int) int {
	visited := initVisited(m, n)
	count := 0
	queue := list.New()
	queue.PushBack([]int{0, 0})
	for queue.Len() > 0 {
		temp := queue.Remove(queue.Front()).([]int)
		i, j := temp[0], temp[1]
		si, sj := i/10+i%10, j/10+j%10
		if i >= m || j >= n || visited[i][j] || si+sj > k {
			continue
		}
		count++
		visited[i][j] = true
		queue.PushBack([]int{i + 1, j})
		queue.PushBack([]int{i, j + 1})
	}
	return count
}

func dfs(m, n, k, i, j int, visited [][]bool) int {
	if i >= m || j >= n || bitSum(i)+bitSum(j) > k || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(m, n, k, i+1, j, visited) + dfs(m, n, k, i, j+1, visited)
}

func initVisited(m, n int) [][]bool {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	return visited
}

func bitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(movingCount(3, 2, 1))
	fmt.Println(movingCount2(3, 2, 1))
}
