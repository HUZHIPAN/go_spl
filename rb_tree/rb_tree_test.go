package rb_tree

import (
	// "debug/macho"

	"math/rand"
	"testing"
	"time"
)

func TestAvl1(t *testing.T) {

	n := 100

	tree := AVLTree[int]{}
	tree.Constructor(func(e1, e2 int) int8 {
		if e1 == e2 {
			return 0
		} else if e1 > e2 {
			return 1
		} else {
			return -1
		}
	})

	treeBst := BstTree[int]{}
	treeBst.Constructor(Compare)

	rbTree := RBTree[int, struct{}]{}
	rbTree.Constructor(Compare)

	var (
		startTime int64
		endTime   int64
	)

	startTime = time.Now().UnixMicro()
	for i := 1; i < n; i++ {
		r := rand.Int()
		// r = i
		tree.Add(r)
		tree.Constants(r)
		tree.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("tree cost：%v \n", endTime-startTime)
	// t.Logf("is balanced：%v \n", tree.IsBalanced(tree.root))
	// t.Logf("is BST：%v \n", tree.IsBst(tree.root))

	// t.Logf("\n")

	startTime = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		r := rand.Int()
		// r := i
		treeBst.Add(r)
		treeBst.Constants(r)
		treeBst.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("treeBst cost：%v \n", endTime-startTime)
	// t.Logf("is balanced：%v \n", treeBst.IsBalanced(treeBst.root))

	startTime = time.Now().UnixMicro()
	for i := 1; i < n; i++ {
		r := rand.Int()
		// r = i
		rbTree.Add(r, struct{}{})
		rbTree.Constants(r)
		rbTree.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("rbTree cost：%v \n", endTime-startTime)

}
