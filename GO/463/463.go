package main

import "fmt"

/*1.自己写的方法
执行用时：96 ms, 在所有 Go 提交中击败了12.88%的用户
内存消耗：7 MB, 在所有 Go 提交中击败了19.49%的用户
*/
func islandPerimeter(grid [][]int) int {
	sum := 0
	for i := range grid {
		for j := range grid[i] {
			if 1 != grid[i][j] {
				continue
			}
			cnt := 4
			if i >= 1 && 1 == grid[i-1][j] {
				cnt--
			}
			if i < len(grid)-1 && 1 == grid[i+1][j] {
				cnt--
			}
			if j >= 1 && 1 == grid[i][j-1] {
				cnt--
			}
			if j < len(grid[i])-1 && 1 == grid[i][j+1] {
				cnt--
			}
			fmt.Printf("grid[%d][%d]=%d\n", i, j, cnt)
			sum += cnt
		}
	}
	return sum
}

/*2.深度优先搜索*/
/*
执行用时：72 ms, 在所有 Go 提交中击败了83.33%的用户
内存消耗：7.6 MB, 在所有 Go 提交中击败了11.02%的用户
*/
type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter2(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			ans++
			return
		}
		if grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2
		for _, d := range dir4 {
			dfs(x+d.x, y+d.y)
		}
	}
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				dfs(i, j)
			}
		}
	}
	return
}

/*3.迭代*/
/*
执行用时：76 ms, 在所有 Go 提交中击败了64.39%的用户
内存消耗：7 MB, 在所有 Go 提交中击败了22.88%的用户
*/
// type pair struct{ x, y int }

// var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// func islandPerimeter3(grid [][]int) (ans int) {
// 	n, m := len(grid), len(grid[0])
// 	for i, row := range grid {
// 		for j, v := range row {
// 			if v == 1 {
// 				for _, d := range dir4 {
// 					if x, y := i+d.x, j+d.y; x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
// 						ans++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return
// }

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(islandPerimeter(grid))
	grid = [][]int{{0, 1}}
	fmt.Println(islandPerimeter(grid))

}
