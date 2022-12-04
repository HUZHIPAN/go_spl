package union_find

type UnionFind struct {
	parent []int // 节点父节点
	height []int // 节点树高度
}

func Constructor(size int) UnionFind {
	uf := UnionFind{}
	uf.parent = make([]int, size)
	uf.height = make([]int, size)
	for i := 0; i < size; i++ {
		uf.parent[i] = i
		uf.height[i] = 1
	}
	return uf
}

func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.findRoot(p)
	qRoot := uf.findRoot(q)
	if pRoot == qRoot {
		return
	}
	pRootHeight := uf.height[pRoot]
	qRootHeight := uf.height[qRoot]
	if pRootHeight <= qRootHeight {
		uf.parent[pRoot] = qRoot
		uf.height[pRoot] += qRootHeight
	} else {
		uf.parent[qRoot] = pRoot
		uf.height[qRoot] += pRootHeight
	}
}

func (uf *UnionFind) IsConnected(p, q int) bool {
	return uf.findRoot(p) == uf.findRoot(q)
}

// 查找元素根节点
func (uf *UnionFind) findRoot(i int) int {
	parentIndex := uf.parent[i]
	if i == parentIndex {
		return i
	}
	return uf.findRoot(parentIndex)
}