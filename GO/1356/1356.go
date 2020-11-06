package main

import (
	"fmt"
	"sort"
)

/*
执行用时：8 ms, 在所有 Go 提交中击败了94.05%的用户
内存消耗：4.4 MB, 在所有 Go 提交中击败了35.44%的用户
*/
func sortByBits(arr []int) []int {
	ans := []int{}
	mapBits := map[int][]int{}
	countBits := func(v int) {
		cnt := 0
		for i := uint(0); i < 32; i++ {
			if 1 == 1&(v>>i) {
				cnt++
			}
		}
		mapBits[cnt] = append(mapBits[cnt], v)
	}
	for _, v := range arr {
		countBits(v)
	}
	keys := []int{}
	for i := range mapBits {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, v := range keys {
		sort.Ints(mapBits[v])
		for _, v1 := range mapBits[v] {
			ans = append(ans, v1)
		}
	}
	return ans
}

//递推预处理
var bit = [1e4 + 1]int{}

func init() {
	for i := 1; i <= 1e4; i++ {
		bit[i] = bit[i>>1] + i&1
	}
}

func sortByBits2(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := bit[x], bit[y]
		return cx < cy || cx == cy && x < y
	})
	return a
}

func main() {
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// fmt.Println(sortByBits(arr))
	arr := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	fmt.Println(sortByBits(arr))

}
