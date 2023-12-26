package rb_tree2

import (
	"fmt"
	"math"
)

// 红黑树性质
// 1、节点为黑色或者红色
// 2、根节点是黑色的
// 3、红色节点的两个子节点是黑色的（不能出现两个连续的红色节点）
// 4、每个叶子结点都是黑色的（此处的叶子结点指的是空结点，也被称为NIL节点）
// 5、任意一个节点到每个叶子节点的路径都包含相同数量的黑色节点

// 根据5条基本性质可以推出其他引申性质

// 红黑树
type RBTree[K any, E any] struct {
	root *node[K, E]
	size int

	// 比较器，自定义比较e1大于e2返回1，相等返回0，e1小于e2返回-1
	Comparator func(e1, e2 K) int8
}

type NODE_COLOR bool

const (
	_COLOR_RED   NODE_COLOR = true  // 红
	_COLOR_BLACK NODE_COLOR = false // 黑
)

type node[K any, E any] struct {
	color  NODE_COLOR
	key    K
	e      E
	left   *node[K, E]
	right  *node[K, E]
	parent *node[K, E]
}

func (t *RBTree[K, E]) Constructor(comparator func(e1, e2 K) int8) {
	t.root = nil
	t.size = 0
	t.Comparator = comparator
}

// Add 向红黑树中添加元素
func (t *RBTree[K, E]) Add(key K, e E) {

	// 当前根节点为空，直接插入节点作为根节点
	if t.root == nil {
		t.root = &node[K, E]{key: key, e: e, color: _COLOR_BLACK}
		t.size++
		return
	}

	// 从根节点开始查找需要插入到哪个位置
	var cur *node[K, E] = t.root

	for {
		r := t.Comparator(key, cur.key)

		// 当前元素存在，覆盖旧的值
		if r == 0 {
			cur.e = e
			return
		}

		// 插入元素大于当前迭代元素，继续向当前迭代元素的右子树插入
		if r > 0 {
			if cur.right == nil {

				cur.right = &node[K, E]{key: key, e: e, color: _COLOR_RED, parent: cur}
				t.size++

				// 规定插入节点为红色，当其父节点为红色的时候会破坏性质3（不能出现连续的红色节点），需要进行修复
				// 其父节点为黑色时，插入红色节点不会破坏红黑树性质，不需要进行处理
				if cur.color == _COLOR_RED {
					t.fixAfterInsert(cur.right)
				}
				return
			}
			cur = cur.right

			// 插入元素小于当前迭代元素，继续向当前迭代元素的左子树插入
		} else {
			if cur.left == nil {
				cur.left = &node[K, E]{key: key, e: e, color: _COLOR_RED, parent: cur}
				t.size++
				if cur.color == _COLOR_RED {
					t.fixAfterInsert(cur.left)
				}
				return
			}
			cur = cur.left
		}
	}

}

// 插入节点后修复
// 规定插入节点是红色的，父节点是红色时需要进行处理
// 父节点为黑色，插入红色节点不会破坏性质
func (t *RBTree[K, E]) fixAfterInsert(node *node[K, E]) {

	for t.isRed(node.parent) {

		// 叔叔节点是红色
		if t.isRed(t.getUncleNode(node)) {
			// 此时父节点和叔叔节点都是红色，根据红黑树性质3（不能有连续的红色节点），爷爷节点一定是黑色的
			// 将父节点和叔叔节点设为黑色，爷爷节点设为红色
			t.flipColors(node.parent.parent)

			// 将爷爷节点看作新插入的红节点，继续进行修复操作
			node = node.parent.parent

			// 叔叔节点是黑色
		} else {

			// 父节点在爷爷节点的左侧
			if node.parent == node.parent.parent.left {

				// 插入节点在其父节点的右侧，父节点在爷爷节点的左侧
				if node == node.parent.right {
					// 对插入节点的父节点进行左旋转（插入节点在右侧，通过旋转调整到左侧）
					node = node.parent
					t.leftRotate(node)
				}

				// 将父节点设为黑色，爷爷节点设为红色，再对爷爷节点进行右旋转
				// 可以论证爷爷节点的当前局部子树黑色节点并没有发生变化，依旧保持平衡
				node.parent.color = _COLOR_BLACK
				node.parent.parent.color = _COLOR_RED
				t.rightRotate(node.parent.parent)
				break

				// 父节点在爷爷节点的右侧
			} else {
				if node == node.parent.left {
					node = node.parent
					t.rightRotate(node)
				}

				node.parent.color = _COLOR_BLACK
				node.parent.parent.color = _COLOR_RED
				t.leftRotate(node.parent.parent)
				break

			}

		}

	}

	// 设置根节点为黑色
	t.root.color = _COLOR_BLACK
}

