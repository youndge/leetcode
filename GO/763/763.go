package main

import "fmt"

/**
1.从左到右遍历字符串，遍历的同时维护开始下标start和结束下标end，初始start=end=0；
2.对于遍历到的每一个字符c，记录其最后一次出现的位置end_c,则当前片段的结束下标end一
定不小于end_c,有end=max(end,end_c);
3.当访问到下标end时，当前片段结束即[start,end]，长度等于end-start+1,然后下一个片
段由end+1开始，有start = end +1;
4.重复上述步骤，直到遍历结束；
*/
/*1.双指针+贪心算法*/
func partitionLabels(S string) []int {
	ans := []int{}
	lastPos := [26]int{}
	for i, b := range S {
		lastPos[b-'a'] = i
	}
	start, end := 0, 0
	for i, b := range S {
		if lastPos[b-'a'] > end {
			end = lastPos[b-'a']
		}
		//当访问到下标为end时，记录第一个片段，开始下一个
		if i == end {
			ans = append(ans, i-start+1)
			start = end + 1
		}
	}
	return ans
}

func main() {
	S := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(S))
}
