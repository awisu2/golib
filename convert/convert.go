package convert

import (
	"strconv"
)

var NumString = [...]string{"０", "１", "２", "３", "４", "５", "６", "７", "８", "９"}

// 文字列の数字を日本語に変換
func JapaneseNumString(s string) (num string, err error) {
	for i := 0; i < len(s); i++ {
		idx, _ := strconv.Atoi(s[i : i+1])
		if err != nil {
			return "", nil
		}
		num += NumString[idx]
	}
	return
}
