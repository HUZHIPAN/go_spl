package main


// 写一个冒泡排序算法函数
func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 写一个选择排序算法函数
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}


// 写一篇关于接入chatgpt的文章
// 1. 什么是chatgpt
// 2. chatgpt的原理
// 3. chatgpt的使用
// 4. chatgpt的优缺点