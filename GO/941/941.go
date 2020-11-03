package main

import "fmt"

func validMountainArray(A []int) bool {
	top, end := 0, len(A)
	for ; top+1 < end && A[top] < A[top+1]; top++ {
	}
	if top == 0 || top == end-1 {
		return false
	}
	for ; top+1 < end && A[top] > A[top+1]; top++ {
	}
	return top == end-1
}
func main() {
	nums := []int{2, 1}
	fmt.Println(validMountainArray(nums))
	nums = []int{3, 5, 5}
	fmt.Println(validMountainArray(nums))
	nums = []int{0, 3, 2, 1}
	fmt.Println(validMountainArray(nums))
}
