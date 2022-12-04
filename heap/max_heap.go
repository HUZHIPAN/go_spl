package heap

// 支持基础数据结构的最大堆实现
type MaxHeap[T ~int | ~float64 | ~float32 | ~string] struct {
	data []T
}

// 获取元素个数
func (h *MaxHeap[T]) GetSize() int {
	return len(h.data)
}
// 判断是否为空
func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.data) == 0
}
// 添加元素
func (h *MaxHeap[T]) Add(e T) {
	h.data = append(h.data, e)
	h.siftUp(h.GetSize()-1)
}
// 查看当前最大元素
func (h *MaxHeap[T]) PeekMax() T {
	maxE := h.data[0]
	return maxE
}
// 提取最大元素
func (h *MaxHeap[T]) ExtractMax() T {
	maxE := h.PeekMax()
	h.swap(0, h.GetSize()-1)
	h.data = h.data[:h.GetSize()-1] // 删除最后一个元素
	h.siftDown(0)
	return maxE
}

// 元素上浮
func (h *MaxHeap[T]) siftUp(currentIndex int) {
	for currentIndex > 0 {
		parentIndex := h.getParentIndex(currentIndex) // 获取父亲节点下标
		if h.data[currentIndex] > h.data[parentIndex] {
			h.swap(currentIndex, parentIndex)
			currentIndex = parentIndex
		} else {
			break
		}
	}
}

// 元素下浮
func (h *MaxHeap[T]) siftDown(currentIndex int) {
	nodeMaxIndex := h.GetSize()-1
	for {
		childIndex := h.getLeftIndex(currentIndex)
		if childIndex > nodeMaxIndex {
			break
		}

		// 存在右节点，并且右节点 > 左节点
		if childIndex+1 <= nodeMaxIndex && h.data[childIndex+1] > h.data[childIndex] {
			childIndex = childIndex+1 
		}

		if h.data[currentIndex] < h.data[childIndex] {
			h.swap(currentIndex, childIndex)
			currentIndex = childIndex
			continue
		}

		break
	}	
}

// 交换两个元素
func (h *MaxHeap[T]) swap(i, j int) {
	tmpE := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = tmpE
}


func (h *MaxHeap[T]) getParentIndex(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap[T]) getLeftIndex(index int) int {
	return index*2 + 1
}

// func (h *MaxHeap[T]) getRightIndex(index int) int {
// 	return index*2 + 2
// }
