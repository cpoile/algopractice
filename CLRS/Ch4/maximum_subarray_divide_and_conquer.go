package Ch4

func maxSubarray(vals []int) (start, end, total int) {
	diffs := make([]int, 1, len(vals))
	for i := 1; i < len(vals); i++ {
		diffs = append(diffs, vals[i]-vals[i-1])
	}
	return maxSub(diffs, 0, len(diffs)-1)
}

func maxSub(diffs []int, lo, hi int) (start, end, total int) {
	if hi == lo {
		return lo, lo, diffs[lo]
	}
	mid := lo + (hi-lo)/2
	ll, lh, lt := maxSub(diffs, lo, mid)
	rl, rh, rt := maxSub(diffs, mid+1, hi)
	cl, ch, ct := maxCross(diffs, lo, mid, hi)
	if lt >= rt && lt >= ct {
		return ll, lh, lt
	} else if rt >= lt && rt >= ct {
		return rl, rh, rt
	} else {
		return cl, ch, ct
	}
}

func maxCross(diffs []int, lo, mid, hi int) (start, end, total int) {
	maxLeft := -1 << 31
	var curLeft, maxLIdx int
	for i := mid; i >= lo; i-- {
		curLeft += diffs[i]
		if curLeft > maxLeft {
			maxLeft, maxLIdx = curLeft, i
		}
	}
	maxRight := -1 << 31
	var curRight, maxRIdx int
	for i := mid + 1; i <= hi; i++ {
		curRight += diffs[i]
		if curRight > maxRight {
			maxRight, maxRIdx = curRight, i
		}
	}
	return maxLIdx, maxRIdx, maxLeft + maxRight
}
