package main

import (
	"fmt"
	"sort"
)

//自己写的
func intersection(nums1 []int, nums2 []int) []int {
	ans := []int{}
	var isExist func(int, []int) bool
	isExist = func(val int, arr []int) bool {
		for _, v := range arr {
			if v == val {
				return true
			}
		}
		return false
	}
	for _, v := range nums1 {
		if true == isExist(v, nums2) && false == isExist(v, ans) {
			ans = append(ans, v)
		}
	}
	return ans
}

//两个集合
func intersection2(nums1 []int, nums2 []int) (intersection []int) {
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	for v := range set1 {
		if _, has := set2[v]; has {
			intersection = append(intersection, v)
		}
	}
	return
}

//排序+双指针
func intersection3(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
