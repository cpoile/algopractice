package ch3_7

var str1, str2 string

func EditDistance(s1, s2 string) int {
	str1, str2 = s1, s2
	return recEdit(len(s1), len(s2))
}

func EditDistanceDP(s1, s2 string) int {
	str1, str2 = s1, s2
	return editDP()
}

func recEdit(i, j int) int {
	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}
	ins := recEdit(i, j-1) + 1
	del := recEdit(i-1, j) + 1
	subsVal := 1
	if str1[i-1] == str2[j-1] {
		subsVal = 0
	}
	subs := recEdit(i-1, j-1) + subsVal
	return min(ins, del, subs)
}

func editDP() int {
	iMax, jMax := len(str1)+1, len(str2)+1

	tab := make([][]int, iMax)
	for i := 0; i < iMax; i++ {
		tab[i] = make([]int, jMax)
	}
	// base case when i = 0
	for j := 0; j < jMax; j++ {
		tab[0][j] = j
	}
	// base case when j = 0
	for i := 0; i < iMax; i++ {
		tab[i][0] = i
	}

	for i := 1; i < iMax; i++ {
		for j := 1; j < jMax; j++ {
			ins := tab[i][j-1] + 1
			del := tab[i-1][j] + 1
			subsVal := 1
			if str1[i-1] == str2[j-1] {
				subsVal = 0
			}
			subs := tab[i-1][j-1] + subsVal
			tab[i][j] = min(ins, del, subs)
		}
	}
	return tab[len(str1)][len(str2)]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}
