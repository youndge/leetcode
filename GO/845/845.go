package main

import "fmt"

func longestMountain(A []int) int {
	ans, n := 0, len(A)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if A[i-1] < A[i] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if A[i+1] < A[i] {
			right[i] = right[i+1] + 1
		}
	}
	for i, l := range left {
		r := right[i]
		if l > 0 && r > 0 && l+r+1 > ans {
			ans = l + r + 1
		}
	}
	return ans
}
func main() {
	arr := []int{2, 1, 4, 7, 3, 2, 5}
	fmt.Println(longestMountain(arr))
	arr = []int{2, 2, 2}
	fmt.Println(longestMountain(arr))

}
