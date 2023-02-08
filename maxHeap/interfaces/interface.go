package interfaces

type MaxHeapInterface interface {
	Heapify([]int)     // to heapify an array
	Insert(data int)   // to insert a new element to heap
	Remove()           // to reomve the max value from heap (from top)
	Peek() int         // to get the max value from max heap (from top)
	DisplayAllvalues() // to display all values in the heap
}
