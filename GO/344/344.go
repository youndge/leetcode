package main

import (
	"fmt"
)

func reverseString(s []byte) {
	h, t := 0, len(s)-1
	var rever func(int, int, []byte)
	rever = func(i, j int, s []byte) {
		if i >= j {
			return
		}
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
		i++
		j--
		rever(i, j, s)
	}
	rever(h, t, s)
}

func main() {
	nums := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(nums)
	reverseString(nums)
	fmt.Println(nums)
	return
}
