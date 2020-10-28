package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	times := []int{}
	m := map[int]int{}
	for _, v := range arr {
		m[v]++
	}
	var find func(int, []int) bool
	find = func(b int, a []int) bool {
		for _, v := range a {
			if b == v {
				return true
			}
		}
		return false
	}
	for _, v := range m {
		if true == find(v, times) {
			return false
		}
		times = append(times, v)
	}
	return true
}
func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
	arr = []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr))

}
