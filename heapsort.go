package dac

//时间O(nlogn) 空间O(1)
func HeapSort(nums []int) {
	heapSize := len(nums)
	buildMaxHead(nums, heapSize)
	for i := len(nums) - 1; i >= 0; i-- {
		//swap
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapfy(nums, 0, heapSize)
	}
}

func buildMinHead(nums []int, heapSize int) {
	for i := heapSize/2 - 1; i >= 0; i-- { //只需要从树的倒数第二层开始
		minHeapfy(nums, i, heapSize)
	}
}

func buildMaxHead(nums []int, heapSize int) {
	for i := heapSize/2 - 1; i >= 0; i-- { //只需要从树的倒数第二层开始
		maxHeapfy(nums, i, heapSize)
	}
}

func maxHeapfy(nums []int, i int, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i //左孩子，右孩子，根
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	//swap
	if i != largest {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapfy(nums, largest, heapSize)
	}
}

func minHeapfy(nums []int, i int, heapSize int) {
	l, r, minest := i*2+1, i*2+2, i //左孩子，右孩子，根
	if l < heapSize && nums[l] < nums[minest] {
		minest = l
	}
	if r < heapSize && nums[r] < nums[minest] {
		minest = r
	}
	//swap
	if i != minest {
		nums[i], nums[minest] = nums[minest], nums[i]
		minHeapfy(nums, minest, heapSize)
	}
}
