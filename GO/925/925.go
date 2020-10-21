package main

import "fmt"

//1.双指针
func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j-1] == typed[j] {
			j++
		} else {
			return false
		}
	}
	return i == len(name)
}
func main() {
	name, typed := "alex", "aaleex"
	fmt.Println(isLongPressedName(name, typed))
	name, typed = "saeed", "ssaaedd"
	fmt.Println(isLongPressedName(name, typed))
}
