package Sort

// 选择排序
//
// 首先在未排序序列中找到最小(大)元素, 存放到排序序列的起始位置
// 然后, 再从剩余未排序元素中继续寻找最小(大)元素, 然后放到已排序序列的末尾
// 以此类推, 直到所有元素均排序完毕
//
// 平均时间复杂度  О(n²)
// 最坏时间复杂度  О(n²)
// 最优时间复杂度  О(n²)
// 额外空间复杂度  O(1)
func selectionSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums)-1; i++ { // i ~ N-1
		minIndex := i
		for j := i + 1; j < len(nums); j++ { // i ~ N上寻找最小下标
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		swap(nums, i, minIndex)
	}
}
