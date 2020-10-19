package main

import "fmt"

/**
执行用时：48 ms, 在所有 Go 提交中击败了12.89%的用户
内存消耗：6.7 MB, 在所有 Go 提交中击败了28.39%的用户
*/
//1.自己写的递归
func search(nums []int, target int) int {
	var find func(int, int) int
	find = func(left, right int) int {
		if left == right {
			if target != nums[left] {
				return -1
			}
			return left
		}
		if left+1 == right {
			if target == nums[left] {
				return left
			} else if target == nums[right] {
				return right
			} else {
				return -1
			}
		}
		if target <= nums[(left+right)/2] {
			return find(left, (left+right)/2)
		} else {
			return find((left+right)/2, right)
		}
	}
	return find(0, len(nums)-1)
}
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
	//nums2 := []int{2, 3, 4, 4, 4, 7, 7, 8, 10, 10, 11, 12, 13, 14, 15, 15, 17, 18, 19, 23, 24, 24, 24, 24, 25, 26, 26, 26, 27, 27, 28, 29, 29, 30, 33, 36, 38, 38, 40, 40, 41, 43, 43, 43, 44, 46, 46, 47, 51, 52, 52, 53, 54, 56, 57, 57, 57, 58, 58, 61, 61, 61, 62, 64, 64, 66, 66, 67, 67, 67, 70, 72, 74, 74, 74, 75, 75, 78, 78, 78, 79, 79, 80, 83, 83, 83, 83, 84, 84, 86, 88, 89, 89, 90, 91, 91, 92, 93, 93, 96}
	//fmt.Println(upper_bound_(100, 1, nums2))
	nums3 := []int{5}
	fmt.Println(search(nums3, 5))
}
