package Sort

// 冒泡排序
//
// 重复地走访过要排序的数列, 一次比较两个元素, 如果它们的顺序错误就把它们交换过来
// 走访数列的工作是重复地进行直到没有再需要交换, 也就是说该数列已经排序完成
//
// 平均时间复杂度  О(n²)
// 最坏时间复杂度  О(n²)
// 最优时间复杂度  О(n²)
// 额外空间复杂度  O(1)
func bubbleSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	for e := len(nums) - 1; e > 0; e-- { // 0 ~ e
		for i := 0; i < e; i++ {
			if nums[i] > nums[i+1] {
				swap(nums, i, i+1)
			}
		}
	}
}
