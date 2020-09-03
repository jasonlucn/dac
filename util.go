package dac

func max(m int, n int) int {
	if m >= n {
		return m
	}
	return n
}

func swap(a, b interface{}) {
	a, b = b, a
}
