package main

import "fmt"

func commonChars(A []string) []string {
	ans := []string{}
	strmap := map[byte]int{}
	for i := 'a'; i <= 'z'; i++ {
		strmap[byte(i)] = 0
	}
	var find func(string)
	find = func(s string) {
		for _, b := range s {
			strmap[byte(b)]++
		}
	}
	for _, str := range A {
		find(str)
	}
	for i := 'a'; i <= 'z'; i++ {
		if 0 < strmap[byte(i)] {
			for j := 0; j < strmap[byte(i)]/3 && strmap[byte(i)]/3 < 4; j++ {
				ans = append(ans, string(i))
			}
		}
	}
	return ans
}
func main() {
	s := []string{}
	s = append(s, "bella", "label", "roller")
	fmt.Println(commonChars(s))
	s1 := []string{}
	s1 = append(s1, "acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd")
	//["a","a","a","a","b","b","b","b","c","c","c","c","c","c","d","d","d","d","d","d"]
	fmt.Println(commonChars(s1))
	return
}
