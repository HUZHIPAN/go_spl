package sequential

import (
	"fmt"
	"testing"
)

type NumArray struct {
	st SegmentTree[int]
}

func Constructor(nums []int) NumArray {
	st := SegmentTree[int]{}
	st.InitSegmentTree(nums, func(leftVal, rightVal int) int { return leftVal + rightVal })
	return NumArray{
		st: st,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.st.Query(left, right)
}

func (this *NumArray) Update(index int, val int) {
	this.st.Set(index, val)
}

func TestSeq(t *testing.T) {

	obj := Constructor([]int{9, -8});
	v0 := obj.SumRange(0, 1)
	obj.Update(0, 3)
	v1 := obj.SumRange(1,1)
	v2 := obj.SumRange(0,1)
	obj.Update(1, -3)
	v3 := obj.SumRange(0,1)

	fmt.Println(v0,v1,v2,v3)
	
}