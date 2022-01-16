package Sort

// 插入排序
//
// 通过构建有序序列, 对于未排序数据
// 在已排序序列中从后向前扫描, 找到相应位置并插入
//
// 平均时间复杂度  О(n²)
// 最坏时间复杂度  О(n²)
// 最优时间复杂度  О(n)
// 额外空间复杂度  O(1)
func insertionSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	// 0 ~ 0已经有序, 0 ~ i想有序
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			swap(nums, j-1, j)
		}
	}
}
