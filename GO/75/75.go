package main

import (
	"fmt"
)

func sortColors(nums []int) {
	r, w := 0, 0
	for _, v := range nums {
		if 0 == v {
			r++
		} else if 1 == v {
			w++
		}
	}
	ans := []int{}
	for k, _ := range nums {
		if r > k {
			ans = append(ans, 0)
			continue
		}
		if r+w > k {
			ans = append(ans, 1)
			continue
		}
		ans = append(ans, 2)
	}
	copy(nums, ans)
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
	return
}
