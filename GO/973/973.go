package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//1.排序
func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:k]
}

//2.优先队列（大根堆）
type pair struct {
	dist  int
	point []int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].dist > h[j].dist }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kClosest2(points [][]int, k int) (ans [][]int) {
	h := make(hp, k)
	for i, p := range points[:k] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}
	heap.Init(&h) // O(k) 初始化堆
	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0) // 效率比 pop 后 push 要快
		}
	}
	for _, p := range h {
		ans = append(ans, p.point)
	}
	return
}

//3.快速排序
// type pair struct {
// 	dist  int
// 	point []int
// }
// type hp []pair

// func (h hp) Len() int            { return len(h) }
// func (h hp) Less(i, j int) bool  { return h[i].dist > h[j].dist }
// func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
// func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// func kClosest3(points [][]int, k int) (ans [][]int) {
// 	h := make(hp, k)
// 	for i, p := range points[:k] {
// 		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
// 	}
// 	heap.Init(&h) // O(k) 初始化堆
// 	for _, p := range points[k:] {
// 		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].dist {
// 			h[0] = pair{dist, p}
// 			heap.Fix(&h, 0) // 效率比 pop 后 push 要快
// 		}
// 	}
// 	for _, p := range h {
// 		ans = append(ans, p.point)
// 	}
// 	return
// }

func main() {
	k, points := 1, [][]int{[]int{1, 3}, []int{-2, 2}}
	fmt.Println(kClosest(points, k))
	k, points = 2, [][]int{[]int{3, 3}, []int{5, -1}, []int{-2, 4}}
	fmt.Println(kClosest(points, k))
}
