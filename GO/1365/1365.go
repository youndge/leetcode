package main

import "fmt"

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
func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
	nums = []int{6, 5, 4, 8}
	fmt.Println(smallerNumbersThanCurrent(nums))

}
