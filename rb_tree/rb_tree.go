package rb_tree

// 红黑树
type RBTree[K any, E any] struct {
	root *node[K, E]
	size uint

	// 比较器，自定义比较e1大于e2返回1，相等返回0，e1小于e2返回-1
	Comparator func(e1, e2 K) int8
}

type NODE_COLOR bool

const (
	_COLOR_RED   NODE_COLOR = true  // 红色节点
	_COLOR_BLACK NODE_COLOR = false // 黑色节点
)

type node[K any, E any] struct {
	key   K
	e     E
	color NODE_COLOR
	left  *node[K, E]
	right *node[K, E]
}

func (t *RBTree[K, E]) Constructor(comparator func(e1, e2 K) int8) {
	t.root = nil
	t.size = 0
	t.Comparator = comparator
}

func (t *RBTree[K, E]) Add(key K, e E) {
	t.root = t.addNode(t.root, key, e)
	t.root.color = _COLOR_BLACK
}

// 向以root为根节点的树中添加元素
func (t *RBTree[K, E]) addNode(root *node[K, E], key K, e E) *node[K, E] {
	if root == nil {
		t.size++
		// 递归到底，将空子树置为节点
		// 根据红黑树的定义，最后添加的节点一定是红色的
		return &node[K, E]{key: key, e: e, color: _COLOR_RED}
	}

	r := t.Comparator(key, root.key)
	switch r {
	case 0: // 添加元素key等于当前节点key，覆盖节点值
		root.e = e
		return root

	case 1: // 添加元素大于当前节点
		root.right = t.addNode(root.right, key, e)

	case -1: // 添加元素小于当前节点
		root.left = t.addNode(root.left, key, e)
	}

	// 对照2-3树进行理解，红黑树等价于2-3树
	// 黑色节点对应2-3树的2节点
	// 红色节点表示该节点与它的父节点是结合的（3节点）
	// 添加后红黑树节点形状，对于每一次添加节后导致不符合红黑树性质，只可能存在以下情况
	// 依次监测，执行对应操作子过程

	if t.isRed(root.left) && t.isRed(root.right) {
		t.flipColors(root)
	}

	/*
	* eg.1
	*
	*  黑
	*    \
	*     红
	 */
	if t.isRed(root.right) && !t.isRed(root.left) {
		// 进行左旋转
		root = t.leftRotate(root)
	}

	if t.isRed(root.left) && t.isRed(root.left.right) {
		root.left = t.leftRotate(root.left)
	}

	/*
	* eg.2
	*
	*        黑
	*       /
	*     红
	*    /
	*   红
	 */
	if t.isRed(root.left) && t.isRed(root.left.left) {
		// 对当前节点进行右旋转
		// 转变为eg.3树形状

		root.color = _COLOR_RED
		root.left.color = _COLOR_BLACK

		root = t.rightRotate(root)
	}

	return root

}

// 该节点是否红节点
func (t *RBTree[K, E]) isRed(node *node[K, E]) bool {
	if node == nil {
		return false
	}
	return node.color == _COLOR_RED
}

/*
* 颜色翻转等价于2-3树中的分裂操作
*
*     黑           颜色翻转            红
*   /   \       ------------->      /   \
*  红     红                        黑    黑
 */
func (t *RBTree[K, E]) flipColors(node *node[K, E]) {
	node.color = _COLOR_RED
	node.left.color = _COLOR_BLACK
	node.right.color = _COLOR_BLACK
}

/*
*     node                                 x
*    /    \            左旋转             /  \
*   T1     x       ------------->      node   T3
*         / \                         /    \
*       T2   T3                      T1     T2
 */
func (t *RBTree[K, E]) leftRotate(node *node[K, E]) *node[K, E] {
	x := node.right
	node.right = x.left
	x.left = node
	return x
}

/*
*        node                                 x
*       /    \            右旋转             /  \
*      x      T2       ------------->      y   node
*    /  \                                      /   \
*   y     T1                                 T1     T2
 */
func (t *RBTree[K, E]) rightRotate(node *node[K, E]) *node[K, E] {
	x := node.left
	node.left = x.right
	x.right = node
	return x
}

func (t *RBTree[K, E]) Get(key K) (E, bool) {
	if t.root == nil {
		return *new(E), false
	}

	curRoot := t.root
	for {
		r := t.Comparator(key, curRoot.key)
		if r == 0 {
			return curRoot.e, true
		} else if r > 0 { // 当前添加元素大于
			curRoot = curRoot.right
		} else {
			curRoot = curRoot.left
		}
		if curRoot == nil {
			break
		}
	}
	return *new(E), false
}

func (t *RBTree[K, E]) Constants(key K) bool {
	_, exist := t.Get(key)
	return exist
}
