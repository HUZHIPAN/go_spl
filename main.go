package main

import (
	"data_struct/rb_tree"
	"data_struct/rb_tree/rb_tree2"
	"fmt"
	"math/rand"

	// "net/http"
	"runtime"
	"time"

	_ "net/http/pprof"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

func main() {

	// i := 1

	// i++
	// Test2()
	TestAvl1()
}

func Test2() {

	var (
		s1 int64
		s2 int64
	)

	n := 1000000000

	s1 = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		// if i == i {
		// }
	}
	s2 = (time.Now().UnixMicro() - s1) / 1000
	fmt.Printf("cost：%vms \n", s2)

	s1 = time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		// if Eq(&i, &i) {
		// }
		B()
	}
	s2 = (time.Now().UnixMicro() - s1) / 1000
	fmt.Printf("cost：%vms \n", s2)

}

func Eq(e1, e2 *int) bool {
	return *e1 == *e2
}
func B() {

}

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

func TestAvl1() {

	// go func() {
	// 	http.ListenAndServe("localhost:8080", nil)
	// }()

	PrintfMem()
	avlTree := rb_tree.AVLTree[int]{}
	avlTree.Constructor(Compare)

	treeBst := rb_tree.BstTree[int]{}
	treeBst.Constructor(Compare)

	rbTree := rb_tree.RBTree[int, struct{}]{}
	rbTree.Constructor(Compare)

	rbTree2 := rb_tree2.RBTree[int, struct{}]{}
	rbTree2.Constructor(Compare)

	rbt := rbt.NewWith(utils.IntComparator)

	hashMap1 := map[int]struct{}{}

	var (
		startTime int64
		endTime   int64

		s1 int64
		s2 int64
	)

	n := 10000000000
	var r int

	// PrintfMem()
	// dataQueue := make([]int, n)

	// startTime = time.Now().UnixMicro()
	// for i := 0; i < n; i++ {
	// 	r = rand.Int()
	// 	// r = i
	// 	avlTree.Add(r)

	// 	ok := avlTree.Constants(r)
	// 	if !ok {
	// 		panic("异常，avlTree， index:" + string(rune(r)))
	// 	}
	// 	avlTree.Constants(r + 1)
	// }
	// endTime = time.Now().UnixMicro()
	// fmt.Printf("AVLcost：%vms \n", (endTime-startTime)/1000)
	// fmt.Printf("is balanced：%v \n", tree.IsBalanced(tree.root))
	// fmt.Printf("is BST：%v \n", tree.IsBst(tree.root))

	// fmt.Printf("\n")

	// startTime = time.Now().UnixMicro()
	// for i := 0; i < n; i++ {
	// 	r = rand.Int()
	// 	r = i
	// 	treeBst.Add(r)
	// 	// treeBst.Constants(r)
	// 	// treeBst.Constants(r + 1)
	// }
	// endTime = time.Now().UnixMicro()
	// fmt.Printf("BST cost：%vms \n", (endTime-startTime)/1000)

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

	// startTime = time.Now().UnixMicro()
	// for i := 0; i < n; i++ {
	// 	r = rand.Int()
	// 	// r = i
	// 	rbTree.Add(r, struct{}{})
	// 	ok := rbTree.Constants(r)
	// 	if !ok {
	// 		panic("异常，rbTree， index:" + string(rune(r)))
	// 	}
	// 	rbTree.Constants(r + 1)
	// }
	// endTime = time.Now().UnixMicro()
	// fmt.Printf("TBtree cost：%vms \n", (endTime-startTime)/1000)

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

	// fmt.Printf("%v, %v", r, hashMap1)

}

// RangeRandom 返回指定范围内的随机整数。[min,max),不包含max
func RangeRandom(min, max int) (number int) {

	// return 150
	//创建随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	number = r.Intn(max-min) + min
	return number
}
