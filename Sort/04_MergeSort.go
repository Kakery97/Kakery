package Sort

// 归并排序
//
// 采用分治法将已有序的子序列合并，得到完全有序的序列
// 即先使每个子序列有序，再使子序列段间有序
//
// 平均时间复杂度  О(nlogn)
// 最坏时间复杂度  О(nlogn)
// 最优时间复杂度  О(nlogn)
// 额外空间复杂度  O(n)
func mergeSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	mergeSortProcess(nums, 0, len(nums)-1)
}

// 使用递归实现的归并排序主处理流程
// 使得数组nums的l~r这一范围内变得有序
func mergeSortProcess(nums []int, l, r int) {
	if l == r {
		return
	}
	m := l + ((r - l) >> 1)      // 位操作取中点
	mergeSortProcess(nums, l, m) // 分治-分割: 递归地把当前序列平均分割成两半
	mergeSortProcess(nums, m+1, r)
	merge(nums, l, m, r) // 分治-集成: 在保持元素顺序的同时将上一步得到的子序列集成到一起
}

// 归并操作: 合并nums数组中已经有序的两个部分l~m以及m+1~r
// 使得数组nums的l~r这一范围内整体变得有序
func merge(nums []int, l, m, r int) {
	pl, pr := l, m+1
	var tempArray []int
	for pl <= m && pr <= r {
		if nums[pl] <= nums[pr] {
			tempArray = append(tempArray, nums[pl])
			pl++
		} else {
			tempArray = append(tempArray, nums[pr])
			pr++
		}
	}
	for pl <= m {
		tempArray = append(tempArray, nums[pl])
		pl++
	}
	for pr <= r {
		tempArray = append(tempArray, nums[pr])
		pr++
	}
	for i, value := range tempArray {
		nums[l+i] = value
	}
}
