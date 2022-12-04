package avl2

// 平衡二叉树

type BstTree[E any] struct {
	root       *BstNode[E]
	size       int
	Comparator func(e1, e2 E) int8 // 比较器，e1大于e2返回1，相等返回0，e1小于e2返回-1
}
type BstNode[E any] struct {
	e      E
	height int // 以当前元素为根节点的树高
	left   *BstNode[E]
	right  *BstNode[E]
}

func (t *BstTree[E]) Constructor(comparator func(e1, e2 E) int8) {
	t.root = nil
	t.size = 0
	t.Comparator = comparator
}

func (t *BstTree[E]) Add(e E) {
	if t.root == nil {
		t.root = &BstNode[E]{e: e, height: 1}
		return
	}

	parentRoot := t.root
	for {
		r := t.Comparator(e, parentRoot.e)
		if r == 0 {
			return
		}

		if r > 0 { // 当前添加元素大于当前遍历节点
			if parentRoot.right == nil {
				parentRoot.right = &BstNode[E]{e: e, height: 1}
				break
			} else {
				parentRoot = parentRoot.right
				continue
			}
		} else {
			if parentRoot.left == nil {
				parentRoot.left = &BstNode[E]{e: e, height: 1}
				break
			} else {
				parentRoot = parentRoot.left
				continue
			}
		}

	}
}

func (t *BstTree[E]) Constants(e E) bool {
	if t.root == nil {
		return false
	}

	parentRoot := t.root
	for {
		r := t.Comparator(e, parentRoot.e)
		if r == 0 {
			return true
		} else if r > 0 { // 当前添加元素大于
			parentRoot = parentRoot.right
		} else {
			parentRoot = parentRoot.left
		}
		if parentRoot == nil {
			break
		}
	}
	return false
}