// Del 删除元素
// 删除元素成功返回true，不存在则返回false
func (t *RBTree[K, E]) Del(key K) bool {

	// 根据key查找删除节点
	node := t.findNode(key)
	if node == nil {
		return false
	}

	// 二叉树的删除，找到删除节点的替换节点，可将本次删除操作看作删除替换节点，而替代结点最后总是在树末
	// 替换节点只可能是 有一个左孩子的前驱节点 和 没有孩子节点的后继节点
	replaceNode := t.findReplaceNode(node)

	// 当前删除节点替换节点是根节点（当前红黑树中只剩下一个根节点）
	if replaceNode == t.root {
		t.root = nil
		t.size = 0
		return true
	}

	// （替换节点就是删除节点）删除节点为叶子节点
	if replaceNode == node {

		// 当删除的节点为黑色时，当前节点所在子树的黑色节点个数会减少一个
		// 黑色节点个数减少破环性质5（任意一个节点到每个叶子节点的路径都包含相同数量的黑色节点）
		// 需要进行旋转、变色等操作保持红黑树性质
		if replaceNode.color == _COLOR_BLACK {
			t.fixAfterDelete(replaceNode)
		}

		// 因为删除节点是叶子节点，所以不需要做额外操作
		// 判断删除节点是其父节点的左孩子还是右孩子，直接移除即可
		if replaceNode == replaceNode.parent.left {
			replaceNode.parent.left = nil
		} else {
			replaceNode.parent.right = nil
		}

	} else {
		// 当删除节点存在子节点作为替换节点时，需要其替换节点的值替换到删除节点
		node.key = replaceNode.key
		node.e = replaceNode.e

		if replaceNode.color == _COLOR_BLACK {
			t.fixAfterDelete(replaceNode)
		}

		// 替换节点在其父节点的左侧
		if replaceNode == replaceNode.parent.left {

			// 需要判断替换节点存在左孩子的情况，直接将左孩子放到替换节点的位置（完成移除替换节点）
			// 由于前面的逻辑，到此处替换节点只可能存在 有一个左孩子 和 没有孩子 两种情况
			if replaceNode.left != nil {
				replaceNode.left.parent = replaceNode.parent
				replaceNode.parent.left = replaceNode.left
			} else {
				// 替换节点没有孩子，并且替换节点在其父节点的左侧，直接将替换节点父节点的左孩子置空（完成移除替换节点）
				replaceNode.parent.left = nil
			}

			// 替换节点在其父节点的右侧
		} else {
			if replaceNode.left != nil {
				replaceNode.left.parent = replaceNode.parent
				replaceNode.parent.right = replaceNode.left
			} else {
				replaceNode.parent.right = nil
			}
		}
	}

	t.size--
	return true
}

