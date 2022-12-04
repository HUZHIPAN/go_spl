package sequential

import (
	"fmt"
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	RangeBitwiseAnd(1, 10000)

}

func TestTrap(t *testing.T) {
	Trap([]int{4, 2, 0, 3, 2, 5})
	Min(1, 1)
}

func TestSearchInsert(t *testing.T) {
	searchInsert([]int{1, 3, 5, 6}, 0)
}

func TestSegmentTree(t *testing.T) {
	st := SegmentTree[int]{}

	st.InitSegmentTree([]int{-2,0,3,-5,2,-1}, func(leftVal, rightVal int) int { return leftVal + rightVal })

	val := st.Query(0, 1)
	val2 := st.Query(2,3)
	val3 := st.Query(0,5)

	st.Set(1, 9)

	val4 := st.Query(0, 5)

	fmt.Println(val, val2, val3, val4)

}


