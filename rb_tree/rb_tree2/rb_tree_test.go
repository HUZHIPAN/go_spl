package rb_tree2

import (
	// "debug/macho"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"

	"fmt"
	"math/rand"
	"runtime"
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
func Compare2(e1, e2 int) int {
	if e1 == e2 {
		return 0
	} else if e1 > e2 {
		return 1
	} else {
		return -1
	}
}
func TestAvl1(t *testing.T) {

	n := 10000

	var r int

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
	treeBst.Constructor(func(e1, e2 int) int8 {
		if e1 == e2 {
			return 0
		} else if e1 > e2 {
			return 1
		} else {
			return -1
		}
	})

	rbTree := RBTree[int, struct{}]{}
	rbTree.Constructor(func(e1, e2 int) int8 {
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
		endTime   int64
	)

	startTime = time.Now().UnixMicro()
	for i := 1; i < n; i++ {
		r = rand.Int()
		// r = i
		tree.Add(r)
		tree.Constants(r)
		tree.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("AVLcost：%vms \n", (endTime-startTime)/1000)
	t.Logf("is balanced：%v \n", tree.IsBalanced(tree.root))
	t.Logf("is BST：%v \n", tree.IsBst(tree.root))

	t.Logf("\n")

	// startTime = time.Now().UnixMicro()
	// for i := 0; i < n; i++ {
	// 	r = rand.Int()
	// 	// r = i
	// 	treeBst.Add(r)
	// 	treeBst.Constants(r)
	// 	treeBst.Constants(r + 1)
	// }
	// endTime = time.Now().UnixMicro()
	// t.Logf("BST cost：%vms \n", (endTime-startTime)/1000)

	startTime = time.Now().UnixMicro()
	for i := 1; i < n; i++ {
		r = rand.Int()
		// r = i
		rbTree.Add(r, struct{}{})
		rbTree.Constants(r)
		rbTree.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("TBtree cost：%vms \n", (endTime-startTime)/1000)

}

func TestRB1(t *testing.T) {

	n := 1000000
	var r int

	rbTree := RBTree[int, struct{}]{}
	rbTree.Constructor(func(e1, e2 int) int8 {
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
		endTime   int64
	)

	startTime = time.Now().UnixMicro()
	for i := 1; i < n; i++ {
		// r = rand.Int()
		r = i
		rbTree.Add(r, struct{}{})
		// rbTree.PrintTree(5)
		rbTree.Constants(r)
		rbTree.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("TBtree cost：%vms \n", (endTime-startTime)/1000)

}

func TestRBPrint(t *testing.T) {

	rbTree := RBTree[int, struct{}]{}
	rbTree.Constructor(func(e1, e2 int) int8 {
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
		endTime   int64
	)

	n := 50000
	var r int

	// testList := []int{9604, 3719, 7620, 2222, 9999}

	startTime = time.Now().UnixMicro()
	for i := 0; i <= n; i++ {
		// r = GenerateRandInt(0, 999999999)
		// r = testList[i]
		r = i
		// fmt.Printf("try add: %v \n", r)
		rbTree.Add(r, struct{}{})
		// rbTree.Add(-r, struct{}{})

		// fmt.Printf("size: %v \n", rbTree.size)

		// rbTree.PrintTree(6, rbTree.root)
		// rbTree.Constants(r)
		// rbTree.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()

	fmt.Printf("size: %v \n", rbTree.size)

	fmt.Printf("TBtree cost：%vms \n", (endTime-startTime)/1000)

	// defer func() {
	// 	fmt.Printf("size: %v", rbTree.size)
	// }()

	// rbTree.PrintTree(6, rbTree.root, nil)

}

func GenerateRandInt(min, max int) int {
	r := rand.Intn(9999-1111) + 1111
	rand.Seed(time.Now().Local().UnixNano() + rand.Int63() + int64(r)) //随机种子

	return rand.Intn(max-min) + min
}

func TestRB2(t *testing.T) {

	n := 10
	var r int

	rbTree := RBTree[int, struct{}]{}
	rbTree.Constructor(func(e1, e2 int) int8 {
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
		endTime   int64
	)

	startTime = time.Now().UnixMicro()
	for i := 1; i < n; i++ {
		// r = rand.Int()
		r = i
		rbTree.Add(r, struct{}{})
		// rbTree.PrintTree(5)
		rbTree.Constants(r)
		rbTree.Constants(r + 1)
	}
	endTime = time.Now().UnixMicro()
	t.Logf("TBtree cost：%vms \n", (endTime-startTime)/1000)

	rbTree.PrintTree(6, rbTree.root, nil)

	rbTree.Del(5)

	rbTree.PrintTree(6, rbTree.root, nil)
	rbTree.Del(2)
	rbTree.PrintTree(6, rbTree.root, nil)

	rbTree.Del(4)
	rbTree.PrintTree(6, rbTree.root, nil)

}

func TestRB3(t *testing.T) {

	var r int
	rbTree := RBTree[int, struct{}]{}
	rbTree.Constructor(func(e1, e2 int) int8 {
		if e1 == e2 {
			return 0
		} else if e1 > e2 {
			return 1
		} else {
			return -1
		}
	})

	rbTree2 := RBTree[int, struct{}]{}
	rbTree2.Constructor(func(e1, e2 int) int8 {
		if e1 == e2 {
			return 0
		} else if e1 > e2 {
			return 1
		} else {
			return -1
		}
	})

	hashMap1 := map[int]struct{}{}

	n := 20000
	for i := 0; i < n; i++ {
		r = GenerateRandInt(0, n)
		// r = i
		// r = n - (i + 1)
		rbTree.Add(r, struct{}{})
		rbTree2.Add(r, struct{}{})
		hashMap1[r] = struct{}{}

		msg := rbTree.IsRbTree(nil)
		if msg != "" {
			fmt.Println(msg)
		}
	}

	for j := 0; j < n; j++ {
		r = GenerateRandInt(0, n)
		// r = j
		// r = n - (j + 1)

		// fmt.Println(hashMap1)
		// rbTree.PrintTree(6, nil, nil)
		_, e1 := rbTree.Get(r)
		_, e2 := hashMap1[r]
		if e1 != e2 {
			// panic("异常")
			fmt.Println("一")
		}

		if rbTree.Constants(r) {

			fmt.Printf("删除%v \n", r)

			n1 := rbTree.findNode(r)
			// var n2 = n1
			if n1.parent != nil && n1.parent.parent != nil {
				// n2 = n1.parent.parent
			}
			// rbTree.PrintTree(6, n2, nil)
			rbTree.Del(r)
			delete(hashMap1, r)

			// rbTree.PrintTree(6, n2, nil)
			// rbTree.PrintTree(6, nil, nil)

			msg := rbTree.IsRbTree(nil)
			if msg != "" {
				fmt.Println(msg)
			}

			for k1, _ := range hashMap1 {
				e1 := rbTree.Constants(k1)
				if !e1 {
					fmt.Println("not t")
					rbTree2.Del(r)
				}
			}

		}

		if !rbTree.IsBst(rbTree.root) {
			fmt.Println("not bst")
		}

		msg := rbTree.IsRbTree(nil)
		_, e1 = rbTree.Get(r)
		_, e2 = hashMap1[r]
		if msg != "" || e1 != e2 {
			fmt.Println(msg)

			fmt.Println("异常")
			bugNode := rbTree2.findNode(r)
			rbTree2.PrintTree(6, bugNode.parent.parent, nil)

			rbTree2.Del(r)
			rbTree2.PrintTree(6, bugNode.parent.parent, nil)
		}

		rbTree2.Del(r)
	}

	rbTree.PrintTree(6, nil, nil)
	fmt.Printf("元素个数：%v -- %v \n", rbTree.size, len(hashMap1))

}

func TestRB4(t *testing.T) {

	rbTree := RBTree[int, int]{}
	rbTree.Constructor(func(e1, e2 int) int8 {
		if e1 == e2 {
			return 0
		} else if e1 > e2 {
			return 1
		} else {
			return -1
		}
	})

	rbTree2 := RBTree[int, int]{}
	rbTree2.Constructor(func(e1, e2 int) int8 {
		if e1 == e2 {
			return 0
		} else if e1 > e2 {
			return 1
		} else {
			return -1
		}
	})

	hashMap1 := map[int]int{}
	r := 0
	n := 2000000

	for i := 0; i < n; i++ {
		r = rand.Int()
		// r = i
		// r = n - (i + 1)
		r = GenerateRandInt(0, n)
		rbTree.Add(r, r)
		hashMap1[r] = r

		v1, e1 := rbTree.Get(r)
		v2, e2 := hashMap1[r]
		if v1 != v2 || e1 != e2 {
			panic("fff")
		}

		r = GenerateRandInt(0, n)
		rbTree.Del(r)
		delete(hashMap1, r)

		v1, e1 = rbTree.Get(r)
		v2, e2 = hashMap1[r]
		if v1 != v2 || e1 != e2 {
			panic("fff")
		}

		if rbTree.size != len(hashMap1) {
			panic("cc")
		}
	}

	// fmt.Println(hashMap1)
	// rbTree.PrintTree(6, nil, nil)

	fmt.Printf("元素个数：%v \n", rbTree.size)
}

func TestRB5(t *testing.T) {

	rbTree := RBTree[int, int]{}
	rbTree.Constructor(func(e1, e2 int) int8 {
		if e1 == e2 {
			return 0
		} else if e1 > e2 {
			return 1
		} else {
			return -1
		}
	})

	rbTree2 := RBTree[int, int]{}
	rbTree2.Constructor(func(e1, e2 int) int8 {
		if e1 == e2 {
			return 0
		} else if e1 > e2 {
			return 1
		} else {
			return -1
		}
	})

	hashMap1 := map[int]int{}
	r := 0
	n := 2000

	for i := 0; i < n; i++ {
		r = rand.Int()
		// r = i
		// r = n - (i + 1)
		r = GenerateRandInt(0, n)
		rbTree.Add(r, r)
		hashMap1[r] = r

		v1, e1 := rbTree.Get(r)
		v2, e2 := hashMap1[r]
		if v1 != v2 || e1 != e2 {
			panic("fff")
		}

		r = GenerateRandInt(0, n)
		rbTree.Del(r)
		delete(hashMap1, r)

		v1, e1 = rbTree.Get(r)
		v2, e2 = hashMap1[r]
		if v1 != v2 || e1 != e2 {
			panic("fff")
		}

		if rbTree.size != len(hashMap1) {
			panic("cc")
		}
	}

	fmt.Println(hashMap1)
	rbTree.PrintTree(6, nil, nil)

	fmt.Printf("元素个数：%v \n", rbTree.size)
}

func Test6(t *testing.T) {

	rbTree := RBTree[int, struct{}]{}
	rbTree.Constructor(Compare)

	// hashMap1 := map[int]struct{}{}

	var (
		startTime int64
		endTime   int64
	)

	n := 100000
	var r int

	startTime = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		r = RangeRandom(0, n)
		// r = i
		rbTree.Add(r, struct{}{})
		_, ok := rbTree.Get(r)
		if !ok {
			// panic("异常，rbTree2， index:" + string(rune(r)))
		}
		// rbTree2.Get(r + 1)
	}
	for i := 0; i < n; i++ {
		r = RangeRandom(0, n)
		// r = i
		rbTree.Del(r)
		_, ok := rbTree.Get(r)
		if !ok {
			// panic("异常，rbTree2， index:" + string(rune(r)))
		}
		// rbTree2.Get(r + 1)
	}
	endTime = time.Now().UnixMicro()
	fmt.Printf("TBtree cost：%vms \n", (endTime-startTime)/1000)

	// fmt.Printf("%v, %v", r, hashMap1)
}

// RangeRandom 返回指定范围内的随机整数。[min,max),不包含max
func RangeRandom(min, max int) (number int) {
	//创建随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	number = r.Intn(max-min) + min
	return number
}

func getMemStats() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m
}
func PrintfMem() {
	runtime.GC()

	m := getMemStats()
	fmt.Printf("Alloc %v; TotalAlloc %v; Sys %v; NumGC %v \n", m.Alloc, m.TotalAlloc, m.Sys, m.NumGC)
}

func TestDemo1(t *testing.T) {
	rbTree2 := RBTree[int, struct{}]{}
	rbTree2.Constructor(Compare)

	var (
		startTime int64
		endTime   int64

		s1 int64
		s2 int64
	)

	n := 100000000
	var r int

	startTime = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		// r = RangeRandom(0, n)
		r = i
		rbTree2.Add(r, struct{}{})
		_, _ = rbTree2.Get(r)
	}
	s1 = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		// r = RangeRandom(0, n)
		r = i
		rbTree2.Del(r)
		rbTree2.Get(r)
	}
	s2 = (time.Now().UnixMicro() - s1) / 1000
	endTime = time.Now().UnixMicro()
	fmt.Printf("TBtree2 cost：%vms, del %vms \n", (endTime-startTime)/1000, s2)

	PrintfMem()
}

func TestDemo2(t *testing.T) {
	hashMap1 := map[int]struct{}{}

	var (
		startTime int64
		endTime   int64

		s1 int64
		s2 int64
	)

	n := 100000000
	var r int

	startTime = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		// r = RangeRandom(0, n)
		r = i
		hashMap1[r] = struct{}{}
		_, _ = hashMap1[r]

		// _, _ = hashMap1[r+1]
	}
	s1 = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		// r = RangeRandom(0, n)
		r = i
		delete(hashMap1, r)
		_, _ = hashMap1[r]
	}
	s2 = (time.Now().UnixMicro() - s1) / 1000
	endTime = time.Now().UnixMicro()
	fmt.Printf("hashMap1 cost：%vms, del %vms \n", (endTime-startTime)/1000, s2)

	PrintfMem()
}

func TestDemo3(t *testing.T) {
	rbt := rbt.NewWith(utils.IntComparator)

	var (
		startTime int64
		endTime   int64

		s1 int64
		s2 int64
	)

	n := 100000000
	var r int

	startTime = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		// r = RangeRandom(0, n)
		r = i
		rbt.Put(r, struct{}{})
		_, _ = rbt.Get(r)
	}
	s1 = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		// r = RangeRandom(0, n)
		r = i
		rbt.Remove(r)
		rbt.Get(r)
	}
	s2 = (time.Now().UnixMicro() - s1) / 1000
	endTime = time.Now().UnixMicro()
	fmt.Printf("rbt cost：%vms, del %vms \n", (endTime-startTime)/1000, s2)

	PrintfMem()
}
