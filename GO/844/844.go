package main

import "fmt"

//1.栈
func backspaceCompare(S string, T string) bool {
	var recreate func(string) string
	recreate = func(str string) string {
		stack := []byte{}
		for _, b := range str {
			if '#' != byte(b) {
				stack = append(stack, byte(b))
			} else if 0 < len(stack) {
				stack = stack[:len(stack)-1]
			}
		}
		return string(stack)
	}
	return recreate(S) == recreate(T)
}

//2.双指针
func backspaceCompare2(S string, T string) bool {
	skipS, skipT := 0, 0
	i, j := len(S)-1, len(T)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if '#' == S[i] {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if '#' == T[j] {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if S[i] != T[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
func main() {
	S := "ab#c"
	T := "ab#c"
	fmt.Println(backspaceCompare(S, T))
	fmt.Println(backspaceCompare2(S, T))

	S = "ab##"
	T = "c#d#"
	fmt.Println(backspaceCompare(S, T))
	fmt.Println(backspaceCompare2(S, T))

}
