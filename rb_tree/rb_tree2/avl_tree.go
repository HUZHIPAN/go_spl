package rb_tree2

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

// 返回当前节点树高
func (t *AVLTree[E]) getHeight(node *Node[E]) uint {
	if node == nil {
		return 0
	}

	return node.height
}

func (t *AVLTree[E]) Add(e E) {
	if t.root == nil {
		t.root = &Node[E]{e: e, height: 1}
		t.size++
		return
	}

	t.root = t.addNode(t.root, e)
}

// 向以root为根节点的树中添加元素
func (t *AVLTree[E]) addNode(root *Node[E], e E) *Node[E] {
	if root == nil {
		t.size++
		return &Node[E]{e: e, height: 1} // 递归到底，将空子树置为节点
	}

	r := t.Comparator(e, root.e)
	switch r {
	case 0: // 添加元素等于当前节点
		return root

	case 1: // 添加元素大于当前节点
		root.right = t.addNode(root.right, e)

	case -1: // 添加元素小于当前节点
		root.left = t.addNode(root.left, e)
	}

	root.height = getMax(t.getHeight(root.left), t.getHeight(root.right)) + 1

	bf := t.getBalanceFactor(root)
	if bf > 1 && t.getBalanceFactor(root.left) >= 0 { // LL
		// 左子树高度减右子树高度大于1，向左倾斜，需要向右旋转
		return t.rightRotate(root)
	}
	if bf < -1 && t.getBalanceFactor(root.right) <= 0 { // RR
		return t.leftRotate(root)
	}

	if bf > 1 && t.getBalanceFactor(root.left) < 0 { // LR
		root.left = t.leftRotate(root.left)
		return t.rightRotate(root)
	}

	if bf < -1 && t.getBalanceFactor(root.right) > 0 { // RL
		root.right = t.rightRotate(root.right)
		return t.leftRotate(root)
	}

	return root
}

func (t *AVLTree[E]) rightRotate(y *Node[E]) *Node[E] {
	x := y.left
	t3 := x.right

	x.right = y
	y.left = t3

	y.height = getMax(t.getHeight(y.left), t.getHeight(y.right)) + 1
	x.height = getMax(t.getHeight(x.left), t.getHeight(x.right)) + 1
	return x
}

func (t *AVLTree[E]) leftRotate(y *Node[E]) *Node[E] {
	x := y.right
	t2 := x.left

	x.left = y
	y.right = t2

	y.height = getMax(t.getHeight(y.left), t.getHeight(y.right)) + 1
	x.height = getMax(t.getHeight(x.left), t.getHeight(x.right)) + 1
	return x
}

func getMax(h1, h2 uint) uint {
	if h1 > h2 {
		return h1
	}
	return h2
}

// 获取节点平衡因子
func (t *AVLTree[E]) getBalanceFactor(node *Node[E]) int8 {
	if node == nil {
		return 0
	}
	return int8(t.getHeight(node.left) - t.getHeight(node.right))
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
