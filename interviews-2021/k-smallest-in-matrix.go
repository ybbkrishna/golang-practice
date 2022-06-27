package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	mat := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	fmt.Println(kthSmallest(mat, 8))
}

func kthSmallest(matrix [][]int, k int) int {
	idx := 0
	prevValue := math.MaxInt64
	pq := make(PriorityQueue, len(matrix))
	for i := 0; i < len(matrix); i++ {
		pq[i] = &Item{i, 0, matrix[i][0], i}
	}
	heap.Init(&pq)

	for pq.Len() > 0 && idx < k {
		val := heap.Pop(&pq).(*Item)
		prevValue = val.Value
		idx++
		if val.YIdx < len(matrix)-1 {
			heap.Push(&pq, &Item{val.XIdx, val.YIdx + 1, matrix[val.XIdx][val.YIdx+1], -1})
		}
	}
	return prevValue
}

func PrettyPrint(x interface{}, err1 error) {
	val, err := json.MarshalIndent(x, " ", "  ")
	if err != nil {
		fmt.Println("Error :", err)
	}
	if err1 != nil {
		fmt.Println("Error1 :", err)
	}
	fmt.Println(string(val))
}

type Item struct {
	XIdx  int
	YIdx  int
	Value int

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Value < pq[j].Value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
