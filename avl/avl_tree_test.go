package avl

import (
	// "debug/macho"
	"math/rand"
	"testing"
	"time"
)

func TestAvl1(t *testing.T) {

	n := 1000000_0

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

	treeBst := AVLTree[int]{}
	treeBst.Constructor(func(e1, e2 int) int8 {
		if e1 == e2 {
			return 0
		} else if e1 > e2 {
			return 1
		} else {
			return -1
		}
	})

	var (
		startTime int64
		endTime int64
	)

	startTime = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		r := rand.Int()
		tree.Add(r)
		tree.Constants(r)
		tree.Constants(r+1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("cost：%v \n", endTime-startTime)
	t.Logf("is balanced：%v \n", tree.IsBalanced(tree.root))


	startTime = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		r := rand.Int()
		treeBst.Add(r)
		treeBst.Constants(r)
		treeBst.Constants(r+1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("cost：%v \n", endTime-startTime)


}
