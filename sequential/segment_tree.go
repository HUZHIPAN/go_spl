package sequential

import (
	"fmt"
)



type treeNode[E any] struct{
	l int // 左边界
	r int // 右边界
	e E // 节点
}

// 线段树
type SegmentTree[E any] struct {
	data []E
	tree []treeNode[E]

	merger func(leftVal, rightVal E) E // 合并方法
}

func (st *SegmentTree[E]) InitSegmentTree(data []E, merger func(leftVal, rightVal E) E) {
	st.data = data
	st.tree = make([]treeNode[E], st.GetSize()*4)
	st.merger = merger
	st.buildSegmentTree(0, 0, st.GetSize()-1)
}

// 生成节点index 范围l...r范围线段树节点值
func (st *SegmentTree[E]) buildSegmentTree(index, l, r int) {
	if l == r { // 当节点为叶子节点
		st.tree[index].l = l
		st.tree[index].r = r
		st.tree[index].e = st.data[l]
		return
	}

	leftChildIndex := st.getLeftIndex(index)
	rightChildIndex := st.getRightIndex(index)

	mid := l + (r-l)/2
	st.buildSegmentTree(leftChildIndex, l, mid)
	st.buildSegmentTree(rightChildIndex, mid+1, r)

	st.tree[index].l = l
	st.tree[index].r = r
	st.tree[index].e = st.merger(st.tree[leftChildIndex].e, st.tree[rightChildIndex].e)
}

// 将index的值替换为e
func (st *SegmentTree[E]) Set(index int, e E) {
	st.data[index] = e
	st.setIndex(0, 0, st.GetSize()-1, index,e)
}

// 更新index下 区间l...r的值
func (st *SegmentTree[E]) setIndex(treeIndex, l, r int, index int,e E) {
	if l == r { // 当节点为叶子节点
		st.tree[treeIndex].e = st.data[l]
		return
	}

	leftChildIndex := st.getLeftIndex(treeIndex)
	rightChildIndex := st.getRightIndex(treeIndex)
	mid := l + (r-l)/2

	if index >= mid+1 { // right
		st.setIndex(rightChildIndex, mid+1, r, index,e)
	} else { // left
		st.setIndex(leftChildIndex, l, mid, index,e)
	}

	st.tree[treeIndex].e = st.merger(st.tree[leftChildIndex].e, st.tree[rightChildIndex].e)
}

func (st *SegmentTree[E]) Query(queryL, queryR int) E {
	if queryL < 0 || queryR >= st.GetSize() || queryR < queryL {
		panic(fmt.Sprintf("边界错误：%d -- %d", queryL, queryR))
	}
	return st.queryBetween(0, queryL, queryR)
}

// 在线段树节点index下查找 区间queryL到queryR
func (st *SegmentTree[E]) queryBetween(index , queryL, queryR int) E {
	l := st.tree[index].l
	r := st.tree[index].r
	if queryL == l && queryR == r {
		return st.tree[index].e
	}

	leftChildIndex := st.getLeftIndex(index)
	rightChildIndex := st.getRightIndex(index)
	mid := l + (r-l)/2

	if queryL >= mid+1 {
		return st.queryBetween(rightChildIndex, queryL, queryR)
	} else if queryR <= mid {
		return st.queryBetween(leftChildIndex, queryL, queryR)
	}

	leftResult := st.queryBetween(leftChildIndex, queryL, mid)
	rightResult :=  st.queryBetween(rightChildIndex, mid+1, queryR)

	return st.merger(leftResult, rightResult)
}

func (st *SegmentTree[E]) GetSize() int {
	return len(st.data)
}

func (st *SegmentTree[E]) getLeftIndex(index int) int {
	return index*2 + 1
}

func (st *SegmentTree[E]) getRightIndex(index int) int {
	return index*2 + 2
}
