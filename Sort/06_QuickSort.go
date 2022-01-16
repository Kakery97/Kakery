package Sort

import "math/rand"

// 快速排序
//
// 采用分治法将一个序列分为较小和较大的2个子序列, 然后递归地排序两个子序列
//
// 平均时间复杂度  О(nlogn)
// 最坏时间复杂度  О(n²)
// 最优时间复杂度  О(nlogn)
// 额外空间复杂度  O(logn)
func quickSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	quickSortProcess(nums, 0, len(nums)-1)
}

// 使用递归实现的快速排序处理流程
// 使得数组nums的l~r这一范围内变得有序
func quickSortProcess(nums []int, l, r int) {
	if l < r {
		swap(nums, l+rand.Intn(r-l+1), r) // 从数组中随机挑出一个元素作为"基准"
		p := partition(nums, l, r)        // 分治-分割: 将所有比基准值小的元素摆放在基准前面, 所有比基准值大的元素摆在基准后面
		quickSortProcess(nums, l, p[0]-1) // 递归排序子序列: 递归地将小于基准值元素的子序列和大于基准值元素的子序列分别排序
		quickSortProcess(nums, p[1]+1, r)
	}
}

// 默认以nums[r]作为基准值, 将l~r这一范围依次划分为小于/等于/大于nums[r]的三个部分
// 将等于区域的左右两个边界下标作为数组返回
func partition(nums []int, l, r int) [2]int {
	less, more := l-1, r
	for l < more { // l表示当前遍历的位置, nums[r]作为基准值
		if nums[l] < nums[r] { // 当前位置 < 基准值:
			swap(nums, less+1, l) // 1.当前位置的数和小于区的下一个数做交换
			less++                // 2.小于区边界右移
			l++                   // 3.当前位置跳下一个
		} else if nums[l] == nums[r] { // 当前位置 = 基准值:
			l++ // 1.当前位置跳下一个
		} else { // 当前位置 > 基准值:
			swap(nums, more-1, l) // 1.当前位置的数和大于区的前一个数做交换
			more--                // 2.大于区边界左移
		}
	}
	swap(nums, more, r)
	return [2]int{less + 1, more}
}
