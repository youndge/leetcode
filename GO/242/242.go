package main

import (
	"fmt"
	"sort"
)

//1.排序
func isAnagram(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	return string(s1) == string(s2)
}

//2.哈希表
func isAnagram2(s, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
	s, t = "rat", "car"
	fmt.Println(isAnagram(s, t))
}
