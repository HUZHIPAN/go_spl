package sequential

import (
	"fmt"
	"testing"
)






/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */



func TestTrie(t *testing.T) {

	trie := GetTrie()

	trie.Insert("abc")
	r2 := trie.Search("abc")
	r3 := trie.StartsWith("ab")

	r7 := trie.SearchExpression("ab.")
	r4 := trie.SearchExpression(".b")
	r5 := trie.SearchExpression("...")
	r6 := trie.SearchExpression("ab.")


	fmt.Println(r2,r3,r4, r5,r6,r7)

	// trie.Add("abc")
	// trie.Add("123")
	// trie.Add("好呀")
	// trie.Add("好的134")
}