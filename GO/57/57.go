package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	var min func(int, int) int
	var max func(int, int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	left, right, merged := newInterval[0], newInterval[1], false
	for _, interval := range intervals {
		if interval[0] > right {
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			ans = append(ans, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return ans
}
func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	fmt.Println(insert(intervals, newInterval))
}
