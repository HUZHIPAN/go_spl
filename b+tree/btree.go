package btree


type BTree[K, E any] struct {
	m int    // m阶btree
	size int
	root *Node[K,E]

	Comparator func(e1, e2 K) int8 // 比较器，e1大于e2返回1，相等返回0，e1小于e2返回-1
}

type Node[K,E any] struct {
	pointer []*Node[K,E]
	element  []*Element[K,E]
}


func (t *BTree[K, E]) insertElement(node *Node[K,E], key K, e E) *Node[K,E] {
	node.element = append(node.element, &Element[K, E]{key: key, e: e})
	var k int = len(node.element) -1
	for i := k; i > 0; i-- {
		if t.Comparator(node.element[i].key, node.element[i-1].key) < 0 {
			tmp := node.element[i-1]
			node.element[i-1] = node.element[i]
			node.element[i] = tmp
		}
	}
	return node
} 
type Element[K, E any] struct {
	key K
	e E
}

func (t *BTree[K, E]) Add(key K, e E) {
	if t.root == nil {
		t.root = &Node[K,E]{
			pointer: make([]*Node[K, E], 0),     // 每个节点至多m个分支指针
			element: make([]*Element[K, E], 0), // 每个节点至多m-1个元素
		}
		t.root.element = append(t.root.element, &Element[K,E]{key: key, e: e})
		return
	}

	t.root = t.insert(t.root, key, e)
}

func (t *BTree[K, E]) insert(node *Node[K, E], key K, e E) *Node[K,E] {

	if node.isLeafNode() { // 是叶子节点
		node.element = append(node.element, &Element[K, E]{key: key, e: e})
		var k int = len(node.element) -1
		for i := k; i > 0; i-- {
			if t.Comparator(node.element[i].key, node.element[i-1].key) < 0 {
				tmp := node.element[i-1]
				node.element[i-1] = node.element[i]
				node.element[i] = tmp
			}
		}
		return node
	} 

	var nextChild *Node[K,E]
	for i := 0; i < len(node.element); i++ {
		r := t.Comparator(key, node.element[i].key)
		if r < 0 {
			node.pointer[i] = t.insert(nextChild, key, e)
			break
		} else if r == 0 {
			node.element[i].e = e
			return node 
		}
	}
	
	node.pointer[len(node.element)] = t.insert(node.pointer[len(node.element)], key, e)




	return nil

	

	

	// r := t.Comparator(key, node.)

}


func (t *BTree[K, E]) spitNode(root *Node[K,E]) {
	// mid := t.m / 2
	return



}

// 是否叶子节点
func (n *Node[K, E]) isLeafNode() bool {
	return len(n.pointer) == 0
}