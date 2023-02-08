package main

import (
	"fmt"

	"github.com/nikhilnarayanan623/ds-heap/maxHeap"
)

func main() {

	heap := maxHeap.GetMaxHeap()

	heap.Insert(10)
	heap.Insert(1)
	heap.Insert(40)

	fmt.Println(" peek ", heap.Peek())

	heap.DisplayAllvalues()
	heap.Heapify([]int{3, 1, 890, 390, 3})
	heap.Remove()

	heap.DisplayAllvalues()
}
