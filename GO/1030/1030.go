package main

import (
	"fmt"
	"sort"
)

//1.直接排序
func allCellsDistOrder(n, m, r0, c0 int) [][]int {
	ans := make([][]int, 0, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = append(ans, []int{i, j})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		a, b := ans[i], ans[j]
		return abs(a[0]-r0)+abs(a[1]-c0) < abs(b[0]-r0)+abs(b[1]-c0)
	})
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//2.桶排序
// func allCellsDistOrder(n, m, r0, c0 int) [][]int {
// 	maxDist := max(r0, n-1-r0) + max(c0, m-1-c0)
// 	buckets := make([][][]int, maxDist+1)

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			dist := abs(i-r0) + abs(j-c0)
// 			buckets[dist] = append(buckets[dist], []int{i, j})
// 		}
// 	}

// 	ans := make([][]int, 0, n*m)
// 	for _, bucket := range buckets {
// 		ans = append(ans, bucket...)
// 	}
// 	return ans
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

//3.几何法
// var dir4 = [][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

// func allCellsDistOrder(n, m, r0, c0 int) [][]int {
// 	ans := make([][]int, 1, n*m)
// 	ans[0] = []int{r0, c0}
// 	maxDist := max(r0, n-1-r0) + max(c0, m-1-c0)
// 	row, col := r0, c0
// 	for dist := 1; dist <= maxDist; dist++ {
// 		row--
// 		for i, dir := range dir4 {
// 			for i%2 == 0 && row != r0 || i%2 == 1 && col != c0 {
// 				if 0 <= row && row < n && 0 <= col && col < m {
// 					ans = append(ans, []int{row, col})
// 				}
// 				row += dir[0]
// 				col += dir[1]
// 			}
// 		}
// 	}
// 	return ans
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func main() {
	R, C, r0, c0 := 1, 2, 0, 0
	fmt.Println(allCellsDistOrder(R, C, r0, c0))
	R, C, r0, c0 = 2, 2, 0, 1
	fmt.Println(allCellsDistOrder(R, C, r0, c0))
}
