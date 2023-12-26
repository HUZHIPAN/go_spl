package rb_tree

import (
	// "debug/macho"

	"math/rand"
	"testing"
	"time"
)

func Compare(e1, e2 int) int8 {
	if e1 == e2 {
		return 0
	} else if e1 > e2 {
		return 1
	} else {
		return -1
	}
}

func TestAvl1(t *testing.T) {

	n := 20000

	var r int

	tree := AVLTree[int]{}
	tree.Constructor(Compare)

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
		r = rand.Int()
		r = i
		tree.Add(r)
		tree.Constants(r)
		tree.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("AVLcost：%vms \n", (endTime-startTime)/1000)
	// t.Logf("is balanced：%v \n", tree.IsBalanced(tree.root))
	// t.Logf("is BST：%v \n", tree.IsBst(tree.root))

	// t.Logf("\n")

	startTime = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		r = rand.Int()
		r = i
		treeBst.Add(r)
		treeBst.Constants(r)
		treeBst.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("BST cost：%vms \n", (endTime-startTime)/1000)

	startTime = time.Now().UnixMicro()
	for i := 1; i < n; i++ {
		r = rand.Int()
		r = i
		rbTree.Add(r, struct{}{})
		rbTree.Constants(r)
		rbTree.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("TBtree cost：%vms \n", (endTime-startTime)/1000)

}
