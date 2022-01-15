package Sort

import "math"

// 在数组中交换指定下标的数字
func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

// 找出最大值
func max(nums ...int64) int64 {
	var maxNum int64 = math.MinInt64
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

// 找出最小值
func min(nums ...int64) int64 {
	var minNum int64 = math.MaxInt64
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

// 求和
func sum(nums ...int64) int64 {
	var sumNum int64 = 0
	for _, num := range nums {
		sumNum += num
	}
	return sumNum
}
