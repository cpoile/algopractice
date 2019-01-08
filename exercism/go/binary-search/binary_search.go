package binarysearch

func SearchInts(a []int, k int) int {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if k < a[mid] {
			hi = mid - 1
		} else if k > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
