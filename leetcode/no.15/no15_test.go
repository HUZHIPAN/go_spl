package no15

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {

	res := insertSort([]int{4, 3, 9, 8})

	fmt.Println(res)

}
