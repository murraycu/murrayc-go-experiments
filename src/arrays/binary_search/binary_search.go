package main

func binary_search_lo_hi(a []int, lo int, hi int, value int) (int, bool) {
	for lo <= hi {
		mid := lo + ((hi - lo) / 2)
		midval := a[mid]
		if midval == value {
			return mid, true
		} else if value < midval {
			hi = mid
		} else {
			lo = mid
		}
	}

	return -1, false
}

func binary_search(a []int, value int) (int, bool) {
	hi := len(a)
	if hi < 0 {
		return -1, false
	}

	return binary_search_lo_hi(a, 0, hi-1, value)
}
