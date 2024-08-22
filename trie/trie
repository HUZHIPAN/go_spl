package trie

import "sync"

// Trie 前缀树实现
type Trie struct {
	root *trieNode
	size int64
	lock sync.RWMutex
}

type trieNode struct {
	value    rune
	children map[rune]*trieNode
	isEnd    bool
}

// NewTrie 创建trie
func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{
			children: nil,
			isEnd:    false,
		},
		lock: sync.RWMutex{},
		size: 0,
	}
}

// Insert 添加词
func (t *Trie) Insert(word string) {
	if word == "" {
		return
	}
	sliceList := []rune(word)

	//var parent = t.root
	var node = t.root

	t.lock.Lock()
	defer t.lock.Unlock()

	for _, e := range sliceList {
		//parent = node
		node = node.addChildren(e)
	}

	if !node.isEnd {
		node.isEnd = true
		t.size++
	}
}

// Delete 删除词
func (t *Trie) Delete(word string) bool {
	if word == "" {
		return false
	}
	sliceList := []rune(word)

	t.lock.Lock()
	defer t.lock.Unlock()

	var stack []*trieNode
	var node = t.root
	stack = append(stack, node)
	for _, e := range sliceList {
		child := node.children[e]
		if child == nil {
			return false
		}
		node = child
		stack = append(stack, node)
	}

	if node.isEnd {
		node.isEnd = false
		t.size--

		var tmpNode *trieNode
		var parent *trieNode

		parent = node
		stack = stack[:len(stack)-1]

		// 回溯，删除节点
		for len(stack) > 0 {
			tmpNode = parent
			parent = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(tmpNode.children) == 0 {
				if tmpNode.isEnd {
					delete(parent.children, tmpNode.value)
				} else {
					break
				}
			} else {
				break
			}
		}

		return true
	}

	return false
}

// RemoveAll 清空词库
func (t *Trie) RemoveAll() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.root = &trieNode{
		children: nil,
		isEnd:    false,
	}
	t.size = 0
}

func (n *trieNode) addChildren(e rune) *trieNode {
	if n.children == nil {
		n.children = make(map[rune]*trieNode)
	}

	if childrenNode, ok := n.children[e]; ok {
		return childrenNode
	} else {
		n.children[e] = &trieNode{
			value:    e,
			children: nil,
			isEnd:    false,
		}

		return n.children[e]
	}
}

// Contains 匹配字符串是否在包含前缀树中的词
// O(n+m)
func (t *Trie) Contains(word string) bool {
	if word == "" {
		return false
	}
	sliceList := []rune(word)
	node := t.root

	num := len(sliceList)

	t.lock.RLock()
	defer t.lock.RUnlock()

	st := 0  // 子串开始位置
	i := num // 当前子串最大循环次数
	for st < num {
		i = num - st

		var j int
		for j = st; j < st+i; j++ {
			if childrenNode, ok := node.children[sliceList[j]]; ok {
				node = childrenNode
				if node.isEnd {
					return true
				}
			} else {
				break
			}
		}

		if node.isEnd {
			return true
		}

		node = t.root
		st++
	}

	return false
}

// Match 匹配某个词是否在前缀树中
func (t *Trie) Match(word string) bool {
	if word == "" {
		return false
	}
	list := []rune(word)
	var node = t.root

	t.lock.RLock()
	defer t.lock.RUnlock()

	for _, v := range list {
		if childrenNode, ok := node.children[v]; ok {
			node = childrenNode
		} else {
			return false
		}
	}
	return node.isEnd
}
