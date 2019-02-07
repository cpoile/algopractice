package treesort

type tree struct {
	val         string
	left, right *tree
}

func treeSort(a []string) {
	var root *tree
	for _, s := range a {
		root = add(root, s)
	}
	appendTo(a[:0], root)
}

func add(x *tree, val string) *tree {
	if x == nil {
		return &tree{val: val}
	}
	if val <= x.val {
		x.left = add(x.left, val)
	} else {
		x.right = add(x.right, val)
	}
	return x
}

func appendTo(a []string, x *tree) []string {
	if x != nil {
		a = appendTo(a, x.left)
		a = append(a, x.val)
		a = appendTo(a, x.right)
	}
	return a
}
