package main

import (
	"fmt"
)

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	sz := make([]int, n)
	dp := make([]int, n)
	var dfs func(u, f int)
	dfs = func(u, f int) {
		sz[u] = 1
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			dfs(v, u)
			dp[u] += dp[v] + sz[v]
			sz[u] += sz[v]
		}
	}
	dfs(0, -1)

	ans := make([]int, n)
	var dfs2 func(u, f int)
	dfs2 = func(u, f int) {
		ans[u] = dp[u]
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			pu, pv := dp[u], dp[v]
			su, sv := sz[u], sz[v]

			dp[u] -= dp[v] + sz[v]
			sz[u] -= sz[v]
			dp[v] += dp[u] + sz[u]
			sz[v] += sz[u]

			dfs2(v, u)

			dp[u], dp[v] = pu, pv
			sz[u], sz[v] = su, sv
		}
	}
	dfs2(0, -1)
	return ans
}

func main() {
	N := 6
	e1 := []int{0, 1}
	e2 := []int{0, 2}
	e3 := []int{2, 3}
	e4 := []int{2, 4}
	e5 := []int{2, 5}

	edges := [][]int{}
	edges = append(edges, e1)
	edges = append(edges, e2)
	edges = append(edges, e3)
	edges = append(edges, e4)
	edges = append(edges, e5)

	fmt.Println(sumOfDistancesInTree(N, edges))
	return
}
