package lc_14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var ret []byte
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for _, s := range strs[1:] {
			if i == len(s) || s[i] != c {
				return string(ret)
			}
		}
		ret = append(ret, c)
	}
	return string(ret)
}
