package sublist

type Relation string

// Sublist uses backtracking to see if l1 is a sublist of l2
func Sublist(l1, l2 []int) Relation {
	if equalSlice(l1, l2) {
		return "equal"
	}
	if recSublist(l1, l2) {
		return "sublist"
	}
	if recSublist(l2, l1) {
		return "superlist"
	}
	return "unequal"
}

func recSublist(l1, l2 []int) bool {
	if len(l1) == 0 || equalSlice(l1, l2) {
		return true
	}
	if len(l1) > len(l2) {
		return false
	}
	return recSublist(l1, l2[1:]) || recSublist(l1, l2[:len(l2)-1])
}

func equalSlice(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}
