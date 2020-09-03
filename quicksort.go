package dac

//平均Ologn,最快n^2, 空间O(1)
func QuickSort(values []int, left int, right int) {
	pivot := values[left]
	l, r := left, right
	forward := 1 //right->left=1, left->right=2
	for l < r {
		if forward == 1 { //right->left
			if values[r] < pivot {
				values[l] = values[r]
				l++
				forward = 2
			} else {
				r--
			}
		} else { //left->right
			if values[l] >= pivot {
				values[r] = values[l]
				r--
				forward = 1
			} else {
				l++
			}
		}
	}
	values[l] = pivot
	if l-left > 1 {
		QuickSort(values, left, l-1)
	}
	if right-l > 1 {
		QuickSort(values, l+1, right)
	}
}
