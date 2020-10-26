package main

import (
	"fmt"
	"sort"
)

/*1.暴力*/
func smallerNumbersThanCurrent(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				cnt++
			}
		}
		ans = append(ans, cnt)
	}
	return ans
}

type pair struct{ v, pos int }

/*2.快速排序*/
func smallerNumbersThanCurrent2(nums []int) []int {
	n := len(nums)
	data := make([]pair, n)
	for i, v := range nums {
		data[i] = pair{v, i}
	}
	sort.Slice(data, func(i, j int) bool { return data[i].v < data[j].v })
	ans := make([]int, n)
	prev := -1
	for i, d := range data {
		if prev == -1 || d.v != data[i-1].v {
			prev = i
		}
		ans[d.pos] = prev
	}
	return ans
}

/*3.计数排序*/
func smallerNumbersThanCurrent3(nums []int) []int {
	cnt := [101]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for i := 0; i < 100; i++ {
		cnt[i+1] += cnt[i]
	}
	ans := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			ans[i] = cnt[v-1]
		}
	}
	return ans
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
	nums = []int{6, 5, 4, 8}
	fmt.Println(smallerNumbersThanCurrent2(nums))

}
