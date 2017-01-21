package strings

import (
	"strings"
)

func IndexNum(s, sep string, num int) int {
	index := -1
	sum := 0
	for count := 1; count <= num; count++ {
		i := strings.Index(s, sep)
		if i < 0 {
			break
		}
		sum += i
		if count == num {
			index = sum
			break
		}
		s = s[i+1:]
		sum += 1
	}
	return index
}
