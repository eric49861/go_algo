package offer

import (
	"math"
	"strings"
)

func strToInt(str string) int {
	//忽略开头的空格
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	//如果存在正负号
	i := 0
	isPositive := 1
	if str[i] == '+' || str[i] == '-' {
		if str[i] == '-' {
			isPositive = -1
		}
		i++
	}
	var ans int = 0
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		temp := int(str[i] - '0')
		//判断溢出
		if ans > math.MaxInt32/10 || ((ans == math.MaxInt32/10) && (temp > math.MaxInt32%10)) {
			ans = math.MaxInt32
			if isPositive < 0 {
				ans += 1
			}
			break
		}
		ans *= 10
		ans += temp
		i++
	}
	return ans * isPositive
}
