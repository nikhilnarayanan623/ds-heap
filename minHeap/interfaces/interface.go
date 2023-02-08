package interfaces

type MinHeapinterface interface {
	Heapify([]int)   // to heapify an array
	Insert(data int) // to insert a new value to the heap
	Remove()         // to remove a value from heap (remove from top)
	Peek() int       // to get min value from the heap (its top of the heap)
	DisplayAllvalues()
}