// 当红黑树删除黑色节点时，需要进行旋转变色操作维护性质
func (t *RBTree[K, E]) fixAfterDelete(replaceNode *node[K, E]) {

	// replaceNode看作删除节点（只是假设删除来进行平衡维护，实际删除在自平衡完成后）
	// 删除黑色节点会导致所在子树路径上的黑色节点个数减1，所以需要和兄弟子树保持黑色节点数一致
	for replaceNode.color == _COLOR_BLACK && replaceNode != t.root {

		// 删除节点是其父节点的左孩子
		if replaceNode == replaceNode.parent.left {

			// 情况1 删除节点的兄弟节点是红色的
			// 根据红黑树性质3（红节点的两个子节点都是黑色），兄弟结点的父结点和子结点肯定为黑色，只存在这一种情况
			if replaceNode.parent.right.color == _COLOR_RED {

				// 将兄弟节点设为黑色，将父节点设为红色，对父节点进行左旋转；
				// 可以论证此时局部子树的黑色节点并不会发生改变
				// 转变为情况2.2，继续迭代处理
				replaceNode.parent.right.color = _COLOR_BLACK
				replaceNode.parent.color = _COLOR_RED
				t.leftRotate(replaceNode.parent)

			} else {
				// 情况2 删除节点的兄弟节点是黑色的
				// 父节点和子节点的颜色无法确定，所以需要分情况处理

				// 情况2.1 兄弟节点的右孩子是红色的
				if t.isRed(replaceNode.parent.right.right) {
					// 将兄弟节点设为父节点的颜色，父节点设为黑色，兄弟节点的右孩子设为黑色，对父节点进行左旋转；
					// 经过这一步处理，子树已经是平衡的（删除节点看作需要移除的），每棵子树都保持平衡状态，最终整棵树肯定是平衡的
					// 完成平衡，将这一步作为迭代处理的退出条件
					replaceNode.parent.right.color = replaceNode.parent.color
					replaceNode.parent.color = _COLOR_BLACK
					replaceNode.parent.right.right.color = _COLOR_BLACK
					t.leftRotate(replaceNode.parent)
					break

					// 情况2.2 兄弟节点的左孩子是红色的
					// 由于先判断了右孩子是红色的情况，所以到这里右孩子一定是黑色
				} else if t.isRed(replaceNode.parent.right.left) {
					// 将兄弟节点设为红色，兄弟节点的左孩子设为黑色，再对兄弟节点进行右旋转；
					// 转变为情况2.1，继续迭代处理
					replaceNode.parent.right.color = _COLOR_RED
					replaceNode.parent.right.left.color = _COLOR_BLACK
					t.rightRotate(replaceNode.parent.right)

					// 情况2.3 兄弟节点的孩子都是黑色的
				} else {
					// 将兄弟节点变成红色，由于自身是需要删除的，将兄弟节点染红，保持父节点左右子树黑色节点数一致
					// 由于父节点左右子树黑色都减1了，父节点的兄弟节点黑色节点会多一个
					// 此时将父节点看作需要删除节点，再进行删除修复平衡逻辑
					replaceNode.parent.right.color = _COLOR_RED
					replaceNode = replaceNode.parent
				}
			}

			// 删除节点是其父节点的右孩子
			// 跟删除节点是其父节点的左孩子处理完全是镜像的
		} else {
			if replaceNode.parent.left.color == _COLOR_RED {
				replaceNode.parent.left.color = _COLOR_BLACK
				replaceNode.parent.color = _COLOR_RED
				t.rightRotate(replaceNode.parent)
			} else {
				if t.isRed(replaceNode.parent.left.left) {
					replaceNode.parent.left.color = replaceNode.parent.color
					replaceNode.parent.color = _COLOR_BLACK
					replaceNode.parent.left.left.color = _COLOR_BLACK
					t.rightRotate(replaceNode.parent)
					break
				} else if t.isRed(replaceNode.parent.left.right) {
					replaceNode.parent.left.color = _COLOR_RED
					replaceNode.parent.left.right.color = _COLOR_BLACK
					t.leftRotate(replaceNode.parent.left)
				} else {
					replaceNode.parent.left.color = _COLOR_RED
					replaceNode = replaceNode.parent
				}

			}

		}

	}

	replaceNode.color = _COLOR_BLACK
}

// Get 查找一个元素
func (t *RBTree[K, E]) Get(key K) (E, bool) {
	e := t.findNode(key)
	if e == nil {
		return *new(E), false
	}
	return e.e, true
}

// Constants 是否包含某个元素
func (t *RBTree[K, E]) Constants(key K) bool {
	e := t.findNode(key)
	return e != nil
}

