package heap

import (
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	var heap1 MaxHeap[int] = MaxHeap[int]{}

	for i := 0; i < 1000000; i++ {
		heap1.Add(rand.Int())
	}

	sortData := []int{}
	for !heap1.IsEmpty() {
		sortData = append(sortData, heap1.ExtractMax())
	}
	
	for k := range sortData {
		if k == len(sortData)-1 {
			break
		}
		if sortData[k] < sortData[k+1] {
			panic("发生异常")
		}
	}

}
