package main

import (
	"fmt"
	"sort"
)

/*
执行用时：60 ms, 在所有 Go 提交中击败了6.04%的用户
内存消耗：8.1 MB, 在所有 Go 提交中击败了5.52%的用户
*/
//自己写的
func sortArrayByParityII(A []int) []int {
	ans, arr := []int{}, []int{}
	i, j := 0, 1
	m := map[int]int{}
	for _, v := range A {
		if 0 == v%2 {
			m[i] = v
			i += 2
		} else {
			m[j] = v
			j += 2
		}
	}
	for k := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for _, v := range arr {
		ans = append(ans, m[v])
	}
	return ans
}

//双指针
func sortArrayByParityII2(a []int) []int {
	for i, j := 0, 1; i < len(a); i += 2 {
		if a[i]%2 == 1 {
			for a[j]%2 == 1 {
				j += 2
			}
			a[i], a[j] = a[j], a[i]
		}
	}
	return a
}

func main() {
	A := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(A))
}