// 获取删除当前节点的替换节点
// 优先使用前驱节点，不存在前驱节点则使用后继节点
// 只可能返回 有一个左孩子的前驱节点 和 没有孩子节点的后继节点
// 没有前驱或后继节点时返回当前节点（当前节点为叶子节点）
func (t *RBTree[K, E]) findReplaceNode(root *node[K, E]) *node[K, E] {

	var cur *node[K, E]

	if root.left != nil {
		cur = root.left

		for cur.right != nil {
			cur = cur.right
		}

	} else if root.right != nil {
		cur = root.right

		for cur.right != nil {
			cur = cur.right
		}

	} else {
		return root
	}

	return cur
}

// 返回一个节点的叔叔节点
func (t *RBTree[K, E]) getUncleNode(node *node[K, E]) *node[K, E] {
	if node == nil || node.parent == nil || node.parent.parent == nil {
		return nil
	}

	if node.parent == node.parent.parent.left {
		return node.parent.parent.right
	} else {
		return node.parent.parent.left
	}
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
func (t *RBTree[K, E]) leftRotate(node *node[K, E]) {

	x := node.right
	x.parent = node.parent

	if node.parent == nil {
		t.root = x
	} else {
		if node == node.parent.left {
			node.parent.left = x
		} else {
			node.parent.right = x
		}
	}

	node.right = x.left
	if x.left != nil {
		x.left.parent = node
	}
	x.left = node
	node.parent = x
}

/*
*        node                                 x
*       /    \            右旋转             /  \
*      x      T2       ------------->      y   node
*    /  \                                      /   \
*   y     T1                                 T1     T2
 */
func (t *RBTree[K, E]) rightRotate(node *node[K, E]) {
	x := node.left
	x.parent = node.parent

	if node.parent == nil {
		t.root = x
	} else {
		if node == node.parent.left {
			node.parent.left = x
		} else {
			node.parent.right = x
		}
	}

	node.left = x.right
	if x.right != nil {
		x.right.parent = node
	}

	x.right = node
	node.parent = x
}

// 查找节点
func (t *RBTree[K, E]) findNode(key K) *node[K, E] {
	if t.root == nil {
		return nil
	}

	curRoot := t.root
	for {
		r := t.Comparator(key, curRoot.key)
		if r == 0 {
			return curRoot
		} else if r > 0 { // 当前添加元素大于
			curRoot = curRoot.right
		} else {
			curRoot = curRoot.left
		}
		if curRoot == nil {
			break
		}
	}
	return nil
}

func (t *RBTree[K, E]) PrintTree(height int, root *node[K, E], formatFunc func(e *node[K, E]) string) {

	height = height * 2
	var maxWidthSize = int(math.Pow(2, float64(height/2))) * 2

	var mapArr = [][]*node[K, E]{}
	for i := 0; i < height; i++ {
		tmp1 := []*node[K, E]{}
		for j := 0; j < maxWidthSize; j++ {
			tmp1 = append(tmp1, nil)
		}
		mapArr = append(mapArr, tmp1)
	}

	centerIndex := maxWidthSize/2 - 1

	var pRoot *node[K, E]
	if root != nil {
		pRoot = root
	} else {
		pRoot = t.root
	}
	if pRoot == nil {
		return
	}

	mapArr[0][centerIndex] = pRoot

	t.printTree(pRoot.left, 0, centerIndex, mapArr)
	t.printTree(pRoot.right, 0, centerIndex, mapArr)

	modeFmt := "%1v"

	for i := 0; i < height; i++ {
		for j := 0; j < maxWidthSize; j++ {
			e := mapArr[i][j]
			if e != nil {

				var showWord string
				if formatFunc != nil {
					showWord = formatFunc(e)
				} else {
					showWord = fmt.Sprintf("%v", e.key)
				}

				tmp1 := fmt.Sprintf(modeFmt, showWord)
				tCo := 30 // 黑色
				if e.color {
					tCo = 31 // 红色
				}

				fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, tCo, 47, tmp1, 0x1B)

			} else {
				s := ""
				// if i >= 1 && i < height-1 && j >= 1 && j < maxWidthSize-1 {
				// 	if mapArr[i+1][j-1] != nil && mapArr[i-1][j+1] != nil {
				// 		s = "/"
				// 	}
				// 	if mapArr[i+1][j+1] != nil && mapArr[i-1][j-1] != nil {
				// 		s = "\\"
				// 	}
				// } else {
				// 	s = ""
				// }

				fmt.Printf(modeFmt, s)
			}

		}
		fmt.Println("|") // 右边界
	}

	return

}

