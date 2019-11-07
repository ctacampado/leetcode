package searchsort

// BSearchRange implements binary search for a sorted array of integers given range.
func BSearchRange(a []int, t, l, r int) int {
	m := 0
	for l <= r {
		m = (l + r) / 2
		if a[m] < t {
			l = m + 1
		} else if a[m] > t {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

// BSearch implements binary search for a sorted array of integers.
func BSearch(a []int, t int) int {
	r := len(a) - 1
	l, m := 0, 0
	for l <= r {
		m = (l + r) / 2
		if a[m] < t {
			l = m + 1
		} else if a[m] > t {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

// SortTwoSorted implements sorting two unmerged arrays of integer
func SortTwoSorted(a, b []int) []int {
	i, j, k := 0, 0, 0
	length := (len(a) + len(b)) - 1
	res := make([]int, length+1)

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		res[k] = a[i]
		k++
		i++
	}
	for j < len(b) {
		res[k] = b[j]
		k++
		j++
	}
	return res
}
