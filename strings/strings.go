package strings

import (
	"strings"
)

// exec strings.Index some time by num
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

func RuneLen(s string) int {
	return len([]rune(s))
}

func RuneSubString(s string, start int, end int) string {
	r := []rune(s)
	r = r[start:end]
	return string(r)
}