func (t *RBTree[K, E]) printTree(node *node[K, E], h, w int, dataMap [][]*node[K, E]) {
	if node == nil || node.parent == nil {
		return
	}

	var wOffect int

	level := (h/2 + 1) + 1

	if level > 2 {
		level += (level - 2)
	}

	levelOffset := len(dataMap[0]) / (level * 2)

	if node == node.parent.left {
		// wOffect = w - 2
		wOffect = w - levelOffset + 1
	} else {
		// wOffect = w + 2
		wOffect = w + levelOffset - 1

	}

	// fmt.Printf("%v -- %v \n", o, (w/2)+1)

	if wOffect >= len(dataMap[0]) || h+2 >= len(dataMap) || wOffect < 0 {
		fmt.Printf("当前层级过深，不打印子节点\n")
		return
	}

	dataMap[h+2][wOffect] = node

	t.printTree(node.left, h+2, wOffect, dataMap)
	t.printTree(node.right, h+2, wOffect, dataMap)

}

func (t *RBTree[K, E]) IsRbTree(node *node[K, E]) string {
	root := node
	if root == nil {
		root = t.root
	}
	if root == nil {
		return ""
	}

	if root.color != _COLOR_BLACK {
		msg := "违反性质2：根节点必须为黑色！"
		if msg != "" {
			return msg
		}

	}

	blackHeight := 0

	msg := t.isRbTree(root, 0, &blackHeight)

	if msg != "" {
		return msg
	}

	return ""
}

func (t *RBTree[K, E]) isRbTree(root *node[K, E], blackNum int, benchmark *int) string {

	if root == nil {
		if *benchmark == 0 {
			*benchmark = blackNum
			return ""
		} else if blackNum != *benchmark {
			return "违反性质4：从任意节点到每个叶子节点的所有路径都包含相同数目的黑色节点！"
		} else {
			return ""
		}
	}

	if root.color == _COLOR_RED && (t.isRed(root.left) || t.isRed(root.right)) {
		return "违反性质3：路径上有两个连续的红色节点！"
	}

	if root.color == _COLOR_BLACK {
		blackNum++
	}

	return t.isRbTree(root.left, blackNum, benchmark) + t.isRbTree(root.right, blackNum, benchmark)
}

// 检查二分查找性质
func (t *RBTree[K, E]) IsBst(root *node[K, E]) bool {
	if root == nil {
		return true
	}

	if root.left != nil && t.Comparator(root.left.key, root.key) > 0 {
		return false
	} else if root.right != nil && t.Comparator(root.right.key, root.key) < 0 {
		return false
	}

	return t.IsBst(root.left) && t.IsBst(root.right)
}

func (t *RBTree[K, E]) checkLevelRelation(cnode *node[K, E], eStack map[*node[K, E]]struct{}) string {

	r := ""
	if cnode == nil {
		return r
	}

	// isTop := false
	if eStack == nil {
		eStack = map[*node[K, E]]struct{}{}
		// isTop = true
		defer func() {
			if r != "" {
				t.PrintTree(6, cnode, nil)
			}

		}()
	}

	_, ok := eStack[cnode]

	if ok {
		r = fmt.Sprintf("当前节点%v的孩子节点包含循环引用", cnode.key)
		return r
	}

	eStack[cnode] = struct{}{}

	if cnode.left != nil {
		if cnode.left.parent != cnode {
			r = fmt.Sprintf("当前节点%v的左节点的父节点不是当前节点", cnode.key)
			return r
		} else {
			r = t.checkLevelRelation(cnode.left, eStack)
			if r != "" {
				return r
			}
		}
	}

	if cnode.right != nil {
		if cnode.right.parent != cnode {
			r = fmt.Sprintf("当前节点%v的右节点的父节点不是当前节点", cnode.key)
			return r
		} else {
			r = t.checkLevelRelation(cnode.right, eStack)
			if r != "" {
				return r
			}
		}
	}

	return r

}
