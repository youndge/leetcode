package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

func main() {
	k := 3
	num := "1432219"
	fmt.Println(removeKdigits(num, k))
	k = 1
	num = "10200"
	fmt.Println(removeKdigits(num, k))
	k = 2
	num = "10"
	fmt.Println(removeKdigits(num, k))
}
