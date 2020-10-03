package main

import "fmt"

//暴力
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				res := []int{i, j}
				return res
			}
		}
	}
	return nil
}

//哈希表
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if k1, ok := m[target-v]; ok {
			return []int{k, k1}
		}
		m[v] = k
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
}
