package offer

func replaceSpace(s string) string {
	//return strings.Replace(s, " ", "%20", -1)
	ans := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans = append(ans, "%20"...)
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
