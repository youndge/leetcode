package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}

func main() {
	gas, cost := []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
