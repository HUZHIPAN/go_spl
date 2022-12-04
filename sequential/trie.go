package sequential

type TrieNode struct {
	IsWord bool
	TreeMap map[byte]*TrieNode
}
func generateTrieNode() *TrieNode {
	tn := TrieNode{}
	tn.TreeMap = map[byte]*TrieNode{}
	return &tn
}

type Trie struct {
	size int
	root *TrieNode
}
func GetTrie() Trie {
	t := Trie{}
	t.root = generateTrieNode()
	t.size = 0
	return t
}


func (t *Trie) Insert(word string) {
	var curNode = t.root
	for i:= 0; i < len(word); i++ {
		c := word[i]
		if _,exist := curNode.TreeMap[c]; !exist {
			curNode.TreeMap[c] = generateTrieNode()
		}
		curNode = curNode.TreeMap[c]
	}

	if !curNode.IsWord {
		curNode.IsWord = true
		t.size++
	}
}

func (t *Trie) Search(word string) bool {
	var curNode = t.root
	for i:= 0; i < len(word); i++ {
		c := word[i]
		if _,exist := curNode.TreeMap[c]; !exist {
			return false
		}
		curNode = curNode.TreeMap[c]
	}
	return curNode.IsWord
}

func (t *Trie) SearchExpression(word string) bool {
	return t.searchWithNode(t.root, word)
}

func (t *Trie) searchWithNode(node *TrieNode, word string) bool {
	var curNode = node
	wordLenght := len(word)
	for i:= 0; i < wordLenght; i++ {
		c := word[i]
		if c == '.' {
			if i == wordLenght-1 {
				for _,tmp := range curNode.TreeMap {
					if tmp.IsWord {
						return true
					}
				}
				return false
			}
			for _,n := range curNode.TreeMap {
				if t.searchWithNode(n, word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			if _,exist := curNode.TreeMap[c]; !exist {
				return false
			}
			curNode = curNode.TreeMap[c]
		}
	}
	return curNode.IsWord
}


func (t *Trie) StartsWith(word string) bool {
	var curNode = t.root
	for i:= 0; i < len(word); i++ {
		c := word[i]
		if _,exist := curNode.TreeMap[c]; !exist {
			return false
		}
		curNode = curNode.TreeMap[c]
	}
	return true
}