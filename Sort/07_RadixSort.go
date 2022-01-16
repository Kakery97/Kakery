package Sort

import "math"

// 基数排序(仅适用于非负数)
//
// 基数排序是一种非比较型整数排序算法, 其原理是将整数按位数切割成不同的数字，然后按每个位数分别比较
// 将所有待比较数值(正整数)统一为同样的数位长度, 数位较短的数前面补零
// 然后, 从最低位开始, 依次进行一次排序, 这样从最低位排序一直到最高位排序完成以后, 数列就变成一个有序序列
//
// 最坏时间复杂度(k为数字位数)  О(kn)
// 额外空间复杂度(k为数字位数)  O(k)
func radixSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	bucket := make([]int, len(nums))
	for d := 1; d <= maxDigits(nums); d++ {
		var count [RADIX]int
		for _, num := range nums {
			count[getDigit(num, d)]++
		}
		for i := 1; i < RADIX; i++ {
			count[i] += count[i-1]
		}
		for i := len(nums) - 1; i >= 0; i-- {
			digit := getDigit(nums[i], d)
			bucket[count[digit]-1] = nums[i]
			count[digit]--
		}
		copy(nums, bucket)
	}
}

func maxDigits(nums []int) (res int) {
	maxNum := math.MinInt
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	for maxNum != 0 {
		res++
		maxNum /= RADIX
	}
	return
}

func getDigit(x, digit int) int {
	return (x / int(math.Pow(RADIX, float64(digit-1)))) % RADIX
}
