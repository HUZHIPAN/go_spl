package no11

func maxArea(height []int) int {
	left,right := 0, len(height)-1
	max := 0
	for right >= left {
		leftHeight := height[left]
		rightHeight := height[right]

		width := right-left
		
		var capacity int
		if leftHeight > rightHeight {
			capacity = rightHeight * width
		} else {
			capacity = leftHeight * width
		}


		if leftHeight > rightHeight {
			right--
		} else {
			left++
		}

		if capacity > max {
			max = capacity
		}
	}
	return max
}