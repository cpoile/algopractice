package lc_4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mid1, mid2 int
	N := len(nums1) + len(nums2)
	if N%2 == 0 {
		mid1, mid2 = N/2-1, N/2
	} else {
		mid1, mid2 = N/2, N/2
	}
	var cur1, cur2 int

	cur1 = len(nums1) / 2
	cur2 = len(nums2) / 2
	if nums1[cur1] > nums2[cur2] {

	}
}
