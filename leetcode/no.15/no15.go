package no15

import "sync"

func insertSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	seg := 1 // 划分，左边已排序，右边未排序
	for ; seg < len(nums); seg++ {
		j := seg - 1
		value := nums[seg]
		for ; j >= 0; j-- {
			if value < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}
	return nums
}

func threeSum(nums []int) [][]int {
	c := sync.Mutex{}

	c.Lock()

	sortNums := insertSort(nums)

	var (
		left  int
		right int
	)

	for i := 0; i < len(sortNums); i++ {

	}

	return nil
}
