package Sort

// 堆排序
//
// 若以升序排序说明, 把数组转换成最大堆(也叫优先级队列)
// 这是一种满足最大堆性质的二叉树：对于除了根之外的每个节点i, A[parent(i)] ≥ A[i]
// 重复从最大堆取出数值最大的结点(把根结点和最后一个结点交换, 把交换后的最后一个结点移出堆)
// 并让残余的堆维持最大堆性质, 最终不断削减堆的数量使数组有序
//
// 平均时间复杂度  О(nlogn)
// 最坏时间复杂度  О(nlogn)
// 最优时间复杂度  О(nlogn)
// 额外空间复杂度  O(1)
func heapSort(nums []int) {
	if nums == nil || len(nums) < 2 {
		return
	}
	//for i := len(nums) - 1; i >= 0; i-- { // 第一步: 使用最大堆化, 时间复杂度О(n)
	//	maxHeapify(nums, i, len(nums))
	//}
	for i := 1; i < len(nums); i++ { // 第一步: 使用堆插入使数组堆化, 时间复杂度О(nlogn)
		heapInsert(nums, i)
	}
	for heapSize := len(nums) - 1; heapSize > 0; heapSize-- { // 第二步: 对堆化数据排序, 时间复杂度О(nlogn)
		swap(nums, 0, heapSize)
		maxHeapify(nums, 0, heapSize)
	}
}

// 堆插入
// 条件: 数组索引为index之前的数据已经是最大堆
func heapInsert(nums []int, index int) {
	for nums[index] > nums[(index-1)/2] {
		swap(nums, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// 最大堆化: 调整索引为index处的数据, 使其符合最大堆的特性
// index 需要堆化处理的数据的索引
// heapSize 未排序的堆(数组)的长度
func maxHeapify(nums []int, index, heapSize int) {
	left := 2*index + 1   // 左子节点索引
	for left < heapSize { // 左子节点索引超出计算范围, 直接返回
		biggest := left // 子节点值最大索引默认为左子节点
		if left+1 < heapSize && nums[left+1] > nums[left] {
			biggest = left + 1 // 判断左右子节点哪个较大
		}
		if nums[biggest] > nums[index] { // 若子节点最大值大于父节点, 则父节点被子节点调换
			swap(nums, index, biggest)
			index = biggest
			left = 2*index + 1
		} else {
			break // 子节点最大值不大于父节点, 直接返回
		}
	}
}
