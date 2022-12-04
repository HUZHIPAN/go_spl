package heap

type PriorityQueue[T ~int | ~float64 | ~float32 | ~string] struct {
	maxHeap MaxHeap[T]
}

func (q *PriorityQueue[T]) Enqueue(e T) {
	q.maxHeap.Add(e)
}

// func (q *PriorityQueue[T]) Dequeue() T {
// 	q.maxHeap.ExtractMax()
// 	return 
// }