package sequential

type MapSum struct {
	size int
	root *MapSumNode
}
type MapSumNode struct {
	IsWord bool
	Val int
	TreeMap map[byte]*MapSumNode
}
func generateMapSumNode() *MapSumNode {
	tn := MapSumNode{}
	tn.TreeMap = map[byte]*MapSumNode{}
	return &tn
}


func Constructor_() MapSum {
	m := MapSum{}
	m.root = generateMapSumNode()
	return m
}


func (this *MapSum) Insert(key string, val int)  {
	var curNode = this.root
	for i:= 0; i < len(key); i++ {
		c := key[i]
		_,exist := curNode.TreeMap[c]
		if !exist {
			curNode.TreeMap[c] = generateMapSumNode()
		}
		curNode = curNode.TreeMap[c]
	}

	curNode.Val = val
	if !curNode.IsWord {
		curNode.IsWord = true
		this.size++
	}
}

func (t *MapSum) Sum(prefix string) int {
	var curNode = t.root
	for i:= 0; i < len(prefix); i++ {
		c := prefix[i]
		if _,exist := curNode.TreeMap[c]; !exist {
			return 0
		}
		curNode = curNode.TreeMap[c]
	}
	return t.sumWithNode(curNode)
}
func (t *MapSum) sumWithNode(node *MapSumNode) int {
	sum := node.Val
	for _,n := range node.TreeMap {
		sum += t.sumWithNode(n)
	}
	return sum
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */