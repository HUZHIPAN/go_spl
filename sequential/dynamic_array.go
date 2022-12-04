package sequential


type ArraySlice struct {
	lenght int
	data   [2]int
}

func RangeBitwiseAnd(left int, right int) int {

	// f := int32(1 << 31)

	// fmt.Print(f)

	var inLeft int32 = int32(left)
	var inRight int32 = int32(right)
	var bitAnd int32

	var leftBitDiff int32 = 0
	for n := 31; n >= 0; n-- {
		isnEq := inLeft&(1<<n) == 0
		if isnEq {
			t := (1 << n)
			leftBitDiff += int32(t)
		}
	}

	var rightBitDiff int32 = 0
	for n := 0; n <= 31; n++ {
		isnEq := inRight&(1<<n) == 0
		if isnEq {
			rightBitDiff += 1 << n
		}
	}

	// return left

	inLeft -= leftBitDiff
	inRight -= rightBitDiff

	bitAnd = -1
	var i int32
	for i = inLeft; i <= inRight; i++ {
		if bitAnd == -1 {
			bitAnd = i
		}

		bitAnd = bitAnd & i
	}

	return int(bitAnd)
}

func search(nums []int, target int) int {

	var centerIndex int
	var arrLen int = len(nums)
	if arrLen%2 != 0 {
		if nums[arrLen-1] == target {
			return arrLen - 1
		}
		arrLen = arrLen - 1
	}

	// var loopNum = 0

	for {
		centerIndex = (arrLen / 2) - 1
		if nums[centerIndex] == target {
			return centerIndex
		}
		if target > nums[centerIndex] {

		}

	}

}

func isMatch(s1, s2 byte) bool {
	if s1 == '(' && s2 == ')' {
		return true
	}
	return false
}

func longestValidParentheses(s string) int {

	// var sLenght int = len(s)
	// var dp []int = make([]int, sLenght)

	// dp[sLenght-1] = 0

	// if isMatch(s[sLenght-1], s[sLenght-2]) {
	// 	dp[sLenght-2] = 2
	// } else {
	// 	dp[sLenght-2] = 0
	// }

	// for i := sLenght-3; i >= 0; i-- {

	// }

	// return 1

	// stack := []string{}
	// stackTopIndex
	for i := 0; i < len(s); i++ {

	}
	return 1
}

// 后一个结果依赖前一个步骤的决策
// 1、确定边界条件：只存在1或2个房子
// 2、决策：偷 = 向前跳两次偷到总数 + 本次偷得   不偷 = 上一次偷到总数   取最大
func rob(every []int) int {
	length := len(every)
	if length == 1 {
		return every[0]
	}
	if length == 2 {
		return Max(every[0], every[1])
	}

	pre2 := every[0]
	pre1 := Max(every[0], every[1])
	curr := pre1
	for i := 2; i < length; i++ {
		curr = Max(pre2+every[i], pre1)
		pre2 = pre1
		pre1 = curr
	}

	return curr
}

func Max(e1, e2 int) int {
	if e1 > e2 {
		return e1
	}
	return e2
}
func Min(e1, e2 int) int {
	if e1 < e2 {
		return e1
	}
	return e2
}
func Trap(height []int) int {
	lenght := len(height)

	var maxLeftDp = make([]int, lenght)
	var maxRightDp = make([]int, lenght)

	if lenght <= 2 {
		return 0
	}
	maxLeftDp[0] = 0
	maxLeftDp[1] = height[0]
	for i := 2; i < lenght; i++ {
		maxLeftDp[i] = Max(maxLeftDp[i-1], height[i-1]) // 左边最高
	}

	maxRightDp[lenght-1] = 0
	maxRightDp[lenght-2] = height[lenght-1]
	for i := lenght - 3; i >= 0; i-- {
		maxRightDp[i] = Max(maxRightDp[i+1], height[i+1]) // 右边最高
	}

	dp := make([]int, lenght)
	dp[0] = 0 // 第一个由于左边界为空，所以一定为0

	for i := 1; i < lenght-1; i++ {
		minHeight := Min(maxLeftDp[i], maxRightDp[i])
		// 左右边界最低长度减去当前 边高 = 可接雨水的单位
		if currUnit := minHeight - height[i]; currUnit > 0 {
			dp[i] = dp[i-1] + currUnit
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[lenght-2]
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
/func isBadVersion(version int) bool;
*/

func searchInsert(nums []int, target int) int {
	n := len(nums)

	left,right := 0,n-1

	for left <= right {
		ans := (right-left)/2 + left

		if nums[ans] == target {
			return ans
		} else if nums[ans] > target {
			right = ans-1
		} else {
			left = ans+1
		}
	}
	return left
}









func searchInsertV(nums []int, target int) int {
    n := len(nums)
    left, right := 0, n - 1
    ans := n
    for left <= right {
        mid := (right - left) >> 1 + left
        if target <= nums[mid] {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}


func sortedSquares(nums []int) []int {
	nums = []int{-9,-8,-1,0,1,2,3}

	n := len(nums)
	arr := make([]int, n)

	for i := 0; true; i++ {
		if nums[i] >= 0 {
			break
		}
		arr[i] = nums[i] * nums[i]
	}

	return arr
}