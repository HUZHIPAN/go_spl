package avl

// 平衡二叉树

type AVLTree[E any] struct {
	root       *Node[E]
	size       uint
	Comparator func(e1, e2 E) int8 // 比较器，e1大于e2返回1，相等返回0，e1小于e2返回-1
}
type Node[E any] struct {
	e      E
	height uint // 以当前元素为根节点的树高
	left   *Node[E]
	right  *Node[E]
}

func (t *AVLTree[E]) Constructor(comparator func(e1, e2 E) int8) {
	t.root = nil
	t.size = 0
	t.Comparator = comparator
}

// 获取节点平衡因子
func (t *AVLTree[E]) getBalanceFactor(node *Node[E]) int8 {
	if node == nil {
		return 0
	}
	return int8(t.getChildNodeHeight(node.left) - t.getChildNodeHeight(node.right))
}	

// 返回当前节点树高
func (t *AVLTree[E]) getChildNodeHeight(node *Node[E]) uint {
	if node == nil {
		return 0
	}
	var (
		leftChildHeight uint
		rightChildHeight uint
	)
	if node.left != nil {
		leftChildHeight = node.left.height
	} else {
		leftChildHeight = 0
	}
	if node.right != nil {
		rightChildHeight = node.right.height
	} else {
		rightChildHeight = 0
	}

	if leftChildHeight > rightChildHeight {
		return leftChildHeight
	} else {
		return rightChildHeight
	}
}

func (t *AVLTree[E]) Add(e E) {
	if t.root == nil {
		t.root = &Node[E]{e: e, height: 1}
		t.size++
		return
	}

	t.addNode(t.root, e)
}

// 向以root为根节点的树中添加元素
func (t *AVLTree[E]) addNode(root *Node[E], e E) {

	r := t.Comparator(e, root.e)
	switch r {
	case 0: // 添加元素等于当前节点
		return

	case 1: // 添加元素大于当前节点
		if root.right != nil {
			t.addNode(root.right, e)
		} else {
			root.right = &Node[E]{e: e, height: 1} // 递归到底，将空子树置为节点
			t.size++
			return
		}

	case -1: // 添加元素小于当前节点
		if root.left != nil {
			t.addNode(root.left, e)
		} else {
			root.right = &Node[E]{e: e, height: 1}
			t.size++
			return
		}
	}

	root.height = t.getChildNodeHeight(root) + 1
}

func (t *AVLTree[E]) Constants(e E) bool {
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


// 检查树平衡性质
func (t *AVLTree[E]) IsBalanced(root *Node[E]) bool {
	if root == nil {
		return true
	}
	
	bf := t.getBalanceFactor(root)
	if bf > 1 || bf < -1 {
		return false
	}
	
	return t.IsBalanced(root.left) && t.IsBalanced(root.right)
}

// 检查二分查找性质
func (t *AVLTree[E]) IsBst(root *Node[E]) bool {
	if root == nil {
		return true
	}

	if root.left != nil && t.Comparator(root.left.e, root.e) > 0 {
		return false
	} else if root.right != nil && t.Comparator(root.right.e, root.e) < 0 {
		return false
	}

	return t.IsBst(root.left) && t.IsBst(root.right)
}

