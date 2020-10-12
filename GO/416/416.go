package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	if 2 > len(nums) {
		return false
	}
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if 0 != sum%2 || max > sum/2 {
		return false
	}
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, sum/2+1)
	}
	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		for j := 1; j < sum/2+1; j++ {
			if j > v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][sum/2]
}

func main() {
	nums := []int{1, 5, 11, 5}
	nums2 := []int{1, 2, 3, 5}

	fmt.Println(canPartition(nums))
	fmt.Println(canPartition(nums2))
	return
}
