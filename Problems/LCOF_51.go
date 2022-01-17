package Problems

// 剑指 Offer 51. 数组中的逆序对
// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// 分治 + 归并排序
func reversePairs(nums []int) int {
	// 数组为空或长度为1, 直接返回0
	if nums == nil || len(nums) < 2 {
		return 0
	}

	return mergeSort(nums, 0, len(nums)-1)
}

// 使数组nums的l~r下标范围内有序, 使用递归实现
// 从大到小排序的同时统计逆序对的数量
func mergeSort(nums []int, l, r int) int {
	// 递归退出条件: 数组左右的范围重叠
	if l < r {
		// 计算中点将数组分割为左右两个部分
		m := l + ((r - l) >> 1)
		// 采用分治法将已有序的子数组合并
		return mergeSort(nums, l, m) + mergeSort(nums, m+1, r) + merge(nums, l, m, r)
	} else {
		return 0
	}
}

// 归并操作 - 合并nums数组中已经有序的两个部分l~m以及m+1~r
// 使得数组nums的l~r这一范围内整体变得有序
func merge(nums []int, l, m, r int) (res int) {
	// 分别定义指向左半部分和右半部分数组的指针
	p1, p2 := l, m+1
	var helpArray []int
	// 情况一: 当两个数组都有元素的时候, 遍历比较两边数组每个元素大小
	for p1 <= m && p2 <= r {
		// 取较大的元素加入到临时数组中, 并使指针右移
		if nums[p1] > nums[p2] {
			helpArray = append(helpArray, nums[p1])
			p1++
			// 若左侧数组数字较大, 则说明这个数字大于右边数组所有的元素, 属于逆序对的情况, 累计逆序对结果
			res += r - p2 + 1
		} else {
			helpArray = append(helpArray, nums[p2])
			p2++
		}
	}
	// 情况二: 当其中一个数组都有剩余元素, 另外一个为空的时候, 直接将剩余元素加入到临时数组后面
	// 因为此时都是较小的元素, 所以不存在逆序对的情况, 不用统计最终结果
	for p1 <= m {
		helpArray = append(helpArray, nums[p1])
		p1++
	}
	for p2 <= r {
		helpArray = append(helpArray, nums[p2])
		p2++
	}
	// 将零时数组中的元素覆盖nums旧数组中的元素, 使得整体有序
	for i, num := range helpArray {
		nums[l+i] = num
	}
	return
}
