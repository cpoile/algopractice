package binarysearchtree

type SearchTreeData struct {
	left, right *SearchTreeData
	data        int
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{nil, nil, data}
}

func (bst *SearchTreeData) Insert(data int) {
	if data <= bst.data {
		if bst.left == nil {
			bst.left = Bst(data)
		} else {
			bst.left.Insert(data)
		}
	} else {
		if bst.right == nil {
			bst.right = Bst(data)
		} else {
			bst.right.Insert(data)
		}
	}
}

func (bst *SearchTreeData) MapString(f func(int) string) []string {
	if bst == nil {
		return []string{}
	}
	return append(append(bst.left.MapString(f), f(bst.data)), bst.right.MapString(f)...)
}

func (bst *SearchTreeData) MapInt(f func(int) int) []int {
	if bst == nil {
		return []int{}
	}
	return append(append(bst.left.MapInt(f), f(bst.data)), bst.right.MapInt(f)...)
}
